package dto

import "github.com/gogf/gf/v2/os/gtime"

type AlbumListDTO struct {
	AlbumUuid       string      `json:"album_uuid"`
	Title           string      `json:"title"`
	ReleaseDate     *gtime.Time `json:"release_date"`
	TotalSongs      int         `json:"total_songs"`
	TotalDuration   string      `json:"total_duration"`
	CoverImage      string      `json:"cover_image"`
	Description     string      `json:"description"`
	BackgroundStory string      `json:"background_story"`
	Producer        string      `json:"producer"`
	CreatedAt       *gtime.Time `json:"created_at"`
	UpdatedAt       *gtime.Time `json:"updated_at"`
}
