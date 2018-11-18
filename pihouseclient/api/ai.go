package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	wit "github.com/jsgoecke/go-wit"
)

func PostAIIntent(outcomes []wit.Outcome) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(outcomes)
	res, _ := http.Post("http://"+APIAddress+"/v1/api/ai/outcomes", "application/json; charset=utf-8", b)
	io.Copy(os.Stdout, res.Body)
}
