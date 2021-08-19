package tui

import (
	"github.com/charmbracelet/lipgloss"
)

var activeBorderColor = lipgloss.Color("243")
var inactiveBorderColor = lipgloss.Color("236")

var hiddenBorder = lipgloss.Border{
	TopLeft:     " ",
	Top:         " ",
	TopRight:    " ",
	BottomLeft:  " ",
	Bottom:      " ",
	BottomRight: " ",
}

var appBoxStyle = lipgloss.NewStyle()

var menuStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(inactiveBorderColor).
	Padding(1, 2).
	MarginRight(1).
	Width(24)

var menuActiveStyle = menuStyle.Copy().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(activeBorderColor)

var contentBoxStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(inactiveBorderColor).
	PaddingRight(1).
	MarginBottom(1)

var contentBoxActiveStyle = contentBoxStyle.Copy().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(activeBorderColor)

var headerStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("61")).
	Align(lipgloss.Right).
	Bold(true)

var footerStyle = lipgloss.NewStyle().
	MarginTop(1)

var helpKeyStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("241"))

var helpValueStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("239"))

var menuItemStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("252")).
	PaddingLeft(2)

var selectedMenuItemStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("207")).
	PaddingLeft(1)

var menuCursor = lipgloss.NewStyle().
	Foreground(lipgloss.Color("213")).
	SetString(">")

var errorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF00000"))

var helpDivider = lipgloss.NewStyle().
	Foreground(lipgloss.Color("237")).
	SetString(" • ")