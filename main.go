package main

import (
	"flag"
	"github.com/s1591/rgo/stations"
	"os"
)

var (
	random         *bool
	dumpStations   *bool
	printStreamUrl *bool
)

func init() {
	random = flag.Bool("r", false, "play a random station.\n")
	dumpStations = flag.Bool("d", false, "dumps stations to a json(in the current directory).\n")
	printStreamUrl = flag.Bool("u", false, "print the selected station stream url and exit.\n")

	flag.Parse()
}

func main() {

	if *dumpStations {
		DumpStations()
	} else if *printStreamUrl {
		station, err := RunSelect()
		if err != nil {
			if err.Error() == "user aborted" {
				os.Exit(0)
			} else {
				panic(err)
			}
		}
		println(station.Url)
	} else if *random {
		Play(stations.RandomStation())
	} else {
		station, err := RunSelect()
		if err != nil {
			if err.Error() == "user aborted" {
				os.Exit(0)
			} else {
				panic(err)
			}
		}
		Play(station)
	}

}
