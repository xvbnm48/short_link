package service

import (
	"api-service/internal/model"
	"api-service/internal/repository"
	"crypto/rand"
	"math/big"
)

type LinkService interface {
	CreateShortLink(linkRequest model.LinkCreateRequest) (model.LinkCreateResponse, error)
	GetShortLink(id int) (model.Link, error)
}

type linkService struct {
	repo repository.LinkRepository
}

// CreateShortLink implements LinkService.
func (l *linkService) CreateShortLink(linkRequest model.LinkCreateRequest) (model.LinkCreateResponse, error) {
	codeRand, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		return model.LinkCreateResponse{}, err
	}
	baseCode := "localhost:3000/" + codeRand.String()
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
