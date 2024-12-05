package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/jacobshade/lbuc-admin/server/model"
)

// GetEmail of user
func GetEmail(token string) string {
	reqURL, err := url.Parse("https://www.googleapis.com/oauth2/v1/userinfo")
	if err != nil {
		panic(err)
	}
	ptoken := fmt.Sprintf("Bearer %s", token)
	res := &http.Request{
		Method: "GET",
		URL:    reqURL,
		Header: map[string][]string{
			"Authorization": {ptoken},
		},
	}
	req, err := http.DefaultClient.Do(res)
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var data model.GoogleResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}
	return data.Email
}
