package lease

import (
	namingproxy "distributed-platforms/internal/services/naming/proxy"
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
	mu                 sync.Mutex
	Leases             map[string]Lease
	NamingServiceProxy namingproxy.NamingProxy
	LeaseType int
}

func NewLeaseManager() *LeaseManager {
	return &LeaseManager{
		Leases:             make(map[string]Lease),
		NamingServiceProxy: namingproxy.New(shared.LocalHost, shared.NamingPort),
	}
}

func (lm *LeaseManager) NewLease(id string, duration time.Duration) {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	expiration := time.Now().Add(duration)
	lm.Leases[id] = Lease{id: id, expiresAt: expiration}
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

// Remove Leases expirados
func (lm *LeaseManager) CleanupExpiredLeases() {
	for {
		time.Sleep(1 * time.Second)
		lm.mu.Lock()
		for id, lease := range lm.Leases {
			fmt.Println("Lease expiring in", time.Until(lease.expiresAt).Seconds())
			if time.Now().After(lease.expiresAt) {
				fmt.Printf("Lease %s expirou e foi removido\n", id)
				delete(lm.Leases, id)
				lm.NamingServiceProxy.Unbind(id)
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
