package anime

type Model struct {
	Id          uint16   `json:"id"`            // 2 bytes
	Name        string   `json:"name"`          // varchar(60)  120 bytes
	Tags        uint8    `json:"tags"`          // 1 byte
	InLive      bool     `json:"in_live"`       // 1 byte
	NextNewCap  uint8    `json:"next_new_cap"`  // 1 byte
	MaxCaps     uint16   `json:"max_caps"`      // 2 bytes
	LastViewCap uint16   `json:"last_view_cap"` // 2 bytes
	Related     []uint16 `json:"related"`       // varchar(60) 120 bytes
	ImgPath     string   `json:"img_path"`      // varchar(34) 68 bytes
	// Path -> Anime/$Id.(jpg, jpeg, png)
	// Class size in bytes: 318 bytes || 0.31 KB
}

// func _main() {
// 	// adfa := getDataChunksByIds([]int16{1, 1, 1}, "asdad")

// }

// func _getDataChunksByIds(chunk []int16, field string) []string {

// 	asdf := []string{"", "asdf"}
// 	for i := 0; i < len(chunk); i++ {
// 		//Search Items By Id's
// 		vasdv := fmt.Sprintf(`SELECT FROM animes (%s)  WHERE VALUE (%d)

// 		`, field, chunk[i])
// 		asdf = append(asdf, vasdv)
// 	}
// 	return asdf
// }
