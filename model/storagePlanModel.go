package model

import (
	"gin_work/wrap/database"
	"time"
)

type storagePlanModel struct {
	Id             int           `gorm:"primaryKey"`
	Name           string        `gorm:"column:name;"`
	GoodsId        int           `gorm:"column:goods_id;"`
	CreateTime     time.Duration `gorm:"column:create_time;"`
	PlanType       int8          `gorm:"column:type;"`
	Count          int           `gorm:"column:count;"`
	OrderNo        string        `gorm:"column:order_no;"`
	storeId        int           `gorm:"column:shop_id;"`
	OrderId        int           `gorm:"column:order_id;"`
	OrderSn        string        `gorm:"column:order_sn;"`
	Linkman        string        `gorm:"column:linkman;"`
	Phone          string        `gorm:"column:phone;"`
	DriverName     string        `gorm:"column:driver_name;"`
	DriverPhone    string        `gorm:"column:driver_phone;"`
	PlateNumber    string        `gorm:"column:plate_number;"`
	DrivingLicense string        `gorm:"column:driving_license;"`
	DriverSign     string        `gorm:"column:driver_sign;"`
}

var StoragePlanModel *storagePlanModel

func (*storagePlanModel) TableName() string {
	return "jmk_storage_plan"
}

func (*storagePlanModel) FindJoinById(id int) (*storagePlanModel, error) {
	result := database.DB.Find(&StoragePlanModel, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return StoragePlanModel, nil
}
