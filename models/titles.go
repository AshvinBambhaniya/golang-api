package models

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/rs/xid"
)

// This boilerplate we are storing password in plan format!

// UserTable represent table name
const TitleTable = "titles"

// Title model

type Title struct {
	ID                string   `json:"id" db:"id"`
	Title             string   `json:"title" db:"title" validate:"required"`
	Type              string   `json:"type" db:"type" validate:"required"`
	Description       string   `json:"description" db:"description"`
	ReleaseYear       int      `json:"release_year" db:"release_year" validate:"required"`
	AgeCertification  string   `json:"age_certification" db:"age_certification" validate:"required"`
	Runtime           int      `json:"runtime" db:"runtime" validate:"required"`
	Genres            []string `json:"genres" db:"genres" validate:"required"`
	ProductionCountry []string `json:"production_countries" db:"production_countries" validate:"required"`
	Seasons           int      `json:"seasons" db:"seasons"`
	IMDBID            string   `json:"imdb_id" db:"imdb_id"`
	IMDBScore         float64  `json:"imdb_score" db:"imdb_score"`
	IMDBVotes         float64  `json:"imdb_votes" db:"imdb_votes"`
	TMDBPopularity    float64  `json:"tmdb_popularity" db:"tmdb_popularity"`
	TMDBScore         float64  `json:"tmdb_score" db:"tmdb_score"`
	CreatedAt         string   `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt         string   `json:"updated_at,omitempty" db:"updated_at,omitempty"`
}

// UserModel implements user related database operations
type TitleModel struct {
	db *goqu.Database
}

// InitUserModel Init model
func InitTitleModel(goqu *goqu.Database) (TitleModel, error) {
	return TitleModel{
		db: goqu,
	}, nil
}

// GetUsers list all users
func (model *TitleModel) GetTitles() ([]Title, error) {
	var titles []Title
	if err := model.db.From(UserTable).ScanStructs(&titles); err != nil {
		return nil, err
	}
	return titles, nil
}

// GetUser get user by id
func (model *TitleModel) GetTitleByID(id string) (Title, error) {
	title := Title{}
	found, err := model.db.From(TitleTable).Where(goqu.Ex{"id": id}).ScanStruct(&title)

	if err != nil {
		return title, err
	}

	if !found {
		return title, sql.ErrNoRows
	}

	return title, err
}

func (model *TitleModel) InsertTitle(title Title) (Title, error) {
	_, err := model.db.Insert(TitleTable).Rows(title).Executor().Exec()

	if err != nil {
		return title, err
	}
	return title, err
}

// UpdateTitle updates an existing title
func (model *TitleModel) UpdateTitle(title Title) error {
	_, err := model.db.Update(TitleTable).Set(title).Where(goqu.Ex{"id": title.ID}).Executor().Exec()
	return err
}

// DeleteTitle deletes a title by its ID
func (model *TitleModel) DeleteTitle(id string) error {
	_, err := model.db.Delete(TitleTable).Where(goqu.Ex{"id": id}).Executor().Exec()
	return err
}
