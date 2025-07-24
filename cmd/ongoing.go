package cmd

import (
	"log"
	"strconv"
	"strings"

	"github.com/faizisyellow/lima/movie"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

	pos := args[0]
	dur := args[1]

	movies, err := movie.ReadMovies(viper.GetString(EnvFile))
	if err != nil {
		log.Fatal(err)
	}

	p, err := strconv.Atoi(pos)
	if err != nil {
		log.Fatalf("%v position is not valid", pos)
	}

	if p == 0 || p > len(movies) {
		log.Println(p, "doesn't match any movies")
		return
	}

	if !strings.Contains(dur, ":") {
		log.Println("duration format is invalid. (H:M:S)")
		return
	}

	if movies[p-1].Status != "ongoing" {
		log.Println("status movie not ongoing.")
		return
	}

	err = movies[p-1].SetRecentWatch(dur)
	if err != nil {
		log.Fatal(err)
	}

	err = movie.SaveMovie(viper.GetString(EnvFile), movies)
	if err != nil {
		log.Fatal(err)
	}

	color.Green("Update Latest Watch successfully")
}
