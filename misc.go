package main

import (
	"encoding/json"
	"github.com/s1591/rgo/stations"
	"os"
)

func DumpStations() {

	d := []stations.RadioStations{}

	for _, s := range stations.SMap {
		d = append(d, s)
	}

	data, err := json.MarshalIndent(d, "", "\t")
	if err != nil {
		panic(err)
	}

	writeStationsToFile(data)

}

// helper func to write data to file
func writeStationsToFile(data []byte) {

	f, err := os.Create("stations.json")
	if err != nil {
		os.Remove("stations.json")
		_, err := f.Write(data)
		if err == nil {
			println("dump successful")
		} else {
			panic(err)
		}
		return
	}
	defer f.Close()

	f.Write(data)
	println("dump successful")

}
