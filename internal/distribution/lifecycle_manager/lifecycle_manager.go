package lifecyclemanager

import (
	"distributed-platforms/internal/app/calculator"
	"distributed-platforms/internal/lease"
	"distributed-platforms/internal/shared"
	"fmt"
	"time"
)

type LifecycleManager struct {
	Lm *lease.LeaseManager
}

func NewLifecycleManager() *LifecycleManager {
	return &LifecycleManager{
		Lm: lease.NewLeaseManager(),
	}
}

func (lcm LifecycleManager) Lease(d time.Duration, c **calculator.Calculator) {
	// If lease for calculator exists, renew it and return
	if lcm.HasLease() {
		if lcm.Lm.LeaseType == 0 {
			fmt.Println("\t Como lease pertence a tipo 0, lease renovado a cada chamada")
			lcm.RenewLease()
		}
		return
	}

	// Creates new lease if no lease is currently active
	if (lcm.Lm.LeaseType == 0) && (!lcm.Lm.LeaseOkayFlag) {
		lcm.CreateLease(d, c)
		lcm.MaybeAddRemoteObjectBinding()
	}
	// TODO: is this condition okay?

}

func (lcm LifecycleManager) MaybeAddRemoteObjectBinding() {
	rs := lcm.Lm.NamingServiceProxy.Find("calculator")
	if rs.LeaseName != "calculator" {
		lcm.Lm.NamingServiceProxy.Bind("calculator", shared.IOR{Host: shared.LocalHost, Port: shared.ClientServerPort})
	}
}

func (lcm LifecycleManager) CreateLease(d time.Duration, c **calculator.Calculator) {
	fmt.Println("Creating new lease!")
	lcm.Lm.NewLease("calculator", d)
	*c = lcm.CreateObject()
}

func (lcm LifecycleManager) CreateObject() *calculator.Calculator {
	fmt.Println("Creating remote object")
	return &calculator.Calculator{}
}

func (lcm LifecycleManager) DestroyObject(c **calculator.Calculator) {
	fmt.Println("Destroying remote object")
	*c = nil
}

func (lcm LifecycleManager) RenewLease() {
	lcm.Lm.UpdateLease("calculator", time.Duration(shared.DefaultLeasingTimeSeconds*float64(time.Second)))
}

func (lcm LifecycleManager) RenewLease_v2(T int) {
	lcm.Lm.UpdateLease("calculator", time.Duration(float64(T)*float64(time.Second)))
}

func (lcm LifecycleManager) LeaseTypeSet(leaseType int) {
	lcm.Lm.LeaseTypeSet(leaseType)
}

func (lcm LifecycleManager) HasLease() bool {
	// fmt.Println("Checking if lease exists")
	return lcm.Lm.LeaseExists("calculator")
}

func (lcm LifecycleManager) ProcessLeases() {
	fmt.Println("Processing leases")
}
