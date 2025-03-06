package manga

type Model struct {
	Id          uint16   `json:"id"`            // 2 bytes
	Name        string   `json:"name"`          //varchar(60) 120 bytes
	Tags        uint8    `json:"tags"`          // 1 byte
	OnGoing     bool     `json:"on_going"`      // 1 byte
	MaxCaps     uint16   `json:"max_caps"`      // 2 bytes
	LastViewCap uint16   `json:"last_view_cap"` // 2 bytes
	Related     []uint16 `json:"related"`       // varchar(60) 120 bytes
	ImgPath     string   `json:"img_path"`      // varchar(34) 68 bytes
	// Path -> Anime/$Id.(jpg, jpeg, png)
	// Class size in bytes: 316 bytes || 0.31 KB
}
