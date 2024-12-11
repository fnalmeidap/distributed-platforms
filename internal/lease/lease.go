package lease

import (
	"distributed-platforms/internal/distribution/requestor"
	"distributed-platforms/internal/shared"
	"fmt"
	"sync"
	"time"
)

type Lease struct {
	id        string
	expiresAt time.Time
}

type LeaseManager struct {
	mu            sync.Mutex
	Leases        map[string]Lease
	LeaseType     int
	LeaseOkayFlag bool
}

func NewLeaseManager() *LeaseManager {
	return &LeaseManager{
		Leases: make(map[string]Lease),
	}
}

func (lm *LeaseManager) NewLease(id string, duration time.Duration) {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	expiration := time.Now().Add(duration)
	lm.Leases[id] = Lease{id: id, expiresAt: expiration}
	lm.LeaseType = 2        //default value
	lm.LeaseOkayFlag = true //as it is being created, its treated as first lease
	fmt.Printf("Lease %s criado. Expira em: %v\n", id, expiration)
}

func (lm *LeaseManager) UpdateLease(id string, duration time.Duration) {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	expiration := time.Now().Add(duration)
	lm.Leases[id] = Lease{id: id, expiresAt: expiration}
	fmt.Printf("Lease %s renovado. Agora expira em: %v\n", id, expiration)
}

func (lm *LeaseManager) LeaseTypeSet(leaseType int) {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	lm.LeaseType = leaseType
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

// Remove Leases expirados
func (lm *LeaseManager) CleanupExpiredLeases(iorToServer shared.IOR) {
	for {
		time.Sleep(1 * time.Second)
		lm.mu.Lock()
		for id, lease := range lm.Leases {
			fmt.Println("Lease expiring in", time.Until(lease.expiresAt).Seconds())
			if lm.LeaseType == 2 {
				if time.Now().After(lease.expiresAt.Add(-6*time.Second)) && time.Now().Before(lease.expiresAt.Add(-5*time.Second)) {
					fmt.Println("Warning Client that resource will be deleted in 5 s ")
					clientSendMsg(iorToServer, id)
				}
			}
			if time.Now().After(lease.expiresAt) {
				fmt.Printf("Lease %s expirou e foi removido\n", id)
				lm.LeaseOkayFlag = false
				delete(lm.Leases, id)
			}
		}
		lm.mu.Unlock()
	}
}

func (lm *LeaseManager) LeaseExists(op string) bool {
	exists := false
	for key := range lm.Leases {
		if op == key {
			exists = true
		} else {
			exists = false
		}
	}
	return exists
}
