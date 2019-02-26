package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Short comment for show sub-command",
	Long:  `Long comment for show sub-command`,
	RunE: func(cmd *cobra.Command, args []string) error {

		i, err := cmd.Flags().GetInt("int2")
		if err != nil {
			return err
		}

		b, err := cmd.Flags().GetBool("bool2")
		if err != nil {
			return err
		}

		s, err := cmd.Flags().GetString("str2")
		if err != nil {
			return err
		}

		fmt.Println("Integer option value:", i)
		fmt.Println("String option value:", s)
		fmt.Println("Boolean option value:", b)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	showCmd.Flags().IntP("int2", "i", 0, "integer option")
	showCmd.Flags().BoolP("bool2", "b", false, "boolean option")
	showCmd.Flags().StringP("str2", "s", "", "string option")
}
