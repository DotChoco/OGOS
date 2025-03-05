package manga

type Model struct {
	Id          uint16   `json:"id"`
	Name        string   `json:"name"` //varchar(60)
	Tags        uint8    `json:"tags"`
	OnGoing     bool     `json:"on_going"`
	MaxCaps     uint16   `json:"max_caps"`
	LastViewCap uint16   `json:"last_view_cap"`
	Related     []uint16 `json:"related"`
	ImgPath     string   `json:"img_path"` // varchar(34)
	// Path -> Anime/$Id.(jpg, jpeg, png)
}
