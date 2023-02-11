package types

type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
}

type GeneralResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
