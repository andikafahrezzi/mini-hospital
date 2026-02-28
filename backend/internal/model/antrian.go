package model

type Antrian struct {
	ID         int    `json:"id"`
	NamaPasien string `json:"nama_pasien"`
	NoAntrian  int    `json:"no_antrian"`
	PoliID     int    `json:"poli_id"`
	Status	  string  `json:"status"`
	DokterID   int    `json:"dokter_id"`
	NamaPoli   string `json:"nama_poli"`
	NamaDokter string `json:"nama_dokter"`
}

type CreateAntrianRequest struct {
	NamaPasien string `json:"nama_pasien"`
	PoliID     int    `json:"poli_id"`
	DokterID   int    `json:"dokter_id"`
}

type Poli struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

type Dokter struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}