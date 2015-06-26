package steam

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const steamEndpoint string = "http://api.steampowered.com"

type Param struct {
	key   string
	value string
}

type SteamRequest struct {
	Interface string
	Endpoint  string
	Version   int
	Params    []Param
}

// API Key
var steamAPIKey string = ""

func SetAPIKey(key string) {
	steamAPIKey = key
}

func GetAPIKey() string {
	return steamAPIKey
}

// Get Steam API Object
func (r *SteamRequest) Get(result interface{}) {
	requestUri := fmt.Sprintf("%s/%s/%s/v%04d/?format=json&%s", steamEndpoint, r.Interface, r.Endpoint, r.Version, buildQueryString(r.Params))

	res, _ := http.Get(requestUri)

	var body []byte
	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)

	json.Unmarshal(body, result)
}

func buildQueryString(params []Param) string {
	var queries []string
	for _, v := range params {
		queries = append(queries, fmt.Sprintf("%s=%s", v.key, v.value))
	}
	return strings.Join(queries, "&")
}
