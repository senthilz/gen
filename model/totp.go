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


CREATE TABLE `totp` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(120) NOT NULL DEFAULT '' COMMENT 'Name of the app',
  `email` varchar(120) NOT NULL,
  `seed` varchar(32) NOT NULL,
  `deleted` enum('Y','N') DEFAULT 'N',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_email` (`email`),
  UNIQUE KEY `idx_seed` (`seed`),
  UNIQUE KEY `idx_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci

JSON Sample
-------------------------------------
{    "id": 31,    "name": "KGVspSMrhrTfIIKhrcNFGvcEZ",    "email": "rTFfNrHiTyaKrIqjCcjGyKbaZ",    "seed": "RPGfUqhOPHjtGGCJBphrpJcUD",    "deleted": "PyEbVZVKErgmHgACWNSnvksFB"}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// Totp struct is a row record of the totp table in the zen database
type Totp struct {
	//[ 0] id                                             ubigint              null: false  primary: true   isArray: false  auto: true   col: ubigint         len: -1      default: []
	ID uint64 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:ubigint;" form:"id" json:"id"`
	//[ 1] name                                           varchar(120)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 120     default: []
	Name string `gorm:"column:name;type:varchar;size:120;" form:"name" json:"name"` // Name of the app
	//[ 2] email                                          varchar(120)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 120     default: []
	Email string `gorm:"column:email;type:varchar;size:120;" form:"email" json:"email"`
	//[ 3] seed                                           varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	Seed string `gorm:"column:seed;type:varchar;size:32;" form:"seed" json:"seed"`
	//[ 4] deleted                                        char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [N]
	Deleted sql.NullString `gorm:"column:deleted;type:char;size:1;default:N;" form:"deleted" json:"deleted"`
}

var totpTableInfo = &TableInfo{
	Name: "totp",
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
			Comment:            `Name of the app`,
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
			Name:               "email",
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
			GoFieldName:        "Email",
			GoFieldType:        "string",
			JSONFieldName:      "email",
			ProtobufFieldName:  "email",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "seed",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "Seed",
			GoFieldType:        "string",
			JSONFieldName:      "seed",
			ProtobufFieldName:  "seed",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "deleted",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "Deleted",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "deleted",
			ProtobufFieldName:  "deleted",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *Totp) TableName() string {
	return "totp"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *Totp) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *Totp) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *Totp) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *Totp) TableInfo() *TableInfo {
	return totpTableInfo
}
