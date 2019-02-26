package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "GormCobraTest",
	Short: "SQLite DB 操作のCLI(Command Line Interface) Short",
	Long:  `SQLite DB 操作で Cobra と Gorm を使ったCLI(Command Line Interface) Long`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
