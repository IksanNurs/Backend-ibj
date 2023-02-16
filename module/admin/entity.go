package admin

import "database/sql"

type Admin struct {
	ID       sql.NullInt64
	Name     sql.NullString
	Email    sql.NullString
	Password sql.NullString
}
