package GormTest

import (
	"gorm.io/gorm"
	"testing"
)

type User struct {
	gorm.Model
	Name      string
	OrderList []Order `json:"-"`
}

type Order struct {
	gorm.Model
	OrderName string
	UserID    uint
	//User      User `gorm:"foreignKey:UserID" json:"-"`
}

func TestOrderTest(t *testing.T) {
	DBConnection()

	//var users User
	//tx := DB.Where("id = ?", 1).Preload("OrderList", "id = ?", 1).Find(&users)
	//if tx.Error != nil {
	//	t.Error(tx.Error)
	//}

	//fmt.Println(users)
}
