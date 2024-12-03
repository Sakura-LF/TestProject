package LoadBalance

import "sync"

type LoadBalance interface {
	Add(string)
	Get() string
}

type RoundRobinLoadBalance struct {
	current int
	servers []string

	lock sync.Mutex
}

func NewRoundRobinLoadBalance() *RoundRobinLoadBalance {
	return &RoundRobinLoadBalance{
		current: 0,
		servers: make([]string, 0, 10),
	}
}

func (r *RoundRobinLoadBalance) Add(server string) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.servers = append(r.servers, server)
}

func (r *RoundRobinLoadBalance) Get() string {
	r.lock.Lock()
	defer r.lock.Unlock()
	if r.current >= len(r.servers) {
		r.current = 0
	}
	server := r.servers[r.current]
	r.current++
	return server
}
