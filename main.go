package main

import (
	"fmt"
	"simple-cache-proxy/cmd"
)

func main() {
	rootCmd := cmd.NewRootCMD()

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
