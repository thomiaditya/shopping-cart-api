package model

import "github.com/thomiaditya/shop-api/internal/database"

// Database connection instance
var db = database.GetDatabaseInstance().GetConnection()
