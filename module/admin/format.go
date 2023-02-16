package admin

type FormatKey struct {
	Token string      `json:"token"`
	Admin interface{} `json:"admin"`
}

type SubFormatKeyValue struct {
	ID    int    `json:"id" sql:"unique"`
	Name  string `json:"name"`
	Email string `json:"email"`
}


func FormatAdmin(admin Admin, token string) FormatKey {
	
	formatter:=FormatKey{
		Token: token,
		Admin: SubFormatKeyValue{
			ID: int(admin.ID.Int64),
			Email: admin.Email.String,
			Name: admin.Name.String,
		},
	}

	return formatter
}