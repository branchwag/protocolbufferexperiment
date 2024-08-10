package main

import (
	"fmt"
	"log"

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
	
	data, err := proto.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}