package main

import (
	"fmt"

	"github.com/alijkdkar/AI-File-DB-Go-Core/pkg/entity"
)

func main() {

	// file := entity.File{}
	// if hashValue, file, err := file.Load("A.jpg").GetHash(); err == nil {
	// 	fmt.Println("lent :", len(hashValue))
	// 	file.SaveHash("Hashed.txt")
	// }

	file1 := entity.File{}

	txtplan, f, er := file1.Load("Hashed.txt").GetTextPlain()
	if er == nil {
		fmt.Println(txtplan)
		f.Save("b.jpg")
	} else {
		fmt.Println(er)
	}

}
