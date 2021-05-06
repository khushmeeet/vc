package cmd

import (
	"github.com/khushmeeet/vc/vc"
	"github.com/spf13/cobra"
)

// readTreeCmd represents the readTree command
var readTreeCmd = &cobra.Command{
	Use:   "readTree",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		oid := vc.GetOid(args[0])
		vc.ReadTree(oid)
	},
}

func init() {
	rootCmd.AddCommand(readTreeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readTreeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readTreeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
