package pixela

import (
	"testing"
)

func TestCreateGraph(t *testing.T) {
	c := NewClient()
	//Initialize test account
	c.DeleteUser("test-gainings", "testtest")
	c.DeleteUser("test-gainings", "testhogehoge")
	c.RegisterUser("test-gainings", "testtest", "yes", "yes")

	gi := GraphInfo{
		ID:   "hoge",
		Name: "fuga",
		Unit: "Kg",
		//Invalid unit type
		UnitType: "string",
		//invalid color type
		Color: "skyblue",
	}
	err := gi.Validate()
	if err == nil {
		t.Fatalf("want invalid unit type error")
	}
	gi.UnitType = "float"
	err = gi.Validate()
	if err == nil {
		t.Fatalf("want invalid color error")
	}
	gi.Color = "shibafu"
	err = gi.Validate()
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}

	err = c.CreateGraph("test-gainings", "testtest", gi)
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}
}
