package models

import (
    "time"
)

type Anime struct {
    ID             uint64    `gorm:"primary_key"`
    Title          string    `gorm:"type:varchar(512);not  null"`
    KanaTitle      string    `gorm:"type:varchar(512);not  null"`
    imageUrl       string    `gorm:"type:varchar(512);not  null"`
    OriginalMedia  string    `gorm:"type:varchar(512);not  null"`
    ReleaseDate    string    `gorm:"type:varchar(512);not  null"`
    ReleaseWay     string    `gorm:"type:varchar(512);not  null"`
    RunningTime    uint8     `gorm:"type:uint;default:0"`
    Episodes       uint8     `gorm:"type:uint;default:0"`
    OriginalAuthor string    `gorm:"type:varchar(512);not  null"`
    Director       string    `gorm:"type:varchar(512);not  null"`
    Production     string    `gorm:"type:varchar(512);not  null"`
    CreatedAt      time.Time
    UpdatedAt      time.Time `gorm:"<-:create"`
}

func InsertAnime() {

}

func BatchInsertAnime() {

}

func UpdateAnime() {

}

func GetAnime() {

}

func DeleteAnime() {

}
