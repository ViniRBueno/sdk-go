package auth

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"github.com/fatih/color"
)

// Token represents a token oauth information
type Token struct {
	AccessToken  string `yaml:"access_token"`
	State        string
	SessionState string
}

var (
	server *http.Server
)

// Authenticate uses the Ad OAUTH2.0 routine
func Authenticate(clientID string) (*Token, error) {

	// defer func() {
	// 	//go func() {
	// 	if err := server.Shutdown(context.Background()); err != nil {

	// 		yellow := color.New(color.FgYellow).SprintFunc()
	// 		red := color.New(color.FgRed).SprintFunc()

	// 		fmt.Printf("%s %s", yellow("error stopping server:"), red(err.Error()))
	// 	}
	// 	//}()
	// }()

	masterChannel := startLogin(clientID)

	go timeOutSoonAlert()

	select {
	case success := <-masterChannel:
		return success, nil

	case <-time.After(60 * time.Second):
		return nil, fmt.Errorf("TimedOut!%s", "")

	}
}

func timeOutSoonAlert() {
	time.Sleep(30 * time.Second)
	color.Red("...30 seconds remaining!")
}

func startLogin(clientID string) <-chan *Token {

	tChannel := make(chan *Token)

	if validLogin := initFrontEnd(clientID); !validLogin {
		tChannel <- nil
		return tChannel
	}

	go initServerCallBack(tChannel)

	return tChannel
}

func initServerCallBack(channel chan *Token) {

	m := http.NewServeMux()
	server = &http.Server{Addr: ":8081", Handler: m}

	m.Handle("/callback", handleImplicitGrant())
	m.Handle("/token", handleToken(channel))

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}

func handleImplicitGrant() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		printBodyCallback(w, `<h2>Generating token..</h2><strong>please wait.</strong>
		<script type='text/javascript'>
			if(window.location.hash) {
				window.location.href='http://localhost:8081/token' + window.location.hash.replace("#","?")
			} else {
				window.location.href='http://localhost:8081/token'
			}
		</script>
		`)

	})
}

func handleToken(channel chan *Token) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		idtoken := r.URL.Query().Get("id_token")

		if idtoken != "" {

			printBodyCallback(w, "<h1>Login successfuly.</h1> <h2>You can close this window.</h2>")

			channel <- &Token{
				AccessToken:  idtoken,
				State:        r.URL.Query().Get("state"),
				SessionState: r.URL.Query().Get("session_state"),
			}
		} else {
			http.Error(w, "Access token not found", http.StatusUnauthorized)
			channel <- nil
			return // don't call original handler
		}
	})
}

func initFrontEnd(clientID string) bool {
	url := fmt.Sprintf(`https://login.microsoftonline.com/common/oauth2/v2.0/authorize?client_id=%s&response_type=%s&redirect_uri=%s&scope=%s&response_mode=%s&state=%s&nonce=%s`,
		clientID, "id_token", "http://localhost:8081/callback", "openid", "fragment", randomString(20), randomString(6))

	info := color.New(color.FgGreen).SprintFunc()

	fmt.Printf("Openning login page at: %s\n", info(url))

	err := OpenBrowser(url)

	color.Blue("About 1 minute for login")

	return err == nil
}

// OpenBrowser opens a current os webbrowser with an specific url
func OpenBrowser(url string) error {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func printBodyCallback(w http.ResponseWriter, info string) {
	responseString := fmt.Sprintf(`<!DOCTYPE html>
	<html>
	<head>
	<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
	<title>%s</title>
	</head>
	<body>%s</body>
	</html>`, "GRUPO LTM CLI", info)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, responseString)
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}
