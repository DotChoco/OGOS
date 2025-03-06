package anime

type Thumbnail struct {
	Id          uint16 `json:"id"`            // 2 bytes
	Name        string `json:"name"`          // varchar(60)  120 bytes
	Tags        uint8  `json:"tags"`          // 1 byte
	InLive      bool   `json:"in_live"`       // 1 byte
	NextNewCap  uint8  `json:"next_new_cap"`  // 1 byte
	LastViewCap uint16 `json:"last_view_cap"` // 2 bytes
	ImgPath     string `json:"img_path"`      // varchar(34) 68 bytes
	// Path -> Anime/$Id.(jpg, jpeg, png)
	// Class size in bytes: 195 bytes || 0.19 KB
}
