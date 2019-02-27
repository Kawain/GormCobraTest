package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/Kawain/GormCobraTest/db"
	"github.com/spf13/cobra"
)

//コマンド例
//GormCobraTest insert -1 鈴木一郎 -2 スズキイチロウ -3 1 -4 0900909090 -5 mail@mail.com -6 1945/01/01 -7 80 -8 1 -9 1

var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Short insert",
	Long:  `Long insert`,
	RunE: func(cmd *cobra.Command, args []string) error {
		user := db.User{}
		var err error
		user.Name, err = cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		user.Katakana, err = cmd.Flags().GetString("kaka")
		if err != nil {
			return err
		}

		user.GenderID, err = cmd.Flags().GetUint("sex")
		if err != nil {
			return err
		}

		user.Tel, err = cmd.Flags().GetString("tel")
		if err != nil {
			return err
		}

		user.Mail, err = cmd.Flags().GetString("mail")
		if err != nil {
			return err
		}

		user.Birthday, err = cmd.Flags().GetString("birthday")
		if err != nil {
			return err
		}

		user.Age, err = cmd.Flags().GetUint("age")
		if err != nil {
			return err
		}

		user.HometownID, err = cmd.Flags().GetUint("hometown")
		if err != nil {
			return err
		}

		user.BloodTypeID, err = cmd.Flags().GetUint("bloodtype")
		if err != nil {
			return err
		}

		err = db.Insert(&user)
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
	rootCmd.AddCommand(insertCmd)

	insertCmd.Flags().StringP("name", "1", "", "Name string option")
	insertCmd.Flags().StringP("kaka", "2", "", "Katakana string option")
	insertCmd.Flags().UintP("sex", "3", 0, "GenderID integer option")
	insertCmd.Flags().StringP("tel", "4", "", "Tel string option")
	insertCmd.Flags().StringP("mail", "5", "", "Mail string option")
	insertCmd.Flags().StringP("birthday", "6", "", "Birthday string option")
	insertCmd.Flags().UintP("age", "7", 0, "Age integer option")
	insertCmd.Flags().UintP("hometown", "8", 0, "HometownID integer option")
	insertCmd.Flags().UintP("bloodtype", "9", 0, "BloodTypeID integer option")
}
