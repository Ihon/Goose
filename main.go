package main

import (
    "fmt"
    "log"
    "math/rand"
    "os"
    "time"
    "gomymath/models"
    "gomymath/service"
    "gomymath/system"

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

    // 100787
    for i := uint32(1); i <= 100787; i++ {
        var data service.Anime = service.CollyAnime(i)

        if (data.Title == "") {
            fmt.Println(i, "Incomplete file data")
        } else {
fmt.Printf("%+v\n", data)
            os.Exit(0)
        }

        coolDown := rand.Intn(15)
        time.Sleep(time.Duration(coolDown) * time.Second)
    }

    // system.DB.AutoMigrate(&models.Anime{})
    // fmt.Println("? Migration complete")
}
