package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"

	speech "cloud.google.com/go/speech/apiv1"
	porcupine "github.com/charithe/porcupine-go"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

type keywordFlags []*porcupine.Keyword

func (kf *keywordFlags) Set(v string) error {
	parts := strings.Split(v, ":")
	if len(parts) != 3 {
		return errors.New("expected flag value to contain the keyword, filepath and sensitivity separated by colon charcters")
	}

	sensitivity, err := strconv.ParseFloat(parts[2], 32)
	if err != nil {
		return err
	}

	*kf = append(*kf, &porcupine.Keyword{Value: parts[0], FilePath: parts[1], Sensitivity: float32(sensitivity)})
	return nil
}

func (kf *keywordFlags) String() string {
	var sb strings.Builder
	for _, k := range *kf {
		sb.WriteString(fmt.Sprintf("%s:%s:%f", k.Value, k.FilePath, k.Sensitivity))
		sb.WriteString(", ")
	}
	return sb.String()
}

func main() {
	log.Printf("Starting!")
	var input string
	var modelPath string
	var keywords keywordFlags

	flag.StringVar(&input, "input", "-", "Path to read input audio from (PCM 16-bit LE)")
	flag.StringVar(&modelPath, "model_path", "", "Path to the Porcupine model")
	flag.Var(&keywords, "keyword", "Colon separated keyword, data file and sensitivity values (Eg. pineapple:pineapple_linux.ppn:0.5)")
	flag.Parse()

	if input == "" || modelPath == "" || len(keywords) == 0 {
		fmt.Fprintln(os.Stderr, "Usage: ./demo -input=<path_to_audio_input_data> -model_path=<path_to_model> -keyword=<keyword:path_to_data_file:sensitivity>")
		os.Exit(2)
	}

	p, err := porcupine.New(modelPath, keywords...)
	if err != nil {
		log.Fatalf("failed to initialize porcupine: %+v", err)
	}
	defer p.Close()

	var audio io.Reader
	if input == "-" {
		audio = bufio.NewReader(os.Stdin)
	} else {
		f, err := os.Open(input)
		if err != nil {
			log.Fatalf("failed to open input [%s]: %+v", input, err)
		}
		defer f.Close()

		audio = bufio.NewReader(f)
	}

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt)

	listen(p, audio, shutdownChan)
}

func listen(p porcupine.Porcupine, audio io.Reader, shutdownChan <-chan os.Signal) {

	// == Setup google voice API
	ctx := context.Background()

	// [START speech_transcribe_streaming_mic]
	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	stream, err := client.StreamingRecognize(ctx)
	if err != nil {
		log.Fatal(err)
	}

	frameSize := porcupine.FrameLength()
	audioFrame := make([]int16, frameSize)
	buffer := make([]byte, frameSize*2)

	log.Printf("listening...")

	for {
		select {
		case <-shutdownChan:
			log.Printf("shutting down")
			return
		default:
			if err := readAudioFrame(audio, buffer, audioFrame); err != nil {
				log.Printf("error: %+v", err)
				return
			}

			word, err := p.Process(audioFrame)
			if err != nil {
				log.Printf("error: %+v", err)
				continue
			}

			if word != "" {
				log.Printf("detected word: \"%s\"", word)
				log.Printf("Initiating up voice recognition")
				// Send the initial configuration message.
				if err := stream.Send(&speechpb.StreamingRecognizeRequest{
					StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
						StreamingConfig: &speechpb.StreamingRecognitionConfig{
							Config: &speechpb.RecognitionConfig{
								Encoding:        speechpb.RecognitionConfig_LINEAR16,
								SampleRateHertz: 16000,
								LanguageCode:    "en-US",
							},
						},
					},
				}); err != nil {
					log.Fatal(err)
				}

				log.Printf("Starting voice stream...")
				apiStreamStopChan := make(chan bool)
				go func() {
					// Pipe stdin to the API.
					buf := make([]byte, 1024)
					for {
						select {
						case <-apiStreamStopChan:
							return
						default:
							n, err := os.Stdin.Read(buf)
							if n > 0 {
								if err := stream.Send(&speechpb.StreamingRecognizeRequest{
									StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
										AudioContent: buf[:n],
									},
								}); err != nil {
									log.Printf("Could not send audio: %v", err)
								}
							}
							if err == io.EOF {
								// Nothing else to pipe, close the stream.
								if err := stream.CloseSend(); err != nil {
									log.Fatalf("Could not close stream: %v", err)
								}
								return
							}
							if err != nil {
								log.Printf("Could not read from stdin: %v", err)
								continue
							}
						}
					}
				}()

				resp, err := stream.Recv()
				if err == io.EOF {
					continue
				}
				if err != nil {
					log.Fatalf("Cannot stream results: %v", err)
				}
				if err := resp.Error; err != nil {
					// Workaround while the API doesn't give a more informative error.
					if err.Code == 3 || err.Code == 11 {
						log.Print("WARNING: Speech recognition request exceeded limit of 60 seconds.")
					}
					log.Fatalf("Could not recognize: %v", err)
				}
				for _, result := range resp.Results {
					fmt.Printf("Result: %+v\n", result)
				}
				// [END speech_transcribe_streaming_mic]
			}
		}
	}
}

func readAudioFrame(src io.Reader, buffer []byte, audioFrame []int16) error {
	_, err := io.ReadFull(src, buffer)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(buffer)
	for i := 0; i < len(audioFrame); i++ {
		if err := binary.Read(buf, binary.LittleEndian, &audioFrame[i]); err != nil {
			return err
		}
	}

	return nil
}
