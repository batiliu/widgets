package main

import (
	"fmt"

	"tencent.com/widgets/framework"
	"tencent.com/widgets/window"
)

func main() {
	w := &window.Window{}
	w.AttachRootWidget(&framework.StatefulWidget{})
	fmt.Printf("w's addr: %+v\n", w)

	w2 := &window.Window{}
	w2.AttachRootWidget(&framework.StatelessWidget{})
	fmt.Printf("w2's addr: %+v\n", w2)

	fmt.Println("Hello World!")
}
