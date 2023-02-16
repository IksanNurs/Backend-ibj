package course_categories

import (
	"database/sql"
)

type Repository interface {
	Create(course_categories CourseCategories) (CourseCategories, error)
	Update(course_categories CourseCategories) (CourseCategories, error)
	Delete(ID int) error
	Read() ([]CourseCategories, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(course_categories CourseCategories) (CourseCategories, error) {

	sqlStmt := "INSERT INTO course_categories (name) VALUES (?);"
	_, err := r.db.Exec(sqlStmt, course_categories.Name.String)
	if err != nil {
		return course_categories, err
	}
	sqlStmt = "SELECT id, name FROM course_categories ORDER BY id DESC LIMIT 1 "
	row := r.db.QueryRow(sqlStmt)
	err = row.Scan(
		&course_categories.ID,
		&course_categories.Name,
	)
	if err != nil {
		return course_categories, err
	}
	return course_categories, nil
}

func (r *repository) Read() ([]CourseCategories, error) {
	var sqlStmt string = "SELECT id, name FROM course_categories"
	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	var course_categories []CourseCategories
	for rows.Next() {
		var coursecategories CourseCategories
		err = rows.Scan(
			&coursecategories.ID,
			&coursecategories.Name,

		)
		
		if err != nil {
			return nil, err
		}
		course_categories = append(course_categories, coursecategories)
	}
	return course_categories, nil
}

func (r *repository) Update(course_categories CourseCategories) (CourseCategories, error) {
	sqlStmt := "UPDATE course_categories SET name=? WHERE id=?"
	_, err := r.db.Exec(sqlStmt, course_categories.Name.String, course_categories.ID.Int64)
	if err != nil {
		return course_categories, err
	}
	sqlStmt = "SELECT id, name FROM course_categories WHERE id=? "
	row := r.db.QueryRow(sqlStmt, course_categories.ID.Int64)
	err = row.Scan(
		&course_categories.ID,
		&course_categories.Name,
	)
	if err != nil {
		return course_categories, err
	}
	return course_categories, nil
}

func (r *repository) Delete(ID int) error {
	sqlStmt := "DELETE FROM course_categories WHERE id=?"
	_, err := r.db.Exec(sqlStmt, ID)
	if err != nil {
		return  err
	}
	return  nil
}


