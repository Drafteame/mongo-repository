package mgorepo

import "testing"

func TestNewSearchOrders(t *testing.T) {
	so := NewSearchOrders()

	if so == nil {
		t.Error("NewSearchOrders() should not return nil")
	}

	if len(so) != 0 {
		t.Errorf("NewSearchOrders() should return empty SearchOrders, got %d elements", len(so))
	}
}
