package spreadsheet

import "fmt"

type SpreadSheet struct {
	SheetID      int
	Rows, Column int
}

func (d SpreadSheet) PageSetup() {

	fmt.Println("SpreadSheet::PageSetup", d.SheetID, d.Rows, d.Column)

}
func (d SpreadSheet) Preview() {
	fmt.Println("SpreadSheet::Preview", d.SheetID, d.Rows, d.Column)
}
func (d SpreadSheet) Print() {
	fmt.Println("SpreadSheet::Print", d.SheetID, d.Rows, d.Column)
}
