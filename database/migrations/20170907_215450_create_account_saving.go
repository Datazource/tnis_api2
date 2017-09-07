package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateAccountSaving_20170907_215450 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateAccountSaving_20170907_215450{}
	m.Created = "20170907_215450"

	migration.Register("CreateAccountSaving_20170907_215450", m)
}

// Run the migrations
func (m *CreateAccountSaving_20170907_215450) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *CreateAccountSaving_20170907_215450) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
