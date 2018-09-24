package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
)

var oddsUrl = "http://odds.500.com/europe_jczq_%s.shtml"
var ouzhiUrl = "http://odds.500.com/fenxi/ouzhi-%s.shtml?ctype=2" //百家欧洲指数和必发
var touzhuUrl = "http://odds.500.com/fenxi/touzhu-%s.shtml"
var yazhiUrl = "http://odds.500.com/fenxi/yazhi-%s.shtml?ctype=2"

type Match struct {
	ID          string //比赛编号
	HName       string //主队名称
	VName       string //客队名称
	Score       string //比分
	LeagueMatch string //联赛
	Rotation    string //轮次
}

type OuZhi struct {
	ID             string
	Name           string //公司名字
	BeginWinRate   string
	BeginWinPD     string
	BeginDrawRate  string
	BeginDrawPD    string
	BeginLossRate  string
	BeginLossPD    string
	BeginReturn    string
	BeginWinKaili  string
	BeginDrawKaili string
	BeginLossKaili string

	EndWinRate   string
	EndWinPD     string
	EndDrawRate  string
	EndDrawPD    string
	EndLossRate  string
	EndLossPD    string
	EndReturn    string
	EndWinKaili  string
	EndDrawKaili string
	EndLossKaili string
}
type YaZhi struct {
	ID          string
	Name        string
	BeginHWater string
	BeginPan    string
	BeginVWater string

	EndHWater          string
	EndHWaterDirection string
	EndPan             string
	EndPanState        string

	EndVWater          string
	EndVWaterDirection string
}
type TouZhu struct {
	ID         string
	DealerWin  string //主胜庄家盈亏
	DealerDraw string
	DealerLoss string

	WinIndex  string //主胜冷热指数
	DrawIndex string
	LossIndex string
}

var MatchData = new(Match)
var cookie = `sdc_session=1537535875941; _jzqc=1; _jzqy=1.1537535876.1537535876.1.jzqsr=baidu.-; __utmc=63332592; __utmz=63332592.1537535877.1.1.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; ck_RegUrl=trade.500.com; appform=1; bdshare_firstime=1537536082256; ck_user=MTU1MjAwMTAwMDk%3D; ck_user2=MTU1MjAwMTAwMDk%3D; ck_RegFromUrl=https%3A//www.baidu.com/link%3Furl%3Dojv-2sXkKqB9vUmBoL0kUfqLJX3yYK2iaPCl_6_WLXS%26wd%3D%26eqid%3Df3646a0500024697000000065ba5913f; seo_key=baidu%7C%7Chttps://www.baidu.com/link?url=ojv-2sXkKqB9vUmBoL0kUfqLJX3yYK2iaPCl_6_WLXS&wd=&eqid=f3646a0500024697000000065ba5913f; Hm_lvt_4f816d475bb0b9ed640ae412d6b42cab=1537535876,1537577300; btn_follow=-1; op_chupan=1; _jzqx=1.1537576510.1537711878.5.jzqsr=odds%2E500%2Ecom|jzqct=/fenxi/rangqiu-730460%2Eshtml.jzqsr=odds%2E500%2Ecom|jzqct=/europe_jczq_2018-09-22%2Eshtml; _jzqckmp=1; usercheck=MzA0NDc4MTEyOGIxYzYzYTU2NGU1MWQwNzFiYmRjMDI4M2JhMTExOA%3D%3D; _jzqa=1.547320212841500200.1537535876.1537711878.1537766696.9; __utma=63332592.1738223105.1537535877.1537711891.1537766697.9; __utmt=1; motion_id=1537767483081_0.881682107095074; WT_FPC=id=undefined:lv=1537767531464:ss=1537766696246; sdc_userflag=1537766696251::1537767531471::6; _qzja=1.1025932337.1537535916302.1537711877086.1537766696284.1537767336167.1537767531514.0.0.0.962.9; _qzjb=1.1537766696284.6.0.0.0; _qzjc=1; _qzjto=94.1.0; _jzqb=1.6.10.1537766696.1; Hm_lpvt_4f816d475bb0b9ed640ae412d6b42cab=1537767532; __utmb=63332592.6.10.1537766697; CLICKSTRN_ID=111.10.136.74-1537535876.226933::8E3BA998695510A86EA5DA5F1DC17237`
var matchs = make([]Match, 0)
var OuZhis = make([]OuZhi, 0)
var YaZhis = make([]YaZhi, 0)
var BiFa = new(TouZhu)

var times = []string{"2018-09-10"}

func main() {
	// MatchData.ID = "762333"

	for _, time := range times {
		GetMatchs(time)
	}
	count := len(matchs)
	fmt.Println("一共", count, "场比赛")

	for i, _ := range matchs {
		fmt.Println("开始抓取%d场,剩余%d", i, count-i)
		matchs[i].OuZhi()
		matchs[i].Bifar()
		matchs[i].YaZhi()
	}

	fmt.Println(matchs)

}

