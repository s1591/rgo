package stations

import "math/rand"

// struct to hold some info about station
type RadioStation struct {
	Name    string
	Url     string
	Website string
}

// holds the all the stations of the corresponding name
type RadioStations struct {
	Name     string
	Stations *[]RadioStation // needs to be comparable for huh
}

type StationMap map[string]RadioStations

var (
	// maps station name to the corresponding station
	SMap StationMap = make(StationMap)
	// holds all the available stations
	Stations []RadioStations = make([]RadioStations, 0) // space not allocated as some init() may still need to run
)

// register registers the main station
//
// it appends the Main Station into Stations and fills the SMap
func (rs *RadioStations) Register() {
	Stations = append(Stations, *rs)
	SMap[rs.Name] = *rs
}

func CreateMainStation(name string, stations ...RadioStation) RadioStations {
	return RadioStations{
		Name:     name,
		Stations: &stations,
	}
}

func RandomStation() RadioStation {

	stations := *Stations[rand.Intn(len(Stations))].Stations
	return stations[rand.Intn(len(stations))]

}
