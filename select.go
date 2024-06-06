package main

import (
	"github.com/s1591/rgo/stations"

	"github.com/charmbracelet/huh"
)

// run the huh select prompt
//
// the options are taken from station.SMap
func RunSelect() (chosenStation stations.RadioStation, err error) { // integrate into bubble tea

	var (
		chosenMainStation stations.RadioStations
		msOptions         []huh.Option[stations.RadioStations]
		sOptions          []huh.Option[stations.RadioStation]
	)

	for _, s := range stations.SMap {
		msOptions = append(msOptions, huh.NewOption(s.Name, s))
	}

	ms := huh.NewSelect[stations.RadioStations]().
		Title("Pick a Station").Options(msOptions...).
		Value(&chosenMainStation).WithTheme(huh.ThemeCatppuccin()).
		WithHeight(10)

	err = huh.NewForm(huh.NewGroup(ms)).Run()
	if err != nil {
		return
	}

	for _, s := range *chosenMainStation.Stations {
		sOptions = append(sOptions, huh.NewOption(s.Name, s))
	}

	s := huh.NewSelect[stations.RadioStation]().
		Title(chosenMainStation.Name).Options(sOptions...).
		Value(&chosenStation).WithTheme(huh.ThemeCatppuccin()).
		WithHeight(10)

	err = huh.NewForm(huh.NewGroup(s)).Run()

	return

}
