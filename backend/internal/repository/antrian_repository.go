package repository

import (
	"backend/internal/model"
	"backend/pkg/database"
)

type AntrianRepository struct{}
type PoliRepository struct{}

func NewAntrianRepository() *AntrianRepository {
	return &AntrianRepository{}
}

func (r *AntrianRepository) GetAll() ([]model.Antrian, error) {
	rows, err := database.DB.Query(`
		SELECT a.id, a.nama_pasien, a.no_antrian,
		       a.poli_id, a.dokter_id,
		       p.nama, d.nama
		FROM antrian a
		JOIN poli p ON a.poli_id = p.id
		JOIN dokter d ON a.dokter_id = d.id
		ORDER BY a.no_antrian ASC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.Antrian

	for rows.Next() {
		var a model.Antrian
		err := rows.Scan(
			&a.ID,
			&a.NamaPasien,
			&a.NoAntrian,
			&a.PoliID,
			&a.DokterID,
			&a.NamaPoli,
			&a.NamaDokter,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, a)
	}

	return result, nil
}

func (r *AntrianRepository) Create(a *model.Antrian) error {
	_, err := database.DB.Exec(
		"INSERT INTO antrian (nama_pasien, no_antrian, poli_id, dokter_id) VALUES (?, ?, ?, ?)",
		a.NamaPasien,
		a.NoAntrian,
		a.PoliID,
		a.DokterID,
	)
	return err
}

func (r *AntrianRepository) Delete(id int) error {
	_, err := database.DB.Exec("DELETE FROM antrian WHERE id = ?", id)
	return err
}

func (r *AntrianRepository) GetLastQueue(poliID int) (int, error) {
	var last int
	err := database.DB.QueryRow(
		"SELECT IFNULL(MAX(no_antrian),0) FROM antrian WHERE poli_id = ?",
		poliID,
	).Scan(&last)
	return last, err
}

func (r *AntrianRepository) GetDoctorByPoli(poliID int) (int, error) {
	var id int
	err := database.DB.QueryRow(
		"SELECT id FROM dokter WHERE poli_id = ? LIMIT 1",
		poliID,
	).Scan(&id)
	return id, err
}


func NewPoliRepository() *PoliRepository {
	return &PoliRepository{}
}

func (r *PoliRepository) GetAll() ([]model.Poli, error) {
	rows, err := database.DB.Query("SELECT id, nama FROM poli")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.Poli

	for rows.Next() {
		var p model.Poli
		rows.Scan(&p.ID, &p.Nama)
		result = append(result, p)
	}

	return result, nil
}