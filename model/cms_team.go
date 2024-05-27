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


CREATE TABLE `cms_team` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT,
  `full_name` char(100) NOT NULL DEFAULT '',
  `short_name` char(12) NOT NULL DEFAULT '',
  `abbreviation` char(5) NOT NULL DEFAULT '',
  `filename` char(20) NOT NULL DEFAULT '',
  `class_id` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `country_id` smallint(5) unsigned NOT NULL DEFAULT 0,
  `continent_id` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `address` char(200) NOT NULL DEFAULT '',
  `team_founded` char(20) NOT NULL DEFAULT '',
  `achievements` char(255) NOT NULL DEFAULT '',
  `player_numbers` char(100) NOT NULL DEFAULT '',
  `club_numbers` char(100) NOT NULL DEFAULT '',
  `is_country` enum('Y','N') DEFAULT 'N',
  `country_icc_status` char(50) NOT NULL DEFAULT '',
  `country_icc_elected` smallint(5) unsigned NOT NULL DEFAULT 0,
  `country_icc_zone` char(30) NOT NULL DEFAULT '',
  `notes` char(255) NOT NULL,
  `logo_object_id` int(10) unsigned NOT NULL DEFAULT 0,
  `logo_alt_id` char(10) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `full_name` (`full_name`(15)),
  KEY `short_name` (`short_name`(5)),
  KEY `abbreviation` (`abbreviation`),
  KEY `filename` (`filename`(5)),
  KEY `class_id` (`class_id`),
  KEY `is_country` (`is_country`),
  KEY `country_id` (`country_id`),
  KEY `continent_id` (`continent_id`),
  KEY `logo_object_id` (`logo_object_id`),
  KEY `logo_alt_id` (`logo_alt_id`),
  FULLTEXT KEY `keywords` (`full_name`,`abbreviation`,`filename`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci

JSON Sample
-------------------------------------
{    "id": 87,    "full_name": "ImUQUkIiRHfkbXTxUHijQCaVX",    "short_name": "NkUeelHXafaTmwryIDwPNvIdV",    "abbreviation": "gJCRLKRoRCjNLfogGjVFlyYWV",    "filename": "hGrpXLeOEIsRlQeKLgLgBBdPn",    "class_id": 20,    "country_id": 3,    "continent_id": 22,    "address": "WHamxOEomKOBFIwRHrlSaZUow",    "team_founded": "nvFhfnTyDoHAeUFaOrkCGICOn",    "achievements": "NLBqbplmaluURSiWnfguUuMdI",    "player_numbers": "GBnjCkYgUNKjWcoXZLRkOQSlL",    "club_numbers": "mBdOuuPlBHHhqFCKbVrnGOjKm",    "is_country": "wSFUnZhhPjMRQjMWxMBnorgOO",    "country_icc_status": "lvbeZpWBVBObmnOwsZJxvxLUQ",    "country_icc_elected": 61,    "country_icc_zone": "IQPhIJUJjjNebejfaHUranKER",    "notes": "vbJDQFSAJiqbnpUNiTtTsDagK",    "logo_object_id": 15,    "logo_alt_id": "lUUZZeNJNGfAKdgjBNbrZOGSh"}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 5] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[15] column is set for unsigned
[18] column is set for unsigned



*/

