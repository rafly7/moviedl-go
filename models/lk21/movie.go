package lk21

type Movie struct {
	Title      string   `json:"title"`
	ImageSrc   string   `json:"image_src"`
	SlugBypass string   `json:"slug"`
	Genres     []string `json:"genres"`
}
