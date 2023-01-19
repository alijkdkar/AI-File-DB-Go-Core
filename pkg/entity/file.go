package entity

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/alijkdkar/AI-File-DB-Go-Core/pkg/utility"
)

type File struct {
	FileName           string
	FileContaint       []byte
	FileStringContaint string
	Extention          string
	HashValue          string
	CreatedTime        time.Time
}

func (f *File) GetHash() (string, *File, error) {

	totlaHash := ""
	for _, v := range strings.Split(f.FileStringContaint, "\n") {
		if v != "" {
			fmt.Println("Line before hash:", v)
			// hash, err := utility.EncryptAES(utility.GetKey(""), string(v))
			hash, err := utility.Encrypt(utility.GetKey(""), string(v))
			fmt.Println("hash Line:", hash)
			totlaHash += (hash + "@@@\n")
			if err != nil {
				return "", f, errors.New("getFileHash Error => " + err.Error())
			}
		}
	}

	f.HashValue = totlaHash
	return totlaHash, f, nil

}

func (f *File) GetTextPlain() (string, *File, error) {

	totoalUnHashed := ""
	for _, v := range strings.Split(f.FileStringContaint, "@@@\n") {
		if strings.TrimSpace(v) != "" && strings.TrimSpace(v) != "\n" {
			// unhashedLine, err := utility.DecryptAES(utility.GetKey(""), v)
			fmt.Println("unhashed line1:", v)
			unhashedLine, err := utility.Decrypt(utility.GetKey(""), string(v))

			fmt.Println("unhashed line", unhashedLine)

			if err != nil {
				return "", f, errors.New("GetPlainText" + err.Error())
			}
			totoalUnHashed += (unhashedLine + "\n")
		}
	}

	f.FileContaint = []byte(totoalUnHashed)
	return string(totoalUnHashed), f, nil
}

func (f *File) Load(filePath string) *File {

	filecontaint, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer filecontaint.Close()
	// if err != nil {
	// 	panic(err.Error())
	// 	//return f, err
	// }
	scanner := bufio.NewScanner(filecontaint)
	for scanner.Scan() {
		fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>")
		f.FileStringContaint += (scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}

	fmt.Println(f.FileStringContaint)

	return f
}

func (f *File) Save(filePath string) error {
	err := ioutil.WriteFile(filePath, []byte(f.FileContaint), 0644) //0660
	if err != nil {
		return errors.New("Error to write file" + err.Error())
	}
	return nil
}

func (f *File) SaveHash(filePath string) error {
	if f.HashValue == "" {
		f.HashValue, _, _ = f.GetHash()
	}

	err := ioutil.WriteFile(filePath, []byte(f.HashValue), 0644) //0660
	if err != nil {
		return errors.New("Error to write file" + err.Error())
	}
	return nil
}
