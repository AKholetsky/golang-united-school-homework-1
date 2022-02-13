package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func hello() {
	mes := emoji.Sprint("Hello :world_map:!")
	fmt.Println(mes)
}
func main() {
	hello()
}
