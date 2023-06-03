package models

import (
    "time"
    "gomymath/system"
    "gorm.io/gorm/clause"
)

type Anime struct {
    ID             uint32    `gorm:"primary_key"`
    Title          string    `gorm:"type:varchar(512);not null"`
    KanaTitle      string    `gorm:"type:varchar(512);not null"`
    ImageUrl       string    `gorm:"type:varchar(512);not null"`
    ReleaseMedia   string    `gorm:"type:varchar(512);not null"`
    OriginalMedia  string    `gorm:"type:varchar(512);not null"`
    ReleaseDate    string    `gorm:"type:varchar(512);not null"`
    ReleaseWay     string    `gorm:"type:varchar(512);not null"`
    RunningTime    int       `gorm:"type:uint;default:0"`
    Episodes       int       `gorm:"type:uint;default:0"`
    OriginalAuthor string    `gorm:"type:varchar(512);not null"`
    Director       string    `gorm:"type:varchar(512);not null"`
    Production     string    `gorm:"type:varchar(512);not null"`
    CreatedAt      time.Time `default:NOW()"`
    UpdatedAt      time.Time `gorm:"<-:create;default:NOW()"`
}

func InsertAnime(anime Anime) int64 {
    result := system.DB.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(&anime)
    return result.RowsAffected
}

func BatchInsertAnime(anime []*Anime) int64 {
    result := system.DB.Clauses(clause.Insert{Modifier: "IGNORE"}).Create(anime)
    return result.RowsAffected
}

func UpdateAnime() {

}

func GetAnime() {

}

func DeleteAnime() {

}
