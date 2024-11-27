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

func (lcm LifecycleManager) Lease(lm *lease.LeaseManager, d time.Duration, c **calculator.Calculator) {
	// If lease for calculator exists, renew it and return
	if lcm.HasLease(lm) {
		fmt.Println("HAS LEASE")
		lcm.RenewLease(lm)
		return
	}

	// Creates new lease if no lease is currently active
	fmt.Println("DOES NOT HAS LEASE")
	lcm.CreateLease(lm, d, c)
}

func (lcm LifecycleManager) CreateLease(lm *lease.LeaseManager, d time.Duration, c **calculator.Calculator) {
	fmt.Println("Creating new lease!")
	lm.NewLease("calculator", d)
	*c = lcm.CreateObject()
}

func (lcm LifecycleManager) CreateObject() *calculator.Calculator {
	fmt.Println("Creating remote object")
	return &calculator.Calculator{}
}

func (lcm LifecycleManager) DestroyObject(lm *lease.Lease, c **calculator.Calculator) {
	fmt.Println("Destroying remote object")
	*c = nil
}

func (lcm LifecycleManager) RenewLease(lm *lease.LeaseManager) {
	lm.UpdateLease("calculator", time.Duration(shared.DefaultLeasingTimeSeconds*float64(time.Second)))
}

func (lcm LifecycleManager) HasLease(lm *lease.LeaseManager) bool {
	fmt.Println("Checking if lease exists")
	return lm.LeaseExists("calculator")
}
