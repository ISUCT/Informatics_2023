package lab_4

// Kozlov Egor 1/278: Variant 11

import "fmt"

type document struct {
	documentPath string
	documentType string
	documentSize int64
}

func NewDocument(path, documentType string, documentSize int64) (*document, error) {
	document := document{
		documentPath: path,
		documentType: documentType,
	}

	error := document.SetDocumentSize(documentSize)

	if error != nil {
		return nil, error
	}

	return &document, error
}

func (d *document) DocumentPath() string {
	return d.documentPath
}

func (d *document) DocumentType() string {
	return d.documentType
}

func (d *document) DocumentSize() int64 {
	return d.documentSize
}

func (d *document) SetDocumentSize(size int64) error {
	if size < 0 {
		return fmt.Errorf("Error: size must be >= 0. Current value: %v", size)
	}

	d.documentSize = size
	return nil
}
