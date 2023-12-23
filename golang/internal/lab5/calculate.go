package lab5

import "fmt"

type document struct {
	document_type    string
	size             int
	year_of_creation int
}

// size - размер документа в килобайтах
func NewDocument(document_type_var string, size_var, year_of_creation_var int) (document, error) {
	var d document = document{
		document_type:    document_type_var,
		year_of_creation: year_of_creation_var,
	}
	err := d.SetSize(size_var)
	return d, err
}

func (d *document) SetSize(size int) error {
	if size > 0 {
		d.size = size
		return nil
	}
	return fmt.Errorf("Document of type %s has invalid size", d.GetDocumentType())
}

func (d *document) GetDocumentType() string {
	return d.document_type
}

func (d *document) GetDocumentSize() int {
	return d.size
}

func (d *document) GetDocumentDate() int {
	return d.year_of_creation
}

func (d *document) DeleteDocument() {
	fmt.Printf("Document of type %s created %d was sent to the recycle bin", d.GetDocumentType(), d.GetDocumentDate())
}
