package cmd

import (
	"github.com/spf13/cobra"
)

// ongoingCmd represents the ongoing command
var ongoingCmd = &cobra.Command{
	Use:     "ongoing",
	Aliases: []string{"og"},
	Short:   "ongoing update movie's recent watch duration that still to be watching: args (id, duration)",
	Long:    `ongoing update movie's recent watch duration with format (h:m:s)`,
	Run:     OnGoingRun,
}

func init() {
	rootCmd.AddCommand(ongoingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ongoingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ongoingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func OnGoingRun(cmd *cobra.Command, args []string) {

}
