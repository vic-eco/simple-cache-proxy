package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func NewRootCMD() *cobra.Command {
	var port int
	var origin string

	rootCMD := &cobra.Command{
		Use:   "caching-proxy",
		Short: "A simple caching proxy",
		Long:  "A simple in-memory caching proxy for web requests",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Works!")
			return nil
		},
	}

	rootCMD.Flags().IntVarP(&port, "port", "p", 3000, "port that the proxy server will run on")
	rootCMD.Flags().StringVarP(&origin, "origin", "o", "", "origin URL from which the server will cache requests")
	rootCMD.MarkFlagRequired("origin")

	return rootCMD
}
