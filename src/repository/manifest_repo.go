package repository

import (
	"github.com/WilhelmWeber/iiif_annotation_server_v2/src/model"
	"gorm.io/gorm"
)

type ManifestRepository interface {
	GetById(id uint) (*model.Manifest, error)
	GetByUrl(url string) *model.Manifest
	GetAll() ([]*model.Manifest, error)
	Create(user *model.Manifest) error
	Update(user *model.Manifest) error
	Delete(user *model.Manifest) error
}

type ManifestService struct {
	db *gorm.DB
}

func NewManifestService(db *gorm.DB) *ManifestService {
	return &ManifestService{db}
}

func (m *ManifestService) GetById(id uint) (*model.Manifest, error) {
	var manifest model.Manifest
	if err := m.db.First(&manifest, id).Error; err != nil {
		return nil, err
	}
	return &manifest, nil
}

func (m *ManifestService) GetByUrl(url string) *model.Manifest {
	var manifest model.Manifest
	m.db.Where("url = ?", url).First(&manifest)
	return &manifest
}

func (m *ManifestService) GetAll() ([]*model.Manifest, error) {
	var manifests []*model.Manifest
	if err := m.db.Find(&manifests).Error; err != nil {
		return nil, err
	}
	return manifests, nil
}

func (m *ManifestService) Create(manifest *model.Manifest) error {
	return m.db.Create(manifest).Error
}

func (m *ManifestService) Update(manifest *model.Manifest) error {
	return m.db.Save(manifest).Error
}

func (m *ManifestService) Delete(manifest *model.Manifest) error {
	return m.db.Delete(manifest).Error
}
