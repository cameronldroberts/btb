package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/go-rod/bypass"
	"github.com/go-rod/rod"
)

var (
	gamecount     int
	aocount       int
	argoscount    int
	shoptocounter int
)

func main() {
	fmt.Println("Beat the bot running....")
	client := &http.Client{}

	for {
		argos(*client)
		ao(*client)
		game(*client)
		shopto(*client)
		// smyths(*client)
		// asda(*client)
		// currys(*client)
		// amazon(*client)

		time.Sleep(5 * time.Second)
	}

}
func smyths(client http.Client) error {
	// time.Sleep(time.Duration(rand.Intn(15)) * time.Second)
	browser := rod.New().Timeout(time.Minute).MustConnect()
	defer browser.MustClose()
	page := bypass.MustPage(browser)
	page.Timeout(time.Minute).MustNavigate("https://www.smythstoys.com/uk/en-gb/video-games-and-tablets/playstation-5/playstation-5-consoles/playstation-5-console/p/191259")

	element, err := page.Timeout(10 * time.Second).Element("body > div:nth-child(10) > section > div > div > div:nth-child(2) > div.detail_right.margn_tp_n > div.yCmsContentSlot.page-details-variants-select.js-addtocart-channal > div > div > div > div.AddToCart-PickUpInStoreAction > div.row > div.col-xs-6.js-changeDeliveryChannel.radioControl.option-selected > p")
	if err != nil {
		fmt.Println("smyths ", err)
		page.MustScreenshot("")
		return err
	}

	elementText, err := element.Text()
	regx := regexp.MustCompile("Out Of Stock")
	matches := regx.FindAllStringIndex(elementText, -1)

	if len(matches) != 1 {
		twilio(client, "SMYTHS", "https://www.smythstoys.com/uk/en-gb/video-games-and-tablets/playstation-5/playstation-5-consoles/playstation-5-console/p/191259")
		fmt.Println(time.Now().Clock())
		fmt.Println("https://www.smythstoys.com/uk/en-gb/video-games-and-tablets/playstation-5/playstation-5-consoles/playstation-5-console/p/191259 #######################################################################################")
		time.Sleep(1 * time.Minute)

	}
	return nil
}
func asda(client http.Client) error {
	browser := rod.New().Timeout(time.Minute).MustConnect()
	defer browser.MustClose()
	page := bypass.MustPage(browser)
	page.MustNavigate("https://direct.asda.com/george/toys-character/gaming/gaming-consoles/playstation5-console/050887006,default,pd.html")
	element, err := page.Timeout(10 * time.Second).Element("#main-header")
	if err != nil {
		fmt.Println("asda", err)
		page.MustScreenshot("")
		return err
	}
	elementText, err := element.Text()
	if elementText != "In a spin?" {
		fmt.Println("in stock ")
		twilio(client, "ASDA", "https://direct.asda.com/george/toys-character/gaming/gaming-consoles/playstation5-console/050887006,default,pd.html")
		fmt.Println(time.Now().Clock())
		fmt.Println("https://direct.asda.com/george/toys-character/gaming/gaming-consoles/playstation5-console/050887006,default,pd.html ######################################################################################################")

	}
	return nil
}

func shopto(client http.Client) error {

	time.Sleep(time.Duration(rand.Intn(15)) * time.Second)
	browser := rod.New().Timeout(time.Minute).MustConnect()
	defer browser.MustClose()
	page := bypass.MustPage(browser)
	page.Timeout(time.Minute).MustNavigate("https://www.shopto.net/en/search/?input_search=PlayStation%205%20Console&utm_source=Website&utm_campaign=PS5%20Console%20Xl")

	element, err := page.Timeout(10 * time.Second).Element("#shop_content > div")
	if err != nil {
		fmt.Println("shopto ", err)
		page.MustScreenshot("")
		return err
	}

	elementText, err := element.Text()
	regx := regexp.MustCompile("Sold out")
	matches := regx.FindAllStringIndex(elementText, -1)
	if len(matches) != 4 {
		if shoptocounter == 0 {
			twilio(client, "SHOPTO", "https://www.shopto.net/en/search/?input_search=PlayStation%205%20Console&utm_source=Website&utm_campaign=PS5%20Console%20Xl")
			fmt.Println(time.Now().Clock())
			fmt.Println("https://www.shopto.net/en/search/?input_search=PlayStation%205%20Console&utm_source=Website&utm_campaign=PS5%20Console%20Xl #######################################################################################")
			time.Sleep(1 * time.Minute)
			shoptocounter++
		}

	}
	return nil
}
func amazon(client http.Client) error {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	browser := rod.New().Timeout(time.Minute).MustConnect()
	defer browser.MustClose()
	page := bypass.MustPage(browser)
	page.MustNavigate("https://www.amazon.co.uk/dp/B08H95Y452/ref=twister_B08J4RCVXW?_encoding=UTF8&psc=1")
	element, err := page.Timeout(10 * time.Second).Element("#availability")
	if err != nil {
		fmt.Println("check screenshots ", err)
		page.MustScreenshot("")
		return err
	}
	availability, err := element.Text()
	if strings.Contains(availability, "In stock.") {
		fmt.Println("in stock ")
		twilio(client, "AMAZON", "https://www.amazon.co.uk/dp/B08H95Y452/ref=twister_B08J4RCVXW?_encoding=UTF8&psc=1")
		fmt.Println(time.Now().Clock())
		fmt.Println("https://www.amazon.co.uk/dp/B08H95Y452/ref=twister_B08J4RCVXW?_encoding=UTF8&psc=1 ######################################################################################################")
	}
	return nil
}

