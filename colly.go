package main

import (
    "strings"
    "strconv"
    "github.com/gocolly/colly"
)

type anime struct {
    AnimedbId int
    Title string
    KanaTitle string
    imageUrl string
    OriginalMedia string
    ReleaseDate string
    ReleaseWay string
    RunningTime int
    Episodes int
    OriginalAuthor string
    Director string
    Production string
}

func collyAnime(id uint32) anime {
    informationMap := make(map[int]string)
    informationType := map[string]string {
        "初出日": "ReleaseDate",
        "初出方法": "ReleaseWay",
        "分数": "RunningTime",
        "話数": "Episodes",
        "原作": "OriginalAuthor",
        "監督": "Director",
        "制作": "Production",
    }
    countN, countD := 0, 0
    rawData := anime{}
    // 在colly中使用 Collector 這類物件 來做事情
    c := colly.NewCollector()

    // 當Visit訪問網頁後，網頁響應(Response)時候執行的事情
    // c.OnResponse(func(r *colly.Response) {
    //     // 返回的Response物件r.Body 是[]Byte格式，要再轉成字串
    //     fmt.Println(string(r.Body))
    // })

    c.OnHTML(".title > h1:nth-child(1)", func(e *colly.HTMLElement) {
        rawData.Title = strings.TrimSpace(e.Text)
    })

    c.OnHTML(".title > h2:nth-child(1)", func(e *colly.HTMLElement) {
        rawData.KanaTitle = strings.TrimSpace(e.Text)
    })

    c.OnHTML(".pick > img:nth-child(1)", func(e *colly.HTMLElement) {
        rawData.imageUrl = strings.TrimSpace(e.Attr("src"))
    })

    c.OnHTML("img.sakuhin_icon:nth-child(2)", func(e *colly.HTMLElement) {
        rawData.imageUrl = strings.TrimSpace(e.Attr("src"))
    })

    c.OnHTML("img.sakuhin_icon:nth-child(3)", func(e *colly.HTMLElement) {
        rawData.imageUrl = strings.TrimSpace(e.Attr("src"))
    })

    c.OnHTML(".kihon > dt", func(e *colly.HTMLElement) {
        countN++
        informationName := strings.TrimSpace(e.Text)
        informationName = strings.Replace(informationName, "■", "", -1)
        name, isExist := informationType[informationName]
        if (isExist) {
            informationMap[countN] = name
        }
    })

    c.OnHTML(".kihon > dd", func(e *colly.HTMLElement) {
        countD++
        informationData := strings.TrimSpace(e.Text)
        name := informationMap[countD]
        switch {
            case name == "ReleaseDate":
                rawData.ReleaseDate = informationData
            case name == "ReleaseWay":
                rawData.ReleaseWay = informationData
            case name == "RunningTime":
                informationData = strings.Replace(informationData, "分", "", -1)
                rawData.RunningTime, _ = strconv.Atoi(informationData)
            case name == "Episodes":
                informationData = strings.Replace(informationData, "話", "", -1)
                rawData.Episodes, _ = strconv.Atoi(informationData)
            case name == "OriginalAuthor":
                rawData.OriginalAuthor = informationData
            case name == "Director":
                rawData.Director = informationData
            case name == "Production":
                rawData.Production = informationData
            default:
                return
        }
    })

    c.OnRequest(func(r *colly.Request) {
        r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
    })

    c.Visit("https://db.animedb.jp/index.php/searchdata/?mode=view&id=100611")

    return rawData
}
