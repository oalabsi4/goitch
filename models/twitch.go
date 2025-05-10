package models

type ChannelData struct {
	Id           string   `json:"id"`
	UserId       string   `json:"user_id"`
	UserLogin    string   `json:"user_login"`
	UserName     string   `json:"user_name"`
	GameId       string   `json:"game_id"`
	GameName     string   `json:"game_name"`
	Type         string   `json:"type"`
	Title        string   `json:"title"`
	ViewerCount  int      `json:"viewer_count"`
	StartedAt    string   `json:"started_at"`
	Language     string   `json:"language"`
	ThumbnailUrl string   `json:"thumbnail_url"`
	TagIds       []string `json:"tag_ids"`
	Tags         []string `json:"tags"`
	IsMature     bool     `json:"is_mature"`
}
type ChanelPagination struct {
	Curser string `json:"cursor"`
}

type ChannelResponse struct {
	Data       []ChannelData    `json:"data"`
	Pagination ChanelPagination `json:"pagination"`
}
