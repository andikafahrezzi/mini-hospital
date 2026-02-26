package controllers

import (
    "backend/database"
    "net/http"

    "github.com/gin-gonic/gin"
)

type InputAntrian struct {
    NamaPasien string `json:"nama_pasien"`
    PoliID     int    `json:"poli_id"`
}

// POST /antrian
func AddAntrian(c *gin.Context) {
    var input InputAntrian
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var lastNo int
    database.DB.QueryRow("SELECT IFNULL(MAX(no_antrian),0) FROM antrian").Scan(&lastNo)
    noAntrian := lastNo + 1

    var dokterID int
    err := database.DB.QueryRow("SELECT id FROM dokter WHERE poli_id=? AND tersedia=1 LIMIT 1", input.PoliID).Scan(&dokterID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Tidak ada dokter tersedia"})
        return
    }

    _, err = database.DB.Exec(
        "INSERT INTO antrian (nama_pasien, no_antrian, poli_id, dokter_id) VALUES (?, ?, ?, ?)",
        input.NamaPasien, noAntrian, input.PoliID, dokterID,
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":     "Antrian berhasil ditambahkan",
        "no_antrian":  noAntrian,
        "nama_pasien": input.NamaPasien,
        "poli_id":     input.PoliID,
        "dokter_id":   dokterID,
    })
}

// GET /antrian
func GetAntrian(c *gin.Context) {
    query := `
        SELECT a.id, a.nama_pasien, a.no_antrian, 
               p.nama AS nama_poli, d.nama AS nama_dokter
        FROM antrian a
        JOIN poli p ON a.poli_id = p.id
        JOIN dokter d ON a.dokter_id = d.id
        ORDER BY a.no_antrian ASC
    `
    rows, err := database.DB.Query(query)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var list []map[string]interface{}
    for rows.Next() {
        var id, no int
        var nama_pasien, nama_poli, nama_dokter string
        rows.Scan(&id, &nama_pasien, &no, &nama_poli, &nama_dokter)

        a := map[string]interface{}{
            "id":          id,
            "nama_pasien": nama_pasien,
            "no_antrian":  no,
            "nama_poli":   nama_poli,
            "nama_dokter": nama_dokter,
        }
        list = append(list, a)
    }

    c.JSON(200, gin.H{"antrian": list})
}

// DELETE /antrian/:id
func DeleteAntrian(c *gin.Context) {
    id := c.Param("id")
    _, err := database.DB.Exec("DELETE FROM antrian WHERE id=?", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Antrian dihapus"})
}