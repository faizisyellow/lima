package movie

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestMovie(t *testing.T) {

	t.Run("should success get a list of movies", func(t *testing.T) {

		tadd := time.Now().Local()

		input := []Movie{
			{
				Title:       "All The Bright Places",
				Year:        "2020",
				Status:      "watched",
				Category:    "movie",
				IsGoTo:      true,
				RecentWatch: "",
				AddAt:       tadd,
			},
			{
				Title:       "He's all that",
				Year:        "2020",
				Status:      "watched",
				Category:    "movie",
				IsGoTo:      true,
				RecentWatch: "",
				AddAt:       tadd,
			},
			{
				Title:       "virgin River",
				Year:        "2016",
				Status:      "watched",
				Category:    "series",
				IsGoTo:      true,
				RecentWatch: "",
				AddAt:       tadd,
			},
		}

		file, err := os.CreateTemp("./", "test.*.json")
		if err != nil {
			t.Errorf("error create temporary file")
		}

		defer os.Remove(file.Name())

		b, err := json.Marshal(input)
		if err != nil {
			t.Errorf("error marshal a type to json :%v", err)
		}

		err = os.WriteFile(file.Name(), b, 0644)
		if err != nil {
			t.Errorf("error write a  temporary file: %v", err)
		}

		expected := []Movie{
			{
				Position:    1,
				Title:       "All The Bright Places",
				Year:        "2020",
				Status:      "watched",
				Category:    "movie",
				IsGoTo:      true,
				RecentWatch: "",
				AddAt:       tadd,
			},
			{
				Position:    2,
				Title:       "He's all that",
				Year:        "2020",
				Status:      "watched",
				Category:    "movie",
				IsGoTo:      true,
				RecentWatch: "",
				AddAt:       tadd,
			},
			{
				Position:    3,
				Title:       "virgin River",
				Year:        "2016",
				Status:      "watched",
				Category:    "series",
				IsGoTo:      true,
				RecentWatch: "",
				AddAt:       tadd,
			},
		}

		movies, err := ReadMovies(file.Name())
		if err != nil {
			t.Errorf("expected nil but got error :%v", err)
		}

		if !reflect.DeepEqual(expected, movies) {
			t.Error("expected to be match but got different")
		}
	})

	t.Run("should failed because there's no movies added yet", func(t *testing.T) {

		file, err := os.CreateTemp("./", "test.*.json")
		if err != nil {
			t.Errorf("error create temporary file")
		}

		defer os.Remove(file.Name())

		b, err := json.Marshal([]Movie{})
		if err != nil {
			t.Errorf("error marshal a type to json :%v", err)
		}

		err = os.WriteFile(file.Name(), b, 0644)
		if err != nil {
			t.Errorf("error write a  temporary file: %v", err)
		}

		movie, err := ReadMovies(file.Name())
		if err == nil || len(movie) != 0 {
			t.Error("expected error message but got nil")
		}
	})

	t.Run("should save a movie successfully", func(t *testing.T) {

		input := Movie{
			Title:       "All The Bright Places",
			Year:        "2020",
			Status:      "watched",
			Category:    "movie",
			IsGoTo:      true,
			RecentWatch: "",
			AddAt:       time.Now().Local(),
		}

		file, err := os.CreateTemp("./", "test.*.json")
		if err != nil {
			t.Errorf("error create temporary file")
		}

		defer os.Remove(file.Name())

		err = SaveMovie(file.Name(), []Movie{input})
		if err != nil {
			t.Errorf("expected nil but got err: %v", err)
		}
	})

	t.Run("should fail save a movie", func(t *testing.T) {

		err := SaveMovie("", nil)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("should success remove a movie", func(t *testing.T) {

		input := Movie{
			Title:       "All The Bright Places",
			Year:        "2020",
			Status:      "watched",
			Category:    "movie",
			IsGoTo:      true,
			RecentWatch: "",
			AddAt:       time.Now().Local(),
		}

		file, err := os.CreateTemp("./", "test.*.json")
		if err != nil {
			t.Errorf("error create temporary file: %v", err)
		}

		defer os.Remove(file.Name())

		b, err := json.Marshal([]Movie{input})
		if err != nil {
			t.Errorf("error marshal a type to json :%v", err)
		}

		err = os.WriteFile(file.Name(), b, 0644)
		if err != nil {
			t.Errorf("error write a  temporary file: %v", err)
		}

		err = DeleteMovies(file.Name(), 1)
		if err != nil {
			t.Errorf("expected nil but got error: %v", err)
		}
	})

	t.Run("shoud fail remove a movie because movie there's no movies", func(t *testing.T) {

		file, err := os.CreateTemp("./", "test.*.json")
		if err != nil {
			t.Errorf("error create temporary file: %v", err)
		}

		defer os.Remove(file.Name())

		b, err := json.Marshal([]Movie{})
		if err != nil {
			t.Errorf("error marshal a type to json :%v", err)
		}

		err = os.WriteFile(file.Name(), b, 0644)
		if err != nil {
			t.Errorf("error write a  temporary file: %v", err)
		}

		err = DeleteMovies(file.Name(), 10)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})

	t.Run("shoud fail remove a movie because movie not found", func(t *testing.T) {

		input := Movie{
			Title:       "All The Bright Places",
			Year:        "2020",
			Status:      "watched",
			Category:    "movie",
			IsGoTo:      true,
			RecentWatch: "",
			AddAt:       time.Now().Local(),
		}

		file, err := os.CreateTemp("./", "test.*.json")
		if err != nil {
			t.Errorf("error create temporary file: %v", err)
		}

		defer os.Remove(file.Name())

		b, err := json.Marshal([]Movie{input})
		if err != nil {
			t.Errorf("error marshal a type to json :%v", err)
		}

		err = os.WriteFile(file.Name(), b, 0644)
		if err != nil {
			t.Errorf("error write a  temporary file: %v", err)
		}

		err = DeleteMovies(file.Name(), 2)
		if err == nil {
			t.Errorf("expected error but got nil")
		}
	})
}
