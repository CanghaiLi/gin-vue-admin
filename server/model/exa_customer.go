package model

import (
	"gorm.io/gorm"
	"time"
)

type ExaCustomer struct {
	gorm.Model
	CustomerName       string  `json:"customerName" form:"customerName" gorm:"comment:'客户名'"`
	CustomerPhoneData  string  `json:"customerPhoneData" form:"customerPhoneData" gorm:"comment:'客户手机号'"`
	SysUserID          uint    `json:"sysUserId" form:"sysUserId" gorm:"comment:'管理ID'"`
	SysUserAuthorityID string  `json:"sysUserAuthorityID" form:"sysUserAuthorityID" gorm:"comment:'管理角色ID'"`
	SysUser            SysUser `json:"sysUser" form:"sysUser" gorm:"comment:'管理详情'"`
}

func ExaCustomerData() []ExaCustomer {
	return []ExaCustomer{
		{Model: gorm.Model{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, CustomerName: "测试客户", CustomerPhoneData: "1761111111", SysUserID: 10, SysUserAuthorityID: "888"},
	}
}
