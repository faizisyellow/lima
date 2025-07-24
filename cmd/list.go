package cmd

import (
	"log"
	"os"
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
	Short:   "list lists All Movies	",
	Long:    `list lists a movie with additional sort, filter, and searching`,
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
	listCmd.Flags().StringVar(&sort, "sort", "latest", "to sorting list of movies by the time the movie added to the list")
}

func ListRun(cmd *cobra.Command, args []string) {

	movies, err := movie.ReadMovies(viper.GetString(EnvFile))
	if err != nil {
		log.Fatal(err)
	}

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 2, ' ', 0)
	pret := color.New(color.FgBlue)

	switch {

	// Only filter by status
	case statusMovie != "" && categoryMovie == "" && !goTo && search == "":

		for _, movie := range movies {

			if statusMovie == movie.Status {
				pret.Fprintln(w,
					movie.Label()+"\t",
					utils.ToUpperFirst(movie.Title)+"\t",
					movie.Year+"\t",
					utils.ToUpperFirst(movie.Category)+"\t",
					utils.ToUpperFirst(movie.Status)+"\t",
					movie.PrettyRW(),
				)
			}

		}

	// Only filter by category
	case categoryMovie != "" && statusMovie == "" && !goTo && search == "":

		for _, movie := range movies {

			if categoryMovie == movie.Category {
				pret.Fprintln(w,
					movie.Label()+"\t",
					utils.ToUpperFirst(movie.Title)+"\t",
					movie.Year+"\t",
					utils.ToUpperFirst(movie.Category)+"\t",
					utils.ToUpperFirst(movie.Status)+"\t",
					movie.PrettyRW(),
				)
			}

		}

	// Only filter by go-to
	case goTo && statusMovie == "" && categoryMovie == "" && search == "":

		for _, movie := range movies {

			if movie.IsGoTo {
				pret.Fprintln(w,
					movie.Label()+"\t",
					utils.ToUpperFirst(movie.Title)+"\t",
					movie.Year+"\t",
					utils.ToUpperFirst(movie.Category)+"\t",
					utils.ToUpperFirst(movie.Status)+"\t",
					movie.PrettyRW(),
				)
			}

		}

	// Only filter by searching
	case search != "" && statusMovie == "" && categoryMovie == "" && !goTo:

		for _, movie := range movies {

			if strings.Contains(movie.Title, search) {
				pret.Fprintln(w,
					movie.Label()+"\t",
					utils.ToUpperFirst(movie.Title)+"\t",
					movie.Year+"\t",
					utils.ToUpperFirst(movie.Category)+"\t",
					utils.ToUpperFirst(movie.Status)+"\t",
					movie.PrettyRW(),
				)
			}

		}

	// filter by status movie and category movie
	case statusMovie != "" && categoryMovie != "":

		for _, movie := range movies {

			if statusMovie == movie.Status && categoryMovie == movie.Category {
				pret.Fprintln(w,
					movie.Label()+"\t",
					utils.ToUpperFirst(movie.Title)+"\t",
					movie.Year+"\t",
					utils.ToUpperFirst(movie.Category)+"\t",
					utils.ToUpperFirst(movie.Status)+"\t",
					movie.PrettyRW(),
				)
			}

		}

	// filter by status movie and go-to
	case statusMovie != "" && goTo:

		for _, movie := range movies {

			if statusMovie == movie.Status && movie.IsGoTo {
				pret.Fprintln(w,
					movie.Label()+"\t",
					utils.ToUpperFirst(movie.Title)+"\t",
					movie.Year+"\t",
					utils.ToUpperFirst(movie.Category)+"\t",
					utils.ToUpperFirst(movie.Status)+"\t",
					movie.PrettyRW(),
				)
			}

		}

	// filter by category movie and go-to
	case categoryMovie != "" && goTo:

		for _, movie := range movies {

			if categoryMovie == movie.Category && movie.IsGoTo {
				pret.Fprintln(w,
					movie.Label()+"\t",
					utils.ToUpperFirst(movie.Title)+"\t",
					movie.Year+"\t",
					utils.ToUpperFirst(movie.Category)+"\t",
					utils.ToUpperFirst(movie.Status)+"\t",
					movie.PrettyRW(),
				)
			}

		}

	// filter by searching and status movie
	case search != "" && statusMovie != "":

		for _, movie := range movies {

			if strings.Contains(movie.Title, search) && statusMovie == movie.Status {
				pret.Fprintln(w,
					movie.Label()+"\t",
					utils.ToUpperFirst(movie.Title)+"\t",
					movie.Year+"\t",
					utils.ToUpperFirst(movie.Category)+"\t",
					utils.ToUpperFirst(movie.Status)+"\t",
					movie.PrettyRW(),
				)
			}

		}

	// filter by searching and category movie
	case search != "" && categoryMovie != "":

		for _, movie := range movies {

			if strings.Contains(movie.Title, search) && categoryMovie == movie.Category {
				pret.Fprintln(w,
					movie.Label()+"\t",
					utils.ToUpperFirst(movie.Title)+"\t",
					movie.Year+"\t",
					utils.ToUpperFirst(movie.Category)+"\t",
					utils.ToUpperFirst(movie.Status)+"\t",
					movie.PrettyRW(),
				)
			}

		}

	// filter by searching and go-to
	case search != "" && goTo:

		for _, movie := range movies {

			if strings.Contains(movie.Title, search) && movie.IsGoTo {
				pret.Fprintln(w,
					movie.Label()+"\t",
					utils.ToUpperFirst(movie.Title)+"\t",
					movie.Year+"\t",
					utils.ToUpperFirst(movie.Category)+"\t",
					utils.ToUpperFirst(movie.Status)+"\t",
					movie.PrettyRW(),
				)
			}

		}

	default:
		for _, movie := range movies {
			pret.Fprintln(w,
				movie.Label()+"\t",
				utils.ToUpperFirst(movie.Title)+"\t",
				movie.Year+"\t",
				utils.ToUpperFirst(movie.Category)+"\t",
				utils.ToUpperFirst(movie.Status)+"\t",
				movie.PrettyRW(),
			)
		}

	}

	w.Flush()

}
