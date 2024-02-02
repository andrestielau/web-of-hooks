package auth

type Manager struct {
	keys map[string]HMAC
}

func NewManager() *Manager {
	return &Manager{
		keys: make(map[string]HMAC),
	}
}
