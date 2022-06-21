package repositories

import (
	"github.com/lsflk/gig-sdk/libraries"
	"github.com/lsflk/gig-sdk/models"
	"gopkg.in/mgo.v2"
	"log"
)

type iNormalizedNameRepository interface {
	AddNormalizedName(m models.NormalizedName) (normalizedName models.NormalizedName, err error)
	GetNormalizedNames(searchString string, limit int) ([]models.NormalizedName, error)
	GetNormalizedName(id string) (models.NormalizedName, error)
	GetNormalizedNameBy(attribute string, value string) (models.NormalizedName, error)
}

type NormalizedNameRepository struct {
	iNormalizedNameRepository
}

// AddNormalizedName insert a new NormalizedName into database and returns
// last inserted normalized_name on success.
func (n NormalizedNameRepository) AddNormalizedName(m models.NormalizedName) (normalizedName models.NormalizedName, err error) {
	return repositoryHandler.normalizedNameRepository.AddNormalizedName(m.NewNormalizedName())
}

// GetNormalizedNames Get all NormalizedNames from database and returns
// list of NormalizedName on success
func (n NormalizedNameRepository) GetNormalizedNames(searchString string, limit int) ([]models.NormalizedName, error) {
	return repositoryHandler.normalizedNameRepository.GetNormalizedNames(libraries.ProcessNameString(searchString), limit)
}

// GetNormalizedName Get a NormalizedName from database and returns
// a NormalizedName on success
func (n NormalizedNameRepository) GetNormalizedName(id string) (models.NormalizedName, error) {
	return repositoryHandler.normalizedNameRepository.GetNormalizedName(id)
}

/*
GetNormalizedNameBy Get an Entity from database and returns
a models.Entity on success
*/
func (n NormalizedNameRepository) GetNormalizedNameBy(attribute string, value string) (models.NormalizedName, error) {
	return repositoryHandler.normalizedNameRepository.GetNormalizedNameBy(attribute, value)
}

func (n NormalizedNameRepository) AddTitleToNormalizationDatabase(entityTitle string, normalizedName string) {
	// perform save in async
	go func(entityTitle string, normalizedName string) {
		_, err := repositoryHandler.normalizedNameRepository.AddNormalizedName(
			*new(models.NormalizedName).SetSearchText(entityTitle).SetNormalizedText(normalizedName),
		)
		if err != nil && !mgo.IsDup(err) {
			log.Println("error while saving normalized title:", err)
		}
	}(entityTitle, normalizedName)
}
