package main

import (
    "fmt"
    "log"
    // "gomymath/libraries"
    // "github.com/kataras/iris/v12"
    // "github.com/kataras/iris/v12/mvc"
    // "gorm.io/driver/postgres"
    // "gorm.io/gorm"
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
    // var animeId uint32 = 1
    // 100787
    // var data anime = collyAnime(animeId)
    config, err := LoadConfig("config")
fmt.Printf("%+v\n", config)
    if err != nil {
        log.Fatal("? Could not load environment variables", err)
    }

    ConnectDB(&config)

    // fmt.Printf("%+v\n", data)
}
