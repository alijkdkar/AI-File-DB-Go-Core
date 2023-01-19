package main

import (
	"fmt"
	"os"

	"github.com/alijkdkar/AI-File-DB-Go-Core/pkg/entity"
)

func main() {

	var b []byte = make([]byte, 1)
	for {

		os.Stdin.Read(b)
		if string(b) == "a" {

			go func() {
				file := entity.File{}
				if hashValue, file, err := file.Load("A.jpg").GetHash(); err == nil {
					fmt.Println("lent :", len(hashValue))
					file.SaveHash("Hashed.txt")
				}
			}()

		} else if string(b) == "b" {

			go func() {
				file1 := entity.File{}
				_, f, er := file1.Load("Hashed.txt").GetTextPlain()
				if er == nil {
					//fmt.Println(txtplan)
					f.Save("b.jpg")
				} else {
					fmt.Println(er)
				}
			}()

		} else if string(b) == "q" {
			panic("exit")
		} else {
			fmt.Print(">>>")
		}
	}
}
