package repository

import (
	"backend/internal/model"
	"database/sql"
)

func GetUserByUsername(db *sql.DB, username string) (*model.User, string, error) {

	query := `
		SELECT u.id, u.username, u.password_hash, u.role_id, u.dokter_id, r.name
		FROM users u
		JOIN roles r ON u.role_id = r.id
		WHERE u.username = ?
	`

	row := db.QueryRow(query, username)

	var user model.User
	var roleName string

	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.PasswordHash,
		&user.RoleID,
		&user.DokterID,
		&roleName,
	)

	if err != nil {
		return nil, "", err
	}

	return &user, roleName, nil
}