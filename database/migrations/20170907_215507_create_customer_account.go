package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateCustomerAccount_20170907_215507 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateCustomerAccount_20170907_215507{}
	m.Created = "20170907_215507"

	migration.Register("CreateCustomerAccount_20170907_215507", m)
}

// Run the migrations
func (m *CreateCustomerAccount_20170907_215507) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *CreateCustomerAccount_20170907_215507) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
