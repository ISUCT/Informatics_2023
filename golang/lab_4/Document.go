package lab_4

// Kozlov Egor 1/278: Variant 11

import "fmt"

type Document struct {
	documentPath string
	documentType string
	documentSize int64
}

func NewDocument(path, documentType string, documentSize int64) (*Document, error) {
	if documentSize < 0 {
		return nil, fmt.Errorf("Size is less than zero.")
	}

	document := Document{
		documentPath: path,
		documentType: documentType,
	}

	error := document.SetDocumentSize(documentSize)

	if error != nil {
		return nil, error
	}

	return &document, error
}

func (d *Document) DocumentPath() string {
	return d.documentPath
}

func (d *Document) DocumentType() string {
	return d.documentType
}

func (d *Document) DocumentSize() int64 {
	return d.documentSize
}

func (d *Document) SetDocumentSize(size int64) error {
	if size < 0 {
		return fmt.Errorf("Error: size must be >= 0. Current value: %v", size)
	}

	d.documentSize = size
	return nil
}
