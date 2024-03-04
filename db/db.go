package db

import "gorm.io/gorm"

var dbUser string = "root"
var dbPass string = "***"
var dbName string = "word"
var dbHost string = "localhost"
var dbPort string = "3306"

var DB *gorm.DB
