package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmdは、サブコマンドなしで呼び出されたときの基本コマンドを表します。
var rootCmd = &cobra.Command{
	Use:   "GormCobraTest",
	Short: "SQLite DB 操作のCLI(Command Line Interface)",
	Long:  `CobraとGormを使ったCLI(Command Line Interface)`,
	// 裸のアプリケーションならば、次の行のコメントを外してください。
	// それに関連付けられたアクションがあります。
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Executeはすべての子コマンドをrootコマンドに追加し、適切にフラグを設定します。
// これはmain.main（）によって呼び出されます。 rootCmdに対して一度だけ発生する必要があります。
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// ここでフラグと設定を定義します。
	// Cobraは永続的なフラグをサポートしています。ここで定義されていれば、
	// アプリケーションに対してグローバルになります。
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.GormCobraTest.yaml)")

	// Cobraはローカルフラグもサポートしています。
	// このアクションが直接呼び出されたとき
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
