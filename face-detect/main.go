package main

import (
	"fmt"
	"os"

	. "github.com/kkdai/oxford-face"
)

func main() {
	key := os.Getenv("MSFT_KEY")
	if key == "" {
		fmt.Println("Please export your key to environment first, `export MSFT_KEY=12234`")
		return
	}
	f := NewFace(key)

	//Detect
	ret, err := f.DetectFile(nil, "./1.jpg")
	fmt.Println("ret:", ret, " err=", err)
}
