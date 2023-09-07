package main

import (
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Announcement struct {
	ID             string
	Content        string
	TimePublished  time.Time
	TimeEdited     time.Time
	HasAttachments bool
	Reactions      []Reaction   `gorm:"foreignKey:MessageID"`
	Attachments    []Attachment `gorm:"foreignKey:MessageID"`
}

type Reaction struct {
	ID        uint `gorm:"primarykey"`
	MessageID string
	EmojiID   string
	EmojiName string
	UserID    string
}

type Attachment struct {
	ID          uint `gorm:"primarykey"`
	MessageID   string
	URL         string
	Filename    string
	ContentType string
	Size        int
}

func MigrateDB() {
	db.AutoMigrate(&Announcement{})
	db.AutoMigrate(&Reaction{})
}

var db *gorm.DB

func InitDB() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	dbl, err := gorm.Open(sqlite.Open(".db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	db = dbl
	MigrateDB()
}
