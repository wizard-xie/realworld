package data

import (
	"testing"

	"github.com/wizard-xie/realworld/internal/conf"
)

func TestNewGorm(t *testing.T) {
	NewGorm(&conf.Data{
		Database: &conf.Data_Database{
			Source: "root:159zxcvbnm@@tcp(47.97.122.186:3306)/realworld?charset=utf8mb4&parseTime=True&loc=Local",
		},
	})
}
