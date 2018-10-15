# pixela-go-client

pixela-go-client is unofficial api client for [pixe.la](https://pixe.la/)

[GoDoc](https://godoc.org/github.com/gainings/pixela-go-client)

I need review. plz check code and make Issue and PullRequest!

## example

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gainings/pixela-go-client"
)

func main() {
	user := "test-pixela-client"
	pass := "testtest"
	c := pixela.NewClient(user, pass)

	//regiser as new user
	err := c.RegisterUser("yes", "yes")
	if err != nil {
		log.Fatal(err)
	}
	defer c.DeleteUser()

	//create new graph
	gi := pixela.GraphInfo{
		ID:       "hoge1",
		Name:     "hogehoge",
		Unit:     "Commit",
		UnitType: "int",
		Color:    "shibafu",
	}
	err = c.CreateGraph(gi)
	if err != nil {
		log.Fatal(err)
	}

	items, err := c.ListGraph()
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}

	//register pixel with quantity in graph
	today := time.Now().Format("20060102")
	fmt.Println(today)
	err = c.RegisterPixel("hoge1", today, "5")
	if err != nil {
		log.Fatal(err)
	}

	err = c.IncrementPixelQuantity("hoge1")
	if err != nil {
		log.Fatal(err)
	}

	err = c.DecrementPixelQuantity("hoge1")
	if err != nil {
		log.Fatal(err)
	}

	q, err := c.GetPixelQuantity(gi.ID, today)
	fmt.Println(q)
}
```


