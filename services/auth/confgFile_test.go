package auth

import "testing"

func TestSaveYamlConfigToken(t *testing.T) {
	t.Parallel()

	var token = getToken()

	var outputs = []string{"./test.yaml", "./test2.yaml"}

	for _, v := range outputs {

		error := SaveYamlConfigToken(token, v)

		if error != nil {
			t.Error("token is null")
			return
		}
	}
}

func getToken() *Token {
	return &Token{
		AccessToken:  "123test",
		State:        "1test",
		SessionState: "123test",
	}
}
