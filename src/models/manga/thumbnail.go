package manga

type Thumbnail struct {
	Id          uint16 `json:"id"`            // 2 bytes
	Name        string `json:"name"`          //varchar(60) 120 bytes
	Tags        uint8  `json:"tags"`          // 1 byte
	OnGoing     bool   `json:"on_going"`      // 1 byte
	LastViewCap uint16 `json:"last_view_cap"` // 2 bytes
	ImgPath     string `json:"img_path"`      // varchar(34) 68 bytes
	// Class size  in bytes: 194 bytes || 0.19 KB
}