func currys(client http.Client) error {
	browser := rod.New().Timeout(time.Minute).MustConnect()
	defer browser.MustClose()
	page := bypass.MustPage(browser)
	page.MustNavigate("https://www.currys.co.uk/gbuk/playstation-5-sony-1714-commercial.html")
	page.MustElement("#onetrust-accept-btn-handler").MustClick()

	element, err := page.Timeout(10 * time.Second).Element("#dsgContent > div.sold-out-banner")
	if err != nil {
		fmt.Println("check screenshots ", err)
		page.MustScreenshot("")
		return err
	}
	if element == nil {
		fmt.Println("in stock ")
		twilio(client, "CURRYS", "https://www.currys.co.uk/gbuk/playstation-5-sony-1714-commercial.html")
		fmt.Println(time.Now().Clock())
		fmt.Println("https://www.currys.co.uk/gbuk/playstation-5-sony-1714-commercial.html ######################################################################################################")

	} else {
		fmt.Println("CURRYS")
		fmt.Println("error with request\n", err)
	}
	return nil
}

func ao(client http.Client) error {
	url := "https://ao.com/brands/playstation?cmredirectionvalue=ps5"
	req, _ := http.NewRequest("GET", url, nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		bodyString := string(bodyBytes)
		regx := regexp.MustCompile("Sorry, PlayStation 5 is currently unavailable")
		matches := regx.FindAllStringIndex(bodyString, -1)
		if len(matches) != 1 {
			if aocount == 0 {
				twilio(client, "AO.COM", "https://ao.com/brands/playstation?cmredirectionvalue=ps5")
				fmt.Println(time.Now().Clock())
				fmt.Println("https://ao.com/brands/playstation?cmredirectionvalue=ps5 ######################################################################################################")
				aocount++
			}
		}
	} else {
		fmt.Println("ao.com")
		fmt.Println("error with request\n", err)
	}
	return nil
}

func argos(client http.Client) error {
	url := "https://www.argos.co.uk/product/8349000/"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"87\", \" Not;A Brand\";v=\"99\", \"Chromium\";v=\"87\"")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 11_1_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		regx := regexp.MustCompile("Sorry, PlayStation®5 is currently unavailable.")
		matches := regx.FindAllStringIndex(bodyString, -1)
		if len(matches) != 1 {
			if argoscount == 0 {
				twilio(client, "ARGOS", "https://www.argos.co.uk/product/8349000")
				fmt.Println(time.Now().Clock())
				fmt.Println("https://www.argos.co.uk/product/8349000 ######################################################################################################")
				argoscount++
			}
		}
	} else {
		fmt.Println("ARGOS")
		fmt.Println("error with request\n", err)

	}
	return nil
}

func game(client http.Client) error {
	url := "https://assets.game.net/_master/hardwarePages/playstation5/merch/LaunchCopyAvail/ps5.js?_=1610833181756"
	req, _ := http.NewRequest("GET", url, nil)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if res.StatusCode >= 200 && res.StatusCode < 300 {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		bodyString := string(bodyBytes)
		regx := regexp.MustCompile("Out of stock")
		matches := regx.FindAllStringIndex(bodyString, -1)
		if len(matches) != 2 {

			if gamecount == 0 {
				twilio(client, "GAME", "https://www.game.co.uk/playstation-5")
				fmt.Println(time.Now().Clock())
				fmt.Println("https://www.game.co.uk/playstation-5 ######################################################################################################")
				fmt.Println("https://www.game.co.uk/en/hardware/playstation-5/?sortBy=PRICE_DESC&cm_sp=PlayStationHardwarePage-_-Hardware-_-825gbconsole ######################################################################################################")
				gamecount++
			}
		}
	} else {
		fmt.Println("game")
		fmt.Println("error with request\n", err)
	}
	return nil
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

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("twilio", err)
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println("twilio", resp.Status)
	}
}
