// https://youtu.be/K8w4xCX18Cg

package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main2() {
	var g errgroup.Group

	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
	}

	for _, url := range urls {
		url := url

		g.Go(func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}

			defer resp.Body.Close()

			fmt.Printf("%s: %s\n", url, resp.Status)

			return nil
		})
	}

	err := g.Wait() // faz trabalho de waitGroup
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Todos os sites foram acessados com sucesso\n")
}

/*
type Response struct {
	Status string
	Error  error
}

func main() {
	var wg sync.WaitGroup

	urls := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.twitter.com",
	}

	googleCh := make(chan Response)
	facebookCh := make(chan Response)
	twitterCh := make(chan Response)

	for _, url := range urls {
		url := url

		go func(url string) {
			wg.Add(1)
			defer wg.Done()

			resp, err := http.Get(url)

			defer resp.Body.Close()

			switch url {
			case "https://www.google.com":
				googleCh <- Response{Status: resp.Status, Error: err}
			case "https://www.facebook.com":
				facebookCh <- Response{Status: resp.Status, Error: err}
			case "https://www.twitter.com":
				twitterCh <- Response{Status: resp.Status, Error: err}
			}

			fmt.Printf("%s: %s\n", url, resp.Status)
		}(url)
	}

	for {
		select {
		case resp := <-googleCh:
			fmt.Printf("Google: %s\n", resp.Status)
		case resp := <-facebookCh:
			fmt.Printf("Facebook: %s\n", resp.Status)
		case resp := <-twitterCh:
			fmt.Printf("Twitter: %s\n", resp.Status)
		}
	}
}
*/
