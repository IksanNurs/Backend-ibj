package courses

type Service interface {
	Create(input InputCourses) (Courses, error)
	Read() ([]Courses,error)
	Update(input UpdateCourses) (Courses, error)
	Delete(ID int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input InputCourses) (Courses, error) {
	var courses Courses
	courses.Title.String=input.Title
	courses.Courses_category_id.Int64=int64(input.Course_category_id)
	data, err := s.repository.Create(courses)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *service) Read() ([]Courses,error) {
	data, err := s.repository.Read()
	if err != nil {
		return data, err
	}
	return  data, nil
}

func (s *service) Update(input UpdateCourses) (Courses, error) {
	var courses Courses
	courses.ID.Int64 = int64(input.ID)
	courses.Title.String=input.Title
	courses.Courses_category_id.Int64=int64(input.Course_category_id)
	data, err := s.repository.Update(courses)
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
	return  nil
}

