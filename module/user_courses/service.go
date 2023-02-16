package user_courses

type Service interface {
	Create(input InputUserCourses) (UserCourses, error)
	Read() ([]UserCourses, error)
	Update(input UpdateUserCourses) (UserCourses, error)
	Delete(ID int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input InputUserCourses) (UserCourses, error) {
	var user_courses UserCourses
	user_courses.Course_id.Int64=int64(input.Course_id)
	user_courses.Users_id.Int64=int64(input.Users_id)
	data, err := s.repository.Create(user_courses)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *service) Read() ([]UserCourses, error) {
	data, err := s.repository.Read()
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *service) Update(input UpdateUserCourses) (UserCourses, error) {
	var user_courses UserCourses
	user_courses.ID.Int64 = int64(input.ID)
	user_courses.Course_id.Int64=int64(input.Course_id)
	user_courses.Users_id.Int64=int64(input.Users_id)
	data, err := s.repository.Update(user_courses)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *service) Delete(ID int) error {
	err := s.repository.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
