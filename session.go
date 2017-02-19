package pb_api
//api documentation https://link.privatbank.ua/console/wiki/client_auth

import (
	"net/http"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"errors"
)

type SessionCreateSuccess struct {
	ID		string		`json:"id"`
	ClientID	string		`json:"clientId"`
	ExpiresIn	int		`json:"expiresIn"`
	Roles		[]string	`json:"roles"`
}

func SessionCreate() (*SessionCreateSuccess, error) {
	createURL := url + "auth/createSession"

	bts, err := json.Marshal(cred)
	if err != nil {
		return nil, err
	}

	rq, err := http.NewRequest("POST", createURL, bytes.NewReader(bts))
	if err != nil {
		return nil, err
	}
	rq.Header.Add("Content-Type", "application/json")
	rq.Header.Add("Accept", "application/json")

	rsp, err := http.DefaultClient.Do(rq)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, errors.New(string(b))
	}

	success := SessionCreateSuccess{}
	err = json.NewDecoder(rsp.Body).Decode(&success)
	if err != nil {
		return nil, err
	}

	return &success, nil
}

func SessionValidate(sessionID string) (*SessionCreateSuccess, error) {
	createURL := url + "auth/validateSession"

	rqBody := make(map[string]string)
	rqBody["sessionId"] = sessionID
	bts, err := json.Marshal(rqBody)
	if err != nil {
		return nil, err
	}

	rq, err := http.NewRequest("POST", createURL, bytes.NewReader(bts))
	if err != nil {
		return nil, err
	}
	rq.Header.Add("Content-Type", "application/json")
	rq.Header.Add("Accept", "application/json")

	rsp, err := http.DefaultClient.Do(rq)
	if err != nil {
		return nil, err
	}

	if rsp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(rsp.Body)
		return nil, errors.New(string(b))
	}

	success := SessionCreateSuccess{}
	err = json.NewDecoder(rsp.Body).Decode(&success)
	if err != nil {
		return nil, err
	}

	return &success, nil
}

func SessionRemove(sessionID string) error {
	createURL := url + "auth/removeSession"

	rqBody := make(map[string]string)
	rqBody["sessionId"] = sessionID
	bts, err := json.Marshal(rqBody)
	if err != nil {
		return err
	}

	rq, err := http.NewRequest("POST", createURL, bytes.NewReader(bts))
	if err != nil {
		return err
	}
	rq.Header.Add("Content-Type", "application/json")
	rq.Header.Add("Accept", "application/json")

	rsp, err := http.DefaultClient.Do(rq)
	if err != nil {
		return err
	}

	if rsp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(rsp.Body)
		return errors.New(string(b))
	}

	return nil
}
