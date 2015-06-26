package steam

import (
	"fmt"
	"time"
)

// App News
type AppNews struct {
	AppId     uint32
	NewsItems []NewsItem
}

type NewsItem struct {
	GID           uint64
	Title         string
	URL           string
	IsExternalURL bool
	Author        string
	Contents      string
	FeedLabel     string
	Date          *Timestamp
	FeedName      string
}

type AppNewsJSON struct {
	AppNews AppNews
}

func GetNewsForApp(appId uint32, maxlength uint32, enddate *time.Time, count uint32) AppNews {
	var params []Param
	var result AppNewsJSON
	params = append(params, Param{key: "appid", value: fmt.Sprint(appId)})
	if maxlength != 0 {
		params = append(params, Param{key: "maxlength", value: fmt.Sprint(maxlength)})
	}
	if enddate != nil {
		params = append(params, Param{key: "enddate", value: fmt.Sprint(enddate.Unix())})
	}
	if count != 0 {
		params = append(params, Param{key: "count", value: fmt.Sprint(count)})
	}

	req := SteamRequest{"ISteamNews", "GetNewsForApp", 2, params}
	req.Get(&result)

	return result.AppNews
}
