package main

import (
	"fmt"
	"log"

	"github.com/alijkdkar/AI-File-DB-Go-Core/pkg/entity"
	"github.com/alijkdkar/AI-File-DB-Go-Core/pkg/utility"
)

func main() {

	// 	if string(b) == "a" {

	// 		go func() {
	// 			file := entity.File{}
	// 			if hashValue, file, err := file.Load("A.jpg").GetHash(); err == nil {
	// 				fmt.Println("lent :", len(hashValue))
	// 				file.SaveHash("", "Hashed.txt")
	// 			}
	// 		}()

	// 	} else if string(b) == "b" {

	// 		go func() {
	// 			file1 := entity.File{}
	// 			_, f, er := file1.Load("Hashed.txt").GetTextPlain()
	// 			if er == nil {
	// 				//fmt.Println(txtplan)
	// 				f.Save("", "b.jpg")
	// 			} else {
	// 				fmt.Println(er)
	// 			}
	// 		}()

	fmt.Println("Core runnig...")
	fmt.Println("Core ready to do mission ..")
	fmt.Println("Stop Core CTRL+C ...")

	redDB, err := utility.GetInstance()
	if err != nil {
		fmt.Printf("Error Occerd %v", err.Error())
		panic("Error Occerd")
	}
	chann := redDB.Subcribe("HashChannel")

	for m := range chann {
		log.Printf("read m: %#v", string(m.Message))
		//TODO: you muust to hash file that ricive from redis
		go func() {
			if fileName, err := redDB.Get(string(m.Message)); err == nil {
				fmt.Println("File Id Not Founded")
				file := entity.File{}
				if hashValue, file, err := file.Load(fileName).GetHash(); err == nil {
					fmt.Println("lent :", len(hashValue))
					file.SaveHash("", "Hashed.txt")
				}
			}

		}()

	}
	fmt.Println("app run ended")

}
