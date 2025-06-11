package repository

import (
	"api-service/internal/model"
	"database/sql"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2/log"
)

type LinkRepository interface {
	CreateShortLink(linkRequest model.LinkCreateRequest) (model.Link, error)
	GetShortLink(id int) (model.Link, error)
	GetOriginalURL(shortCode string) (string, error)
}

type linkRepository struct {
	db *sql.DB
}

// GetOriginalURL implements LinkRepository.
func (l *linkRepository) GetOriginalURL(shortCode string) (string, error) {
	log.Info("Executing Func GetOriginalURL with param: ", shortCode)
	query := `SELECT original_url FROM links WHERE short_code = $1`
	fmt.Println("Query:", query)
	var originalURL string
	err := l.db.QueryRow(query, shortCode).Scan(&originalURL)
	if err != nil {
		log.Error("Failed to get original URL:", err)
		return "", err
	}
	log.Info("Original URL retrieved successfully: ", originalURL)
	return originalURL, nil
}

// GetShortLink implements LinkRepository.
func (l *linkRepository) GetShortLink(id int) (model.Link, error) {
	log.Info("Executing Func GetShortLink with param: ", id)
	query := `SELECT short_code, original_url, created_at, updated_at FROM links WHERE id = $1`
	fmt.Println("Query:", query)
	var link model.Link
	err := l.db.QueryRow(query, id).Scan(&link.ShortCode, &link.OriginalURL, &link.CreatedAt, &link.UpdatedAt)
	if err != nil {
		log.Error("Failed to get short link:", err)
		return model.Link{}, err
	}
	return link, nil
}

// CreateShortLink implements LinkRepository.
func (l *linkRepository) CreateShortLink(linkRequest model.LinkCreateRequest) (model.Link, error) {
	log.Info("Executing Func CreateShortLink with param: ", linkRequest)

	// Extract short code from URL (e.g., from "localhost:3000/19191291" get "19191291")
	var shortCode string
	if linkRequest.ShortCode != "" {
		// Split by "/" and get the last part
		splitCode := strings.Split(linkRequest.ShortCode, "/")
		if len(splitCode) > 0 {
			shortCode = splitCode[len(splitCode)-1] // Get the last part
		}
	}

	fmt.Println("Extracted Short Code:", shortCode)

	query := `INSERT INTO links (short_code, original_url, created_at, updated_at) VALUES ($1, $2, NOW(), NOW()) RETURNING id, created_at, updated_at`
	var link model.Link
	err := l.db.QueryRow(query, shortCode, linkRequest.OriginalURL).Scan(&link.ID, &link.CreatedAt, &link.UpdatedAt)
	if err != nil {
		log.Error("Failed to create short link:", err)
		return model.Link{}, err
	}

	// Set the values we know
	link.ShortCode = shortCode
	link.OriginalURL = linkRequest.OriginalURL
	link.NewLink = fmt.Sprintf("localhost:3000/%s", shortCode)

	return link, nil
}

func NewLinkRepository(db *sql.DB) LinkRepository {
	return &linkRepository{db: db}
}
