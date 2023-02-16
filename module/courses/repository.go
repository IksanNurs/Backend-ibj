package courses

import (
	"database/sql"
)

type Repository interface {
	Create(courses Courses) (Courses, error)
	Read() ([]Courses, error)
	Update(courses Courses) (Courses, error)
	Delete(ID int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(courses Courses) (Courses, error) {

	sqlStmt := "INSERT INTO courses (title, course_category_id) VALUES (?, ?);"
	_, err := r.db.Exec(sqlStmt, courses.Title.String, courses.Courses_category_id.Int64)
	if err != nil {
		return courses, err
	}
	sqlStmt = "SELECT co.id, co.title, co.course_category_id, cc.id, cc.name FROM courses co LEFT JOIN course_categories cc ON co.course_category_id = cc.id ORDER BY co.id DESC LIMIT 1 "
	row := r.db.QueryRow(sqlStmt)
	err = row.Scan(
		&courses.ID,
		&courses.Title, 
		&courses.Courses_category_id,
		&courses.Course_categories.ID,
		&courses.Course_categories.Name,
	)
	if err != nil {
		return courses, err
	}
	return courses, nil
}

func (r *repository) Read() ([]Courses, error) {
	var sqlStmt string = "SELECT co.id, co.title, co.course_category_id, cc.id, cc.name FROM courses co LEFT JOIN course_categories cc ON co.course_category_id = cc.id"
	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	var courses []Courses
	for rows.Next() {
		var course Courses
		err = rows.Scan(
			&course.ID,
			&course.Title,
			&course.Courses_category_id,
			&course.Course_categories.ID,
			&course.Course_categories.Name,
		)
		
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (r *repository) Update(courses Courses) (Courses, error) {
	sqlStmt := "UPDATE courses SET title=?, course_category_id=? WHERE id=?"
	_, err := r.db.Exec(sqlStmt,courses.Title.String, courses.Courses_category_id.Int64 , courses.ID.Int64)
	if err != nil {
		return courses, err
	}
	sqlStmt = "SELECT co.id, co.title, co.course_category_id, cc.id, cc.name FROM courses co LEFT JOIN course_categories cc ON co.course_category_id = cc.id WHERE co.id=?"
	row := r.db.QueryRow(sqlStmt, courses.ID.Int64)
	err = row.Scan(
		&courses.ID,
		&courses.Title, 
		&courses.Courses_category_id,
		&courses.Course_categories.ID,
		&courses.Course_categories.Name,
	)
	if err != nil {
		return courses, err
	}
	return courses, nil
}

func (r *repository) Delete(ID int) error {
	sqlStmt := "DELETE FROM courses WHERE id=?"
	_, err := r.db.Exec(sqlStmt, ID)
	if err != nil {
		return  err
	}
	return  nil
}

