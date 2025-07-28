package movie

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/faizisyellow/lima/internal/utils"
)

type Movie struct {
	Position    int
	Title       string
	Year        string
	Status      string // [watchlist, watched, watching]
	Category    string // [movie, series]
	IsGoTo      bool
	RecentWatch string // h:m:s
	AddAt       time.Time
	Season      int
	Episodes    int
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

func (mv *Movie) DisplayDate(opt bool) string {

	if !opt {
		return ""
	}

	return mv.AddAt.Format(time.Stamp)

}

func (mv *Movie) SetRecentWatch(dur string) error {

	if mv.RecentWatch != "" {

		ti, err := time.Parse(time.TimeOnly, dur)
		if err != nil {
			return err
		}

		md, err := time.Parse(time.TimeOnly, mv.RecentWatch)
		if err != nil {
			return err
		}

		if ti.Before(md) {
			return fmt.Errorf("duration can not less then latest watch")
		}
	}

	mv.RecentWatch = dur

	return nil
}

func (mv *Movie) SetWatched() {

	if mv.RecentWatch != "" {
		mv.RecentWatch = ""
	}

	mv.Status = "watched"
}

func (mv *Movie) UpdateProps(title, status, category, year string, episode, season int) error {

	if title != "" {
		mv.Title = title
	}

	if status != "" {

		// Set null if movie still have duration
		if mv.Status == "watching" && status != "watching" {
			mv.RecentWatch = ""
		}

		mv.Status = status
	}

	if category != "" {
		mv.Category = category
	}

	if year != "" {
		mv.Year = year
	}

	if mv.Category == "series" {

		if episode != -1 {
			mv.Episodes = episode
		}

		if season != -1 {
			mv.Season = season
		}
	} else if episode > 0 || season > 0 {
		return fmt.Errorf("can not update episode or season. not a series")
	}

	return nil
}

func (mv *Movie) PrettyCat() string {

	if mv.Category != "series" || mv.Status != "watching" {
		return utils.ToUpperFirst(mv.Category)
	}

	return fmt.Sprintf("%v (Season %v)", utils.ToUpperFirst(mv.Category), mv.Season)
}

func (mv *Movie) PrettyStats() string {

	var ep string

	if mv.Category != "series" || mv.Status != "watching" {
		return utils.ToUpperFirst(mv.Status)
	}

	if mv.Episodes > 1 {
		ep = "Episodes"
	} else {
		ep = "Episode"
	}

	return fmt.Sprintf("%v (%v %v)", utils.ToUpperFirst(mv.Status), ep, mv.Episodes)
}

func New(title, year, category, status string, episode, season int, isGoto bool) (Movie, error) {

	status = strings.ToLower(status)
	category = strings.ToLower(category)

	if category != "movie" && category != "series" {
		return Movie{}, fmt.Errorf("property category: %v not valid", category)
	}

	if status != "watchlist" && status != "watched" && status != "watching" {
		return Movie{}, fmt.Errorf("property status: %v not valid", status)
	}

	if len(year) != 4 {
		return Movie{}, fmt.Errorf("property year: %v not valid", year)
	}

	if category == "series" && status != "watchlist" && season < 1 && episode < 1 {
		return Movie{}, fmt.Errorf("you add a series, need season and episode")
	}

	return Movie{
		Title:    title,
		Year:     year,
		Status:   status,
		Category: category,
		IsGoTo:   isGoto,
		AddAt:    time.Now().Local(),
		Season:   season,
		Episodes: episode,
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
