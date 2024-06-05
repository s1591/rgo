package stations

func init() {
	SomaFM.Register()
}

var SomaFM = CreateMainStation("SomaFM", somaGrooveSalad)

var somaGrooveSalad = RadioStation{
	Name:    "SomaFM (Groove Salad)",
	Url:     "https://somafm.com/groovesalad256.pls",
	Website: "https://somafm.com/",
}
