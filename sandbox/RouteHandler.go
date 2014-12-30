package main

/*
 wrote this to try out different structs in an
 array but working with them via the same interface
*/
import (
	"fmt"
)

type Handler interface {
	Msg() string
}

type Handler1 struct{}

func (h Handler1) Msg() string { return "Handler1" }

type Handler2 struct{}

func (h Handler2) Msg() string { return "Handler2" }

func main() {
	var routeList [2]Handler

	routeList[0] = Handler1{}
	routeList[1] = Handler2{}

	for _, handler := range routeList {
		fmt.Println(handler.Msg())
	}

}
