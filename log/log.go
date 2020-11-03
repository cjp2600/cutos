package log

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

func Fatal(err error) {
	fmt.Printf("[%s] %s \n", promptui.Styler(promptui.FGRed)("ERROR"), err.Error())
	os.Exit(1)
}
