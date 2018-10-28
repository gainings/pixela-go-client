package pixela

import (
	"testing"
	"time"
)

func TestPixel(t *testing.T) {
	c := NewClient("test-gainings", "testtest")
	c.RegisterUser("yes", "yes")
	defer c.DeleteUser()

	now := time.Now().UTC()
	const layout = "20060102"
	today := now.Format(layout)

	gi1 := GraphInfo{
		ID:       "hoge1",
		Name:     "fuga1",
		Unit:     "Kg",
		UnitType: "float",
		Color:    "shibafu",
	}
	err := c.CreateGraph(gi1)
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}
	defer c.DeleteGraph("hoge1")

	err = c.RegisterPixel("hoge1", today, "5.5")
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}
	quantity, err := c.GetPixelQuantity("hoge1", today)
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}
	if quantity != 5.5 {
		t.Fatalf("want 5.5, but got %v", quantity)
	}
	err = c.UpdatePixelQuantity("hoge1", today, "3.14")
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}
	quantity, err = c.GetPixelQuantity("hoge1", today)
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}
	if quantity != 3.14 {
		t.Fatalf("want 3.14, but got %v", quantity)
	}

	err = c.UpdatePixelQuantity("hoge1", today, "1.0")
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}
	err = c.IncrementPixelQuantity("hoge1")
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}
	quantity, err = c.GetPixelQuantity("hoge1", today)
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}
	if quantity != 1.01 {
		t.Fatalf("want 2, but got %v", quantity)
	}
	err = c.DecrementPixelQuantity("hoge1")
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}
	quantity, err = c.GetPixelQuantity("hoge1", today)
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}
	if quantity != 1.0 {
		t.Fatalf("want 1, but got %v", quantity)
	}

	err = c.DeletePixelQuantity("hoge1", today)
	if err != nil {
		t.Fatalf("want nil, but got %v", err)
	}
	quantity, err = c.GetPixelQuantity("hoge1", today)
	if err == nil {
		t.Fatalf("want err, but got nil")
	}
	if quantity != 0 {
		t.Fatalf("want 0, but got %v", quantity)
	}

}
