package dto

type AlbumListDTO struct {
	AlbumUuid       string `json:"album_uuid"`
	Title           string `json:"title"`
	ReleaseDate     string `json:"release_date"`
	CoverImage      string `json:"cover_image"`
	Description     string `json:"description"`
	BackgroundStory string `json:"background"`
	Producer        string `json:"producer"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
