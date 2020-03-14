package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func SendGet(link string) []byte {
	u, _ := url.Parse(link)
	res, err := http.Get(u.String())
	if err != nil {
		fmt.Println(err)
	}
	result, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	return result
}