package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Constraint6Tables_20170907_215608 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Constraint6Tables_20170907_215608{}
	m.Created = "20170907_215608"

	migration.Register("Constraint6Tables_20170907_215608", m)
}

// Run the migrations
func (m *Constraint6Tables_20170907_215608) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Constraint6Tables_20170907_215608) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
