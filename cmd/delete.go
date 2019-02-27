package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/Kawain/GormCobraTest/db"
	"github.com/spf13/cobra"
)

//コマンド例
//GormCobraTest delete -0 55

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Short delete",
	Long:  `Long delete`,
	RunE: func(cmd *cobra.Command, args []string) error {
		user := db.User{}
		var err error
		user.ID, err = cmd.Flags().GetUint("id")
		if err != nil {
			return err
		}

		err = db.Delete(&user)
		if err != nil {
			return err
		}
		m := db.Result{Message: "ok"}
		b, _ := json.Marshal(m)
		fmt.Println(string(b))

		return nil

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().UintP("id", "0", 0, "ID integer option")
}
