package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID    int32          `gorm:"type:int;primarykey;autoIncrement:true"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserName  string         `gorm:"type:varchar(50)"`
	Password  string         `gorm:"type:varchar(50)"`
	PlanID    int32          `gorm:"type:int"`
}

type Dict struct {
	DictID     int32  `gorm:"type:int;primaryKey;autoIncrement:true"`
	DictName   string `gorm:"type:varchar(50)"`
	CoverUrl   string `gorm:"type:varchar(255)"`
	TotalWords int32  `gorm:"type:int"`
	Word       []Word `gorm:"foreignKey:DictID"`
}

type Word struct {
	WordID        int32        `gorm:"type:int;primaryKey;autoIncrement:true"`
	Word          string       `gorm:"type:varchar(50)"`
	Pronunciation string       `gorm:"type:varchar(50)"`
	DictID        int32        `gorm:"type:int"`
	Example       []Example    `gorm:"foreignKey:WordID"`
	Definition    []Definition `gorm:"foreignKey:WordID"`
	Quiz          Quiz         `gorm:"foreignKey:WordID"`
}

type Example struct {
	ExampleID int32  `gorm:"type:int;primaryKey;autoIncrement:true"`
	EnExample string `gorm:"type:varchar(255)"`
	ChExample string `gorm:"type:varchar(255)"`
	WordID    int32  `gorm:"type:int"`
}

type Definition struct {
	DefinitionID int32  `gorm:"type:int;primaryKey;autoIncrement:true"`
	Definition   string `gorm:"type:varchar(255)"`
	PartOfSpeech string `gorm:"type:varchar(25)"`
	WordID       int32  `gorm:"type:int"`
}

type Plan struct {
	PlanID   int32 `gorm:"type:int;primaryKey;autoIncrement:true"`
	Mode     int8  `gorm:"type:tinyint"`
	NLearn   int32 `gorm:"type:int"`
	NReview  int32 `gorm:"type:int"`
	Progress int32 `gorm:"type:int"`
	UserID   int32 `gorm:"type:int"`
	DictID   int32 `gorm:"type:int"`
}

type Quiz struct {
	QuizID        int32  `gorm:"type:int;primaryKey;autoIncrement:true"`
	OptionA       string `gorm:"type:varchar(255)"`
	OptionB       string `gorm:"type:varchar(255)"`
	OptionC       string `gorm:"type:varchar(255)"`
	OptionD       string `gorm:"type:varchar(255)"`
	CorrectOption int32  `gorm:"type:int"`
	WordID        int32  `gorm:"type:int"`
}

type Collection struct {
	CollectionID int32     `gorm:"type:int;primaryKey;autoIncrement:true"`
	WordID       int32     `gorm:"type:int"`
	UserID       int32     `gorm:"type:int"`
	AddTime      time.Time `gorm:"autoCreateTime"`
}

type History struct {
	HistoryID   int32     `gorm:"type:int;primaryKey;autoIncrement:true"`
	WordID      int32     `gorm:"type:int"`
	PlanID      int32     `gorm:"type:int"`
	StartTime   time.Time `gorm:"autoCreateTime"`
	UpdateTime  time.Time `gorm:"autoUpdateTime"`
	Proficiency int32     `gorm:"type:int;Default:0"`
}
