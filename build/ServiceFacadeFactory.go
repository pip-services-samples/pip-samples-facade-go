package build

import (
	bbuild "github.com/pip-services-samples/pip-services-beacons-go/build"
	cbuild "github.com/pip-services3-go/pip-services3-components-go/build"
)

type ServiceFacadeFactory struct {
	cbuild.CompositeFactory
}

func NewServiceFacadeFactory() *ServiceFacadeFactory {
	c := &ServiceFacadeFactory{
		CompositeFactory: *cbuild.NewCompositeFactory(),
	}

	// c.Add(NewAccountsServiceFactory())
	// c.Add(NewSessionsServiceFactory())
	// c.Add(NewPasswordsServiceFactory())
	// c.Add(NewRolesServiceFactory())
	c.Add(bbuild.NewBeaconsServiceFactory())
}