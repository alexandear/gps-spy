package specnomery

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) Check(number string) bool {
	return true
}
