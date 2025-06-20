package main

import (
	"github.com/go-mysql-org/go-mysql/canal"
	"fmt"
)

type MyEventHandler struct {
	canal.DummyEventHandler
}

func (h *MyEventHandler) OnRow(e *canal.RowsEvent) error {
	fmt.Printf("%s %v\n", e.Action, e.Rows)
	return nil
}

func (h *MyEventHandler) String() string {
	return "MyEventHandler"
}

func main() {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = "127.0.0.1:6603"
	cfg.User = "root"
	cfg.Password = "12345678"
	// We only care table canal_test in test db
	cfg.Dump.TableDB = "test"
	cfg.Dump.Tables = []string{"canal_test"}

	c, err := canal.NewCanal(cfg)
	if err != nil {
		panic(err)
	}

	// Register a handler to handle RowsEvent
	c.SetEventHandler(&MyEventHandler{})

	// Start canal
	c.Run()
}
