package pixela

import (
	"testing"
)

func TestPixel(t *testing.T) {
	c := NewClient()
	//Initialize test account
	c.DeleteUser("test-gainings", "testtest")
	c.DeleteUser("test-gainings", "testhogehoge")
	c.DeleteGraph("test-gainings", "testtest", "hoge1")
	c.DeleteGraph("test-gainings", "testtest", "hoge2")
	c.RegisterUser("test-gainings", "testtest", "yes", "yes")
	defer c.DeleteUser("test-gainings", "testtest")

	gi1 := GraphInfo{
		ID:       "hoge1",
		Name:     "fuga1",
		Unit:     "Kg",
		UnitType: "float",
		Color:    "shibafu",
	}
	err := c.CreateGraph("test-gainings", "testtest", gi1)
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}
	defer c.DeleteGraph("test-gainings", "testtest", "hoge1")

	err = c.DrawPixel("test-gainings", "testtest", "hoge1", "20181015", "5.5")
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}
}
