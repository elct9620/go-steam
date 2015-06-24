package api

import "github.com/elct9620/go-steam/client"

// App Struct
type App struct {
	AppId int
	Name  string
}

// App List Struct
type AppList struct {
	AppList struct {
		Apps []App
	}
}

func (list *AppList) Get(params []client.Param) {
	req := client.Request{"ISteamApps", "GetAppList", 2, params}
	req.GetJSON(list)
}

// TODO: GetServersAtAddress Struct

// TODO: UpToDateCheck Struct
