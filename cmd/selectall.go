package cmd

import (
	"github.com/Kawain/GormCobraTest/db"
	"github.com/spf13/cobra"
)

var selectallCmd = &cobra.Command{
	Use:   "selectall",
	Short: "Userモデルの全セレクト Short",
	Long:  `Userモデルの全セレクト Long`,
	Run: func(cmd *cobra.Command, args []string) {
		db.SelectAll()
	},
}

func init() {
	rootCmd.AddCommand(selectallCmd)
}
