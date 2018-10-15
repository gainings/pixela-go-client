package pixela

import (
	"testing"
)

func TestRegisterUser(t *testing.T) {
	c := NewClient()
	//Initialize test account
	c.DeleteUser("test-gainings", "testtest")
	c.DeleteUser("test-gainings", "testhogehoge")

	err := c.RegisterUser("test-gainings", "testtest", "no", "yes")
	if err == nil {
		t.Fatalf("want err, but got nil")
	}

	err = c.RegisterUser("test-gainings", "testtest", "yes", "no")
	if err == nil {
		t.Fatalf("want err, but got nil")
	}

	err = c.RegisterUser("test-gainings", "testtest", "yes", "yes")
	if err != nil {
		t.Fatalf("want nil, got %#v", err)
	}
	err = c.RegisterUser("test-gainings", "testtest", "yes", "yes")
	if err == nil {
		t.Fatalf("want err, but got nil")
	}

	err = c.UpdateToken("test-gainings", "testtest", "testhogehoge")
	if err != nil {
		t.Errorf("want nil, got %v", err)
	}

	err = c.DeleteUser("test-gainings", "testtest")
	if err == nil {
		t.Fatalf("want err, but got nil")
	}

	err = c.DeleteUser("test-gainings", "testhogehoge")
	if err != nil {
		t.Errorf("want nil, got %v", err)
	}
}
