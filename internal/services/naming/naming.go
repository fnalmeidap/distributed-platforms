package naming

import (
	"distributed-platforms/internal/shared"
	"fmt"
)

type NamingService struct {
	Repository map[string]shared.IOR
}

func (n *NamingService) Bind(s string, i shared.IOR) bool {
	fmt.Println("Binding:", s)
	r := false

	// check if repository is already created
	if len(n.Repository) == 0 {
		n.Repository = make(map[string]shared.IOR)
	}
	// check if the service is already registered
	_, ok := n.Repository[s]
	if ok {
		r = false // service already registered
	} else { // service not registered
		n.Repository[s] = shared.IOR{TypeName: i.TypeName, Host: i.Host, Port: i.Port}
		r = true
	}

	fmt.Println("Lookups: ", len(n.Repository))
	return r
}

func (n NamingService) Find(s string) shared.IOR {
	fmt.Println("Finding:", s)
	return n.Repository[s]
}

func (n NamingService) List() map[string]shared.IOR {
	fmt.Println("Listing lookups: ", len(n.Repository))
	return n.Repository
}

func (n NamingService) Unbind(s string) bool {
	fmt.Println("Unbinding:", s)
	r := false
	// check if the resource is already registered
	_, ok := n.Repository[s]
	if ok {

		delete(n.Repository, s)
		fmt.Println("Lookups: ", len(n.Repository))
		r = true
	} else {
		r = false
	}

	return r
}
