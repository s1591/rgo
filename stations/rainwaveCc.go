package stations

func init() {
	RainWaveCC.Register()
}

var RainWaveCC = CreateMainStation("RainWaveCc",
	rainWaveCcAll, rainWaveCcGame,
	rainWaveCcChiptune, rainWaveCcRemix,
	rainWaveCcCovers)

var rainWaveCcAll = RadioStation{
	Url:     "https://rainwave.cc/tune_in/5.ogg.m3u",
	Name:    "RainWaveCc (All)",
	Website: "https://rainwave.cc/",
}

var rainWaveCcGame = RadioStation{
	Name:    "RainWaveCc (Game)",
	Url:     "https://rainwave.cc/tune_in/1.ogg.m3u",
	Website: "https://rainwave.cc/",
}

var rainWaveCcChiptune = RadioStation{
	Name:    "RainWaveCc (Chiptnue)",
	Url:     "https://rainwave.cc/tune_in/4.ogg.m3u",
	Website: "https://rainwave.cc/",
}

var rainWaveCcRemix = RadioStation{
	Name:    "RainWaveCc (Remix)",
	Url:     "https://rainwave.cc/tune_in/2.ogg.m3u",
	Website: "https://rainwave.cc/",
}

var rainWaveCcCovers = RadioStation{
	Name:    "RainWaveCc (Covers)",
	Url:     "https://rainwave.cc/tune_in/3.ogg.m3u",
	Website: "https://rainwave.cc/",
}
