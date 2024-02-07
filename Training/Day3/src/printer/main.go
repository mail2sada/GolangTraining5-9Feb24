package main

import (
	"fmt"
	"printer/document"
	"printer/print"
	"printer/spreadsheet"
)

func main() {
	fmt.Println("Demo: PRinter interface")

	d := document.Document{DocId: 1, Content: "This is test Content"}

	s := spreadsheet.SpreadSheet{SheetID: 1, Rows: 10, Column: 5}

	print.PageSetup(d)
	print.Preview(d)
	print.Print(d)

	print.PageSetup(s)
	print.Preview(s)
	print.Print(s)
}
