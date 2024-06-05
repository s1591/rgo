package stations

func init() {
	Krelez.Register()
}

var Krelez = CreateMainStation("Krelez",
	krelezRandomChipTuneRadio, krelezVapor)

var krelezRandomChipTuneRadio = RadioStation{
	Url:     "http://79.120.11.40:8000/chiptune.ogg.m3u",
	Website: "http://79.120.11.40:8000/",
	Name:    "Random Chiptune Radio by Krelez",
}

var krelezVapor = RadioStation{
	Url:     "http://79.120.11.40:8000/vapor.ogg.m3u",
	Name:    "Vaporfunk Radio by Krelez",
	Website: "http://79.120.11.40:8000/",
}
