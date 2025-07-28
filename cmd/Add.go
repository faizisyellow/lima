package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/faizisyellow/lima/movie"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	year     string
	isGoTo   bool
	category string
	status   string
	episodes int
	season   int
)

// AddCmd represents the Add command
var AddCmd = &cobra.Command{
	Use:     "add [title movie]",
	Example: "add 'title' -y, --year 2025 -s, --status watchlist -c, --category movie -g, --goto",
	Short:   "Adds a movie or series to the list",
	Long:    `Adds a movie or series to the list.`,
	Run:     AddRun,
}

func init() {
	rootCmd.AddCommand(AddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	AddCmd.Flags().StringVarP(&year, "year", "y", strconv.Itoa(time.Now().Year()), "the release year of movie")
	AddCmd.Flags().BoolVarP(&isGoTo, "go-to", "g", false, "go-to is a prop to indicate that the movie is rewatchable")
	AddCmd.Flags().StringVarP(&category, "category", "c", "movie", "the category of the movie")
	AddCmd.Flags().StringVarP(&status, "status", "s", "watchlist", "status prop to indicate the movie is watchlist or still watching or already seen [watchlist,watching,watched]")
	AddCmd.Flags().IntVarP(&episodes, "episode", "e", -1, "if adding a series add an episode")
	AddCmd.Flags().IntVar(&season, "season", -1, "if adding a series add an season")
}

func AddRun(cmd *cobra.Command, args []string) {

	if len(args) == 0 {
		log.Fatal(fmt.Errorf("title required"))
	}

	if (category == "series" && status != "watchlist") && (episodes <= 0 && season <= 0) {
		cobra.CheckErr(fmt.Errorf("episode or season at least 1"))
	}

	nm, err := movie.New(args[0], year, category, status, episodes, season, isGoTo)
	if err != nil {
		log.Fatal(err)
	}

	m, _ := movie.ReadMovies(viper.GetString(EnvFile))

	m = append(m, nm)

	if err := movie.SaveMovie(viper.GetString(EnvFile), m); err != nil {
		log.Fatal(err)
	}

	color.Green("add new movie successfully")
}
