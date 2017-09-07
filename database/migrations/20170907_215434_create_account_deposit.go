package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateAccountDeposit_20170907_215434 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateAccountDeposit_20170907_215434{}
	m.Created = "20170907_215434"

	migration.Register("CreateAccountDeposit_20170907_215434", m)
}

// Run the migrations
func (m *CreateAccountDeposit_20170907_215434) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *CreateAccountDeposit_20170907_215434) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
