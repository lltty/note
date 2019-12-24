package main

import (
	"log"

	"github.com/zieckey/etcdsync"
)

func main() {
	m, err := etcdsync.New("/lock", 10, []string{"http://127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	err = m.Lock()
	if err != nil {
		log.Println("get lock failed:%v", err)
		return
	}

	log.Println("get lock suc")
	err = m.Unlock()
	if err != nil {
		log.Println("unlock lock fail")
	} else {
		log.Printf("unlock succ")
	}
}
