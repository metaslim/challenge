package presenter

import (
	"os"

	"github.com/kataras/tablewriter"
	"github.com/landoop/tableprinter"
	"github.com/metaslim/challenge/lib/model"
)

var _ Output = (*Json)(nil)

//Table is a struct that will allow searh result to be outputted as Table
type Table struct{}

//Flush will output the data
func (output Table) Flush(data model.SearchResult) {
	printer := tableprinter.New(os.Stdout)
	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	printer.HeaderBgColor = tablewriter.BgBlackColor
	printer.HeaderFgColor = tablewriter.FgGreenColor
	printer.AutoWrapText = true
	printer.HeaderAlignment = tableprinter.AlignCenter
	printer.DefaultAlignment = tableprinter.AlignCenter
	printer.NumbersAlignment = tableprinter.AlignCenter
	printer.RowCharLimit = 30

	switch data.(type) {
	case model.OrganizationSearchResult:
		printer.Print(data.(model.OrganizationSearchResult).Items)
	case model.UserSearchResult:
		printer.Print(data.(model.UserSearchResult).Items)
	case model.TicketSearchResult:
		printer.Print(data.(model.TicketSearchResult).Items)
	}
}
