package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"crypto/hmac"
	"crypto/sha256"

	"encoding/hex"
)

const (
	host = "https://coincheck.com"
)

type RateResponse struct {
	Rate string `json:"rate"`
}

func FetchRate(currencyType string) (string, error) {
	result, err := http.Get(host + "/api/rate/" + currencyType)
	if err != nil {
		return "", err
	}
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return "", err
	}

	var rateResponse RateResponse
	err = json.Unmarshal(body, &rateResponse)
	if err != nil {
		return "", err
	}

	return rateResponse.Rate, nil
}

func FetchBalance() ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", host+"/api/accounts/balance", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("ACCESS-KEY", "ACCESS_KEY")

	nonce := strconv.Itoa(int(time.Now().Unix()))
	req.Header.Add("ACCESS-NONCE", nonce)

	mac := hmac.New(sha256.New, []byte("SECRET_KEY"))
	mac.Write([]byte(nonce + host + "/api/accounts/balance"))
	req.Header.Add("ACCESS-SIGNATURE", hex.EncodeToString(mac.Sum(nil)))

	result, err := client.Do(req)
	defer result.Body.Close()

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

func BuyOrder() {

}

func SellOrder() {

}
