package speech

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
)

func Say(speech string) {
	// Instantiates a client.
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Perform the text-to-speech request on the text input with the selected
	// voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: speech},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		// voice gender ("neutral").
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "en-US",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_FEMALE,
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	// The resp's AudioContent is binary.
	filename := "output.mp3"
	err = ioutil.WriteFile(filename, resp.AudioContent, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Audio content written to file: %v\n", filename)

	// f, err := os.Open("output.mp3")

	// // Check for errors when opening the file
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Decode the .mp3 File, if you have a .wav file, use wav.Decode(f)
	// s, format, _ := mp3.Decode(f)

	// time.Sleep(time.Second)

	// // Init the Speaker with the SampleRate of the format and a buffer size of 1/10s
	// speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// time.Sleep(time.Second)

	// // Channel, which will signal the end of the playback.
	// playing := make(chan struct{})

	// // Now we Play our Streamer on the Speaker
	// speaker.Play(beep.Seq(s, beep.Callback(func() {
	// 	// Callback after the stream Ends
	// 	close(playing)
	// })))
	// <-playing

	err = exec.Command("vlc", "--play-and-exit", "output.mp3").Run()

	if err != nil {
		log.Panicln(err)
	}

	log.Println("done!")
}
