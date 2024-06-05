package stations

func init() {
	Tjsla.Register()
}

var Tjsla = CreateMainStation("Tjsla", tjsla1)

var tjsla1 = RadioStation{
	Name:    "tjsla1",
	Url:     "http://cast5.citrus3.com:2199/tunein/-stream/tjsradio.pls",
	Website: "http://www.tjsla.com/",
}
