package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/juju/ansiterm"
)

func main() {
	writer := ansiterm.NewWriter(os.Stdout)

	colors := []ansiterm.Color{
		ansiterm.Black, ansiterm.Red, ansiterm.Green, ansiterm.Yellow,
		ansiterm.Blue, ansiterm.Magenta, ansiterm.Cyan, ansiterm.Gray,
		ansiterm.DarkGray, ansiterm.BrightRed, ansiterm.BrightGreen, ansiterm.BrightYellow,
		ansiterm.BrightBlue, ansiterm.BrightMagenta, ansiterm.BrightCyan, ansiterm.White,
	}

	writer.SetAttribute(ansiterm.Bold)
	fmt.Fprintf(writer, "\nForegrounds\n")
	writer.UnsetAttribute(ansiterm.Bold)
	for i, c := range colors {
		writer.SetForeground(c)
		fmt.Fprintf(writer, "%-15s", c)
		if (i+1)%4 == 0 {
			fmt.Fprintln(writer)
		}
	}
	writer.Reset()

	writer.SetAttribute(ansiterm.Bold)
	fmt.Fprintf(writer, "\nBackgrounds\n")
	writer.UnsetAttribute(ansiterm.Bold)
	for i, c := range colors {
		writer.SetBackground(c)
		fmt.Fprintf(writer, "%-15s", c)
		if (i+1)%4 == 0 {
			writer.SetBackground(ansiterm.Default)
			fmt.Fprintln(writer)
		}
	}

	writer.Reset()
	fmt.Fprintln(writer)

	const (
		// To format things into columns.
		minwidth = 0
		tabwidth = 1
		padding  = 2
		padchar  = ' '
		flags    = 0
	)
	tw := ansiterm.NewColorTabWriter(os.Stdout, minwidth, tabwidth, padding, padchar, flags)
	p := func(values ...string) {
		text := strings.Join(values, "\t")
		fmt.Fprintln(tw, text)
	}
	tw.SetAttribute(ansiterm.Bold)
	p("CLOUD", "TYPE", "REGIONS")
	tw.Reset()

	for i, row := range [][]string{
		{"amazon", "ec2", "lots"},
		{"google", "gce", "less"},
	} {
		tw.SetForeground(colors[i+1])
		fmt.Fprintf(tw, row[0])
		tw.SetForeground(ansiterm.Default)
		fmt.Fprintf(tw, "\t%s\t%s\n", row[1], row[2])
	}
	fmt.Fprintln(tw)
}
