package service

import (
	"api-service/internal/model"
	"api-service/internal/repository"
	"crypto/rand"
	"fmt"
	"math/big"
)

type LinkService interface {
	CreateShortLink(linkRequest model.LinkCreateRequest) (model.LinkCreateResponse, error)
	GetShortLink(id int) (model.Link, error)
	GetOriginalURL(shortCode string) (string, error)
	GetAllLink() ([]model.Link, error)
}

type linkService struct {
	repo repository.LinkRepository
}

// GetAllLink implements LinkService.
func (l *linkService) GetAllLink() ([]model.Link, error) {
	fmt.Println("Executing Func GetAllLink")
	links, err := l.repo.GetAllLink()
	if err != nil {
		return nil, err
	}
	return links, nil
}

// GetOriginalURL implements LinkService.
func (l *linkService) GetOriginalURL(shortCode string) (string, error) {
	originalURL, err := l.repo.GetOriginalURL(shortCode)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}

// CreateShortLink implements LinkService.
func (l *linkService) CreateShortLink(linkRequest model.LinkCreateRequest) (model.LinkCreateResponse, error) {
	codeRand, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return model.LinkCreateResponse{}, err
	}
	baseCode := "localhost:3000/v1/" + codeRand.String()
	linkRequest.ShortCode = baseCode

	createdLink, err := l.repo.CreateShortLink(linkRequest)
	if err != nil {
		return model.LinkCreateResponse{}, err
	}

	return model.LinkCreateResponse{Data: createdLink}, nil
}

// GetShortLink implements LinkService.
func (l *linkService) GetShortLink(id int) (model.Link, error) {
	link, err := l.repo.GetShortLink(id)
	if err != nil {
		return model.Link{}, err
	}
	return link, nil

}

func NewLinkService(repo repository.LinkRepository) LinkService {
	return &linkService{repo: repo}
}
