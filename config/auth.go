package config

type Auth interface {
	Authenticate()
}

type ServiceAccountAuth struct {
}

func (a *ServiceAccountAuth) Authenticate() {

}
