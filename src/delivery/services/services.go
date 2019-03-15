package services

type Service interface {
	Update() bool
	GetLastEvent() string
}

func ServiceFactory(name string, t string) Service {
	switch name {
	case "yanwen":
		return &Yanwen{token: t}
	}
	return nil
}
