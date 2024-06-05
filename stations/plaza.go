package stations

func init() {
	Plaza.Register()
}

var Plaza = CreateMainStation("NightWave Plaza", plazaOne)

var plazaOne = RadioStation{
	Name:    "Nightwave plaza",
	Url:     "https://radio.plaza.one/mp3",
	Website: "https://plaza.one/",
}
