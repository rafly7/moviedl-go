package lk21

type Latest struct {
	Title      string   `json:"title"`
	ImageSrc   string   `json:"image_src"`
	SlugBypass string   `json:"slug"`
	Rating     string   `json:"rating"`
	Genres     []string `json:"genres"`
}
