package main

import (
	"flag"
	"os"
)

var (
	random         *bool
	dumpStations   *bool
	printStreamUrl *bool
)

func init() {
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
