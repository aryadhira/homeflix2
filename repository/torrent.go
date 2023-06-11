package repository

import (
	"encoding/json"
	"fmt"
	"homeflix2/helper"
	"homeflix2/models"
	"homeflix2/settings"
	"log"
	"net/http"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
)

type TorrentRepo struct{}

func (t *TorrentRepo) GetListMovies() (interface{}, error) {
	var err error

	log.Println("Getting Movie List...")
	err = new(settings.Setting).GetConfig()
	helper.ErrorHandler(err)

	baseurl := os.Getenv("BASEURL_GETLIST")
	ytspage := os.Getenv("PAGE")
	limit := os.Getenv("LIMIT")

	totalpage, err := strconv.Atoi(ytspage)
	if err != nil {
		return nil, err
	}

	results := []models.YtsModel{}

	for i := 1; i <= totalpage; i++ {
		client := http.Client{}
		apiresult := models.YtsModel{}
		pagestr := strconv.Itoa(i)
		apiurl := fmt.Sprintf("%v?limit=%v&page=%v", baseurl, limit, pagestr)

		req, err := http.NewRequest("GET", apiurl, nil)
		if err != nil {
			return nil, err
		}

		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}

		err = json.NewDecoder(res.Body).Decode(&apiresult)
		if err != nil {
			return nil, err
		}

		results = append(results, apiresult)

		defer res.Body.Close()
	}

	return results, nil
}

func (t *TorrentRepo) SaveMovies(moviesList []interface{}) error {
	repo := MongoRepo{}

	err := repo.InsertMany(new(models.YtsModel).CollectionName(), moviesList)
	if err != nil {
		return err
	}

	return nil
}

func (t *TorrentRepo) ConstructMovieData(movies []models.Movie) []interface{} {
	moviedata := make([]interface{}, 0)

	for _, each := range movies {
		moviedata = append(moviedata, each)
	}

	return moviedata
}

func (t *TorrentRepo) SyncMovies() error {
	log.Println("Sync Movie Start...")
	db := MongoRepo{}

	log.Println("Deleting current movie list data...")
	err := db.DeleteMany(new(models.YtsModel).CollectionName(), bson.M{})
	if err != nil {
		return err
	}

	ytsraw, err := t.GetListMovies()
	if err != nil {
		return err
	}

	if ytsraw != nil {
		movielist := ytsraw.([]models.YtsModel)

		for _, each := range movielist {
			moviedata := t.ConstructMovieData(each.Data.Movies)
			err = t.SaveMovies(moviedata)
			if err != nil {
				return nil
			}
		}

	}

	log.Println("Sync Movie Finished...")

	return nil
}
