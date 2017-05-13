package dotengine

import "testing"

func TestGetToken(t *testing.T) {

	dotEngine := New("dotcc", "dotcc")

	tokenInfo, _ := dotEngine.Token("room", "userid", DefaultExpires)

	if tokenInfo == nil {
		t.Fail()
	}

}
