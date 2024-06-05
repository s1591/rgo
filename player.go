package main

import (
	"github.com/s1591/rgo/stations"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/gen2brain/go-mpv"
)

var (
	player      *mpv.Mpv
	playerStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Align(lipgloss.Center).Padding(1, 1)
	titleStyle = lipgloss.NewStyle().
			Bold(true)
)

type model struct {
	title          string
	spinner        spinner.Model
	paused         bool
	keys           keyMap
	help           help.Model
	displayWebsite bool
	station        stations.RadioStation
}

type keyMap struct {
	quit          key.Binding
	changeStation key.Binding
	restartStream key.Binding
	pause         key.Binding
	dispWebsite   key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.quit, k.changeStation, k.restartStream, k.pause, k.dispWebsite}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{k.ShortHelp()}
}

func Play(station stations.RadioStation) {
	p := tea.NewProgram(model{
		spinner: spinner.Model{Spinner: speaker()},
		station: station,
		keys: keyMap{
			quit:          key.NewBinding(key.WithKeys("q"), key.WithHelp("q/ctrl-c", "quit")),
			changeStation: key.NewBinding(key.WithKeys("c"), key.WithHelp("c", "change station")),
			restartStream: key.NewBinding(key.WithKeys("r"), key.WithHelp("r", "restart stream")),
			pause:         key.NewBinding(key.WithKeys("p"), key.WithHelp("p/space", "pause")),
			dispWebsite:   key.NewBinding(key.WithKeys("w"), key.WithHelp("w", "website")),
		},
		help: help.New(),
	})
	_, err := p.Run()
	defer player.TerminateDestroy()
	if err != nil {
		panic(err)
	}
}

func (m model) Init() tea.Cmd {
	player = mpv.New()
	err := player.Initialize()
	if err != nil {
		panic(err)
	}
	player.SetProperty("video", mpv.FormatFlag, false)
	err = player.Command([]string{"loadfile", m.station.Url})
	if err != nil {
		panic(err)
	}
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "c":
			newStation, err := RunSelect()
			if err != nil {
				if err.Error() == "user aborted" {
					return m, tea.HideCursor
				}
			}
			m.station = newStation
			player.Command([]string{"loadfile", m.station.Url})
			return m, tea.HideCursor
		case "r":
			player.Command([]string{"loadfile", m.station.Url})
		case "w":
			m.displayWebsite = !m.displayWebsite
		case tea.KeySpace.String(), "p":
			m.togglePause()
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	m.title = player.GetPropertyOsdString("media-title")
	if m.title == player.GetPropertyOsdString("filename") {
		m.title = "media info not available"
	}

	if m.paused == false {
		m.spinner, cmd = m.spinner.Update(msg)
	} else {
		cmd = m.spinner.Tick
	}

	return m, cmd
}

func (m model) View() string {
	var ui string

	ui += titleStyle.Render(m.station.Name)
	if m.displayWebsite {
		ui += titleStyle.Render(" • " + m.station.Website)
	}
	ui += "\n"

	if m.paused {
		ui += "  "
	} else {
		ui += m.spinner.View() + "  "
	}
	ui += m.title + "\n"
	ui += m.help.View(m.keys)

	return playerStyle.Render(ui) + "\n"
}

func (m *model) togglePause() {
	if player != nil {
		m.paused = !m.paused
		player.SetProperty("pause", mpv.FormatFlag, m.paused)
	}
}

func speaker() spinner.Spinner {
	return spinner.Spinner{
		Frames: []string{"", "", ""},
		FPS:    time.Second / 6,
	}
}
