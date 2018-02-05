package garage

import (
	"net/http"
	"fmt"
	"strings"
)

func Notify(body string) {

	url := "https://api.twilio.com/2010-04-01/Accounts/ACeadb68c622751d374cc81198ac17f001/Messages.json"

	payload := strings.NewReader(fmt.Sprintf("------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"To\"\r\n\r\n+12083602737\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"From\"\r\n\r\n+12085381752\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW\r\nContent-Disposition: form-data; name=\"Body\"\r\n\r\n%s\r\n------WebKitFormBoundary7MA4YWxkTrZu0gW--", body))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW")
	req.Header.Add("authorization", "Basic QUNlYWRiNjhjNjIyNzUxZDM3NGNjODExOThhYzE3ZjAwMTpjYzgwYzVhMDdlNWMwMzkzMjk4MTMzMmViNDBjNDcxZQ==")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "2819eb7e-737d-54e3-4a47-6afd0bf9c70c")

	res, _ := http.DefaultClient.Do(req)

	res.Body.Close()
}
