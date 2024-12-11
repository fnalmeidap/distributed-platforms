package lifecyclemanager

import (
	"distributed-platforms/internal/app/calculator"
	"distributed-platforms/internal/distribution/requestor"
	"distributed-platforms/internal/lease"
	"distributed-platforms/internal/shared"
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
			// fmt.Println("\t Como lease pertence a tipo 0, lease renovado a cada chamada")
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
	// fmt.Println("Creating new lease!")
	lcm.Lm.NewLease("calculator", d)
	*c = lcm.CreateObject()
}

func (lcm LifecycleManager) CreateObject() *calculator.Calculator {
	// fmt.Println("Creating remote object")
	return &calculator.Calculator{}
}

func (lcm LifecycleManager) DestroyObject(c **calculator.Calculator) {
	// fmt.Println("Destroying remote object")
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
	// fmt.Println("Processing leases")
}

func clientSendMsg(iorToServer shared.IOR, id string) {
	params := make([]interface{}, 2)
	params[0] = id
	params[1] = 0

	req := shared.Request{Operation: "ReleaseWarn", Params: params}
	inv := shared.Invocation{Ior: iorToServer, Request: req}

	requestor := requestor.Requestor{}
	requestor.Invoke(inv)

}

func (lcm *LifecycleManager) CleanupExpiredLeases(iorToServer shared.IOR, c **calculator.Calculator) {
	for {
		time.Sleep(1 * time.Second)
		lcm.Lm.Mu.Lock()
		for id, lease := range lcm.Lm.Leases {
			// fmt.Println("Lease expiring in", time.Until(lease.ExpiresAt).Seconds())
			if lcm.Lm.LeaseType == 2 {
				if time.Now().After(lease.ExpiresAt.Add(-6*time.Second)) && time.Now().Before(lease.ExpiresAt.Add(-5*time.Second)) {
					// fmt.Println("Warning Client that resource will be deleted in 5 s ")
					clientSendMsg(iorToServer, id)
				}
			}
			if time.Now().After(lease.ExpiresAt) {
				// fmt.Printf("Lease %s expirou e foi removido\n", id)
				lcm.Lm.LeaseOkayFlag = false
				delete(lcm.Lm.Leases, id)
				lcm.Lm.NamingServiceProxy.Unbind(id)
				lcm.DestroyObject(c)
			}
		}
		lcm.Lm.Mu.Unlock()
	}
}
