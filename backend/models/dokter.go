package models

type Dokter struct {
    ID       int    `json:"id"`
    Nama     string `json:"nama"`
    PoliID   int    `json:"poli_id"`
    Tersedia bool   `json:"tersedia"`
}