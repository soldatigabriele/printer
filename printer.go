package printer

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/table"
)

// Blueprint is the main structure printed
type Blueprint struct {
	Colour int
	Title  string
	Data   map[string]interface{}
}

// Colour constants. Determine the Colour of printing
const (
	Blue = iota + 1
	Green
	Red
)

// Print takes a Blueprint struct and prints it in a nice format
func Print(e Blueprint) error {
	// Initialise the table
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// Set the print Colour based on the Colour
	switch e.Colour {
	case Blue:
		t.SetStyle(table.StyleColoredBlackOnBlueWhite)
	case Green:
		t.SetStyle(table.StyleColoredBlackOnGreenWhite)
	case Red:
		t.SetStyle(table.StyleColoredBlackOnRedWhite)
	default:
		t.SetStyle(table.StyleColoredBlackOnBlueWhite)
	}

	t.AppendHeader(table.Row{e.Title, ""})
	for k, v := range e.Data {
		t.AppendRows([]table.Row{
			{k, v},
		})
	}
	fmt.Println()
	t.Render()
	return nil
}
