package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CreateAccountSavingHistory_20170907_215414 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateAccountSavingHistory_20170907_215414{}
	m.Created = "20170907_215414"

	migration.Register("CreateAccountSavingHistory_20170907_215414", m)
}

// Run the migrations
func (m *CreateAccountSavingHistory_20170907_215414) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *CreateAccountSavingHistory_20170907_215414) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
