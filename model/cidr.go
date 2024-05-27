package model

import (
	"database/sql"
	"time"

	"github.com/satori/go.uuid"

	"gorm.io/gorm"
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `cidr` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(120) NOT NULL,
  `ipv4_block` varchar(120) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`),
  UNIQUE KEY `idx_ipv4` (`ipv4_block`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci

JSON Sample
-------------------------------------
{    "id": 86,    "name": "IdPxGsbqxaufoJNJCpZxdnWtY",    "ipv_4_block": "pgfctSPtSjBvYtbFIcdsHvwYX"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// Cidr struct is a row record of the cidr table in the zen database
type Cidr struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" form:"id" json:"id"`
	//[ 1] name                                           varchar(120)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 120     default: []
	Name string `gorm:"column:name;type:varchar;size:120;" form:"name" json:"name"`
	//[ 2] ipv4_block                                     varchar(120)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 120     default: []
	Ipv4Block string `gorm:"column:ipv4_block;type:varchar;size:120;" form:"ipv_4_block" json:"ipv_4_block"`
}

var cidrTableInfo = &TableInfo{
	Name: "cidr",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "ubigint",
			DatabaseTypePretty: "ubigint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "ubigint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint64",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint64",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(120)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       120,
			GoFieldName:        "Name",
			GoFieldType:        "string",
			JSONFieldName:      "name",
			ProtobufFieldName:  "name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "ipv4_block",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(120)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       120,
			GoFieldName:        "Ipv4Block",
			GoFieldType:        "string",
			JSONFieldName:      "ipv_4_block",
			ProtobufFieldName:  "ipv_4_block",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *Cidr) TableName() string {
	return "cidr"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *Cidr) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *Cidr) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *Cidr) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *Cidr) TableInfo() *TableInfo {
	return cidrTableInfo
}
