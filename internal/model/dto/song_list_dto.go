package dto

import "github.com/gogf/gf/v2/os/gtime"

type SongListDTO struct {
	SongUuid    string      `json:"songUuid"`
	AlbumId     string      `json:"albumId"`
	SongTitle   string      `json:"songTitle"`
	Duration    string      `json:"duration"`
	Lyrics      string      `json:"lyrics"`
	Writer      string      `json:"writer"`
	Producer    string      `json:"producer"`
	IsSingle    int         `json:"isSingle"`
	Genre       string      `json:"genre"`
	Label       string      `json:"label"`
	Language    string      `json:"language"`
	Instruments string      `json:"instruments"`
	ReleaseDate *gtime.Time `json:"releaseDate"`
	CreatedAt   *gtime.Time `json:"createdAt"`
	UpdatedAt   *gtime.Time `json:"updatedAt"`
}
