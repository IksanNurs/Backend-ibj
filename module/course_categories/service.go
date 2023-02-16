package course_categories

type Service interface {
	Create(input InputCourseCategories) (CourseCategories, error)
	Read() ([]CourseCategories,error)
	Update(input UpdateCourseCategories) (CourseCategories, error)
	Delete(ID int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input InputCourseCategories) (CourseCategories, error) {
	var coursecategories CourseCategories
	coursecategories.Name.String = input.Nama
	data, err := s.repository.Create(coursecategories)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (s *service) Read() ([]CourseCategories,error) {
	data, err := s.repository.Read()
	if err != nil {
		return data, err
	}
	return  data, nil
}

func (s *service) Update(input UpdateCourseCategories) (CourseCategories, error) {
	var coursecategories CourseCategories
	coursecategories.ID.Int64 = int64(input.ID)
	coursecategories.Name.String = input.Nama
	data, err := s.repository.Update(coursecategories)
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

