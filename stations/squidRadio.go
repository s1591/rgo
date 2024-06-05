package stations

func init() {
	SquidRadio.Register()
}

var SquidRadio = CreateMainStation("SquidRadio", 
    squidRadioPersona, squidRadioSonic)

var squidRadioPersona = RadioStation{
	Name:    "Squid radio (persona)",
	Url:     "https://play.squid-radio.net/persona",
	Website: "https://play.squid-radio.net/",
}

var squidRadioSonic = RadioStation{
	Name:    "Squid radio (sonic)",
	Url:     "https://play.squid-radio.net/sonic",
	Website: "https://play.squid-radio.net/",
}
