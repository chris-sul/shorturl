package link

import (
	"chris-sul/shorturl/database"
	"chris-sul/shorturl/internal/entity"
	"context"
	"database/sql"
)

type Repository interface {
	All(ctx context.Context) ([]entity.LinkEntity, error)
	Create(ctx context.Context, createParams CreateParams) (entity.LinkEntity, error)
	FindByURLCode(ctx context.Context, urlCode string) (string, error)
}

type repository struct {
	queries *database.Queries
	//logger log.Logger
}

func NewRepository(db *sql.DB) Repository {
	queries := database.New(db)
	return repository{queries}
}

func createLinkEntity(link database.Link) entity.LinkEntity {
	var linkEntity entity.LinkEntity

	linkEntity.ID = link.ID
	linkEntity.UrlCode = link.UrlCode
	linkEntity.Destination = link.Destination

	return linkEntity
}

func (r repository) All(ctx context.Context) ([]entity.LinkEntity, error) {
	links, err := r.queries.GetLinks(ctx)
	if err != nil {
		return nil, err
	}

	var linkEntities []entity.LinkEntity
	for _, link := range links {
		linkEntity := createLinkEntity(link)
		linkEntities = append(linkEntities, linkEntity)
	}

	return linkEntities, nil
}

type CreateParams struct {
	URLCode     string
	Destination string
}

func (r repository) Create(ctx context.Context, createParams CreateParams) (entity.LinkEntity, error) {
	createLinkParams := database.CreateLinkParams{
		UrlCode:     createParams.URLCode,
		Destination: createParams.Destination,
	}

	rawLink, err := r.queries.CreateLink(ctx, createLinkParams)
	if err != nil {
		return entity.LinkEntity{}, err
	}
	link := createLinkEntity(rawLink)
	return link, nil
}

func (r repository) FindByURLCode(ctx context.Context, urlCode string) (string, error) {
	destination, err := r.queries.FindLinkDestinationByCode(ctx, urlCode)
	if err != nil {
		return "", err
	}
	
	return destination, nil
}
