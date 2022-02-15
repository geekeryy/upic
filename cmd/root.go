package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "upic",
	Short: "upic 图床工具",
	Long:  "upic 超级便捷的图床工具",
}

func init() {

	rootCmd.AddCommand(uploadCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
