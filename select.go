package main

import (
	"github.com/s1591/rgo/stations"

	"github.com/charmbracelet/huh"
)

func RunSelect() (stations.RadioStation, error) {

	var (
		chosenMainStation stations.RadioStations
		msOptions         []huh.Option[stations.RadioStations]
		chosenStation     stations.RadioStation
		sOptions          []huh.Option[stations.RadioStation]
	)

	for _, s := range stations.SMap {
		msOptions = append(msOptions, huh.NewOption(s.Name, s))
	}

	ms := huh.NewSelect[stations.RadioStations]().
		Title("Pick a Station").Options(msOptions...).
		Value(&chosenMainStation).WithTheme(huh.ThemeCatppuccin()).
		WithHeight(10)

    err := huh.NewForm(huh.NewGroup(ms)).Run()
    if err != nil {
        return chosenStation, err
    }

	for _, s := range *chosenMainStation.Stations {
		sOptions = append(sOptions, huh.NewOption(s.Name, s))
	}

	s := huh.NewSelect[stations.RadioStation]().
		Title(chosenMainStation.Name).Options(sOptions...).
		Value(&chosenStation).WithTheme(huh.ThemeCatppuccin()).
		WithHeight(10)

	err = huh.NewForm(huh.NewGroup(s)).Run()

	return chosenStation, err

}
