package dto

type SocialMediaCreateReq struct {
	UserID         uint64 `json:"user_id"`
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}
