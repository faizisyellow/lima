package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/faizisyellow/lima/movie"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "remove removes a movie from the list by the movie's position",
	Long: `remove removes a movie from the list by the movie's position  
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

	if len(args) == 0 {
		log.Fatal(fmt.Errorf("position movie is required"))
	}

	p, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalf("%v position is not valid", args[0])
	}

	err = movie.DeleteMovies(viper.GetString(EnvFile), p)
	if err != nil {
		log.Fatal(err)
	}

	color.Red("sucess remove movie with id: %v\n", p)
}