//
func GetMatchs(time string) {

	html := GET(fmt.Sprintf(oddsUrl, time), cookie)
	htmlcode := strings.NewReader(GBK2UTF8(html))
	doc, _ := goquery.NewDocumentFromReader(htmlcode)
	doc.Find("#main-tbody").Find("tr").Each(func(i int, tr *goquery.Selection) {
		if i%2 == 0 {
			var match Match

			id, _ := tr.Attr("data-fid")
			match.ID = id
			tr.Find("td").Each(func(j int, td *goquery.Selection) {
				switch j {
				case 1:
					match.LeagueMatch, _ = td.Find("a").Html()
				case 2:
					match.Rotation = td.Text()
				case 4:
					match.HName, _ = td.Find("a").Html()
				case 5:
					match.Score, _ = td.Find("span").Html()
				case 6:
					match.VName, _ = td.Find("a").Html()
				}
			})
			matchs = append(matchs, match)
		}
	})

}

func (match *Match) OuZhi() {
	// var ouzhis = make([]OuZhi, 0)
	html := GET(fmt.Sprintf(ouzhiUrl, match.ID), cookie)
	htmlcode := strings.NewReader(GBK2UTF8(html))
	doc, _ := goquery.NewDocumentFromReader(htmlcode)
	doc.Find("#datatb").Find("tr").Each(func(i int, tr *goquery.Selection) {
		var ouzhi OuZhi
		tr.Find("td").Each(func(j int, td *goquery.Selection) {
			ouzhi.ID = match.ID
			value := td.Text()
			switch j {
			case 1:
				if title, has := td.Attr("title"); has {
					ouzhi.Name = title
				}
			case 3: //初始胜
				// fmt.Println(td.Html())
				ouzhi.BeginWinRate = value
			case 4:
				ouzhi.BeginDrawRate = value
			case 5:
				ouzhi.BeginLossRate = value
			case 6:
				ouzhi.EndWinRate = value
			case 7:
				ouzhi.EndDrawRate = value
			case 8:
				ouzhi.EndLossRate = value
			case 10:
				ouzhi.BeginWinPD = value
			case 11:
				ouzhi.BeginDrawPD = value
			case 12:
				ouzhi.BeginLossPD = value
			case 13:
				ouzhi.EndWinPD = value
			case 14:
				ouzhi.EndDrawPD = value
			case 15:
				ouzhi.EndLossPD = value
			case 17:
				ouzhi.BeginReturn = value
			case 18:
				ouzhi.EndReturn = value
			case 20:
				ouzhi.BeginWinKaili = value
			case 21:
				ouzhi.BeginDrawKaili = value
			case 22:
				ouzhi.BeginLossKaili = value
			case 23:
				ouzhi.EndWinKaili = value
			case 24:
				ouzhi.EndDrawKaili = value
			case 25:
				ouzhi.EndLossKaili = value
			}

		})
		OuZhis = append(OuZhis, ouzhi)
	})
	// match.OuZhis = ouzhis
}

func GBK2UTF8(gbk string) string {
	srcCoder := mahonia.NewDecoder("gbk")
	srcResult := srcCoder.ConvertString(gbk)
	tagCoder := mahonia.NewDecoder("utf-8")
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return string(cdata)
}

func (match *Match) Bifar() {

	html := GET(fmt.Sprintf(touzhuUrl, match.ID), cookie)
	htmlcode := strings.NewReader(GBK2UTF8(html))
	doc, _ := goquery.NewDocumentFromReader(htmlcode)
	doc.Find(".pub_table.pl_table_data.bif-yab").Find("tr").Find("td").Each(func(i int, td *goquery.Selection) {

		value := td.Text()
		switch i {
		case 18:
			BiFa.DealerWin = value
		case 20:
			BiFa.WinIndex = value //主胜冷热指数
		case 29:
			BiFa.DealerDraw = value
		case 32:
			BiFa.DrawIndex = value //主胜冷热指数
		case 40:
			BiFa.DealerLoss = value
		case 42:
			BiFa.LossIndex = value //客胜冷热指数
		}

	})
}

//yazhiUrl
func (match *Match) YaZhi() {
	// var yazhis = make([]YaZhi, 0)
	html := GET(fmt.Sprintf(yazhiUrl, match.ID), cookie)
	htmlcode := strings.NewReader(GBK2UTF8(html))
	doc, _ := goquery.NewDocumentFromReader(htmlcode)
	doc.Find("#datatb").Find("tr").Each(func(i int, tr *goquery.Selection) {
		var yazhi YaZhi

		tr.Find("td").Each(func(j int, td *goquery.Selection) {
			yazhi.ID = match.ID
			switch j {
			case 1: //公司
				yazhi.Name = td.Find(".quancheng").Text()
			case 3:
				ying := td.Text()
				// fmt.Println(td.Html())
				if len(ying) > 4 {
					yazhi.EndHWater = string([]byte(ying)[:5])
					yazhi.EndHWaterDirection = string([]byte(ying[5:]))
				}
			case 4:

				panS := strings.Split(td.Text(), " ")
				if len(panS) > 1 {

					yazhi.EndPan = panS[0]
					yazhi.EndPanState = panS[1]
				}
			case 5:
				ping := td.Text()
				yazhi.EndVWater = string([]byte(ping)[:5])
				yazhi.EndVWaterDirection = string([]byte(ping[5:]))
			case 9:
				yazhi.BeginHWater = td.Text()
			case 10:
				yazhi.BeginPan = td.Text()
			case 11:
				yazhi.BeginVWater = td.Text()
			}

		})
		if "" != yazhi.Name {
			YaZhis = append(YaZhis, yazhi)
		}

	})
	// match.YaZhis = yazhis
}
