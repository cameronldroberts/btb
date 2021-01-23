package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	fmt.Println("Beat the bot running....")
	client := &http.Client{}
	for {
		currys(*client)
		argos(*client)
		ao(*client)
		game(*client)
		time.Sleep(10 * time.Second)
	}

}
func currys(client http.Client) {
	url := "https://www.currys.co.uk/gbuk/playstation-5-sony-1714-commercial.html"
	req, _ := http.NewRequest("GET", url, nil)
	res, err := client.Do(req)
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		bodyString := string(bodyBytes)
		regx := regexp.MustCompile("Playstation 5 sold out")
		matches := regx.FindAllStringIndex(bodyString, -1)
		if len(matches) != 1 {
			twilio(client, "CURRYS", "https://www.currys.co.uk/gbuk/playstation-5-sony-1714-commercial.html")
			for {
				fmt.Println("https://www.currys.co.uk/gbuk/playstation-5-sony-1714-commercial.html ######################################################################################################")
				time.Sleep(50 * time.Millisecond)
			}
		}
	} else {
		fmt.Println("CURRYS")
		fmt.Println("error with request\n", err)
	}

}

func ao(client http.Client) {
	url := "https://ao.com/brands/playstation?cmredirectionvalue=ps5"
	req, _ := http.NewRequest("GET", url, nil)
	res, err := client.Do(req)
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		bodyString := string(bodyBytes)
		regx := regexp.MustCompile("Sorry, PlayStation 5 is currently unavailable")
		matches := regx.FindAllStringIndex(bodyString, -1)
		if len(matches) != 1 {
			twilio(client, "AO.COM", "https://ao.com/brands/playstation?cmredirectionvalue=ps5")
			for {
				fmt.Println("https://ao.com/brands/playstation?cmredirectionvalue=ps5 ######################################################################################################")
				time.Sleep(50 * time.Millisecond)
			}
		}
	} else {
		fmt.Println("ao.com")
		fmt.Println("error with request\n", err)
	}

}

func argos(client http.Client) {
	url := "https://www.argos.co.uk/product/8349000/"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"87\", \" Not;A Brand\";v=\"99\", \"Chromium\";v=\"87\"")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	res, err := client.Do(req)
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		regx := regexp.MustCompile("Sorry, PlayStation®5 is currently unavailable.")
		matches := regx.FindAllStringIndex(bodyString, -1)
		if len(matches) != 1 {
			twilio(client, "ARGOS", "https://www.argos.co.uk/product/8349000")
			for {
				fmt.Println("https://www.argos.co.uk/product/8349000 ######################################################################################################")
				time.Sleep(50 * time.Millisecond)
			}
		}
	} else {
		fmt.Println("ARGOS")
		fmt.Println("error with request\n", err)
	}

}

func game(client http.Client) {
	url := "https://assets.game.net/_master/hardwarePages/playstation5/merch/LaunchCopyAvail/ps5.js?_=1610833181756"
	req, _ := http.NewRequest("GET", url, nil)
	res, err := client.Do(req)
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		bodyString := string(bodyBytes)
		regx := regexp.MustCompile("Out of stock")
		matches := regx.FindAllStringIndex(bodyString, -1)
		if len(matches) != 2 {
			twilio(client, "GAME", "https://www.game.co.uk/playstation-5")
			for {
				fmt.Println("https://www.game.co.uk/playstation-5 ######################################################################################################")
				time.Sleep(50 * time.Millisecond)
			}
		}
	} else {
		fmt.Println("game")
		fmt.Println("error with request\n", err)
	}

}

func twilio(client http.Client, provider string, providerURL string) {
	accountSid := os.Getenv("ACCOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	toNumber := os.Getenv("TO_NUMBER")
	fromNumber := os.Getenv("FROM_NUMBER")
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	body := "\nPS5 stock!!" + "\n" + provider + "\n" + providerURL
	msgData := url.Values{}
	msgData.Set("To", toNumber)
	msgData.Set("From", fromNumber)
	msgData.Set("Body", body)
	msgDataReader := *strings.NewReader(msgData.Encode())

	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
}
