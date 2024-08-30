package html

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

// Titulo obtem título de uma página HTML
func Titulo(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			title := r.FindStringSubmatch(string(html))

			if len(title) > 0 {
				c <- r.FindStringSubmatch(string(html))[1]
			} else {
				mensagem := "Não foi possível encontrar titulo para a url: " + url
				c <- mensagem
			}
		}(url)

	}

	return c
}
