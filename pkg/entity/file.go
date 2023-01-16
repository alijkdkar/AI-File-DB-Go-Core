package entity

import (
	"time"

	"github.com/alijkdkar/AI-File-DB-Go-Core/pkg/utility"
)

type file struct {
	fileName     string
	fileContaint string
	extention    string
	hashValue    string
	CreatedTime  time.Time
}

func (f *file) getFileHash() (string, error) {
	hash, err := utility.EncryptAES(utility.GetKey(""), f.fileContaint)

	if err != nil {
		return "", err
	}
	return hash, nil

}
