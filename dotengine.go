package dotengine

// date     :=  2017-05-13
// auther   :=  notedit

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"

	"github.com/dvsekhvalnov/jose2go"
)

const (
	DefaultExpires = 3600 * 24
	apiUrl         = "https://janus.dot.cc/api/"
	apiCreateToken = "createToken"
)

type DotEngine struct {
	appKey    string
	appSecret string
}

type TokenRes struct {
	Status int    `json:"s"`
	Token  Token  `json:"d"`
	Err    string `json:"e"`
}

type Token struct {
	Token string `json:"token"`
}

func New(appKey, appSecret string) *DotEngine {

	return &DotEngine{appKey: appKey,
		appSecret: appSecret}

}

func (dot *DotEngine) AppKey() string {

	return dot.appKey
}

func (dot *DotEngine) AppSecret() string {

	return dot.appSecret
}

func (dot *DotEngine) Token(room, userID string, expires int) (*Token, error) {

	token := map[string]interface{}{
		"room":    room,
		"user": userID,
		"expires":  expires,
		"role":    "",
		"nonce":   rand.Intn(9999999),
		"appkey": dot.AppKey(),
	}

	payload, err := json.Marshal(token)

	if err != nil {
		return nil, err
	}

	tokenData, err := jose.SignBytes(payload, jose.HS256, dot.AppSecret())

	if err != nil {
		log.Println("jwt token generate error")
		return nil, err
	}

	params := url.Values{}

	params.Add("sign", tokenData)
	params.Add("appkey", dot.AppKey())

	req, err := http.NewRequest("POST", apiUrl+apiCreateToken, strings.NewReader(params.Encode()))

	if err != nil {
		return nil, err
	}

	client := http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("dotengine create token error: %v", res.StatusCode)
	}

	var tokenRes TokenRes

	err = json.NewDecoder(res.Body).Decode(&tokenRes)

	if err != nil {
		return nil, err
	}

	return &tokenRes, nil

}
