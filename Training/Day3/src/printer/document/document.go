package document

import "fmt"

type Document struct {
	DocId   int
	Content string
}

func (d Document) PageSetup() {

	fmt.Println("Document::PageSetup", d.DocId, d.Content)

}
func (d Document) Preview() {
	fmt.Println("Document::Preview", d.DocId, d.Content)
}
func (d Document) Print() {
	fmt.Println("Document::Print", d.DocId, d.Content)
}
