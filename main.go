package main

import (
	"homeflix2/helper"
	"homeflix2/repository"
)

func main() {
	torrent := repository.TorrentRepo{}

	err := torrent.SyncMovies()
	helper.ErrorHandler(err)
}
