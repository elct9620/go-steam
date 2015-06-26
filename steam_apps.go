package steam

// Steam APP Struct
type App struct {
	AppId int
	Name  string
}

// Steam API: AppList
type AppList struct {
	Apps []App
}

type AppListJSON struct {
	AppList AppList
}

func GetAppList() AppList {
	var result AppListJSON
	req := SteamRequest{"ISteamApps", "GetAppList", 2, nil}
	req.Get(&result)

	return result.AppList
}
