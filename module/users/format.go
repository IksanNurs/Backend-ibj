package users

type FormatKey struct {
	Users interface{} `json:"users"`
}

type SubFormatKeyValue struct {
	ID    int    `json:"id" sql:"unique"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FormatUsers(users Users) FormatKey {

	formatter := FormatKey{
		Users: SubFormatKeyValue{
			ID:    int(users.ID.Int64),
			Name:  users.Name.String,
			Email: users.Email.String,
		},
	}

	return formatter
}

func FormatsUsers(users []Users) FormatKey {
	var formatters []SubFormatKeyValue

	for _, user := range users {
		formatter := SubFormatKeyValue{
			ID:    int(user.ID.Int64),
			Name:  user.Name.String,
			Email: user.Email.String,
		}
		formatters = append(formatters, formatter)
	}

	form := FormatKey{
		Users: formatters,
	}
	return form
}
