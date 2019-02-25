package cmd

import (
	"github.com/Kawain/GormCobraTest/db"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "テーブルがあれば削除して初期化する。",
	Long:  `テーブルがあれば削除して初期化する。`,
	Run: func(cmd *cobra.Command, args []string) {
		db.InitializeTable()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
