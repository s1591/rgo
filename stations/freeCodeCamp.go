package stations

func init() {
	FreeCodeCamp.Register()
}

var FreeCodeCamp = CreateMainStation("FreeCodeCamp", freeCodeCamp)

var freeCodeCamp = RadioStation{
	Name:    "FreeCodeCamp Radio",
	Url:     "https://coderadio-admin-v2.freecodecamp.org/listen/coderadio/radio.mp3",
	Website: "https://coderadio.freecodecamp.org/",
}
