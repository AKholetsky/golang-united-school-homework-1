package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	mes := emoji.Sprint("Hello :world_map:!")
	return mes
}
