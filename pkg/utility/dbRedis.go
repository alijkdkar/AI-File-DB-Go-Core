package utility

import (
	"github.com/mediocregopher/radix/v3"
)

type Reddis struct {
	Clint radix.Pool
}

func GetInstance() (*Reddis, error) {

	pool, err := radix.NewPool("tcp", "127.0.0.1:6379", 10)
	if err != nil {
		return nil, err
	}

	return &Reddis{Clint: *pool}, nil
}

func (r Reddis) Get(key string) (string, error) {
	var fileName string
	err := r.Clint.Do(radix.Cmd(&fileName, "GET", key))
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func (r Reddis) Set(key, val string) error {

	return r.Clint.Do(radix.Cmd(nil, "SET", key, val))

}
