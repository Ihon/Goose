package main

import (
	"encoding/json"
	"fmt"
	"gomymath/models"
	"gomymath/service"
	"gomymath/system"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"time"
	// "github.com/kataras/iris/v12"
	// "github.com/kataras/iris/v12/mvc"
)

// mediaType := map[string]string {
//     "gensaku_anime": "anime"
//     "gensaku_manga": "manga"
//     "gensaku_novel": "novel"
//     "gensaku_lite_novel": "lite_novel"
//     "gensaku_ehon": "ehon"
//     "gensaku_game": "game"
//     "gensaku_card": "card"
//     "gensaku_kamisibai": "kamisibai"
//     "gensaku_denki": "denki"
//     "gensaku_gangu": "gangu"
//     "gensaku_chara": "chara"
//     "gensaku_jissya": "jissya"
//     "gensaku_tokusatu": "tokusatu"
//     "gensaku_adult_game": "adult_game"
//     "gensaku_etc": "etc"
// }

func main() {
    config, err := system.LoadConfig("config")
    if err != nil {
        log.Fatal("? Could not load environment variables", err)
    }
    system.ConnectDB(&config)

    pool := gojuonIndex()
    allcount := len(pool)
fmt.Println(allcount)
    animes := []*models.Anime{}
    for _, id := range pool {
        data := service.Anime(service.CollyAnime(id))

        if (data.Title == "") {
            fmt.Println(id, "Incomplete file data")
            continue
        }
fmt.Println("processing:", id)
        values := reflect.ValueOf(data)
        types := values.Type()

        animeData := models.Anime{}
        for i := 0; i < values.NumField(); i++ {
            name := types.Field(i).Name
            switch name {
                case "AnimedbId":
                   animeData.ID = values.Field(i).Interface().(uint32)
                case "Title":
                    animeData.Title = values.Field(i).Interface().(string)
                case "KanaTitle":
                    animeData.KanaTitle = values.Field(i).Interface().(string)
                case "ImageUrl":
                    animeData.ImageUrl = values.Field(i).Interface().(string)
                case "ReleaseMedia":
                    animeData.ReleaseMedia = values.Field(i).Interface().(string)
                case "OriginalMedia":
                    animeData.OriginalMedia = values.Field(i).Interface().(string)
                case "ReleaseDate":
                    animeData.ReleaseDate = values.Field(i).Interface().(string)
                case "ReleaseWay":
                    animeData.ReleaseWay = values.Field(i).Interface().(string)
                case "RunningTime":
                    animeData.RunningTime = values.Field(i).Interface().(int)
                case "Episodes":
                    animeData.Episodes = values.Field(i).Interface().(int)
                case "OriginalAuthor":
                    animeData.OriginalAuthor = values.Field(i).Interface().(string)
                case "Director":
                    animeData.Director = values.Field(i).Interface().(string)
                case "Production":
                    animeData.Production = values.Field(i).Interface().(string)
                default:
                    return
            }
        }
        animes = append(animes, &animeData)
// fmt.Printf("%+v\n", animeData)
        coolDown := rand.Intn(5)
        time.Sleep(time.Duration(coolDown) * time.Second)

        if (allcount >= 100 && len(animes) == 100) {
            RowsAffected := models.BatchInsertAnime(animes)
fmt.Println(RowsAffected)
            animes = []*models.Anime{}
            allcount -= 100
        } else if (allcount < 100 && allcount == len(animes)) {
            RowsAffected := models.BatchInsertAnime(animes)
fmt.Println("last lap:", RowsAffected)
        }
    }
// prettyStruct(animes)
}

func gojuonIndex() []string {
    targetUrl, itemNumber := service.CollyGojuon(16)

    for (itemNumber % 100) > 0 {
        itemNumber--
    }

    time.Sleep(time.Duration(5) * time.Second)

    pool := []string{}
    for (itemNumber >= 0) {
        queryResult := service.CollyGojuonItem(targetUrl, itemNumber)
        pool = append(pool, queryResult...)

        itemNumber -= 100
        time.Sleep(time.Duration(5) * time.Second)
    }

    return pool
}

func alternativeMain() {
    config, err := system.LoadConfig("config")
    if err != nil {
        log.Fatal("? Could not load environment variables", err)
    }
    system.ConnectDB(&config)

    // system.DB.AutoMigrate(&models.Anime{})
    // fmt.Println("? Migration complete")

    animes := []*models.Anime{}
    // 100787
    for id := int(14300); id <= 100787; id += 1000 {
        data := service.Anime(service.CollyAnime(strconv.Itoa(id)))

        if (data.Title == "") {
            fmt.Println(id, "Incomplete file data")
            continue
        }
fmt.Printf("%+v\n", data)
os.Exit(0)
        values := reflect.ValueOf(data)
        types := values.Type()

        animeData := models.Anime{}
        for i := 0; i < values.NumField(); i++ {
            name := types.Field(i).Name
            switch name {
                case "AnimedbId":
                   animeData.ID = values.Field(i).Interface().(uint32)
                case "Title":
                    animeData.Title = values.Field(i).Interface().(string)
                case "KanaTitle":
                    animeData.KanaTitle = values.Field(i).Interface().(string)
                case "ImageUrl":
                    animeData.ImageUrl = values.Field(i).Interface().(string)
                case "ReleaseMedia":
                    animeData.ReleaseMedia = values.Field(i).Interface().(string)
                case "OriginalMedia":
                    animeData.OriginalMedia = values.Field(i).Interface().(string)
                case "ReleaseDate":
                    animeData.ReleaseDate = values.Field(i).Interface().(string)
                case "ReleaseWay":
                    animeData.ReleaseWay = values.Field(i).Interface().(string)
                case "RunningTime":
                    animeData.RunningTime = values.Field(i).Interface().(int)
                case "Episodes":
                    animeData.Episodes = values.Field(i).Interface().(int)
                case "OriginalAuthor":
                    animeData.OriginalAuthor = values.Field(i).Interface().(string)
                case "Director":
                    animeData.Director = values.Field(i).Interface().(string)
                case "Production":
                    animeData.Production = values.Field(i).Interface().(string)
                default:
                    return
            }
        }
        animes = append(animes, &animeData)
// fmt.Printf("%+v\n", animeData)
        coolDown := rand.Intn(5)
        time.Sleep(time.Duration(coolDown) * time.Second)

        // if ((id % 100) == 0) {
            RowsAffected := models.BatchInsertAnime(animes)
fmt.Println(RowsAffected)
            animes = []*models.Anime{}
        // }
    }
// prettyStruct(animes)
os.Exit(0)
}

func prettyStruct(intef interface{}) {
    output, _ := json.MarshalIndent(intef, "", "\t")
    fmt.Printf("%s \n", output)
}
