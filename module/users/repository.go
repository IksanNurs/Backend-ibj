package users

import (
	"database/sql"
)

type Repository interface {
	Create(users Users) (Users, error)
	Read() ([]Users, error)
	Update(users Users) (Users, error)
	Delete(ID int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(users Users) (Users, error) {

	sqlStmt := "INSERT INTO users (name, email, password) VALUES (?, ?, ?);"
	_, err := r.db.Exec(sqlStmt, users.Name.String, users.Email.String, users.Password.String)
	if err != nil {
		return users, err
	}
	sqlStmt = "SELECT id, name, email FROM users ORDER BY id DESC LIMIT 1 "
	row := r.db.QueryRow(sqlStmt)
	err = row.Scan(
		&users.ID,
		&users.Name,
		&users.Email,
	)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *repository) Read() ([]Users, error) {
	var sqlStmt string = "SELECT id, name, email FROM users"
	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	var users []Users
	for rows.Next() {
		var user Users
		err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
		)
		
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *repository) Update(users Users) (Users, error) {
	sqlStmt := "UPDATE users SET name=?, email=?, password=? WHERE id=?"
	_, err := r.db.Exec(sqlStmt, users.Name.String, users.Email.String, users.Password.String, users.ID.Int64)
	if err != nil {
		return users, err
	}
	sqlStmt = "SELECT id, name, email FROM users WHERE id=? "
	row := r.db.QueryRow(sqlStmt, users.ID.Int64)
	err = row.Scan(
		&users.ID,
		&users.Name,
		&users.Email,
	)
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *repository) Delete(ID int) error {
	sqlStmt := "DELETE FROM users WHERE id=?"
	_, err := r.db.Exec(sqlStmt, ID)
	if err != nil {
		return  err
	}
	return  nil
}


