package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("Hello")
	// call api google translate
	url := "https://google-translate1.p.rapidapi.com/language/translate/v2"

	payload := strings.NewReader("q=I love you&target=de&source=en")

	req, _ := http.NewRequest("POST", url, payload)

	fmt.Println(req)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", "0d90c3837dmsh10d2c2a98415ec6p12b042jsn7273754ab8ad")
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
