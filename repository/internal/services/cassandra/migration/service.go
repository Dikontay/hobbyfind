package migration

type service struct {
	configs Configs
}

func NewService(configs Configs) Service {
	return &service{
		configs: configs,
	}
}

func (s service) Init() {
	//TODO implement me
	panic("implement me")
}

func (s service) Stop() {
	//TODO implement me
	panic("implement me")
}
