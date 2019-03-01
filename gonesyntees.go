package gonesyntees

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Voice is the ID of the voice the synthesizer uses
type Voice int

const (
	// Eva is a female voice
	Eva Voice = 0
	// Tõnu is a male voice
	Tõnu Voice = 1
	// Liisi is a female voice
	Liisi Voice = 2
	// Riina is a female voice
	Riina Voice = 3
)

// Response contains the URLs of the audio files gathered from the API response
type Response struct {
	MP3Url string
	WAVUrl string
}

// Request dispatches a request to the API and returns a struct containing URLs to the audio files gathered from the response
func Request(text string, voice Voice, speed int) (*Response, error) {
	if len(text) == 0 {
		return nil, errors.New("the text can not be empty")
	}

	if voice < 0 || voice > 3 {
		return nil, errors.New("voice must be in the range of 0 .. 3")
	}

	if speed < -9 || speed > 9 {
		return nil, errors.New("speed must be in the range of -9 .. 9")
	}

	url := fmt.Sprintf("http://teenus.eki.ee/konesyntees?haal=%d&kiirus=%d&tekst=%s", voice, speed, url.QueryEscape(text))
	httpResponse, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(httpResponse.Body)

	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(strings.NewReader(string(data)))

	var urls map[string]string

	err = decoder.Decode(&urls)

	if err != nil {
		return nil, err
	}

	response := &Response{
		MP3Url: urls["mp3url"],
		WAVUrl: urls["wavurl"],
	}

	return response, nil
}
