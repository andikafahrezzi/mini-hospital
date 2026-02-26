package models

type Antrian struct {
    ID         int    `json:"id"`
    NamaPasien string `json:"nama_pasien"`
    NoAntrian  int    `json:"no_antrian"`
    PoliID     int    `json:"poli_id"`
    DokterID   int    `json:"dokter_id"`
}