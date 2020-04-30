package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"math/rand"
	"net/http"
)

func GetRandomJoke() (joke string) {
	siteData, err := http.Get("https://www.anekdot.ru/random/anekdot/")
	if err != nil {
		fmt.Println(fmt.Errorf("Cannot collect data from the site: %v", err))
		return
	}

	if siteData.StatusCode != 200 {
		fmt.Println(fmt.Sprintf("Cannot collect data from the site: %d %v", siteData.StatusCode, siteData.Status))
		return
	}

	defer siteData.Body.Close()

	bodyData, err := goquery.NewDocumentFromReader(siteData.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("Cannot read the data from body: %v", err))
		return
	}

	allJokes := make([]string, 0, 20)

	bodyData.Find(".text").Each(func(_ int, oneJoke *goquery.Selection) {
		allJokes = append(allJokes, oneJoke.Text())
	})

	joke = allJokes[rand.Intn(20)]

	return
}
