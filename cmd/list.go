package cmd

import (
	"log"
	"os"
	"text/tabwriter"

	"github.com/faizisyellow/lima/internal/utils"
	"github.com/faizisyellow/lima/movie"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "list lists All Movies (default)",
	Long:    `list lists a movie with additional sort and filter`,
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
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func ListRun(cmd *cobra.Command, args []string) {

	movies, err := movie.ReadMovies(viper.GetString(EnvFile))
	if err != nil {
		log.Fatal(err)
	}

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 2, ' ', 0)
	pret := color.New(color.FgBlue)

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

	w.Flush()
}
