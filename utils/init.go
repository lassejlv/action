package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var example string = "https://gist.githubusercontent.com/lassejlv/f6ebe343e4fed8aa2ab2b97b1b95a6d5/raw/903b258b2b576d784d984e0607b38b61fb23eed9/.actions"

func Init() {
	if _, err := os.Stat(".actions"); err == nil {
		fmt.Println("File already exists. Delete it and run 'action --init' again")
		return
	}

	resp, err := http.Get(example)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	readableContent, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(".actions", readableContent, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully created .actions file")
}
