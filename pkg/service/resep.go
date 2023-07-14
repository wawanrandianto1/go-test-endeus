package service

import (
	"endeus/wawan/pkg/endeus"
	"endeus/wawan/pkg/model"
	"endeus/wawan/pkg/repository"
)

type resepService struct {
	*repository.Repo
}

func NewResepService(repo *repository.Repo) endeus.ResepService {
	return &resepService{
		repo,
	}
}

func (s *resepService) GetAll(str string) ([]endeus.ResepCustomResponse, error) {
	var customResp []endeus.ResepCustomResponse
	data, err := s.Repo.Resep.FindAll(str)
	if err != nil {
		return nil, err
	}

	for _, v := range data {
		tmp := endeus.ResepCustomResponse{
			ID:         v.ID,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
			CategoryID: v.CategoryID,
			Judul:      v.Judul,
			Deskripsi:  v.Deskripsi,
			VideoUrl:   v.VideoUrl,
			Publish:    v.Publish,
		}
		customResp = append(customResp, tmp)
	}
	return customResp, nil
}

func (s *resepService) GetAllByCategoryID(catID uint) ([]endeus.ResepCustomResponse, error) {
	var customResp []endeus.ResepCustomResponse
	data, err := s.Repo.Resep.FindAllByCategoryID(catID)
	if err != nil {
		return nil, err
	}

	for _, v := range data {
		tmp := endeus.ResepCustomResponse{
			ID:         v.ID,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
			CategoryID: v.CategoryID,
			Judul:      v.Judul,
			Deskripsi:  v.Deskripsi,
			VideoUrl:   v.VideoUrl,
			Publish:    v.Publish,
		}
		customResp = append(customResp, tmp)
	}
	return customResp, nil
}

func (s *resepService) GetByID(id uint) (model.Resep, error) {
	return s.Repo.Resep.FindByID(id)
}

func (s *resepService) CreateNew(param endeus.ResepParam) error {
	return s.Repo.Resep.Create(param)
}

func (s *resepService) Unpublish(id uint) error {
	return s.Repo.Resep.Unpublish(id)
}

func (s *resepService) Delete(id uint) error {
	return s.Repo.Resep.Delete(id)
}

func (s *resepService) Update(id uint, param endeus.ResepParam) error {
	return s.Repo.Resep.Update(id, param)
}
