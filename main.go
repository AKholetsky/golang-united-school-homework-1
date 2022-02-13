package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func hello() string {
	mes := emoji.Sprint("Hello :world_map:!")
	return mes
}

func main() {
	fmt.Println(hello())
}
