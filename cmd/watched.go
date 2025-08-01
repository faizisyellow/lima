package cmd

import (
	"fmt"
	"strconv"

	"github.com/faizisyellow/lima/movie"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// watchedCmd represents the watched command
var watchedCmd = &cobra.Command{
	Use:     "watched [position]",
	Aliases: []string{"wc"},
	Example: "wc, watched 8",
	Short:   "Update the movie to have been watched",
	Run:     WatchedRun,
}

func init() {
	rootCmd.AddCommand(watchedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// watchedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// watchedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func WatchedRun(cmd *cobra.Command, args []string) {

	if len(args) < 1 {
		cobra.CheckErr(fmt.Errorf("watched needs a position movie for the command"))
	}

	var intArgs []int

	movies, err := movie.ReadMovies(viper.GetString(EnvFile))
	if err != nil {
		cobra.CheckErr(err)
	}

	for _, arg := range args {

		intArg, err := strconv.Atoi(arg)
		if err != nil {
			cobra.CheckErr(fmt.Errorf("%v not valid args, err: %v", arg, err))
		}

		intArgs = append(intArgs, intArg)

	}

	for _, arg := range intArgs {

		if arg <= len(movies) {

			movies[arg-1].SetWatched()
			color.Green("success update movie with id %v to be watched", arg)
		} else {
			color.Yellow("no matching any movies id")
		}
	}

	err = movie.SaveMovie(viper.GetString(EnvFile), movies)
	if err != nil {
		cobra.CheckErr(err)
	}
}
