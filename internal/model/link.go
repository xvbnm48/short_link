package model

import "time"

type Link struct {
	ID          int       `json:"id"`
	ShortCode   string    `json:"short_code"`
	NewLink     string    `json:"new_link"`
	OriginalURL string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Click struct {
	ID        int       `json:"id"`
	LinkID    int       `json:"link_id"`
	IpAddress string    `json:"ip_address"`
	UserAgent string    `json:"user_agent"`
	CreatedAt time.Time `json:"created_at"`
}

type LinkCreateRequest struct {
	OriginalURL string `json:"original_url" validate:"required,url"`
	ShortCode   string `json:"short_code" validate:"required,min=3,max=20"`
}

type LinkCreateResponse struct {
	Data Link `json:"data"`
}
