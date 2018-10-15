package main

import (
	"fmt"

	"github.com/gainings/pixela"
)

func main() {
	c := pixela.NewClient()

	newGi1 := pixela.GraphInfo{
		ID:       "hagehage1",
		Name:     "newFuga1",
		Unit:     "Kg",
		UnitType: "float",
		Color:    "shibafu",
	}
	err := c.CreateGraph("test-gainings", "testtest", newGi1)
	if err != nil {
	}
	gs, err := c.ListGraph("test-gainings", "testtest")
	if err != nil {
	}
	fmt.Println(gs)
	err = c.CreateGraph("test-gainings", "testtest", newGi1)
	if err != nil {
	}
	gs, err = c.ListGraph("test-gainings", "testtest")
	if err != nil {
	}
	fmt.Println(gs)
	newGi1.Name = "newFugaoigagoagiag"
	c.UpdateGraph("test-gainings", "testtest", newGi1)
	gs, err = c.ListGraph("test-gainings", "testtest")
	if err != nil {
	}
	fmt.Println(gs)
}
