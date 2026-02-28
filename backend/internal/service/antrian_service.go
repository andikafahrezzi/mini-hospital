package service

import (
	"backend/internal/model"
	"backend/internal/repository"
	"errors"
)

type AntrianService struct {
	repo *repository.AntrianRepository
}

type PoliService struct {
	repo *repository.PoliRepository
}

type DokterService struct {
	repo *repository.DokterRepository
}

func NewAntrianService(r *repository.AntrianRepository) *AntrianService {
	return &AntrianService{repo: r}
}

func (s *AntrianService) GetAll() ([]model.Antrian, error) {
	return s.repo.GetAll()
}

func (s *AntrianService) Create(req model.CreateAntrianRequest) error {

    // VALIDASI: dokter harus sesuai poli
    valid, err := s.repo.IsDoctorBelongToPoli(req.DokterID, req.PoliID)
    if err != nil {
        return err
    }

    if !valid {
        return errors.New("dokter tidak sesuai dengan poli")
    }

    antrian := model.Antrian{
        NamaPasien: req.NamaPasien,
        PoliID:     req.PoliID,
        DokterID:   req.DokterID, // ⬅ pakai dari frontend
    }

    return s.repo.Create(&antrian)
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

func NewDokterService(r *repository.DokterRepository) *DokterService {
	return &DokterService{repo: r}
}

func (s *DokterService) GetByPoli(poliID int) ([]model.Dokter, error) {
	return s.repo.GetByPoli(poliID)
}

