package garage

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
	"bytes"
	"strings"
)

func ValidateAndGetToken() (string, error) {
	vr := ValidateResponse{}
	client := &http.Client{}

	username := os.Getenv("catmulltrim@gmail.com")
	password := os.Getenv("Im2L84Dinner@jacks")

	validation := Validate{username, password}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(validation)

	req, _ := http.NewRequest("POST", "https://myqexternal.myqdevice.com/api/v4/User/Validate", b)
	req.Header.Set("MyQApplicationId", "OA9I/hgmPHFp9RYKJqCKfwnhh28uqLJzZ9KOJf1DXoo8N2XAaVX6A1wcLYyWsnnv")
	req.Header.Set("Content-Type","application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return "", err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal(resp_body, &vr)

	return vr.SecurityToken, err
}

func OpenGarage(firstname, lastname, phone string) string {
	Notify(fmt.Sprintf("Viewer (%s, %s, %s) Opened Garage Door", firstname, lastname, phone))
	return ChangeStatusOfGarage("1")
}

func CloseGarage(firstname, lastname, phone string) string {
	Notify(fmt.Sprintf("Viewer (%s, %s, %s) Closed Garage Door", firstname, lastname, phone))
	return ChangeStatusOfGarage("0")
}

func ChangeStatusOfGarage(state string) string {
	token,_ := ValidateAndGetToken()

	url := "https://myqexternal.myqdevice.com/api/v4/DeviceAttribute/PutDeviceAttribute"

	payload := strings.NewReader(fmt.Sprintf("{\n\t\"attributeName\": \"desireddoorstate\",\n\t\"myQDeviceId\": \"1702462521\",\n\t\"AttributeValue\": \"%s\"\n}", state))

	req, _ := http.NewRequest("PUT", url, payload)
	req.Header.Set("myqapplicationid", "OA9I/hgmPHFp9RYKJqCKfwnhh28uqLJzZ9KOJf1DXoo8N2XAaVX6A1wcLYyWsnnv")
	req.Header.Set("securitytoken", token)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Set("content-type","application/json")

	resp, _ := http.DefaultClient.Do(req)

	defer resp.Body.Close()

	body,_ := ioutil.ReadAll(resp.Body)

	return string(body)
}