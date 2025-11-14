package main

import "github.com/charmbracelet/lipgloss"

const ( // Catppuccin Mocha
	rosewater = lipgloss.Color("#f5e0dc")
	flamingo  = lipgloss.Color("#f2cdcd")
	pink      = lipgloss.Color("#f5c2e7")
	mauve     = lipgloss.Color("#cba6f7")
	red       = lipgloss.Color("#f38ba8")
	maroon    = lipgloss.Color("#eba0ac")
	peach     = lipgloss.Color("#fab387")
	yellow    = lipgloss.Color("#f9e2af")
	green     = lipgloss.Color("#a6e3a1")
	teal      = lipgloss.Color("#94e2d5")
	sky       = lipgloss.Color("#89dceb")
	sapphire  = lipgloss.Color("#74c7ec")
	blue      = lipgloss.Color("#89b4fa")
	lavander  = lipgloss.Color("#b4befe")

	base  = lipgloss.Color("#1e1e2e")
	text  = lipgloss.Color("#cdd6f4")
	crust = lipgloss.Color("#11111b")
)

var RosewaterStyle = lipgloss.NewStyle().
	Foreground(rosewater)

var FlamingoStyle = lipgloss.NewStyle().
	Foreground(flamingo)

var PinkStyle = lipgloss.NewStyle().
	Foreground(pink)

var MauveStyle = lipgloss.NewStyle().
	Foreground(mauve)

var RedStyle = lipgloss.NewStyle().
	Foreground(red)

var MaroonStyle = lipgloss.NewStyle().
	Foreground(maroon)

var PeachStyle = lipgloss.NewStyle().
	Foreground(peach)

var YellowStyle = lipgloss.NewStyle().
	Foreground(yellow)

var GreenStyle = lipgloss.NewStyle().
	Foreground(green)

var TealStyle = lipgloss.NewStyle().
	Foreground(teal)

var SkyStyle = lipgloss.NewStyle().
	Foreground(sky)

var SapphireStyle = lipgloss.NewStyle().
	Foreground(sapphire)

var BlueStyle = lipgloss.NewStyle().
	Foreground(blue)

var LavanderStyle = lipgloss.NewStyle().
	Foreground(lavander)
