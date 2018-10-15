package pixela

import (
	"testing"
)

func TestRegisterUser(t *testing.T) {
	c := NewClient("test-gainings", "testtest")

	err := c.RegisterUser("no", "yes")
	if err == nil {
		t.Fatalf("want err, but got nil")
	}

	err = c.RegisterUser("yes", "no")
	if err == nil {
		t.Fatalf("want err, but got nil")
	}

	err = c.RegisterUser("yes", "yes")
	if err != nil {
		t.Fatalf("want nil, got %#v", err)
	}
	err = c.RegisterUser("yes", "yes")
	if err == nil {
		t.Fatalf("want err, but got nil")
	}

	err = c.UpdateToken("testhogehoge")
	if err != nil {
		t.Errorf("want nil, got %v", err)
	}

	err = c.DeleteUser()
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}
}
