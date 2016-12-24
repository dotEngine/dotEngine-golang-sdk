package dotengine

import "testing"

func TestGetToken(t *testing.T) {

	dotEngine := New("dotcc", "dotcc")

	token, _ := dotEngine.Token("room", "userid", DefaultExpires)

	if token == nil {
		t.Fail()
	}

}
