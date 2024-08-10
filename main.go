package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"main/model"

	"github.com/golang/protobuf/proto"
)

func main () {
	book := &model.Book{
		Id: 1,
		Title: "The Road to Wigan Pier",
		Authors: []*model.Author{
			{Id: 1, Name: "George Orwell"},
		},
		Category: model.Category_Novel,
	}
	
	//proto
	data, err := proto.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(data)
	os.WriteFile("book.protobuf", data, 0600)
	
	//JSON for performance comparison
	data, err = json.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(data)
	os.WriteFile("book.json", data, 0600)

	
	//decode data from protobuf bytes
	data, err = os.ReadFile("book.protobuf")
	if err != nil {
		log.Fatal(err)
	}
	bookFromFile := model.Book{}
	if err := proto.Unmarshal(data, &bookFromFile); err != nil {
		log.Fatal(err)
	}

	// Safely print the struct without copying the lock
	fmt.Printf("book from protobuf file - ID: %d, Title: %s, Authors: %v, Category: %v\n",
	bookFromFile.Id, bookFromFile.Title, bookFromFile.Authors, bookFromFile.Category)
}