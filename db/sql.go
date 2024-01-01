package db

import (
	"fmt"
	"time"

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
		&Definition{}, &Plan{}, &Quiz{}, &Collection{}, &History{})
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

func FindPlanByID(user_id string, dict_id string) (*Plan, error) {
	var plan Plan
	result := DB.First(&plan, "user_id = ? AND dict_id = ?", user_id, dict_id)
	if result.Error == nil {
		return &plan, nil
	} else {
		return nil, fmt.Errorf("计划不存在")
	}
}

func AddPlan(user_id int, dict_id int, mode int, n_learn int, n_review int) (*Plan, error) {
	plan := Plan{
		UserID:   int32(user_id),
		DictID:   int32(dict_id),
		Mode:     int8(mode),
		NLearn:   int32(n_learn),
		NReview:  int32(n_review),
		Progress: 0,
	}
	DB.Create(&plan)
	return &plan, nil
}

func UpdatePlan(plan Plan, mode int, n_learn int, n_review int) (*Plan, error) {
	plan.Mode = int8(mode)
	plan.NLearn = int32(n_learn)
	plan.NReview = int32(n_review)
	DB.Save(&plan)
	return &plan, nil
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
		Limit(limit).Offset(offset).Where("dict_id = ?", dict_id).Order("word").Find(&words)
	if result.Error == nil {
		return words, nil
	} else {
		return nil, fmt.Errorf("词典不存在")
	}
}

func FindWordByAlphaDesc(dict_id int, limit int, offset int) ([]*Word, error) {
	var words []*Word
	result := DB.Preload("Definition").Preload("Example").Preload("Quiz").
		Limit(limit).Offset(offset).Where("dict_id = ?", dict_id).Order("word DESC").Find(&words)
	if result.Error == nil {
		return words, nil
	} else {
		return nil, fmt.Errorf("词典不存在")
	}
}

func FindWordByRandom(dict_id int, limit int, offset int) ([]*Word, error) {
	var words []*Word
	result := DB.Preload("Definition").Preload("Example").Preload("Quiz").
		Limit(limit).Offset(offset).Where("dict_id = ?", dict_id).Find(&words)
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

func AddUserCollection(user_id int, word_id int) error {
	var collection Collection
	result := DB.First(&collection, "user_id = ? AND word_id = ?", user_id, word_id)
	if result.Error == nil {
		return fmt.Errorf("该生词已存在")
	}
	DB.Create(&Collection{
		UserID: int32(user_id),
		WordID: int32(word_id),
	})
	return nil
}

func DeletCollection(user_id int, word_id int) error {
	var collection Collection
	result := DB.First(&collection, "user_id = ? AND word_id = ?", user_id, word_id)
	if result.Error != nil {
		return fmt.Errorf("该生词不存在")
	}
	DB.Delete(&collection)
	return nil
}

func UpdateUserPlan(user_id int, plan_id int) error {
	result := DB.First(&User{}, "user_id = ?", user_id).Update("plan_id", plan_id)
	if result.Error != nil {
		return fmt.Errorf("用户不存在")
	}
	return nil
}

func AddUserHistory(plan_id int, word_id int, is_know int) error {
	var history History
	var plan Plan
	result := DB.First(&history, "plan_id = ? AND word_id = ?", plan_id, word_id)
	if result.Error == nil {
		DB.Model(&history).Update("Proficiency", history.Proficiency+int32(is_know))
	} else {
		DB.First(&plan, "plan_id = ?", plan_id)
		DB.Model(&plan).Update("progress", plan.Progress+1)
		DB.Create(&History{
			PlanID:      int32(plan_id),
			WordID:      int32(word_id),
			Proficiency: int32(is_know),
		})
	}
	return nil
}

func GetTodyLearn(plan_id int) (int, error) {
	var n int64
	today := time.Now().Format("2006-01-02")
	result := DB.Model(&History{}).Where("plan_id = ?", plan_id).
		Where("DATE(start_time) = ?", today).Count(&n)
	if result.Error != nil {
		return 0, fmt.Errorf("无记录")
	}
	return int(n), nil
}

func GetTodyReview(plan_id int) (int, error) {
	var n int64
	today := time.Now().Format("2006-01-02")
	result := DB.Model(&History{}).Where("plan_id = ?", plan_id).
		Where("DATE(start_time) != ?", today).
		Where("DATE(update_time) = ? ", today).Count(&n)
	if result.Error != nil {
		return 0, fmt.Errorf("无记录")
	}
	return int(n), nil
}
