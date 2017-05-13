package dotengine


import  (
    "testing"
    "log"
)

func TestGetToken(t *testing.T) {

	dotEngine := New("dotcc", "dotcc")

	tokenInfo, err := dotEngine.Token("room", "userid", DefaultExpires)

    log.Println(err)

    log.Println(tokenInfo)

	if err != nil {
		t.Fail()
	}

}
