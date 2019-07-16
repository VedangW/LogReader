package main

import (
	"github.com/jinzhu/gorm"
)

type (
	FilePath struct {
		gorm.Model
		FPath string `json:"file_path"`
	}

	DynJsonEntry struct {
		gorm.Model
		FileId     uint   `json:"file_id"`
		UserId     uint   `json:"user_id"`
		PhoneNo    string `json:"phone_no"`
		Password   string `json:"password"`
		Username   string `json:"username"`
		OrderValid string `json:"order_placed"`
		Extra      string `json:"extra"`
	}
)
