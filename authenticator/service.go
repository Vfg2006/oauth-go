package authenticator

import "fmt"

type Service interface {
	GetToken(text string) (string, error)
}

type service struct {
}

func New() Service {
	return service{}
}

func (s service) GetToken(text string) (string, error) {
	return fmt.Sprintf("TESTE - GetToken - %s", text), nil
}
