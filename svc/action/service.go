package action

// Service handles the processing and routing of actions.
type Service struct {
	sources map[SourceName]Source
}

// NewService creates a new instance of the action service.
func NewService() *Service {
	return &Service{
		sources: make(map[SourceName]Source),
	}
}

// ProcessAction routes the action to its source and processes it.
func (s *Service) ProcessAction(a *Action) error {
	if s.sources[a.Source] == nil {
		return ErrInvalidSourceName
	}
	return s.sources[a.Source].ProcessAction(a)
}

// SetSource sets the source instance associated with a source name.
func (s *Service) SetSource(name SourceName, source Source) {
	s.sources[name] = source
}
