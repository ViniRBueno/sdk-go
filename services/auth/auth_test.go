package auth

import (
	"testing"
	"time"
)

func TestAuthenticate(t *testing.T) {

	t.Parallel()
	expectValue := "123"

	t.Log("this test needs a 10 seconds about")

	go throwCallbackHit()

	result, error := Authenticate("41be4860-940e-45f3-b8dc-1df4884d1463")

	if error != nil {
		t.Error("token is null")
		return
	}

	if result.AccessToken != expectValue {
		t.Errorf("Expected: %s > result: %s", expectValue, result.AccessToken)
		return
	}

}

func throwCallbackHit() {

	time.Sleep(10 * time.Second)

	OpenBrowser("http://localhost:8081/callback#id_token=123&state=123&session_state=123")
}
