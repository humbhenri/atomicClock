// Get the atomic time from http://en.uhrzeit.org/atomic-clock.php
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
    "regexp"
)

func main() {
	resp, err := http.Get("http://en.uhrzeit.org/atomic-clock.php")
	if err != nil {
		fmt.Println("Cannot get time from internet.")
	}
	defer resp.Body.Close()

    regex := regexp.MustCompile("(?:<div id=\"anzeige_zeit\">(.*)</div>)")

	body, err := ioutil.ReadAll(resp.Body)
    matches := regex.FindSubmatch(body)
    fmt.Println(string(matches[1]))

}
