package main

import (
	"log"
	"syscall"

	"github.com/ark-go/fibergo/internal/serverf"
)

func init() {
}

func main() {
	sf := &serverf.ServerFiber{}
	log.Println("Main Вход --------------->", syscall.Getpid())
	sf.Start()

	log.Println("Main Выход --------------->", syscall.Getpid())

}
