package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateUsers_20170907_215058 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateUsers_20170907_215058{}
	m.Created = "20170907_215058"

	migration.Register("CreateUsers_20170907_215058", m)
}

// Run the migrations
func (m *CreateUsers_20170907_215058) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *CreateUsers_20170907_215058) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
