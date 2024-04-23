package event

import (
	"event_api/pkg/domain/convert"
	"testing"
	"time"
)

type testStruct struct {
	Time time.Time
}

func TestJsonUnmarshal(t *testing.T) {
	js := `
	{
		"Time": "2023-04-09 13:00:00"
	}
`
	_, err := convert.FromJSON[testStruct]([]byte(js))
	if err != nil {
		t.Error(err)
	}

	js = `
	{
  		"eventType": "login",
  		"userID": 1,
  		"eventTime": "2023-04-09T13:00:00Z",
  		"payload": "{\"some_field\":\"some_value\"}"
	}
`
	_, err = convert.FromJSON[UnhandledEvent]([]byte(js))
	if err != nil {
		t.Error(err)
	}
}