// CmsTeam struct is a row record of the cms_team table in the zen database
type CmsTeam struct {
	//[ 0] id                                             usmallint            null: false  primary: true   isArray: false  auto: true   col: usmallint       len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:usmallint;" form:"id" json:"id"`
	//[ 1] full_name                                      char(100)            null: false  primary: false  isArray: false  auto: false  col: char            len: 100     default: []
	FullName string `gorm:"column:full_name;type:char;size:100;" form:"full_name" json:"full_name"`
	//[ 2] short_name                                     char(12)             null: false  primary: false  isArray: false  auto: false  col: char            len: 12      default: []
	ShortName string `gorm:"column:short_name;type:char;size:12;" form:"short_name" json:"short_name"`
	//[ 3] abbreviation                                   char(5)              null: false  primary: false  isArray: false  auto: false  col: char            len: 5       default: []
	Abbreviation string `gorm:"column:abbreviation;type:char;size:5;" form:"abbreviation" json:"abbreviation"`
	//[ 4] filename                                       char(20)             null: false  primary: false  isArray: false  auto: false  col: char            len: 20      default: []
	Filename string `gorm:"column:filename;type:char;size:20;" form:"filename" json:"filename"`
	//[ 5] class_id                                       utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	ClassID uint32 `gorm:"column:class_id;type:utinyint;default:0;" form:"class_id" json:"class_id"`
	//[ 6] country_id                                     usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	CountryID uint32 `gorm:"column:country_id;type:usmallint;default:0;" form:"country_id" json:"country_id"`
	//[ 7] continent_id                                   utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	ContinentID uint32 `gorm:"column:continent_id;type:utinyint;default:0;" form:"continent_id" json:"continent_id"`
	//[ 8] address                                        char(200)            null: false  primary: false  isArray: false  auto: false  col: char            len: 200     default: []
	Address string `gorm:"column:address;type:char;size:200;" form:"address" json:"address"`
	//[ 9] team_founded                                   char(20)             null: false  primary: false  isArray: false  auto: false  col: char            len: 20      default: []
	TeamFounded string `gorm:"column:team_founded;type:char;size:20;" form:"team_founded" json:"team_founded"`
	//[10] achievements                                   char(255)            null: false  primary: false  isArray: false  auto: false  col: char            len: 255     default: []
	Achievements string `gorm:"column:achievements;type:char;size:255;" form:"achievements" json:"achievements"`
	//[11] player_numbers                                 char(100)            null: false  primary: false  isArray: false  auto: false  col: char            len: 100     default: []
	PlayerNumbers string `gorm:"column:player_numbers;type:char;size:100;" form:"player_numbers" json:"player_numbers"`
	//[12] club_numbers                                   char(100)            null: false  primary: false  isArray: false  auto: false  col: char            len: 100     default: []
	ClubNumbers string `gorm:"column:club_numbers;type:char;size:100;" form:"club_numbers" json:"club_numbers"`
	//[13] is_country                                     char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [N]
	IsCountry sql.NullString `gorm:"column:is_country;type:char;size:1;default:N;" form:"is_country" json:"is_country"`
	//[14] country_icc_status                             char(50)             null: false  primary: false  isArray: false  auto: false  col: char            len: 50      default: []
	CountryIccStatus string `gorm:"column:country_icc_status;type:char;size:50;" form:"country_icc_status" json:"country_icc_status"`
	//[15] country_icc_elected                            usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	CountryIccElected uint32 `gorm:"column:country_icc_elected;type:usmallint;default:0;" form:"country_icc_elected" json:"country_icc_elected"`
	//[16] country_icc_zone                               char(30)             null: false  primary: false  isArray: false  auto: false  col: char            len: 30      default: []
	CountryIccZone string `gorm:"column:country_icc_zone;type:char;size:30;" form:"country_icc_zone" json:"country_icc_zone"`
	//[17] notes                                          char(255)            null: false  primary: false  isArray: false  auto: false  col: char            len: 255     default: []
	Notes string `gorm:"column:notes;type:char;size:255;" form:"notes" json:"notes"`
	//[18] logo_object_id                                 uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	LogoObjectID uint32 `gorm:"column:logo_object_id;type:uint;default:0;" form:"logo_object_id" json:"logo_object_id"`
	//[19] logo_alt_id                                    char(10)             null: false  primary: false  isArray: false  auto: false  col: char            len: 10      default: []
	LogoAltID string `gorm:"column:logo_alt_id;type:char;size:10;" form:"logo_alt_id" json:"logo_alt_id"`
}

