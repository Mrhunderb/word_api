package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	fmt.Println("Connecting to database...")
	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	}
	DB = db
	DB.AutoMigrate(&Dict{}, &User{}, &Word{}, &Example{},
		&Definition{}, &Plan{}, &Quiz{}, &Collection{})
	fmt.Println("Database connection successful.")
}

func GetDictList() []Dict {
	var dicts []Dict
	DB.Find(&dicts)
	return dicts
}

func FindUser(username string, password string) (*User, error) {
	var user User
	result := DB.First(&user, "user_name = ?", username)
	if result.Error == nil {
		if user.Password == password {
			return &user, nil
		} else {
			return nil, fmt.Errorf("密码错误")
		}
	} else {
		return nil, fmt.Errorf("用户不存在")
	}
}

func FindUserByID(user_id string) (*User, error) {
	var user User
	result := DB.First(&user, "user_id = ?", user_id)
	if result.Error == nil {
		return &user, nil
	} else {
		return nil, fmt.Errorf("用户不存在")
	}
}

func InsertUser(username string, password string) (*User, error) {
	var user User
	result := DB.First(&user, "user_name = ?", username)
	if result.Error == nil {
		return nil, fmt.Errorf("用户已经存在")
	} else {
		DB.Create(&User{
			UserName: username,
			Password: password,
		})
		return &user, nil
	}
}

func FindPlanByUserID(user_id int) (*Plan, error) {
	var user User
	var plan Plan
	result := DB.First(&user, "user_id = ?", user_id)
	if result.Error != nil {
		return nil, fmt.Errorf("用户不存在")
	}
	result = DB.First(&plan, "plan_id = ?", user.PlanID)
	if result.Error == nil {
		return &plan, nil
	} else {
		return nil, fmt.Errorf("计划不存在")
	}
}

func FindPlan(plan_id string) (*Plan, error) {
	var plan Plan
	result := DB.First(&plan, "plan_id = ?", plan_id)
	if result.Error == nil {
		return &plan, nil
	} else {
		return nil, fmt.Errorf("计划不存在")
	}
}

func FindDict(dict_id int) (*Dict, error) {
	var dict Dict
	result := DB.First(&dict, "dict_id = ?", dict_id)
	if result.Error == nil {
		return &dict, nil
	} else {
		return nil, fmt.Errorf("词典不存在")
	}
}

func FindWordByAlpha(dict_id int, limit int, offset int) ([]*Word, error) {
	var words []*Word
	result := DB.Preload("Definition").Preload("Example").Preload("Quiz").
		Limit(limit+1).Offset(offset).Where("dict_id = ?", dict_id).Order("word").Find(&words)
	if result.Error == nil {
		return words, nil
	} else {
		return nil, fmt.Errorf("词典不存在")
	}
}

func FindWordByAlphaDesc(dict_id int, limit int, offset int) ([]*Word, error) {
	var words []*Word
	result := DB.Preload("Definition").Preload("Example").Preload("Quiz").
		Limit(limit+1).Offset(offset).Where("dict_id = ?", dict_id).Order("word DESC").Find(&words)
	if result.Error == nil {
		return words, nil
	} else {
		return nil, fmt.Errorf("词典不存在")
	}
}

func FindWordByRandom(dict_id int, limit int, offset int) ([]*Word, error) {
	var words []*Word
	result := DB.Preload("Definition").Preload("Example").Preload("Quiz").
		Limit(limit+1).Offset(offset).Where("dict_id = ?", dict_id).Find(&words)
	if result.Error == nil {
		return words, nil
	} else {
		return nil, fmt.Errorf("词典不存在")
	}
}

func FindAllWord(dict_id string, offset int) ([]*Word, error) {
	var words []*Word
	result := DB.Preload("Definition").Preload("Example").
		Where("dict_id = ?", dict_id).Limit(50).Offset(offset).Find(&words)
	if result.Error == nil {
		return words, nil
	} else {
		return nil, fmt.Errorf("词典不存在")
	}
}

func FindUserCollection(user_id string) ([]*Word, error) {
	var words []*Word
	subquery := DB.Table("collections").Select("word_id").Where("user_id = ?", user_id)
	result := DB.Preload("Definition").Preload("Example").Where("word_id IN (?)", subquery).Find(&words)
	if result.Error == nil {
		return words, nil
	} else {
		return nil, fmt.Errorf("生词本为空")
	}
}
