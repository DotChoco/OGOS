package anime

type Thumbnail struct {
	Id          uint16 `json:"id"`
	Name        string `json:"name"`
	Tags        uint8  `json:"tags"`
	InLive      bool   `json:"in_live"`
	NextNewCap  uint8  `json:"next_new_cap"`
	LastViewCap uint16 `json:"last_view_cap"`
	ImgPath     string `json:"img_path"`
}
