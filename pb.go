package pb_api

const url string = "https://link.privatbank.ua/api/"

type credentials struct {
	ClientID	string		`json:"clientId"`
	ClientSecret	string		`json:"clientSecret"`
}

var cred credentials

func Init(clientID string, clientSecret string) {
	cred = credentials{
		ClientID:	clientID,
		ClientSecret:	clientSecret,
	}
}
