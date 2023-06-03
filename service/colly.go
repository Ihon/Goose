package service

import (
    // "fmt"
	"log"
    "regexp"
	"net/url"
	"strconv"
	"strings"
	"github.com/gocolly/colly"
)

type Anime struct {
    AnimedbId uint32
    Title string
    KanaTitle string
    ImageUrl string
    ReleaseMedia string
    OriginalMedia string
    ReleaseDate string
    ReleaseWay string
    RunningTime int
    Episodes int
    OriginalAuthor string
    Director string
    Production string
}

func CollyAnime(id string) Anime {
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
    basicSelecter := ".kihon > dd"
    countN, countD := 0, 0
    rawData := Anime{}
    id64, err := strconv.ParseUint(strings.TrimLeft(id, "0*"), 10, 32)
    if err != nil {
        panic(err)
    }
    rawData.AnimedbId = uint32(id64)

    // 在colly中使用 Collector 這類物件 來做事情
    c := colly.NewCollector()

    // 當Visit訪問網頁後，網頁響應(Response)時候執行的事情
    // c.OnResponse(func(r *colly.Response) {
    //     // 返回的Response物件r.Body 是[]Byte格式，要再轉成字串
    //     fmt.Println(string(r.Body))
    // })

    c.OnHTML(".title > h1", func(e *colly.HTMLElement) {
        rawData.Title = strings.TrimSpace(e.Text)
    })

    c.OnHTML(".title > h2", func(e *colly.HTMLElement) {
        rawData.KanaTitle = strings.TrimSpace(e.Text)
    })

    c.OnHTML(".pick > img:nth-child(1)", func(e *colly.HTMLElement) {
        rawData.ImageUrl = strings.TrimSpace(e.Attr("src"))
    })

    c.OnHTML("img.sakuhin_icon:nth-child(2)", func(e *colly.HTMLElement) {
        rawData.ReleaseMedia = strings.TrimSpace(e.Attr("src"))
    })

    c.OnHTML("img.sakuhin_icon:nth-child(3)", func(e *colly.HTMLElement) {
        rawData.OriginalMedia = strings.TrimSpace(e.Attr("src"))
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

    if (countN == 0) {
        c.OnHTML(".kihon > dl > dt", func(e *colly.HTMLElement) {
            countN++
            informationName := strings.TrimSpace(e.Text)
            informationName = strings.Replace(informationName, "■", "", -1)
            name, isExist := informationType[informationName]
            if (isExist) {
                informationMap[countN] = name
            }
        })

        basicSelecter = ".kihon > dl > dd"
    }

    c.OnHTML(basicSelecter, func(e *colly.HTMLElement) {
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

    animedbUrl, err := url.Parse("https://db.animedb.jp/index.php/searchdata/?mode=view&id=100611")
    if err != nil {
        log.Fatal(err)
    }

    paramate := animedbUrl.Query()
    paramate.Set("id", id)
    animedbUrl.RawQuery = paramate.Encode()

    c.Visit(animedbUrl.String())

    return rawData
}

func CollyGojuon(index uint) (string, int)  {
    gojuonList := []string{
        "/index.php/searchdata/?mode=50on&word=ア",
        "/index.php/searchdata/?mode=50on&word=イ",
        "/index.php/searchdata/?mode=50on&word=ウ",
        "/index.php/searchdata/?mode=50on&word=エ",
        "/index.php/searchdata/?mode=50on&word=オ",
        "/index.php/searchdata/?mode=50on&word=カ",
        "/index.php/searchdata/?mode=50on&word=キ",
        "/index.php/searchdata/?mode=50on&word=ク",
        "/index.php/searchdata/?mode=50on&word=ケ",
        "/index.php/searchdata/?mode=50on&word=コ",
        "/index.php/searchdata/?mode=50on&word=ガ",
        "/index.php/searchdata/?mode=50on&word=ギ",
        "/index.php/searchdata/?mode=50on&word=グ",
        "/index.php/searchdata/?mode=50on&word=ゲ",
        "/index.php/searchdata/?mode=50on&word=ゴ",
        "/index.php/searchdata/?mode=50on&word=サ",
        "/index.php/searchdata/?mode=50on&word=シ",
        "/index.php/searchdata/?mode=50on&word=ス",
        "/index.php/searchdata/?mode=50on&word=セ",
        "/index.php/searchdata/?mode=50on&word=ソ",
        "/index.php/searchdata/?mode=50on&word=ザ",
        "/index.php/searchdata/?mode=50on&word=ジ",
        "/index.php/searchdata/?mode=50on&word=ズ",
        "/index.php/searchdata/?mode=50on&word=ゼ",
        "/index.php/searchdata/?mode=50on&word=ゾ",
        "/index.php/searchdata/?mode=50on&word=タ",
        "/index.php/searchdata/?mode=50on&word=チ",
        "/index.php/searchdata/?mode=50on&word=ツ",
        "/index.php/searchdata/?mode=50on&word=テ",
        "/index.php/searchdata/?mode=50on&word=ト",
        "/index.php/searchdata/?mode=50on&word=ダ",
        "/index.php/searchdata/?mode=50on&word=ヂ",
        "/index.php/searchdata/?mode=50on&word=ヅ",
        "/index.php/searchdata/?mode=50on&word=デ",
        "/index.php/searchdata/?mode=50on&word=ド",
        "/index.php/searchdata/?mode=50on&word=ナ",
        "/index.php/searchdata/?mode=50on&word=ニ",
        "/index.php/searchdata/?mode=50on&word=ヌ",
        "/index.php/searchdata/?mode=50on&word=ネ",
        "/index.php/searchdata/?mode=50on&word=ノ",
        "/index.php/searchdata/?mode=50on&word=ハ",
        "/index.php/searchdata/?mode=50on&word=ヒ",
        "/index.php/searchdata/?mode=50on&word=フ",
        "/index.php/searchdata/?mode=50on&word=ヘ",
        "/index.php/searchdata/?mode=50on&word=ホ",
        "/index.php/searchdata/?mode=50on&word=バ",
        "/index.php/searchdata/?mode=50on&word=ビ",
        "/index.php/searchdata/?mode=50on&word=ブ",
        "/index.php/searchdata/?mode=50on&word=べ",
        "/index.php/searchdata/?mode=50on&word=ボ",
        "/index.php/searchdata/?mode=50on&word=パ",
        "/index.php/searchdata/?mode=50on&word=ピ",
        "/index.php/searchdata/?mode=50on&word=プ",
        "/index.php/searchdata/?mode=50on&word=ペ",
        "/index.php/searchdata/?mode=50on&word=ポ",
        "/index.php/searchdata/?mode=50on&word=マ",
        "/index.php/searchdata/?mode=50on&word=ミ",
        "/index.php/searchdata/?mode=50on&word=ム",
        "/index.php/searchdata/?mode=50on&word=メ",
        "/index.php/searchdata/?mode=50on&word=モ",
        "/index.php/searchdata/?mode=50on&word=ヤ",
        "/index.php/searchdata/?mode=50on&word=ユ",
        "/index.php/searchdata/?mode=50on&word=ヨ",
        "/index.php/searchdata/?mode=50on&word=ラ",
        "/index.php/searchdata/?mode=50on&word=リ",
        "/index.php/searchdata/?mode=50on&word=ル",
        "/index.php/searchdata/?mode=50on&word=レ",
        "/index.php/searchdata/?mode=50on&word=ロ",
        "/index.php/searchdata/?mode=50on&word=ワ",
        "/index.php/searchdata/?mode=50on&word=ヲ",
        "/index.php/searchdata/?mode=50on&word=ン",
    }

    var url strings.Builder
    url.WriteString("https://db.animedb.jp")
    url.WriteString(gojuonList[index])

    itemNumber := string("")
    c := colly.NewCollector()
    c.OnHTML(".single > b:nth-child(3)", func(e *colly.HTMLElement) {
        itemNumber = strings.TrimSpace(e.Text)
    })

    c.OnRequest(func(r *colly.Request) {
        r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
    })

    c.Visit(url.String())

    number, _ := strconv.Atoi(itemNumber)

    return url.String(), number
}

func CollyGojuonItem(gojuonUrl string, pageNumber int) []string {
    itemNumbers := []string{}

    animedbUrl, err := url.Parse(gojuonUrl)
    if err != nil {
        log.Fatal(err)
    }

    paramate := animedbUrl.Query()
    paramate.Add("pg", strconv.Itoa(pageNumber))
    animedbUrl.RawQuery = paramate.Encode()

    c := colly.NewCollector()
    c.OnHTML("div.appendifsc > dl > a", func(e *colly.HTMLElement) {
        onclickAttr := strings.TrimSpace(e.Attr("onclick"))
        re := regexp.MustCompile("id=([\\d]+)")
        itemId := re.FindStringSubmatch(onclickAttr)
        itemNumbers = append(itemNumbers, itemId[1])
    })

    c.OnRequest(func(r *colly.Request) {
        r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")
    })

    c.Visit(animedbUrl.String())

    return itemNumbers
}
