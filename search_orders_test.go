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

func TestSearchOrders_Add(t *testing.T) {
	so := NewSearchOrders()

	so = so.Add("name", OrderAsc)

	if len(so) != 1 {
		t.Errorf("Add() should add one element, got %d elements", len(so))
	}

	if so[0].Name != "name" {
		t.Errorf("Add() should add element with name 'name', got '%s'", so[0].Name)
	}

	if so[0].Order != OrderAsc {
		t.Errorf("Add() should add element with order %d, got %d", OrderAsc, so[0].Order)
	}
}
