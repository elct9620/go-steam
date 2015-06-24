package client

import "fmt"
import "strings"
import "net/http"
import "io/ioutil"
import "encoding/json"

type SteamRequest interface {
	Get([]Param)
}

type Request struct {
	Interface string
	Endpoint  string
	Version   int
	Params    []Param
}

type Param struct {
	key   string
	value string
}

func (r *Request) GetBody() []byte {
	requestUri := fmt.Sprintf("%s/%s/%s/v%04d/?format=json&%s", GetEndpoint(), r.Interface, r.Endpoint, r.Version, BuildQueryString(r.Params))

	res, _ := http.Get(requestUri)

	var body []byte
	defer res.Body.Close()
	body, _ = ioutil.ReadAll(res.Body)

	return body
}

func (r *Request) GetJSON(reslut interface{}) {
	json.Unmarshal(r.GetBody(), reslut)
}

func BuildQueryString(params []Param) string {
	var queries []string
	for k, v := range params {
		queries = append(queries, fmt.Sprintf("%s=%s", k, v))
	}
	return strings.Join(queries, "&")
}
