package main

import (
	"fmt"

	"tencent.com/widgets/widget"
	"tencent.com/widgets/window"
)

func main() {
	w := &window.Window{RootWidget: &widget.StatefulWidget{}}
	fmt.Printf("w's addr: %+v\n", w)

	w2 := &window.Window{RootWidget: &widget.StatelessWidget{}}
	fmt.Printf("w2's addr: %+v\n", w2)

	fmt.Println("Hello World!")
}
