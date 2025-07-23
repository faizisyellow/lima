package movie

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Movie struct {
	Position    int
	Title       string
	Year        string
	Status      string // [watchlist, watched, ongoing]
	Category    string // [movie, series]
	IsGoTo      bool
	RecentWatch string // h:m:s
	AddAt       time.Time
}

func (mv *Movie) Label() string {
	return strconv.Itoa(mv.Position) + "."
}

func (mv *Movie) PrettyRW() string {

	if mv.RecentWatch == "" {
		return ""
	} else {
		return mv.RecentWatch
	}

}

func (mv *Movie) SetRecentWatch(dur string) {
	mv.RecentWatch = dur
}

func New(title, year, category, status string, isGoto bool) (Movie, error) {

	status = strings.ToLower(status)
	category = strings.ToLower(category)

	if category != "movie" && category != "series" {
		return Movie{}, fmt.Errorf("property category: %v not valid", category)
	}

	if status != "watchlist" && status != "watched" && status != "ongoing" {
		return Movie{}, fmt.Errorf("property status: %v not valid", status)
	}

	if len(year) != 4 {
		return Movie{}, fmt.Errorf("property year: %v not valid", year)
	}

	return Movie{
		Title:    title,
		Year:     year,
		Status:   status,
		Category: category,
		IsGoTo:   isGoto,
		AddAt:    time.Now().Local(),
	}, nil

}

func SaveMovie(filename string, movies []Movie) error {

	b, err := json.Marshal(movies)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadMovies(filename string) ([]Movie, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		return []Movie{}, err
	}

	var movies []Movie

	err = json.Unmarshal(data, &movies)
	if err != nil {
		return []Movie{}, err
	}

	if len(movies) == 0 {
		return []Movie{}, fmt.Errorf("no movies added yet, please add a movie")
	}

	for i := range movies {
		movies[i].Position = i + 1
	}

	return movies, nil
}

func DeleteMovies(filename string, p int) error {

	m, err := ReadMovies(filename)
	if err != nil {
		return err
	}

	if len(m) == 0 {
		return fmt.Errorf("no movies in the list, add one")
	}

	found := slices.ContainsFunc(m, func(m Movie) bool {
		return m.Position == p
	})

	if !found {
		return fmt.Errorf("no movies with this id: %v in the movie list", p)
	}

	modifiedMovie := slices.DeleteFunc(m, func(m Movie) bool {
		return m.Position == p
	})

	err = SaveMovie(filename, modifiedMovie)
	if err != nil {
		return err
	}

	return nil
}
