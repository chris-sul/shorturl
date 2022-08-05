package link

import (
	"context"
	"net/url"

	"crypto/md5"
	"encoding/hex"

	"chris-sul/shorturl/internal/entity"
)

type Service interface {
	All(ctx context.Context) ([]entity.LinkEntity, error)
	Create(ctx context.Context, createDto CreateDto) (entity.LinkEntity, error)
	
	GetLinkDestination(ctx context.Context, urlCode string) (string, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return service{repo}
}

func (r service) All(ctx context.Context) ([]entity.LinkEntity, error) {
	links, err := r.repo.All(ctx)
	if err != nil {
		return nil, err
	}

	return links, nil
}

type CreateDto struct {
	Destination string `json:"destination" binding:"required"`
}

func (r service) Create(ctx context.Context, createDto CreateDto) (entity.LinkEntity, error) {
	// validate that destination is valid url
	_, err := url.ParseRequestURI(createDto.Destination)
	if err != nil {
		return entity.LinkEntity{}, err
	}

	// Create URL Code
	hash := md5.New()
	hash.Write([]byte(createDto.Destination))
	hashString := hex.EncodeToString(hash.Sum(nil))

	createParams := CreateParams {
		URLCode: hashString[:6],
		Destination: createDto.Destination,
	}

	link, err := r.repo.Create(ctx, createParams)
	if err != nil {
		return entity.LinkEntity{}, err
	}

	return link, nil
}

func (r service) GetLinkDestination(ctx context.Context, urlCode string) (string, error) {
	urlString, err := r.repo.FindByURLCode(ctx, urlCode)	
	if err != nil {
		return "", nil
	}

	return urlString, nil
}
