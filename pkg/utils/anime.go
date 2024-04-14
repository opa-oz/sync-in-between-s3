package utils

import "fmt"

const urlPrefix = "http://db.anime-recommend.ru/anime/covers"

func getAnimeUrl(id int) string {
	return fmt.Sprintf("%s/%d.jpg", urlPrefix, id)
}

func GetUrl(entityType string, id int) string {
	if entityType == "manga" {
		panic("Unexpected manga")
	}

	if entityType == "anime" {
		return getAnimeUrl(id)
	}

	panic("OMG SOMETHING BAD")
}
