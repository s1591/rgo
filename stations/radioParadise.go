package stations

func init() {
	RadioParadise.Register()
}

var RadioParadise = CreateMainStation("RadioParadise",
	radioParadiseM, radioParadiseMM)

var radioParadiseMM = RadioStation{
	Name:    "Radio paradise (main mix)",
	Url:     "http://stream.radioparadise.com/flacm",
	Website: "https://radioparadise.com/",
}

var radioParadiseM = RadioStation{
	Name:    "Radio Paradise (mellow mix)",
	Url:     "http://stream.radioparadise.com/mellow-flacm",
	Website: "https://radioparadise.com/",
}
