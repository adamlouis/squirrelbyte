package secret

import (
	"context"
	"fmt"
)

type Manager interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
}

func NewManger() Manager {
	return &mgr{
		values: map[string]string{},
	}
}

type mgr struct {
	values map[string]string
}

func (m *mgr) Set(ctx context.Context, key, value string) error {
	m.values[key] = value
	return nil
}

func (m *mgr) Get(ctx context.Context, key string) (string, error) {
	v, ok := m.values[key]
	if !ok {
		return "", fmt.Errorf("key not found")
	}
	return v, nil
}
