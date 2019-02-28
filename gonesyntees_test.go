package gonesyntees

import (
	"testing"
)

// TestSuccess tests if a successful response returns correct data
func TestSuccess(t *testing.T) {
	response, err := Request("this is a test", Tõnu, 0)

	if err != nil {
		t.Error(err.Error())
	}

	if len(response.MP3Url) == 0 {
		t.Error("the MP3 url is empty")
	}

	if len(response.WAVUrl) == 0 {
		t.Error("the WAV url is empty")
	}
}

// TestEmptyText tests if the request fails with an empty text
func TestEmptyText(t *testing.T) {
	_, err := Request("", Tõnu, 0)

	if err == nil {
		t.Fail()
	}
}

// TestInvalidVoice tests if the request fails with an invalid voice ID
func TestInvalidVoice(t *testing.T) {
	_, err := Request("this is a test", 4, 0)

	if err == nil {
		t.Fail()
	}

	_, err = Request("this is a test", -1, 0)

	if err == nil {
		t.Fail()
	}
}

// TestInvalidSpeed tests if the request fails with an invalid speed
func TestInvalidSpeed(t *testing.T) {
	_, err := Request("this is a test", Tõnu, 10)

	if err == nil {
		t.Fail()
	}

	_, err = Request("this is a test", Tõnu, -10)

	if err == nil {
		t.Fail()
	}
}
