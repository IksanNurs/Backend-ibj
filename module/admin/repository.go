package admin

import (
	"database/sql"
)

type Repository interface {
	Save(admin Admin) (Admin, error)
	FindByEmail(email string) (Admin, error)
	FindByID(ID int) (Admin, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(admin Admin) (Admin, error) {

	sqlStmt := "INSERT INTO admin (name, email, password) VALUES (?, ?, ?);"
	_, err := r.db.Exec(sqlStmt, admin.Name.String, admin.Email.String, admin.Password.String)
	if err != nil {
		return admin, err
	}
	sqlStmt = "SELECT id, name, email FROM admin ORDER BY id DESC LIMIT 1 "
	row := r.db.QueryRow(sqlStmt)
	err = row.Scan(
		&admin.ID,
		&admin.Name,
		&admin.Email,
	)
	if err != nil {
		return admin, err
	}
	return admin, nil
}

func (r *repository) FindByEmail(email string) (Admin, error) {
	var admin Admin
	var sqlStmt string = "SELECT id, name, password, email FROM admin WHERE email=?"
	row := r.db.QueryRow(sqlStmt, email)
	err := row.Scan(
	&admin.ID,
	&admin.Email,
	&admin.Password,
	&admin.Name,
	)
	if err != nil {
		return admin, err
	}
	return admin, nil
}

func (r *repository) FindByID(ID int) (Admin, error) {
	var admin Admin
	var sqlStmt string = "SELECT id FROM admin WHERE id=?"
	row := r.db.QueryRow(sqlStmt, ID)
	err := row.Scan(
		&admin.ID,
	)
	if err != nil {
		return admin, err
	}
	return admin, nil
}
