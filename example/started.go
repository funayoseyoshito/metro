package main

import (
	"fmt"
	"os"

	"github.com/funayoseyoshito/metro"
)

type user struct {
	name string
}

type Human interface {
	hello()
}

func (h *user) hello() {
	fmt.Println("hello " + h.name)
}

var tmp []Human

func main() {
	//i, _ := tmp.([]user)
	//fmt.Println(i)
	//i, _ := tmp.(*user)
	//fmt.Printf("%T \n", i)
	//fmt.Printf("%T \n", tmp)
	//tmp = i
	//fmt.Printf("%T \n", tmp)
	//panic("================>ok")

	metro := metro.NewMetro(os.Getenv("metoro_token"))
	metro.GetODPTTrainInformation().Execute()
}