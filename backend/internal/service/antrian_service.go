package service

import (
	"backend/internal/model"
	"backend/internal/repository"
)

type AntrianService struct {
	repo *repository.AntrianRepository
}

type PoliService struct {
	repo *repository.PoliRepository
}

func NewAntrianService(r *repository.AntrianRepository) *AntrianService {
	return &AntrianService{repo: r}
}

func (s *AntrianService) GetAll() ([]model.Antrian, error) {
	return s.repo.GetAll()
}

func (s *AntrianService) Create(nama string, poliID int) (*model.Antrian, error) {
	last, err := s.repo.GetLastQueue(poliID)
	if err != nil {
		return nil, err
	}

	dokterID, err := s.repo.GetDoctorByPoli(poliID)
	if err != nil {
		return nil, err
	}

	antrian := &model.Antrian{
		NamaPasien: nama,
		NoAntrian:  last + 1,
		PoliID:     poliID,
		DokterID:   dokterID,
	}

	err = s.repo.Create(antrian)
	if err != nil {
		return nil, err
	}

	return antrian, nil
}

func (s *AntrianService) Delete(id int) error {
	return s.repo.Delete(id)
}



func NewPoliService(r *repository.PoliRepository) *PoliService {
	return &PoliService{repo: r}
}

func (s *PoliService) GetAll() ([]model.Poli, error) {
	return s.repo.GetAll()
}