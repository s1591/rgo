package stations

func init() {
	Wqxr.Register()
}

var Wqxr = CreateMainStation("Wqxr", wqxr105)

var wqxr105 = RadioStation{
	Name:    "wqxr (105.9 FM)",
	Url:     "https://stream.wqxr.org/wqxr-web",
	Website: "https://www.wqxr.org/",
}
