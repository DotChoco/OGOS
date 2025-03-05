package manga

type Thumbnail struct {
	Id          uint16 `json:"id"`
	Name        string `json:"name"`
	Tags        uint8  `json:"tags"`
	OnGoing     bool   `json:"on_going"`
	LastViewCap uint16 `json:"last_view_cap"`
	ImgPath     string `json:"img_path"`
}