var cms_teamTableInfo = &TableInfo{
	Name: "cms_team",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "id",
			ProtobufFieldName:  "id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "full_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       100,
			GoFieldName:        "FullName",
			GoFieldType:        "string",
			JSONFieldName:      "full_name",
			ProtobufFieldName:  "full_name",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "short_name",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(12)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       12,
			GoFieldName:        "ShortName",
			GoFieldType:        "string",
			JSONFieldName:      "short_name",
			ProtobufFieldName:  "short_name",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "abbreviation",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(5)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       5,
			GoFieldName:        "Abbreviation",
			GoFieldType:        "string",
			JSONFieldName:      "abbreviation",
			ProtobufFieldName:  "abbreviation",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "filename",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       20,
			GoFieldName:        "Filename",
			GoFieldType:        "string",
			JSONFieldName:      "filename",
			ProtobufFieldName:  "filename",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "class_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "ClassID",
			GoFieldType:        "uint32",
			JSONFieldName:      "class_id",
			ProtobufFieldName:  "class_id",
			ProtobufType:       "uint32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "country_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "CountryID",
			GoFieldType:        "uint32",
			JSONFieldName:      "country_id",
			ProtobufFieldName:  "country_id",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "continent_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "ContinentID",
			GoFieldType:        "uint32",
			JSONFieldName:      "continent_id",
			ProtobufFieldName:  "continent_id",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "address",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       200,
			GoFieldName:        "Address",
			GoFieldType:        "string",
			JSONFieldName:      "address",
			ProtobufFieldName:  "address",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "team_founded",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       20,
			GoFieldName:        "TeamFounded",
			GoFieldType:        "string",
			JSONFieldName:      "team_founded",
			ProtobufFieldName:  "team_founded",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "achievements",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       255,
			GoFieldName:        "Achievements",
			GoFieldType:        "string",
			JSONFieldName:      "achievements",
			ProtobufFieldName:  "achievements",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "player_numbers",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       100,
			GoFieldName:        "PlayerNumbers",
			GoFieldType:        "string",
			JSONFieldName:      "player_numbers",
			ProtobufFieldName:  "player_numbers",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "club_numbers",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       100,
			GoFieldName:        "ClubNumbers",
			GoFieldType:        "string",
			JSONFieldName:      "club_numbers",
			ProtobufFieldName:  "club_numbers",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "is_country",
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
			GoFieldName:        "IsCountry",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "is_country",
			ProtobufFieldName:  "is_country",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "country_icc_status",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       50,
			GoFieldName:        "CountryIccStatus",
			GoFieldType:        "string",
			JSONFieldName:      "country_icc_status",
			ProtobufFieldName:  "country_icc_status",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "country_icc_elected",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "CountryIccElected",
			GoFieldType:        "uint32",
			JSONFieldName:      "country_icc_elected",
			ProtobufFieldName:  "country_icc_elected",
			ProtobufType:       "uint32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "country_icc_zone",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(30)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       30,
			GoFieldName:        "CountryIccZone",
			GoFieldType:        "string",
			JSONFieldName:      "country_icc_zone",
			ProtobufFieldName:  "country_icc_zone",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "notes",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       255,
			GoFieldName:        "Notes",
			GoFieldType:        "string",
			JSONFieldName:      "notes",
			ProtobufFieldName:  "notes",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "logo_object_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "LogoObjectID",
			GoFieldType:        "uint32",
			JSONFieldName:      "logo_object_id",
			ProtobufFieldName:  "logo_object_id",
			ProtobufType:       "uint32",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "logo_alt_id",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       10,
			GoFieldName:        "LogoAltID",
			GoFieldType:        "string",
			JSONFieldName:      "logo_alt_id",
			ProtobufFieldName:  "logo_alt_id",
			ProtobufType:       "string",
			ProtobufPos:        20,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CmsTeam) TableName() string {
	return "cms_team"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CmsTeam) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CmsTeam) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CmsTeam) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CmsTeam) TableInfo() *TableInfo {
	return cms_teamTableInfo
}
