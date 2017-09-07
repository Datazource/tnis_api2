package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateCustomer_20170907_215519 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateCustomer_20170907_215519{}
	m.Created = "20170907_215519"

	migration.Register("CreateCustomer_20170907_215519", m)
}

// Run the migrations
func (m *CreateCustomer_20170907_215519) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *CreateCustomer_20170907_215519) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
