package db

import (
	"Server/internal/types"
	"fmt"
	"time"
)

func CreateAllTables(db *Database) {
	db.db.AutoMigrate(&types.Admin{})
}

func (db *Database) InsertAdmin() {
	admin := types.Admin{
		FullName:       "Johyunmin",
		LoginId:        "jhm91",
		Password:       "123123",
		RegisteredTime: uint64(time.Now().UnixMilli()),
	}
	db.db.Create(&admin)
}

func (db *Database) SelectAdminList() {
	adminSelected := []types.Admin{}
	db.db.Find(&adminSelected)
	fmt.Println(adminSelected)
}
