package lab5

import (
	"fmt"
	"strings"
)

type document struct {
	type_of_document  string
	document_size     int32
	title_of_document string
}

func (d document) GetType() string {
	return strings.ToUpper(d.type_of_document)
}

func (d document) GetSize() int32 {
	return d.document_size
}

func (d document) GetTitle() string {
	return d.title_of_document
}

func NewDocument(checkType string, checkSize int32, checkTitle string) (document, error) {
	var d document
	err := d.SetType(checkType)
	if err != nil {
		return d, err
	}
	err = d.SetSize(checkSize)
	if err != nil {
		return d, err
	}
	err = d.SetTitle(checkTitle)
	if err != nil {
		return d, err
	}
	return d, nil
}

func (d *document) SetType(type_of_document string) error {
	document_types := []string{"DOCX", "DOC", "DOTX", "DOT", "RTF", "TXT", "HTM", "PDF", "DOCM", "DOTM", "XML", "MHT", "DIC", "THMX"}
	if contains(document_types, strings.ToUpper(type_of_document)) {
		d.type_of_document = type_of_document
		return nil
	}
	return fmt.Errorf("Incorrect type of document")
}

func (d *document) SetSize(document_size int32) error {
	if document_size > 0 {
		d.document_size = document_size
		return nil
	}
	return fmt.Errorf("Wrong document size")
}

func (d *document) SetTitle(title_of_document string) error {
	if len(title_of_document) > 0 {
		d.title_of_document = title_of_document
		return nil
	}
	return fmt.Errorf("Wrong document title")
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
