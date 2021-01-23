### Beat the bots 

It is near impossible to get hold of a PS5 either due to "scalpers" or being too slow! By the time I get a twitter notification I seem to be at the back of a long queue and not having much success. Instead of being stuck refreshing various retailer webpages I decided to write some Golang to do it for me. If the program notices that a search string no longer appears (often "out of stock") then it will display lots of text into the console and send a text message via Twilio. 

Mileage may very running this and lots of room for improvment. 
- Display an alert instead of dumping text into the terminal (just used this as a basic way to grab attention if my machine is open)
- Pretty basic way of checking the stock. Search string on a webpage so may not be the most reliable.
- Not currently deployed anywhere therefore only running when I run it locally 
- Could refactor the code so have a generic website scraper function instead of one for each provider 

### Running the code 

You will need to sign up for a Twilio account which you can do [here](https://www.twilio.com/). Once you have signed up you will be able to grab the `ACCOUNT_SID` , `AUTH_TOKEN` and `FROM_NUMBER`

Export the following values 
```bash
export ACCOUNT_SID="TWILIO_ACCOUNT_SID"
export AUTH_TOKEN="TWILIO_AUTH_TOKEN"
export TO_NUMBER="+447777777777"
export FROM_NUMBER="+447777777777"
```

Run the program 

```bash
go run main.go
```
