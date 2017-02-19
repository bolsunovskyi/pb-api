package pb_api

import (
	"testing"
	"flag"
	"log"
)

var id, secret string

func init() {
	flag.StringVar(&id, "pbid", "", "PB client id")
	flag.StringVar(&secret, "pbsecret", "", "PB client secret")
	flag.Parse()

	if id == "" || secret == "" {
		log.Fatalln("Please provide -pbid and -pbsecret to run tests")
	}
}

func TestSessionCreate(t *testing.T) {
	Init("foo", "bar")
	_, err := SessionCreate()
	if err == nil {
		t.Error("No error on wrong credentials")
		return
	}

	Init(id, secret)
	session, err := SessionCreate()
	if err != nil {
		t.Error(err)
		return
	}

	_, err = SessionValidate(session.ID)
	if err != nil {
		t.Error(err)
		return
	}

	err = SessionRemove(session.ID)
	if err != nil {
		t.Error(err)
		return
	}
}
