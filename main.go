package main

import (
"fmt"
"strings"
"net/http"
"io/ioutil"
)

func main() {

	url := "https://myqexternal.myqdevice.com/api/v4/DeviceAttribute/PutDeviceAttribute"

	payload := strings.NewReader("{\n\t\"attributeName\": \"desireddoorstate\",\n\t\"myQDeviceId\": \"1702462521\",\n\t\"AttributeValue\": \"0\"\n}")

	req, _ := http.NewRequest("PUT", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("myqapplicationid", "OA9I/hgmPHFp9RYKJqCKfwnhh28uqLJzZ9KOJf1DXoo8N2XAaVX6A1wcLYyWsnnv")
	req.Header.Add("securitytoken", "5e1e242a-9933-4fee-afb7-3bd5a6be73ac")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "70317a7e-c6cf-e86f-8dfc-c197b1847f7b")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}