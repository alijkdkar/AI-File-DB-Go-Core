package utility

import (
	"log"

	"github.com/mediocregopher/radix/v3"
)

type Reddis struct {
	// Clint radix.Pool
	Clint radix.Conn
}

func GetInstance() (*Reddis, error) {

	// pool, err := radix.NewPool("tcp", "127.0.0.1:6379", 10)
	pool, err := radix.DefaultConnFunc("tcp", "127.0.0.1:6379")

	if err != nil {
		return nil, err
	}

	return &Reddis{Clint: pool}, nil
}

func (r *Reddis) Get(key string) (string, error) {
	var fileName string
	err := r.Clint.Do(radix.Cmd(&fileName, "GET", key))
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func (r *Reddis) Set(key, val string) error {

	return r.Clint.Do(radix.Cmd(nil, "SET", key, val))

}
func (r *Reddis) Subcribe(chanName string) chan radix.PubSubMessage {

	pstub := radix.PubSub(r.Clint)

	msgCh := make(chan radix.PubSubMessage)
	if err := pstub.Subscribe(msgCh, chanName); err != nil {
		log.Fatal(err)
	}
	return msgCh
}
