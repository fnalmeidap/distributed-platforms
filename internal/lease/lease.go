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
	mu     sync.Mutex
	Leases map[string]Lease
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

// Remove Leases expirados
func (lm *LeaseManager) CleanupExpiredLeases() {
	for {
		time.Sleep(10 * time.Second)
		lm.mu.Lock()
		for id, lease := range lm.Leases {
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
