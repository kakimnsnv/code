package models

type Profile struct {
	ID                uint   `json:"id"`
	UserID            uint   `json:"user_id"`
	Bio               string `json:"bio"`
	ProfilePictureURL string `json:"profile_picture_url"`
}
