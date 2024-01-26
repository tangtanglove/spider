package model

import (
	"time"

	appmodel "github.com/quarkcloudio/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcloudio/quark-go/v2/pkg/dal/db"
	"gorm.io/gorm"
)

// 模型
type Data struct {
	Id        int            `json:"id" gorm:"autoIncrement"`
	Realname  string         `json:"realname" gorm:"size:200;not null"`
	Email     string         `json:"email" gorm:"size:200;not null"`
	Company   string         `json:"company" gorm:"size:200;not null"`
	Status    int            `json:"status" gorm:"size:4;not null;default:1"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Seeder
func (m *Data) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if (&appmodel.Menu{}).IsExist(18) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 18, Name: "数据管理", GuardName: "admin", Icon: "icon-book", Type: 1, Pid: 0, Sort: 0, Path: "/data", Show: 1, IsEngine: 0, IsLink: 0, Status: 1},
		{Id: 19, Name: "数据列表", GuardName: "admin", Icon: "", Type: 2, Pid: 18, Sort: 0, Path: "/api/admin/data/index", Show: 1, IsEngine: 1, IsLink: 0, Status: 1},
	}
	db.Client.Create(&menuSeeders)
}
