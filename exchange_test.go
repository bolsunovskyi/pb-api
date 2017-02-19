package pb_api

import (
	"testing"
)

func TestGetExchangeRate(t *testing.T) {
	Init(id, secret)
	sess, err := SessionCreate()
	if err != nil {
		t.Error(err)
		return
	}
	defer SessionRemove(sess.ID)

	_, err = GetExchangeRate(RATE_PB, sess.ID)
	if err != nil {
		t.Error(err)
	}
}
