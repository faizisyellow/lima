package cmd

import (
	"fmt"
	"strconv"

	"github.com/faizisyellow/lima/movie"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove [positions]",
	Example: "rm, remove 8 9 10 ..",
	Aliases: []string{"rm"},
	Short:   "Removes a movie from the list by the movie's position",
	Long: `removes a movie from the list by the movie's position  
	`,
	Run: RemoveRun,
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func RemoveRun(cmd *cobra.Command, args []string) {

	var intArgs []int

	if len(args) < 1 {
		cobra.CheckErr(fmt.Errorf("remove needs a position movie for the command"))
	}

	for _, arg := range args {

		intArg, err := strconv.Atoi(arg)
		if err != nil {
			cobra.CheckErr(fmt.Errorf("%v not valid args, err: %v", arg, err))
		}

		intArgs = append(intArgs, intArg)
	}

	for _, arg := range intArgs {
		err := movie.DeleteMovies(viper.GetString(EnvFile), arg)
		if err != nil {
			cobra.CheckErr(err)
		}
	}

	color.Red("remove movies sucessfull")
}
