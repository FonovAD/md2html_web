package apiserver

type APIserver struct {
	config *Config
}

func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
	}
}

func (s *APIserver) Start() error {
	return nil
}
