package db

import (
	"ogos/src/models/anime"
	"ogos/src/models/config"
	"ogos/src/models/manga"
)

type DBData struct {
	AnimeList []anime.Model
	MangaList []manga.Model

	AnimeThumbnails []anime.Thumbnail
	MangaThumbnails []manga.Thumbnail

	Conf config.Model
}
