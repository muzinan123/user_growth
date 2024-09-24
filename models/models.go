package models

import (
	"time"
)

type TbCoinDetail struct {
	Id         int        `xorm:"not null pk autoincr UNSIGNED INT(11)"`
	Uid        int        `xorm:"not null default 0 comment('User ID') index(uid) INT(11)"`
	TaskId     int        `xorm:"not null default 0 comment('Task ID') index(uid) INT(11)"`
	Coin       int        `xorm:"not null default 0 comment('Points, positive for rewards, negative for penalties') INT(11)"`
	SysCreated *time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('Creation time') DATETIME"`
	SysUpdated *time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('Modification time') DATETIME"`
}

type TbCoinTask struct {
	Id         int        `xorm:"not null pk autoincr UNSIGNED INT(11)"`
	Task       string     `xorm:"not null default '' comment('Task name, must be unique') unique VARCHAR(255)"`
	Coin       int        `xorm:"not null default 0 comment('Points, positive for rewards, negative for penalties, 0 requires external call to pass value') INT(11)"`
	Limit      int        `xorm:"not null default 0 comment('Daily limit, default 0 for no limit') INT(11)"`
	Start      *time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('Effective start time') DATETIME"`
	SysCreated *time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('Creation time') DATETIME"`
	SysUpdated *time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('Modification time') DATETIME"`
	SysStatus  int        `xorm:"not null default 0 comment('Status, default 0 for active, 1 for deleted') INT(11)"`
}

type TbCoinUser struct {
	Id         int        `xorm:"not null pk autoincr UNSIGNED INT(11)"`
	Uid        int        `xorm:"not null default 0 comment('User ID') unique INT(11)"`
	Coins      int        `xorm:"not null default 0 comment('Total points') INT(11)"`
	SysCreated *time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('Creation time') DATETIME"`
	SysUpdated *time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('Modification time') DATETIME"`
}

type TbGradeInfo struct {
	Id          int        `xorm:"not null pk autoincr UNSIGNED INT(11)"`
	Title       string     `xorm:"not null comment('Grade name') VARCHAR(255)"`
	Description string     `xorm:"not null comment('Grade description') VARCHAR(3000)"`
	Score       int        `xorm:"not null default 0 comment('Maximum growth value for this grade') unique INT(11)"`
	Expired     int        `xorm:"not null default 0 comment('Validity period in days, default 0 for never expire') INT(11)"`
	SysCreated  *time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('Creation time') DATETIME"`
	SysUpdated  *time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('Modification time') DATETIME"`
}

type TbGradePrivilege struct {
	Id          int        `xorm:"not null pk autoincr UNSIGNED INT(11)"`
	GradeId     int        `xorm:"not null default 0 comment('Grade ID') index INT(11)"`
	Product     string     `xorm:"not null comment('Product') VARCHAR(255)"`
	Function    string     `xorm:"not null comment('Function') VARCHAR(255)"`
	Description string     `xorm:"not null comment('Description') VARCHAR(3000)"`
	Expired     int        `xorm:"not null default 0 comment('Validity period in days, default 0 for never expire') INT(11)"`
	SysCreated  *time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('Creation time') DATETIME"`
	SysUpdated  *time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('Modification time') DATETIME"`
	SysStatus   int        `xorm:"not null default 0 comment('Status, default 0 for active, 1 for deleted') INT(11)"`
}

type TbGradeUser struct {
	Id         int        `xorm:"not null pk autoincr UNSIGNED INT(11)"`
	Uid        int        `xorm:"not null default 0 comment('User ID') unique INT(11)"`
	GradeId    int        `xorm:"not null default 0 comment('Grade ID') INT(11)"`
	Expired    *time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('Expiration time') DATETIME"`
	Score      int        `xorm:"not null default 0 comment('Growth value') INT(11)"`
	SysCreated *time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('Creation time') DATETIME"`
	SysUpdated *time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('Modification time') DATETIME"`
}
