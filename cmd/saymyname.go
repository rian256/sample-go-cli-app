package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sayMyNameCmd)
}

var sayMyNameCmd = &cobra.Command{
	Use:   "sebutkan-nama",
	Short: "Gunakan perintah sebutkan-nama untuk menyapa nama kita",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("HALOOOOOO", strings.Join(args, " "))
	},
}
