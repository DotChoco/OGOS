package db

import (
	"ogos/src/models/anime"
	"ogos/src/models/db/userconfig"
	"ogos/src/models/manga"
)

type DBData struct {
	AnimeList []anime.Model
	MangaList []manga.Model

	AnimeThumbnails []anime.Thumbnail
	MangaThumbnails []manga.Thumbnail

	UserConf userconfig.Model
}
