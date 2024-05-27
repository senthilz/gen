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


CREATE TABLE `cms_match` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `international_class_id` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `general_class_id` tinyint(3) unsigned NOT NULL DEFAULT 0,
  `international_number` char(7) DEFAULT NULL,
  `general_number` char(7) DEFAULT NULL,
  `match_title` char(50) DEFAULT NULL,
  `team1_id` smallint(5) unsigned NOT NULL DEFAULT 0,
  `team2_id` smallint(5) unsigned NOT NULL DEFAULT 0,
  `ground_id` smallint(5) unsigned DEFAULT NULL,
  `country_id` smallint(5) unsigned DEFAULT NULL,
  `start_date` date NOT NULL DEFAULT '0000-00-00',
  `end_date` date NOT NULL DEFAULT '0000-00-00',
  `floodlit` enum('Y','N') DEFAULT NULL,
  `season` char(7) DEFAULT NULL,
  `schedule_note` char(255) NOT NULL,
  `cancelled_match` enum('N','Y') NOT NULL DEFAULT 'N',
  `live_match` enum('N','Y','U') NOT NULL DEFAULT 'N',
  `live_companion` enum('N','Y') NOT NULL DEFAULT 'N',
  `live_note` char(255) NOT NULL,
  `start_time_gmt` time DEFAULT NULL,
  `start_datetime_gmt` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `hours_string` char(255) NOT NULL DEFAULT '',
  `date_string` char(75) NOT NULL DEFAULT '',
  `scheduled_days` tinyint(3) unsigned DEFAULT NULL,
  `scheduled_overs` tinyint(3) unsigned DEFAULT NULL,
  `scheduled_innings` tinyint(3) unsigned DEFAULT NULL,
  `keywords_teams` char(255) NOT NULL,
  `keywords_other` char(255) NOT NULL,
  `player_rating` enum('Y','N') NOT NULL DEFAULT 'N',
  `rating_promo` char(255) NOT NULL,
  `team1_name_id` smallint(5) unsigned DEFAULT NULL,
  `team2_name_id` smallint(5) unsigned DEFAULT NULL,
  `ground_name_id` smallint(5) unsigned DEFAULT NULL,
  `country_name_id` smallint(5) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `international_class_id` (`international_class_id`),
  KEY `general_class_id` (`general_class_id`),
  KEY `international_number` (`international_number`),
  KEY `general_number` (`general_number`),
  KEY `team1_id` (`team1_id`),
  KEY `team2_id` (`team2_id`),
  KEY `ground_id` (`ground_id`),
  KEY `country_id` (`country_id`),
  KEY `team1_name_id` (`team1_name_id`),
  KEY `team2_name_id` (`team2_name_id`),
  KEY `ground_name_id` (`ground_name_id`),
  KEY `country_name_id` (`country_name_id`),
  KEY `start_date` (`start_date`),
  KEY `live_match` (`live_match`),
  FULLTEXT KEY `keywords` (`keywords_teams`,`keywords_other`,`match_title`,`season`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb3 COLLATE=utf8mb3_general_ci

JSON Sample
-------------------------------------
{    "id": 29,    "international_class_id": 48,    "general_class_id": 60,    "international_number": "rOyXjkRgreXKTQFeddSdEeJRH",    "general_number": "UDxrRWPQIqTVTEBWFyZZKiFgG",    "match_title": "nrnLSWDHamEkRPnBGGHFlJhWV",    "team_1_id": 17,    "team_2_id": 83,    "ground_id": 86,    "country_id": 34,    "start_date": "2118-09-02T00:27:32.788214798+01:00",    "end_date": "2222-09-26T03:47:28.32087294+01:00",    "floodlit": "VljooCLhcWVqoLmEbQDoTsQcW",    "season": "PhvKUdFbDfpGGHAPYWeREdkuk",    "schedule_note": "vduoCPLWrmdlQfdmHNJDWBkBg",    "cancelled_match": "rVIULFUQYtWCPUVyNOiCclgDS",    "live_match": "XoLERpqJjthLfhpeCayLaqrvD",    "live_companion": "eyASscCKvZTRCqOjHFQcyUaQY",    "live_note": "MGFOBCVsoRZCDnyakspdYVTJA",    "start_time_gmt": "2201-07-19T22:16:09.414869301+01:00",    "start_datetime_gmt": "2185-12-28T23:37:35.860986875Z",    "hours_string": "vSmhIplXWwwUNEHSkgZhUlECI",    "date_string": "RjbSMUysTIAuncIoviOBUKdPB",    "scheduled_days": 31,    "scheduled_overs": 17,    "scheduled_innings": 0,    "keywords_teams": "ZRDjedxYIBcwgiahsZVMPTExg",    "keywords_other": "FBmyaeelLTuAqOOMThZpdfBrr",    "player_rating": "VJKTGKQbNWVbjMPNtyASSrdDT",    "rating_promo": "hhCYyglgFOEBlQXMfhDvSDbcT",    "team_1_name_id": 97,    "team_2_name_id": 98,    "ground_name_id": 70,    "country_name_id": 15}


Comments
-------------------------------------
[ 0] column is set for unsigned
[ 1] column is set for unsigned
[ 2] column is set for unsigned
[ 6] column is set for unsigned
[ 7] column is set for unsigned
[ 8] column is set for unsigned
[ 9] column is set for unsigned
[23] column is set for unsigned
[24] column is set for unsigned
[25] column is set for unsigned
[30] column is set for unsigned
[31] column is set for unsigned
[32] column is set for unsigned
[33] column is set for unsigned



*/

// CmsMatch struct is a row record of the cms_match table in the zen database
type CmsMatch struct {
	//[ 0] id                                             uint                 null: false  primary: true   isArray: false  auto: true   col: uint            len: -1      default: []
	ID uint32 `gorm:"primary_key;AUTO_INCREMENT;column:id;type:uint;" form:"id" json:"id"`
	//[ 1] international_class_id                         utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	InternationalClassID uint32 `gorm:"column:international_class_id;type:utinyint;default:0;" form:"international_class_id" json:"international_class_id"`
	//[ 2] general_class_id                               utinyint             null: false  primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [0]
	GeneralClassID uint32 `gorm:"column:general_class_id;type:utinyint;default:0;" form:"general_class_id" json:"general_class_id"`
	//[ 3] international_number                           char(7)              null: true   primary: false  isArray: false  auto: false  col: char            len: 7       default: [NULL]
	InternationalNumber sql.NullString `gorm:"column:international_number;type:char;size:7;" form:"international_number" json:"international_number"`
	//[ 4] general_number                                 char(7)              null: true   primary: false  isArray: false  auto: false  col: char            len: 7       default: [NULL]
	GeneralNumber sql.NullString `gorm:"column:general_number;type:char;size:7;" form:"general_number" json:"general_number"`
	//[ 5] match_title                                    char(50)             null: true   primary: false  isArray: false  auto: false  col: char            len: 50      default: [NULL]
	MatchTitle sql.NullString `gorm:"column:match_title;type:char;size:50;" form:"match_title" json:"match_title"`
	//[ 6] team1_id                                       usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	Team1ID uint32 `gorm:"column:team1_id;type:usmallint;default:0;" form:"team_1_id" json:"team_1_id"`
	//[ 7] team2_id                                       usmallint            null: false  primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [0]
	Team2ID uint32 `gorm:"column:team2_id;type:usmallint;default:0;" form:"team_2_id" json:"team_2_id"`
	//[ 8] ground_id                                      usmallint            null: true   primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [NULL]
	GroundID sql.NullInt64 `gorm:"column:ground_id;type:usmallint;" form:"ground_id" json:"ground_id"`
	//[ 9] country_id                                     usmallint            null: true   primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [NULL]
	CountryID sql.NullInt64 `gorm:"column:country_id;type:usmallint;" form:"country_id" json:"country_id"`
	//[10] start_date                                     date                 null: false  primary: false  isArray: false  auto: false  col: date            len: -1      default: [0000-00-00]
	StartDate time.Time `gorm:"column:start_date;type:date;default:0000-00-00;" form:"start_date" json:"start_date"`
	//[11] end_date                                       date                 null: false  primary: false  isArray: false  auto: false  col: date            len: -1      default: [0000-00-00]
	EndDate time.Time `gorm:"column:end_date;type:date;default:0000-00-00;" form:"end_date" json:"end_date"`
	//[12] floodlit                                       char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [NULL]
	Floodlit sql.NullString `gorm:"column:floodlit;type:char;size:1;" form:"floodlit" json:"floodlit"`
	//[13] season                                         char(7)              null: true   primary: false  isArray: false  auto: false  col: char            len: 7       default: [NULL]
	Season sql.NullString `gorm:"column:season;type:char;size:7;" form:"season" json:"season"`
	//[14] schedule_note                                  char(255)            null: false  primary: false  isArray: false  auto: false  col: char            len: 255     default: []
	ScheduleNote string `gorm:"column:schedule_note;type:char;size:255;" form:"schedule_note" json:"schedule_note"`
	//[15] cancelled_match                                char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [N]
	CancelledMatch string `gorm:"column:cancelled_match;type:char;size:1;default:N;" form:"cancelled_match" json:"cancelled_match"`
	//[16] live_match                                     char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [N]
	LiveMatch string `gorm:"column:live_match;type:char;size:1;default:N;" form:"live_match" json:"live_match"`
	//[17] live_companion                                 char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [N]
	LiveCompanion string `gorm:"column:live_companion;type:char;size:1;default:N;" form:"live_companion" json:"live_companion"`
	//[18] live_note                                      char(255)            null: false  primary: false  isArray: false  auto: false  col: char            len: 255     default: []
	LiveNote string `gorm:"column:live_note;type:char;size:255;" form:"live_note" json:"live_note"`
	//[19] start_time_gmt                                 time                 null: true   primary: false  isArray: false  auto: false  col: time            len: -1      default: [NULL]
	StartTimeGmt time.Time `gorm:"column:start_time_gmt;type:time;" form:"start_time_gmt" json:"start_time_gmt"`
	//[20] start_datetime_gmt                             datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: [0000-00-00 00:00:00]
	StartDatetimeGmt time.Time `gorm:"column:start_datetime_gmt;type:datetime;default:0000-00-00 00:00:00;" form:"start_datetime_gmt" json:"start_datetime_gmt"`
	//[21] hours_string                                   char(255)            null: false  primary: false  isArray: false  auto: false  col: char            len: 255     default: []
	HoursString string `gorm:"column:hours_string;type:char;size:255;" form:"hours_string" json:"hours_string"`
	//[22] date_string                                    char(75)             null: false  primary: false  isArray: false  auto: false  col: char            len: 75      default: []
	DateString string `gorm:"column:date_string;type:char;size:75;" form:"date_string" json:"date_string"`
	//[23] scheduled_days                                 utinyint             null: true   primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [NULL]
	ScheduledDays sql.NullInt64 `gorm:"column:scheduled_days;type:utinyint;" form:"scheduled_days" json:"scheduled_days"`
	//[24] scheduled_overs                                utinyint             null: true   primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [NULL]
	ScheduledOvers sql.NullInt64 `gorm:"column:scheduled_overs;type:utinyint;" form:"scheduled_overs" json:"scheduled_overs"`
	//[25] scheduled_innings                              utinyint             null: true   primary: false  isArray: false  auto: false  col: utinyint        len: -1      default: [NULL]
	ScheduledInnings sql.NullInt64 `gorm:"column:scheduled_innings;type:utinyint;" form:"scheduled_innings" json:"scheduled_innings"`
	//[26] keywords_teams                                 char(255)            null: false  primary: false  isArray: false  auto: false  col: char            len: 255     default: []
	KeywordsTeams string `gorm:"column:keywords_teams;type:char;size:255;" form:"keywords_teams" json:"keywords_teams"`
	//[27] keywords_other                                 char(255)            null: false  primary: false  isArray: false  auto: false  col: char            len: 255     default: []
	KeywordsOther string `gorm:"column:keywords_other;type:char;size:255;" form:"keywords_other" json:"keywords_other"`
	//[28] player_rating                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [N]
	PlayerRating string `gorm:"column:player_rating;type:char;size:1;default:N;" form:"player_rating" json:"player_rating"`
	//[29] rating_promo                                   char(255)            null: false  primary: false  isArray: false  auto: false  col: char            len: 255     default: []
	RatingPromo string `gorm:"column:rating_promo;type:char;size:255;" form:"rating_promo" json:"rating_promo"`
	//[30] team1_name_id                                  usmallint            null: true   primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [NULL]
	Team1NameID sql.NullInt64 `gorm:"column:team1_name_id;type:usmallint;" form:"team_1_name_id" json:"team_1_name_id"`
	//[31] team2_name_id                                  usmallint            null: true   primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [NULL]
	Team2NameID sql.NullInt64 `gorm:"column:team2_name_id;type:usmallint;" form:"team_2_name_id" json:"team_2_name_id"`
	//[32] ground_name_id                                 usmallint            null: true   primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [NULL]
	GroundNameID sql.NullInt64 `gorm:"column:ground_name_id;type:usmallint;" form:"ground_name_id" json:"ground_name_id"`
	//[33] country_name_id                                usmallint            null: true   primary: false  isArray: false  auto: false  col: usmallint       len: -1      default: [NULL]
	CountryNameID sql.NullInt64 `gorm:"column:country_name_id;type:usmallint;" form:"country_name_id" json:"country_name_id"`
}

var cms_matchTableInfo = &TableInfo{
	Name: "cms_match",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "uint",
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
			Name:               "international_class_id",
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
			GoFieldName:        "InternationalClassID",
			GoFieldType:        "uint32",
			JSONFieldName:      "international_class_id",
			ProtobufFieldName:  "international_class_id",
			ProtobufType:       "uint32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "general_class_id",
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
			GoFieldName:        "GeneralClassID",
			GoFieldType:        "uint32",
			JSONFieldName:      "general_class_id",
			ProtobufFieldName:  "general_class_id",
			ProtobufType:       "uint32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "international_number",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(7)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       7,
			GoFieldName:        "InternationalNumber",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "international_number",
			ProtobufFieldName:  "international_number",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "general_number",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(7)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       7,
			GoFieldName:        "GeneralNumber",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "general_number",
			ProtobufFieldName:  "general_number",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "match_title",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       50,
			GoFieldName:        "MatchTitle",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "match_title",
			ProtobufFieldName:  "match_title",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "team1_id",
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
			GoFieldName:        "Team1ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "team_1_id",
			ProtobufFieldName:  "team_1_id",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "team2_id",
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
			GoFieldName:        "Team2ID",
			GoFieldType:        "uint32",
			JSONFieldName:      "team_2_id",
			ProtobufFieldName:  "team_2_id",
			ProtobufType:       "uint32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "ground_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "GroundID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "ground_id",
			ProtobufFieldName:  "ground_id",
			ProtobufType:       "uint32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "country_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "CountryID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "country_id",
			ProtobufFieldName:  "country_id",
			ProtobufType:       "uint32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "start_date",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "date",
			DatabaseTypePretty: "date",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "date",
			ColumnLength:       -1,
			GoFieldName:        "StartDate",
			GoFieldType:        "time.Time",
			JSONFieldName:      "start_date",
			ProtobufFieldName:  "start_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "end_date",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "date",
			DatabaseTypePretty: "date",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "date",
			ColumnLength:       -1,
			GoFieldName:        "EndDate",
			GoFieldType:        "time.Time",
			JSONFieldName:      "end_date",
			ProtobufFieldName:  "end_date",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "floodlit",
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
			GoFieldName:        "Floodlit",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "floodlit",
			ProtobufFieldName:  "floodlit",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "season",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(7)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       7,
			GoFieldName:        "Season",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "season",
			ProtobufFieldName:  "season",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "schedule_note",
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
			GoFieldName:        "ScheduleNote",
			GoFieldType:        "string",
			JSONFieldName:      "schedule_note",
			ProtobufFieldName:  "schedule_note",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "cancelled_match",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "CancelledMatch",
			GoFieldType:        "string",
			JSONFieldName:      "cancelled_match",
			ProtobufFieldName:  "cancelled_match",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "live_match",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "LiveMatch",
			GoFieldType:        "string",
			JSONFieldName:      "live_match",
			ProtobufFieldName:  "live_match",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "live_companion",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "LiveCompanion",
			GoFieldType:        "string",
			JSONFieldName:      "live_companion",
			ProtobufFieldName:  "live_companion",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "live_note",
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
			GoFieldName:        "LiveNote",
			GoFieldType:        "string",
			JSONFieldName:      "live_note",
			ProtobufFieldName:  "live_note",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "start_time_gmt",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "time",
			DatabaseTypePretty: "time",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "time",
			ColumnLength:       -1,
			GoFieldName:        "StartTimeGmt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "start_time_gmt",
			ProtobufFieldName:  "start_time_gmt",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "start_datetime_gmt",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "StartDatetimeGmt",
			GoFieldType:        "time.Time",
			JSONFieldName:      "start_datetime_gmt",
			ProtobufFieldName:  "start_datetime_gmt",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "hours_string",
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
			GoFieldName:        "HoursString",
			GoFieldType:        "string",
			JSONFieldName:      "hours_string",
			ProtobufFieldName:  "hours_string",
			ProtobufType:       "string",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "date_string",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(75)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       75,
			GoFieldName:        "DateString",
			GoFieldType:        "string",
			JSONFieldName:      "date_string",
			ProtobufFieldName:  "date_string",
			ProtobufType:       "string",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "scheduled_days",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "ScheduledDays",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "scheduled_days",
			ProtobufFieldName:  "scheduled_days",
			ProtobufType:       "uint32",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "scheduled_overs",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "ScheduledOvers",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "scheduled_overs",
			ProtobufFieldName:  "scheduled_overs",
			ProtobufType:       "uint32",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "scheduled_innings",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "utinyint",
			DatabaseTypePretty: "utinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "utinyint",
			ColumnLength:       -1,
			GoFieldName:        "ScheduledInnings",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "scheduled_innings",
			ProtobufFieldName:  "scheduled_innings",
			ProtobufType:       "uint32",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "keywords_teams",
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
			GoFieldName:        "KeywordsTeams",
			GoFieldType:        "string",
			JSONFieldName:      "keywords_teams",
			ProtobufFieldName:  "keywords_teams",
			ProtobufType:       "string",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "keywords_other",
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
			GoFieldName:        "KeywordsOther",
			GoFieldType:        "string",
			JSONFieldName:      "keywords_other",
			ProtobufFieldName:  "keywords_other",
			ProtobufType:       "string",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "player_rating",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "PlayerRating",
			GoFieldType:        "string",
			JSONFieldName:      "player_rating",
			ProtobufFieldName:  "player_rating",
			ProtobufType:       "string",
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "rating_promo",
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
			GoFieldName:        "RatingPromo",
			GoFieldType:        "string",
			JSONFieldName:      "rating_promo",
			ProtobufFieldName:  "rating_promo",
			ProtobufType:       "string",
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "team1_name_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "Team1NameID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "team_1_name_id",
			ProtobufFieldName:  "team_1_name_id",
			ProtobufType:       "uint32",
			ProtobufPos:        31,
		},

		&ColumnInfo{
			Index:              31,
			Name:               "team2_name_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "Team2NameID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "team_2_name_id",
			ProtobufFieldName:  "team_2_name_id",
			ProtobufType:       "uint32",
			ProtobufPos:        32,
		},

		&ColumnInfo{
			Index:              32,
			Name:               "ground_name_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "GroundNameID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "ground_name_id",
			ProtobufFieldName:  "ground_name_id",
			ProtobufType:       "uint32",
			ProtobufPos:        33,
		},

		&ColumnInfo{
			Index:              33,
			Name:               "country_name_id",
			Comment:            ``,
			Notes:              `column is set for unsigned`,
			Nullable:           true,
			DatabaseTypeName:   "usmallint",
			DatabaseTypePretty: "usmallint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "usmallint",
			ColumnLength:       -1,
			GoFieldName:        "CountryNameID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "country_name_id",
			ProtobufFieldName:  "country_name_id",
			ProtobufType:       "uint32",
			ProtobufPos:        34,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CmsMatch) TableName() string {
	return "cms_match"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CmsMatch) BeforeSave(tx *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CmsMatch) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CmsMatch) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CmsMatch) TableInfo() *TableInfo {
	return cms_matchTableInfo
}
