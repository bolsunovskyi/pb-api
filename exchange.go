package pb_api

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"

	"strconv"
)
//api documentation https://link.privatbank.ua/console/wiki/p24business_kurs_valut

const RATE_NBU int = 3
const RATE_PB int = 5

type ExchangeRate struct {
	CCY		string		`json:"@ccy"`
	BaseCCY		string		`json:"@base_ccy"`
	BuySTR		string		`json:"@buy"`
	SaleSTR		string		`json:"@sale"`
}

func (r ExchangeRate) GetBuy() float64 {
	b, _ := strconv.ParseFloat(r.BuySTR, 64)
	return b
}

func (r ExchangeRate) GetSale() float64 {
	b, _ := strconv.ParseFloat(r.SaleSTR, 64)
	return b
}

type Rate struct {
	ExchangeRate	ExchangeRate	`json:"exchangerate"`
}

func GetExchangeRate(rateID int, sessionID string) (*[]Rate, error) {
	rates := make([]Rate, 0)

	getURL := fmt.Sprintf(url + "kursValut?exchange=true&coursid=%d", rateID)

	rq, err := http.NewRequest("GET", getURL, nil)
	if err != nil {
		return nil, err
	}
	rq.Header.Add("Accept", "application/json")
	rq.Header.Add("Authorization", fmt.Sprintf("Token %s", sessionID))

	rsp, err := http.DefaultClient.Do(rq)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, errors.New(string(b))
	}

	err = json.NewDecoder(rsp.Body).Decode(&rates)
	if err != nil {
		return nil, err
	}

	return &rates, nil
}