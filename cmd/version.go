package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version number of Cabty",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println("cabty v0.0.1")
		},
	})
}
