package cmd

import (
	"log"
	"os"
	"slices"
	"strings"
	"text/tabwriter"

	"github.com/faizisyellow/lima/internal/utils"
	"github.com/faizisyellow/lima/movie"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	statusMovie   string
	categoryMovie string
	goTo          bool
	date          bool
	sort          string
	search        string
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Example: "list --status watched --category movie",
	Short:   "Lists all movies and series",
	Long:    `lists a movie with additional sort, filter, and searching`,
	Run:     ListRun,
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	listCmd.Flags().StringVarP(&statusMovie, "status", "s", "", "to filter list of movies by status")
	listCmd.Flags().StringVarP(&categoryMovie, "cat", "c", "", "to filter list of movies by category")
	listCmd.Flags().BoolVarP(&goTo, "go-to", "g", false, "to filter list of movies only go-to movie")
	listCmd.Flags().StringVar(&search, "search", "", "to searching a movie by the movie's title")

	listCmd.Flags().BoolVarP(&date, "date", "d", false, "to display date column when the movie is added to the list")
	listCmd.Flags().StringVar(&sort, "sort", "older", "to sorting list of movies by the time the movie added to the list")
}

func ListRun(cmd *cobra.Command, args []string) {

	var filteredMovies []*movie.Movie

	movies, err := movie.ReadMovies(viper.GetString(EnvFile))
	if err != nil {
		log.Fatal(err)
	}

	w := tabwriter.NewWriter(os.Stdout, 3, 2, 2, ' ', 0)
	pret := color.New(color.FgBlue)

	for _, movie := range movies {

		/*
			If there's no status skip the conditional.
			If there's status and the status is the same with the the movie's status
			Skip the conditional.
			If there's status and the status is not the same skip the current loop immediately.
		*/
		if statusMovie != "" && movie.Status != statusMovie {
			continue
		}

		if categoryMovie != "" && movie.Category != categoryMovie {
			continue
		}

		if goTo && !movie.IsGoTo {
			continue
		}

		if search != "" && !strings.Contains(strings.ToLower(movie.Title), strings.ToLower(search)) {
			continue
		}

		filteredMovies = append(filteredMovies, &movie)
	}

	slices.SortFunc(filteredMovies, func(a *movie.Movie, b *movie.Movie) int {

		// sorted desc
		if sort == "latest" {
			return b.AddAt.Compare(a.AddAt)
		}

		// sorted asc
		return a.AddAt.Compare(b.AddAt)
	})

	for _, movie := range filteredMovies {

		pret.Fprintln(w,
			movie.Label()+"\t",
			utils.ToUpperFirst(movie.Title)+"\t",
			movie.Year+"\t",
			movie.PrettyCat()+"\t",
			movie.PrettyStats()+"\t",
			movie.PrettyRW()+"\t",
			movie.DisplayDate(date),
		)

	}

	if len(filteredMovies) <= 0 {
		color.Yellow("movies not found")
	}

	w.Flush()

}
