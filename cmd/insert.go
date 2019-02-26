package cmd

import (
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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("insert called")

		user := db.User{}
		user.Name, _ = cmd.Flags().GetString("name")
		user.Katakana, _ = cmd.Flags().GetString("kaka")

		sex, _ := cmd.Flags().GetInt("sex")
		user.GenderID = uint(sex)

		user.Tel, _ = cmd.Flags().GetString("tel")
		user.Mail, _ = cmd.Flags().GetString("mail")
		user.Birthday, _ = cmd.Flags().GetString("birthday")

		age, _ := cmd.Flags().GetInt("age")
		user.Age = uint(age)

		hometown, _ := cmd.Flags().GetInt("hometown")
		user.HometownID = uint(hometown)

		bloodtype, _ := cmd.Flags().GetInt("bloodtype")
		user.BloodTypeID = uint(bloodtype)

		db.Insert(&user)
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)

	insertCmd.Flags().StringP("name", "1", "", "Name string option")
	insertCmd.Flags().StringP("kaka", "2", "", "Katakana string option")
	insertCmd.Flags().IntP("sex", "3", 0, "GenderID integer option")
	insertCmd.Flags().StringP("tel", "4", "", "Tel string option")
	insertCmd.Flags().StringP("mail", "5", "", "Mail string option")
	insertCmd.Flags().StringP("birthday", "6", "", "Birthday string option")
	insertCmd.Flags().IntP("age", "7", 0, "Age integer option")
	insertCmd.Flags().IntP("hometown", "8", 0, "HometownID integer option")
	insertCmd.Flags().IntP("bloodtype", "9", 0, "BloodTypeID integer option")
}
