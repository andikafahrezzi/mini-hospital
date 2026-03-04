package repository

import (
	"backend/internal/model"
	"backend/pkg/database"
	"fmt"
	"time"
	"database/sql"
)

type AntrianRepository struct{}
type PoliRepository struct{}
type DokterRepository struct{}


func NewAntrianRepository() *AntrianRepository {
	return &AntrianRepository{}
}

func (r *AntrianRepository) GetAll() ([]model.Antrian, error) {
rows, err := database.DB.Query(`
	SELECT a.id, a.nama_pasien, a.no_antrian,
	       a.poli_id, a.dokter_id, a.status,
	       p.nama, d.nama
	FROM antrian a
	JOIN poli p ON a.poli_id = p.id
	JOIN dokter d ON a.dokter_id = d.id
	WHERE a.tanggal = CURDATE()
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
					&a.Status,
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
	tx, err := database.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	today := time.Now().Format("2006-01-02")


	var prefix string
	err = tx.QueryRow(
		"SELECT prefix FROM poli WHERE id = ?",
		a.PoliID,
	).Scan(&prefix)
	if err != nil {
		return err
	}


	var lastNumber int
	err = tx.QueryRow(`
		SELECT COALESCE(MAX(CAST(SUBSTRING(no_antrian, 2) AS UNSIGNED)), 0)
		FROM antrian
		WHERE poli_id = ? AND tanggal = ?
	`, a.PoliID, today).Scan(&lastNumber)
	if err != nil {
		return err
	}

	nextNumber := lastNumber + 1
	formatted := fmt.Sprintf("%s%03d", prefix, nextNumber)

	_, err = tx.Exec(`
		INSERT INTO antrian
		(nama_pasien, no_antrian, poli_id, dokter_id, status, tanggal)
		VALUES (?, ?, ?, ?, ?, ?)
	`,
		a.NamaPasien,
		formatted,
		a.PoliID,
		a.DokterID,
		"MENUNGGU",
		today,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
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

func NewDokterRepository() *DokterRepository {
	return &DokterRepository{}
}

func (r *DokterRepository) GetByPoli(poliID int) ([]model.Dokter, error) {
	rows, err := database.DB.Query(
		"SELECT id, nama FROM dokter WHERE poli_id = ?",
		poliID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.Dokter

	for rows.Next() {
		var d model.Dokter
		rows.Scan(&d.ID, &d.Nama)
		result = append(result, d)
	}

	return result, nil
}
func (r *AntrianRepository) IsDoctorBelongToPoli(dokterID int, poliID int) (bool, error) {

    var count int

    err := database.DB.QueryRow(
        "SELECT COUNT(*) FROM dokter WHERE id = ? AND poli_id = ?",
        dokterID,
        poliID,
    ).Scan(&count)

    if err != nil {
        return false, err
    }

    return count > 0, nil
}

func (r *AntrianRepository) GetByID(id int) (*model.Antrian, error) {
	var a model.Antrian

	err := database.DB.QueryRow(`
		SELECT id, nama_pasien, poli_id, dokter_id, no_antrian, status
		FROM antrian
		WHERE id = ?
	`, id).Scan(
		&a.ID,
		&a.NamaPasien,
		&a.PoliID,
		&a.DokterID,
		&a.NoAntrian,
		&a.Status,
	)

	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (r *AntrianRepository) UpdateStatus(id int, status string) error {
	_, err := database.DB.Exec(`
		UPDATE antrian
		SET status = ?
		WHERE id = ?
	`, status, id)

	return err
}

func (r *AntrianRepository) GetLastNoAntrian(poliID int) (string, error) {
	var last sql.NullString

	err := database.DB.QueryRow(`
		SELECT MAX(no_antrian)
		FROM antrian
		WHERE poli_id = ? AND tanggal = CURDATE()
	`, poliID).Scan(&last)

	if err != nil {
		return "", err
	}

	if !last.Valid {
		return "", nil
	}

	return last.String, nil
}