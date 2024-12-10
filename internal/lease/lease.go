package lease

import (
	"fmt"
	"sync"
	"time"
)

type Lease struct {
	id        string
	expiresAt time.Time
}

type LeaseManager struct {
	mu        sync.Mutex
	Leases    map[string]Lease
	LeaseType int
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
			if lm.LeaseType == 2 {
				if time.Now().After(lease.expiresAt.Add(-6*time.Second)) && time.Now().Before(lease.expiresAt.Add(-5*time.Second)) {
					fmt.Println("Warning Client that resource will be deleted in 5 s ")
					// TODO: HOW???!!
				}
			}

			if time.Now().After(lease.expiresAt) {
				fmt.Printf("Lease %s expirou e foi removido\n", id)
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
