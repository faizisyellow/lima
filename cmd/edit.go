package cmd

import (
	"fmt"
	"strconv"

	"github.com/faizisyellow/lima/movie"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:     "edit [position]",
	Example: "Edit 4 --status watched --category series --episode 4 --season 3",
	Short:   "updates a movie's property  that only update the given property",
	Run:     EditRun,
}

var (
	TitleUpdateMovie    string
	StatusUpdateMovie   string
	CategoryUpdateMovie string
	GoToUpdateMovie     bool
	YearUpdateMovie     string
	EpisodeUpdateMovie  int
	SeasonUpdateMovie   int
)

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	editCmd.Flags().BoolVarP(&GoToUpdateMovie, "go-to", "g", false, "enable or disable go-to property")
	editCmd.Flags().StringVarP(&StatusUpdateMovie, "status", "s", "", "movie's status watch (watched,ongoing,watchlist)")
	editCmd.Flags().StringVarP(&CategoryUpdateMovie, "cat", "c", "", "movie's category (series,movie)")
	editCmd.Flags().StringVarP(&TitleUpdateMovie, "title", "t", "", "movie's title")
	editCmd.Flags().StringVarP(&YearUpdateMovie, "year", "y", "", "movie's year")
	editCmd.Flags().IntVarP(&EpisodeUpdateMovie, "episode", "e", -1, "movie's episode")
	editCmd.Flags().IntVar(&SeasonUpdateMovie, "season", -1, "movie's season")

}

func EditRun(cmd *cobra.Command, args []string) {

	movies, err := movie.ReadMovies(viper.GetString(EnvFile))
	if err != nil {
		cobra.CheckErr(err)
	}

	p, err := strconv.Atoi(args[0])
	if err != nil {
		cobra.CheckErr(fmt.Errorf("%v arg is invalid: %v", args[0], err))
	}

	if p <= 0 || p > len(movies) {
		cobra.CheckErr(fmt.Errorf("no matching any movies"))
	}

	err = movies[p-1].UpdateProps(TitleUpdateMovie, StatusUpdateMovie, CategoryUpdateMovie, YearUpdateMovie, EpisodeUpdateMovie, SeasonUpdateMovie, GoToUpdateMovie)
	if err != nil {
		cobra.CheckErr(fmt.Errorf("can not update episode or season. not a series"))
	}

	err = movie.SaveMovie(viper.GetString(EnvFile), movies)
	if err != nil {
		cobra.CheckErr(err)
	}

	color.Yellow("update movie with id %v successfull", p)
}
