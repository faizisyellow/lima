package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/faizisyellow/lima/movie"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ongoingCmd represents the ongoing command
var ongoingCmd = &cobra.Command{
	Use:     "ongoing [position duration]",
	Aliases: []string{"og"},
	Example: `og, ongoing 7 "01:00:00"`,
	Short:   "Update movie's recent watch duration that still to be watching",
	Long:    `update movie's recent watch duration with format (h:m:s)`,
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

	if len(args) < 1 {
		cobra.CheckErr(fmt.Errorf("ongoing needs a position movie for the command"))
	}

	pos := args[0]
	dur := args[1]

	movies, err := movie.ReadMovies(viper.GetString(EnvFile))
	if err != nil {
		cobra.CheckErr(err)
	}

	p, err := strconv.Atoi(pos)
	if err != nil {
		cobra.CheckErr(fmt.Errorf("%v position is not valid", pos))
	}

	if p <= 0 || p > len(movies) {
		cobra.CheckErr(fmt.Errorf("%v doesn't match any movies", p))
		return
	}

	if !strings.Contains(dur, ":") {
		cobra.CheckErr(fmt.Errorf("duration format is invalid. (h:m:s)"))
		return
	}

	if movies[p-1].Status != "ongoing" {
		cobra.CheckErr(fmt.Errorf("movie with id %v is not ongoing", p))
		return
	}

	err = movies[p-1].SetRecentWatch(dur)
	if err != nil {
		cobra.CheckErr(err)
	}

	err = movie.SaveMovie(viper.GetString(EnvFile), movies)
	if err != nil {
		cobra.CheckErr(err)
	}

	color.Green("update Latest Watch successfully")
}
