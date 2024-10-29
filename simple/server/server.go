package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Lease struct {
	id        string
	expiresAt time.Time
}

type LeaseManager struct {
	mu     sync.Mutex
	leases map[string]Lease
}

// Cria um novo lease com um tempo de vida específico
func (lm *LeaseManager) NewLease(id string, duration time.Duration) {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	expiration := time.Now().Add(duration)
	lm.leases[id] = Lease{id: id, expiresAt: expiration}
	fmt.Printf("Lease %s criado. Expira em: %v\n", id, expiration)
}

func (lm *LeaseManager) UpdateLease(id string, duration time.Duration) {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	expiration := time.Now().Add(duration)
	lm.leases[id] = Lease{id: id, expiresAt: expiration}
	fmt.Printf("Lease %s renovado. Agora expira em: %v\n", id, expiration)
}

// Remove leases expirados
func (lm *LeaseManager) CleanupExpiredLeases() {
	for {
		time.Sleep(10 * time.Second)
		lm.mu.Lock()
		for id, lease := range lm.leases {
			if time.Now().After(lease.expiresAt) {
				fmt.Printf("Lease %s expirou e foi removido\n", id)
				delete(lm.leases, id)
			}
		}
		lm.mu.Unlock()
	}
}

func main() {
	fmt.Printf("Lease 12\n")
	leaseManager := &LeaseManager{leases: make(map[string]Lease)}

	fmt.Printf("Lease 12\n")
	// Endpoint para criar um lease
	http.HandleFunc("/lease", func(w http.ResponseWriter, r *http.Request) {
		leaseID := r.URL.Query().Get("id")

		duration, _ := time.ParseDuration(r.URL.Query().Get("duration"))

		exists := false
		for key := range leaseManager.leases {
			if leaseID == key {
				exists = true
			} else {
				exists = false
			}
		}

		if exists {
			leaseManager.UpdateLease(leaseID, duration)
		} else {
			leaseManager.NewLease(leaseID, duration)
		}

		w.Write([]byte("Lease criado!\n"))
	})
	fmt.Printf("Lease 12\n")
	// Limpeza de leases expirados
	go leaseManager.CleanupExpiredLeases()
	fmt.Printf("Lease 12\n")
	// Inicializa o servidor
	http.ListenAndServe(":8080", nil)
	fmt.Printf("Lease 12\n")
}
