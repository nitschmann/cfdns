package cmdhelper

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// TableWithHeader create an instance of a table with a header which could be printed to stdout
func TableWithHeader(header []string) *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)

	return table
}
