package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal/lab5"

	"log"
)

func main() {
	fmt.Println("Документ 1")
	doc1, err := lab5.NewDocument("Договор", 1240, 2018)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc1.GetDocumentSize())
	doc1.SetSize(1024)
	fmt.Println(doc1.GetDocumentSize())

	fmt.Println("Документ 2")
	doc2, err := lab5.NewDocument("Нотариальное Согласие", 0, 2022)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc2.GetDocumentSize()) // ошибка DocumentSize = 0
	doc2.DeleteDocument()

}

// 11 вариант
