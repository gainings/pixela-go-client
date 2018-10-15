package pixela

import (
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	c := NewClient()
	err := c.Register("test-user"+time.Now().String(), "testtest", "no", "yes")
	if err == nil {
		t.Fatalf("want err, but got nil")
	}

	err = c.Register("test-user"+time.Now().String(), "testtest", "yes", "no")
	if err == nil {
		t.Fatalf("want err, but got nil")
	}

	err = c.Register("test-user"+time.Now().String(), "testtest", "yes", "yes")
	if err != nil {
		t.Fatalf("want nil, got %#v", err)
	}
}
