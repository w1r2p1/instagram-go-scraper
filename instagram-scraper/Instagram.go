package instagram_scraper

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"log"
	"strings"
)

func GetAccoutByUsername(username string) (account Account) {
	url := fmt.Sprintf(ACCOUNT_JSON_INFO, username)
	info, err := _GetJsonFromUrl(url)
	if err != nil {
		log.Fatal(err)
	}
	account = GetFromAccountPage(info)
	return account
}

func GetMedyaByUrl(url string) (media Media) {
	code := strings.Split(url, "/")[4]
	media = GetMedyaByCode(code)
	return
}

func GetMedyaByCode(code string) (media Media) {
	url := fmt.Sprintf(MEDIA_JSON_INFO, code)

	info, err := _GetJsonFromUrl(url)
	if err != nil {
		log.Fatal(err)
	}
	media = GetFromMediaPage(info)

	return
}

func _GetJsonFromUrl(url string) (json_body map[string]interface{}, err error) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode == 404 {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &json_body)
	if err != nil {
		return nil, err
	}

	return
}