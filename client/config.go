package client

const steamEndpoint string = "http://api.steampowered.com"

var steamAPIKey string = ""

func SetAPIKey(key string) {
	steamAPIKey = key
}

func GetAPIKey() string {
	return steamAPIKey
}

func GetEndpoint() string {
	return steamEndpoint
}
