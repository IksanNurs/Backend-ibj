package user_courses

import (
	"database/sql"
)

type Repository interface {
	Create(user_courses UserCourses) (UserCourses, error)
	Read() ([]UserCourses, error)
	Update(user_courses UserCourses) (UserCourses, error)
	Delete(ID int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(user_courses UserCourses) (UserCourses, error) {

	sqlStmt := "INSERT INTO user_courses (users_id, course_id) VALUES (?, ?);"
	_, err := r.db.Exec(sqlStmt, user_courses.Users_id.Int64, user_courses.Course_id.Int64)
	if err != nil {
		return user_courses, err
	}
	sqlStmt = "SELECT uc.id, uc.users_id, uc.course_id, u.id, u.name, u.email, c.id, c.title, c.course_category_id, cc.id, cc.name FROM user_courses uc LEFT JOIN users u ON uc.users_id=u.id LEFT JOIN courses c ON uc.course_id=c.id LEFT JOIN course_categories cc ON c.course_category_id=cc.id ORDER BY uc.id DESC LIMIT 1"
	row := r.db.QueryRow(sqlStmt)
	err = row.Scan(
		&user_courses.ID,
		&user_courses.Users_id,
		&user_courses.Course_id,
		&user_courses.Users.ID,
		&user_courses.Users.Name,
		&user_courses.Users.Email,
		&user_courses.Courses.ID,
		&user_courses.Courses.Title,
		&user_courses.Courses.Courses_category_id,
		&user_courses.Course_category.ID,
		&user_courses.Course_category.Name,
	)
	if err != nil {
		return user_courses, err
	}
	return user_courses, nil
}

func (r *repository) Read() ([]UserCourses, error) {
	var sqlStmt string = "SELECT uc.id, uc.users_id, uc.course_id, u.id, u.name, u.email, c.id, c.title, c.course_category_id, cc.id, cc.name FROM user_courses uc LEFT JOIN users u ON uc.users_id=u.id LEFT JOIN courses c ON uc.course_id=c.id LEFT JOIN course_categories cc ON c.course_category_id=cc.id"
	rows, err := r.db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	var user_courses []UserCourses
	for rows.Next() {
		var usercourses UserCourses
		err = rows.Scan(
			&usercourses.ID,
			&usercourses.Users_id,
			&usercourses.Course_id,
			&usercourses.Users.ID,
			&usercourses.Users.Name,
			&usercourses.Users.Email,
			&usercourses.Courses.ID,
			&usercourses.Courses.Title,
			&usercourses.Courses.Courses_category_id,
			&usercourses.Course_category.ID,
			&usercourses.Course_category.Name,
		)

		if err != nil {
			return nil, err
		}
		user_courses = append(user_courses, usercourses)
	}
	return user_courses, nil
}

func (r *repository) Update(user_courses UserCourses) (UserCourses, error) {
	sqlStmt := "UPDATE user_courses SET users_id=?, course_id=? WHERE id=?"
	_, err := r.db.Exec(sqlStmt, user_courses.Users_id.Int64, user_courses.Course_id.Int64, user_courses.ID.Int64)
	if err != nil {
		return user_courses, err
	}
	sqlStmt = "SELECT uc.id, uc.users_id, uc.course_id, u.id, u.name, u.email, c.id, c.title, c.course_category_id, cc.id, cc.name FROM user_courses uc LEFT JOIN users u ON uc.users_id=u.id LEFT JOIN courses c ON uc.course_id=c.id LEFT JOIN course_categories cc ON c.course_category_id=cc.id WHERE uc.id=?"
	row := r.db.QueryRow(sqlStmt, user_courses.ID.Int64)
	err = row.Scan(
		&user_courses.ID,
		&user_courses.Users_id,
		&user_courses.Course_id,
		&user_courses.Users.ID,
		&user_courses.Users.Name,
		&user_courses.Users.Email,
		&user_courses.Courses.ID,
		&user_courses.Courses.Title,
		&user_courses.Courses.Courses_category_id,
		&user_courses.Course_category.ID,
		&user_courses.Course_category.Name,
	)
	if err != nil {
		return user_courses, err
	}
	return user_courses, nil
}

func (r *repository) Delete(ID int) error {
	sqlStmt := "DELETE FROM user_courses WHERE id=?"
	_, err := r.db.Exec(sqlStmt, ID)
	if err != nil {
		return err
	}
	return nil
}
