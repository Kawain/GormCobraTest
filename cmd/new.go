package cmd

import (
	"github.com/Kawain/GormCobraTest/db"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "テーブル初期 Short",
	Long:  `テーブルがあれば削除、なければ新規作成して初期化する。Long`,
	Run: func(cmd *cobra.Command, args []string) {
		db.InitializeTable()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
