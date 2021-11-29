package factory

import "fmt"

type API interface {
	Say(name string) string
}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}

	return nil
}

type hiAPI struct{}
type helloAPI struct{}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
