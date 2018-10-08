package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/html"

	"github.com/Luxurioust/excelize"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"

	"github.com/robertkrimen/otto"
)

var KailiHost = "http://vip.win0168.com/AsianOdds_n.aspx?id="
var ScoreHost = "http://zq.win007.com/jsData/matchResult/"
var GoalHost = "http://info.win0168.com/analysis/"
var BFIndex = "http://info.win0168.com/analysis/odds/"

var BFHost = "http://vip.win0168.com/betfa/single.aspx?id=1557502"
var Teams = make(map[string]string, 0)
var Copms = make([]Competition, 0)
var StopTime = time.Now().Add(24 * time.Hour)

type Competition struct {
	VSID           string //比赛编号
	VSState        string //是否已经比赛
	HTeamID        string //主队编号
	HTeamName      string //主队名称
	HTeamIDRanking string //主队排名
	HGetGoal       string //进球
	HLossGoal      string //失球
	VSHGoal        string //主队进球

	HEveyGoals   string //主队场均进球
	HWinPercent  string //主队胜率
	HDrawPercent string //主队和率
	HLossPercent string //主队负率

	VTeamID        string //客队编号
	VTeamName      string //客队名称
	VTeamIDRanking string //客队排名
	VGetGoal       string //进球
	VLossGoal      string //失球
	VSVGoal        string //客队进球

	VEveyGoals   string //主队场均进球
	VWinPercent  string //主队胜率
	VDrawPercent string //主队和率
	VLossPercent string //主队负率

	VSTime   string //比赛时间
	VSResult string
	BfWin    string //必发胜
	BfHe     string //必发平
	BfLoss   string //必发输

	BfWinWantBuy  string //必发主队买
	BfDrawWantBuy string //必发主队买
	BfLossWantBuy string //必发主队买

	BfWinWantSell  string //必发主队卖
	BfDrawWantSell string //必发主队卖
	BfLossWantSell string //必发主队卖

	BFWinBuyCount  string //主队买单
	BFDrawBuyCount string //和买单
	BFLossBuyCount string //客队买单

	BFWinSellCount  string //主队买单
	BFDrawSellCount string //和买单
	BFLossSellCount string //客队买单

	KailiWin  string
	KailHe    string
	KailiLoss string

	ErrorNumber int //错误次数
}

var vm = otto.New()
var xlsx = excelize.NewFile()
var errorNumer = 0
var hasError = false

func main() {
	//GetTeams()
	// var dates = []string{"20180915", "20180914", "20180913", "20180912", "20180911", "20180910", "20180909", "20180908", "20180907"}
	// var dates = []string{"20180916"}
	// for _, date := range dates {
	// 	url := "http://www.win0168.com/football/hg/Over_" + date + ".htm"

	// 	GetScore1(url)

	// }
	// // GetScore1()
	// // GetTodayVS()
	// fmt.Println("一共有%d场记录", len(Copms))
	// // fmt.Println(Copms)
	// //
	// //
	// // Copms = Copms[:1]
	// // Copms = append(Copms, Competition{VSID: "1503897"})
	// GetData(0)
	// WriteExcel()
	// fmt.Println(Copms[0].VSID, Copms[0].HWinPercent)
	// fmt.Println(com.BfWin, com.BfWinWantBuy, com.BfWinWantSell)
	// fmt.Println(com.BfHe, com.BfDrawWantBuy, com.BfDrawWantSell)

}

func GetTodayVS() {
	url := "http://www.win0168.com/vbsxml/bfdata.js?r=0071537188542001"
	cookie := `UM_distinctid=165c3525acc854-09c2457c51505e-34677908-13c680-165c3525acd67c; Bet007EuropeIndex_Cookie=null; Bet007live_hiddenID=_1552192_1552185_1583226_1583221_1583227_1583225_1560947_1560944_1585085_1585088_1585090_1585093_1585089_1585087_1549717_1549712_1549715_1511797_1511798_1600021_1600024_1600019_1562117_1562120_1562121_1584031_1584027_1584025_1584030_1584028_1584026_1555672_1555671_1555670_1551888_1551889_1551890_1551891_1551887_1559002_1558995_1558997_1559327_1559328_1559326_1600877_1600874_1600997_1600879_1489037_1489038_1489040_1489042_1489043_1489036_1489260_1489261_1489266_1496838_1500427_1500478_1500433_1500425_1500431_1500430_1500429_1500428_1584503_1584494_1584493_1584504_1584505_1510018_1510010_1510012_1510013_1510015_1510014_1553298_1582377_1488037_1488039_1500680_1500678_1500679_1489576_1489578_1489581_1489582_1489577_1489579_1489580_1503894_1503893_1503895_1550738_1550735_1550737_1550734_1565864_1565865_1565866_1565867_1565868_1565869_1559645_1517503_1551332_1551331_1560702_1560701_1550922_1550919_1562810_1562812_1562808_1557067_1557068_1557070_1557071_1555927_1555929_1567333_1567329_1503527_1503528_1503529_1503524_1570104_1570101_1613270_1613273_1613271_1613272_1613269_1510873_1525964_1525968_1525962_1572756_1572757_1572749_1572751_1572754_1563222_1574388_1574394_1574392_1574393_1574702_1574704_1574696_1574699_1618654_1555470_1557303_1557301_1617427_1617428_1617429_1575627_1575628_1575630_1575632_1575633_1588401_1588402_1621381_1572198_1572199_1577169_1577168_1577174_1577172_1577173_1577170_1541675_1541674_1621397_1621398_1621399_1601447_1601460_1601480_1591236_1591235_1554740_1554736_1554735_1554738_1554737_1554739_1581848_1581849_1581850_1593306_1593307_1567857_1567863_1567856_1567860_1567861_1567858_1567859_1605294_1605293_1605295_1605291_1605289_1575316_1575319_1500972_1500973_1500970_1500971_1576497_1621662_1621663_; win007BfCookie=0^0^1^1^1^1^1^0^0^0^0^0^1^2^1^1^0^1^1; FS007Filter=1^0^_36_31_8_34_11_60_40_9_33_5_16_10_3_32_22_26_25_21_2_4_17_138_122_15_123_284_7_127_128_61_766_118_6_119_137_129_124_166_159_242_358_230_250_135_30_130_133_131_505_240_308_354_391_460_466_469_504_593_613_140_423_596_700_847_1366_; bfWin007FirstMatchTime=2018,8,17,08,00,00`
	jsCode := GET(url, cookie)
	srcCoder := mahonia.NewDecoder("gbk")
	srcResult := srcCoder.ConvertString(jsCode)
	tagCoder := mahonia.NewDecoder("utf-8")
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)

	vm.Run(cdata)
	bfdata, _ := vm.Get("A")

	for _, key := range bfdata.Object().Keys() {
		match, _ := bfdata.Object().Get(key)
		var comps Competition

		//排除已经比赛了的场
		matchState, _ := match.Object().Get("13")
		if "-1" != matchState.String() {
			comps.VSState = matchState.String()

		}

		for j, itemKey := range match.Object().Keys() {
			item, _ := match.Object().Get(itemKey)

			switch j {
			case 0:
				comps.VSID = item.String()
			case 5:
				comps.HTeamName = item.String()
			case 8:
				comps.VTeamName = item.String()
			case 14:
				comps.VSHGoal = item.String() //主队进球
			case 15:
				comps.VSVGoal = item.String() //客队进球
			}
		}
		if "" != comps.VSID {
			Copms = append(Copms, comps)
		}
	}
}

//获取已经比赛完的场次信息
func GetScore1(url string) {
	defer func(url string) {
		if r := recover(); r != nil {
			fmt.Printf("Get URL Error %s\n", url)
			// debug.PrintStack()
			// WriteExcel() //先写一部分
			GetScore1(url)
		}
	}(url)

	// var url = "http://www.win0168.com/football/Next_20180915.htm"
	cookie := "UM_distinctid=165c3525acc854-09c2457c51505e-34677908-13c680-165c3525acd67c; bfWin007FirstMatchTime=2018,8,14,08,00,00; win007BfCookie=2^0^1^1^1^1^1^0^0^0^0^0^1^2^1^1^0^1^1; Bet007EuropeIndex_Cookie=null"
	html := GET(url, cookie)
	htmlcode := strings.NewReader(html)
	doc, _ := goquery.NewDocumentFromReader(htmlcode)
	doc.Find("#table_live").Find("tr").Each(func(tri int, tr *goquery.Selection) {
		infoid, _ := tr.Attr("infoid")
		//47  韩国
		//1 英格兰
		//46 日本
		//4 德国
		//7苏格兰
		//41 美国
		//39 巴西
		//38 阿根廷
		//3 西班牙
		//43墨西哥
		//5法国
		if "46" == infoid || "4" == infoid || "1" == infoid || "47" == infoid || "41" == infoid || "39" == infoid {
			var comp Competition

			tr.Find("td").Each(func(i int, td *goquery.Selection) {

				if i == 3 {
					comp.HTeamName = td.Text()
				}
				if 4 == i {
					td.Find("font").Each(func(i int, font *goquery.Selection) {
						if 0 == i {
							comp.VSHGoal = font.Text()
						} else {
							comp.VSVGoal = font.Text()
						}
					})
				}
				if i == 5 {
					comp.VTeamName = td.Text()
				}
				//over9 next 7
				if (i == 7 || i == 9) && "" != td.Find("a").Text() {

					AsianOdds := td.Find("a").Get(0).Attr[1].Val
					comp.VSID = strings.Split(AsianOdds, "(")[1][0:7]

				}

			})
			// if comp.VSHGoal == comp.VSVGoal { //平局
			Copms = append(Copms, comp)
			// }
			// Copms = append(Copms, comp)

		}

	})

}

func GetData(LoopNumber int) {

	defer func(LoopNumber *int) {
		if r := recover(); r != nil {
			fmt.Printf("第%d个错误\n", *LoopNumber)
			// debug.PrintStack()
			// WriteExcel() //先写一部分
			errorNumer = *LoopNumber
			Copms[errorNumer].ErrorNumber++
			if Copms[errorNumer].ErrorNumber > 2 {
				errorNumer++
			}
			GetData(errorNumer)
		}
	}(&LoopNumber)

	for ; LoopNumber < len(Copms); LoopNumber++ {
		if errorNumer > 0 {
			fmt.Printf("从第%d个恢复\n", errorNumer)
			LoopNumber = errorNumer
			errorNumer = 0
		}
		fmt.Println("剩", len(Copms)-LoopNumber)
		Copms[LoopNumber].GetBFData()
		Copms[LoopNumber].GetGoal()
		// Copms[LoopNumber].GetKaili()

		if "" == Copms[LoopNumber].BfWin {
			Copms[LoopNumber].GetBFData()
		}
		if "" == Copms[LoopNumber].HGetGoal {
			Copms[LoopNumber].GetGoal()
		}

	}
}

func WriteExcel() {
	xlsx.SetCellValue("Sheet1", "A1", "比赛编号")
	xlsx.SetCellValue("Sheet1", "B1", "比赛状态")
	xlsx.SetCellValue("Sheet1", "C1", "比赛时间")
	xlsx.SetCellValue("Sheet1", "D1", "主队编号")
	xlsx.SetCellValue("Sheet1", "E1", "主队名称")
	xlsx.SetCellValue("Sheet1", "F1", "主队排名")
	xlsx.SetCellValue("Sheet1", "G1", "客队排名")
	xlsx.SetCellValue("Sheet1", "H1", "客队编号")
	xlsx.SetCellValue("Sheet1", "I1", "客队名称")
	xlsx.SetCellValue("Sheet1", "J1", "必发胜")
	xlsx.SetCellValue("Sheet1", "K1", "必发平")
	xlsx.SetCellValue("Sheet1", "L1", "必发负")
	xlsx.SetCellValue("Sheet1", "M1", "凯利胜")
	xlsx.SetCellValue("Sheet1", "N1", "凯利平")
	xlsx.SetCellValue("Sheet1", "O1", "凯利输")
	xlsx.SetCellValue("Sheet1", "P1", "主队进球")
	xlsx.SetCellValue("Sheet1", "Q1", "主队失球")
	xlsx.SetCellValue("Sheet1", "R1", "客队进球")
	xlsx.SetCellValue("Sheet1", "S1", "客队失球")
	xlsx.SetCellValue("Sheet1", "T1", "主队比赛得分") //主队进球
	xlsx.SetCellValue("Sheet1", "U1", "客队比赛得分")
	xlsx.SetCellValue("Sheet1", "V1", "主队场均进球") //主队进球
	xlsx.SetCellValue("Sheet1", "W1", "客队场均进球") //主队进球
	xlsx.SetCellValue("Sheet1", "X1", "主队胜率")   //主队进球
	xlsx.SetCellValue("Sheet1", "Y1", "客队胜率")
	xlsx.SetCellValue("Sheet1", "Z1", "主队平率")      //主队进球
	xlsx.SetCellValue("Sheet1", "AA1", "客队平率")     //主队进球
	xlsx.SetCellValue("Sheet1", "AB1", "主队负率")     //主队进球
	xlsx.SetCellValue("Sheet1", "AC1", "客队负率")     //主队进球1
	xlsx.SetCellValue("Sheet1", "AD1", "必发购买倾向")   //主队进球
	xlsx.SetCellValue("Sheet1", "AE1", "必发和购买倾向")  //主队进球
	xlsx.SetCellValue("Sheet1", "AF1", "必发客队购买倾向") //主队进球
	xlsx.SetCellValue("Sheet1", "AG1", "必发卖出倾向")   //主队进球
	xlsx.SetCellValue("Sheet1", "AH1", "必发卖出和倾向")  //主队进球
	xlsx.SetCellValue("Sheet1", "AI1", "必发卖出客队倾向")
	xlsx.SetCellValue("Sheet1", "AJ1", "必发买入主队总量") //主队进球
	xlsx.SetCellValue("Sheet1", "AK1", "必发买入和总量")  //主队进球
	xlsx.SetCellValue("Sheet1", "AL1", "必发买入客队总量") //主队进球
	xlsx.SetCellValue("Sheet1", "AM1", "必发卖出主队总量") //主队进球
	xlsx.SetCellValue("Sheet1", "AN1", "必发卖出和总量")  //主队进球
	xlsx.SetCellValue("Sheet1", "AO1", "必发卖出客队总量") //主队进球

	// line := 'A'
	for linenumber := 1; linenumber < len(Copms)+1; linenumber++ {
		comps := Copms[linenumber-1]
		xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(linenumber+1), comps.VSID)
		xlsx.SetCellValue("Sheet1", "B"+strconv.Itoa(linenumber+1), comps.VSState)
		xlsx.SetCellValue("Sheet1", "C"+strconv.Itoa(linenumber+1), comps.VSTime)
		xlsx.SetCellValue("Sheet1", "D"+strconv.Itoa(linenumber+1), comps.HTeamID)
		xlsx.SetCellValue("Sheet1", "E"+strconv.Itoa(linenumber+1), comps.HTeamName)
		xlsx.SetCellValue("Sheet1", "F"+strconv.Itoa(linenumber+1), comps.HTeamIDRanking)
		xlsx.SetCellValue("Sheet1", "G"+strconv.Itoa(linenumber+1), comps.VTeamIDRanking)
		xlsx.SetCellValue("Sheet1", "H"+strconv.Itoa(linenumber+1), comps.VTeamID)
		xlsx.SetCellValue("Sheet1", "I"+strconv.Itoa(linenumber+1), comps.VTeamName)
		xlsx.SetCellValue("Sheet1", "J"+strconv.Itoa(linenumber+1), comps.BfWin)
		xlsx.SetCellValue("Sheet1", "K"+strconv.Itoa(linenumber+1), comps.BfHe)
		xlsx.SetCellValue("Sheet1", "L"+strconv.Itoa(linenumber+1), comps.BfLoss)
		xlsx.SetCellValue("Sheet1", "M"+strconv.Itoa(linenumber+1), comps.KailiWin)
		xlsx.SetCellValue("Sheet1", "N"+strconv.Itoa(linenumber+1), comps.KailHe)
		xlsx.SetCellValue("Sheet1", "O"+strconv.Itoa(linenumber+1), comps.KailiLoss)
		xlsx.SetCellValue("Sheet1", "P"+strconv.Itoa(linenumber+1), comps.HGetGoal)
		xlsx.SetCellValue("Sheet1", "Q"+strconv.Itoa(linenumber+1), comps.HLossGoal)
		xlsx.SetCellValue("Sheet1", "R"+strconv.Itoa(linenumber+1), comps.VGetGoal)
		xlsx.SetCellValue("Sheet1", "S"+strconv.Itoa(linenumber+1), comps.VLossGoal)
		xlsx.SetCellValue("Sheet1", "T"+strconv.Itoa(linenumber+1), comps.VSHGoal) //主队进球
		xlsx.SetCellValue("Sheet1", "U"+strconv.Itoa(linenumber+1), comps.VSVGoal) //主队进球

		xlsx.SetCellValue("Sheet1", "V"+strconv.Itoa(linenumber+1), comps.HEveyGoals)  //主队进球
		xlsx.SetCellValue("Sheet1", "W"+strconv.Itoa(linenumber+1), comps.VEveyGoals)  //主队进球
		xlsx.SetCellValue("Sheet1", "X"+strconv.Itoa(linenumber+1), comps.HWinPercent) //主队进球
		xlsx.SetCellValue("Sheet1", "Y"+strconv.Itoa(linenumber+1), comps.VWinPercent) //主队进球

		xlsx.SetCellValue("Sheet1", "Z"+strconv.Itoa(linenumber+1), comps.HDrawPercent)  //主队进球
		xlsx.SetCellValue("Sheet1", "AA"+strconv.Itoa(linenumber+1), comps.VDrawPercent) //主队进球
		xlsx.SetCellValue("Sheet1", "AB"+strconv.Itoa(linenumber+1), comps.HLossPercent) //主队进球
		xlsx.SetCellValue("Sheet1", "AC"+strconv.Itoa(linenumber+1), comps.VLossPercent) //主队进球

		xlsx.SetCellValue("Sheet1", "AD"+strconv.Itoa(linenumber+1), comps.BfWinWantBuy)   //主队进球
		xlsx.SetCellValue("Sheet1", "AE"+strconv.Itoa(linenumber+1), comps.BfDrawWantBuy)  //主队进球
		xlsx.SetCellValue("Sheet1", "AF"+strconv.Itoa(linenumber+1), comps.BfLossWantBuy)  //主队进球
		xlsx.SetCellValue("Sheet1", "AG"+strconv.Itoa(linenumber+1), comps.BfWinWantSell)  //主队进球
		xlsx.SetCellValue("Sheet1", "AH"+strconv.Itoa(linenumber+1), comps.BfDrawWantSell) //主队进球
		xlsx.SetCellValue("Sheet1", "AI"+strconv.Itoa(linenumber+1), comps.BfLossWantSell) //主队进球

		xlsx.SetCellValue("Sheet1", "AJ"+strconv.Itoa(linenumber+1), comps.BFWinBuyCount)   //主队进球
		xlsx.SetCellValue("Sheet1", "AK"+strconv.Itoa(linenumber+1), comps.BFDrawBuyCount)  //主队进球
		xlsx.SetCellValue("Sheet1", "AL"+strconv.Itoa(linenumber+1), comps.BFLossBuyCount)  //主队进球
		xlsx.SetCellValue("Sheet1", "AM"+strconv.Itoa(linenumber+1), comps.BFWinSellCount)  //主队进球
		xlsx.SetCellValue("Sheet1", "AN"+strconv.Itoa(linenumber+1), comps.BFDrawSellCount) //主队进球
		xlsx.SetCellValue("Sheet1", "AO"+strconv.Itoa(linenumber+1), comps.BFLossSellCount) //主队进球
	}
	err := xlsx.SaveAs("./0918.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

//获取联赛的所有球队比分
func GetScore() {
	// var url = ScoreHost + year + "/" + league + ".js?version=2018090813"
	// jscode := GET(url)
	f, _ := os.Open("orgdata.js")

	jscode, _ := ioutil.ReadAll(f)
	vm.Run(`var jh = new Object();`)
	vm.Run(string(jscode))
	teams, _ := vm.Get("arrTeam")
	len, _ := teams.Object().Get("length")
	count, _ := len.ToInteger()
	for index := 0; index < int(count); index++ {
		teamOrg, _ := teams.Object().Get(strconv.Itoa(index))
		teamStr, _ := teamOrg.ToString()
		team := strings.Split(teamStr, ",")
		Teams[team[0]] = team[1]
	}
	// fmt.Println(Teams)
	//Get Score
	jh, _ := vm.Get("jh")
	len, _ = jh.Object().Get("length")
	count, _ = len.ToInteger()

	for _, key := range jh.Object().Keys() {
		match, _ := jh.Object().Get(key)

		for _, inkey := range match.Object().Keys() {
			content, _ := match.Object().Get(inkey)
			contentOrg, _ := content.ToString()
			// fmt.Println(contentOrg)
			contentStr := strings.Split(contentOrg, ",")
			// if "0" == contentStr[2] {
			// 	continue
			// }
			vsid := contentStr[0]
			// vsState := contentStr[2]
			vstime := contentStr[3]
			matchTime, _ := time.Parse("2006-01-02 15:04", vstime)

			if matchTime.After(StopTime) {
				break
			}
			hteamid := contentStr[4]
			vteamid := contentStr[5]
			vsresult := contentStr[6]

			hteamranking := contentStr[8]
			vteamranking := contentStr[9]

			Copms = append(Copms, Competition{VSID: vsid, VSTime: vstime,
				VSResult: vsresult, HTeamID: hteamid, VTeamID: vteamid,
				VTeamName: Teams[hteamid], HTeamName: Teams[vteamid],
				VTeamIDRanking: vteamranking, HTeamIDRanking: hteamranking,
			})
		}
	}
	// for index := 0; index < int(count); index++ {
	// 	teamOrg, _ := teams.Object().Get(strconv.Itoa(index))
	// 	teamStr, _ := teamOrg.ToString()
	// 	team := strings.Split(teamStr, ",")
	// 	Teams[team[0]] = team[1]
	// }

}

func (c *Competition) GetGoal() {
	var url = GoalHost + c.VSID + ".htm"
	cookie := `UM_distinctid=165c3525acc854-09c2457c51505e-34677908-13c680-165c3525acd67c; fAnalyCookie=null; analysis_set=0#1#2#3#4#5,6#7#8#9#10#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#28^3#1#8#12#4; Bet007EuropeIndex_Cookie=null; bfWin007FirstMatchTime=2018,8,15,08,00,00; win007BfCookie=0^0^1^1^1^1^1^0^0^0^0^0^1^2^1^1^0^1^1; CNZZDATA1831853=cnzz_eid%3D915603233-1537056114-http%253A%252F%252Fwww.win0168.com%252F%26ntime%3D1537056114; CNZZDATA1274632882=187442473-1536973314-null%7C1537060011`
	html := GET(url, cookie)
	htmlcode := strings.NewReader(html)

	doc, _ := goquery.NewDocumentFromReader(htmlcode)
	// fmt.Println(doc.Find("script").Text())
	jshtml := doc.Find("script").Text()
	var exp, _ = regexp.Compile(`var h_data=([\s\S]*]];\n )`)
	var h2h_homeexp, _ = regexp.Compile(`var h2h_home = ([\d]*;\n)`)
	var h2h_awayexp, _ = regexp.Compile(`var h2h_away = ([\d]*;\n)`)

	jsdata := exp.FindStringSubmatch(jshtml)
	h2h_homejsdata := h2h_homeexp.FindStringSubmatch(jshtml)
	h2h_awayjsdata := h2h_awayexp.FindStringSubmatch(jshtml)

	vm.Run(jsdata[0])
	vm.Run(h2h_homejsdata[0])
	vm.Run(h2h_awayjsdata[0])

	// fmt.Print(vm.Get("h_data"))

	vm.Run(`
    var HEveryGoal = 0,
        VEveryGoal = 0,
        HWinPercent = 0,
        HDrawPercent = 0,
        HLossPercent = 0,
        VWinPercent = 0,
        VDrawPercent = 0,
        VLossPercent = 0;


    var count = "10"
    for (var index = 1; index < 3; index++) {
		var data = [];
		var id = ""
        if (index == 1) {
			id='h'
            data = h_data;
        } else {
			id='a'
            data = a_data;
        }
        if (data.length < 10) {
            count = data.length
        }

        var allGoal = 0,
            allLossGoal = 0,
            winNum = 0,
            drawNum = 0,
            lossNum = 0,
            hgCount = 0,
            hgGoal = 0,
            hgLossGoal = 0,
            hgWin = 0,
            hgDraw = 0,
            hgLoss = 0;
        for (var a = 0; a < count; a++) {
            if ((id == "h" && h2h_home == data[a][4]) || (id == "a" && h2h_away == data[a][4])) {
                allGoal += parseInt(data[a][8]);
                allLossGoal += parseInt(data[a][9]);
                if (parseInt(data[a][8]) > parseInt(data[a][9])) {
                    winNum += 1;
                    if (id == "h")
                        hgWin += 1;
                } else if (parseInt(data[a][8]) == parseInt(data[a][9])) {
                    drawNum += 1;
                    if (id == "h")
                        hgDraw += 1;
                } else {
                    lossNum += 1;
                    if (id == "h")
                        hgLoss += 1;
                }
                if (id == "h") {
                    hgCount += 1;
                    hgGoal += parseInt(data[a][8]);
                    hgLossGoal += parseInt(data[a][9]);
                }
            } else if ((id == "h" && h2h_home == data[a][6]) || (id == "a" && h2h_away == data[a][6])) {
                allGoal += parseInt(data[a][9]);
                allLossGoal += parseInt(data[a][8]);
                if (parseInt(data[a][9]) > parseInt(data[a][8])) {
                    winNum += 1;
                    if (id == "a")
                        hgWin += 1;
                } else if (parseInt(data[a][8]) == parseInt(data[a][9])) {
                    drawNum += 1;
                    if (id == "a")
                        hgDraw += 1;
                } else {
                    lossNum += 1;
                    if (id == "a")
                        hgLoss += 1;
                }
                if (id == "a") {
                    hgCount += 1;
                    hgGoal += parseInt(data[a][9]);
                    hgLossGoal += parseInt(data[a][8]);
                }
            }
        }
        // var tr = document.getElementById("tr_com_" + id);
        // tr.cells[1].innerHTML = allGoal;
        // tr.cells[2].innerHTML = allLossGoal;
        // tr.cells[3].innerHTML = allGoal - allLossGoal;
        if (index == 1) {
            HEveryGoal = count > 0 ? Math.round(allGoal / count * 100) / 100 : 0;
            HWinPercent = count > 0 ? Math.round(winNum / count * 1000) / 10 + "%" : "";
            HDrawPercent = count > 0 ? Math.round(drawNum / count * 1000) / 10 + "%" : "";
            HLossPercent = count > 0 ? Math.round(lossNum / count * 1000) / 10 + "%" : "";

        } else {
            VEveryGoal = count > 0 ? Math.round(allGoal / count * 100) / 100 : 0;
            VWinPercent = count > 0 ? Math.round(winNum / count * 1000) / 10 + "%" : "";
            VDrawPercent = count > 0 ? Math.round(drawNum / count * 1000) / 10 + "%" : "";
            VLossPercent = count > 0 ? Math.round(lossNum / count * 1000) / 10 + "%" : "";
        }

        // tr.cells[7].innerHTML = count > 0 ? Math.round(lossNum / count * 1000) / 10 + "%" : "";
        // tr.cells[8].innerHTML = hgGoal;
        // tr.cells[9].innerHTML = hgLossGoal;
        // tr.cells[10].innerHTML = hgGoal - hgLossGoal;
        //VEveryGoal = count > 0 ? Math.round(hgGoal / hgCount * 100) / 100 : "0";
        //tr.cells[12].innerHTML = count > 0 ? Math.round(hgWin / hgCount * 1000) / 10 + "%" : "";
        //tr.cells[13].innerHTML = count > 0 ? Math.round(hgDraw / hgCount * 1000) / 10 + "%" : "";
        //tr.cells[14].innerHTML = count > 0 ? Math.round(hgLoss / hgCount * 1000) / 10 + "%" : "";

	}
	`)

	// fmt.Println(vm.Get("HEveryGoal"))

	//排名和丢失球
	doc.Find("#porlet_5").Find("tr").Find("tr").Find("td").Each(func(i int, td *goquery.Selection) {
		if 17 == i {
			c.HGetGoal = td.Text()
		}
		if 18 == i {
			c.HLossGoal = td.Text()
		}
		if 21 == i {
			c.HTeamIDRanking = td.Text()
		}
		if 73 == i {
			c.VGetGoal = td.Text()
		}
		if 74 == i {
			c.VLossGoal = td.Text()
		}
		if 77 == i {
			c.VTeamIDRanking = td.Text()
		}
	})
	// fmt.Println(doc.Find("#porlet_11").Html())
	//胜率
	if HEveyGoals, err := vm.Get("HEveryGoal"); err == nil {
		c.HEveyGoals = HEveyGoals.String()
	}
	if HWinPercent, err := vm.Get("HWinPercent"); err == nil {
		c.HWinPercent = HWinPercent.String()
	}
	if HLossPercent, err := vm.Get("HLossPercent"); err == nil {
		c.HLossPercent = HLossPercent.String()
	}
	if HDrawPercent, err := vm.Get("HDrawPercent"); err == nil {
		c.HDrawPercent = HDrawPercent.String()
	}

	if VEveyGoals, err := vm.Get("VEveryGoal"); err == nil {
		c.VEveyGoals = VEveyGoals.String()
	}
	if VWinPercent, err := vm.Get("VWinPercent"); err == nil {
		c.VWinPercent = VWinPercent.String()
	}
	if VLossPercent, err := vm.Get("VLossPercent"); err == nil {
		c.VLossPercent = VLossPercent.String()
	}
	if VDrawPercent, err := vm.Get("VDrawPercent"); err == nil {
		c.VDrawPercent = VDrawPercent.String()
	}

}

//获取必发指数
// func (c *Competition) GetBFIndex() {
// 	// url := MatchDetailHost + VSID + ".htm"
// 	// doc, _ := goquery.NewDocument(url)
// 	seq := strconv.FormatInt(time.Now().UnixNano(), 10)
// 	url := BFIndex + c.VSID + ".htm?" + seq
// 	cookie := `UM_distinctid=165b7a309d24d9-06976110957a49-34677908-13c680-165b7a309d31c8; CNZZDATA1261430177=1607465902-1536382542-http%253A%252F%252Fzq.win007.com%252F%7C1536382542; fAnalyCookie=null; CNZZDATA1831853=cnzz_eid%3D715785739-1536386345-http%253A%252F%252Fzq.win007.com%252F%26ntime%3D1536386345; analysis_set=0#1#2#3#4#5,6#7#8#9#10#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#28^3#1#8#12#4`
// 	htmlCode := GET(url, cookie)
// 	htmlReader, _ := html.Parse(strings.NewReader(htmlCode))
// 	doc := goquery.NewDocumentFromNode(htmlReader)
// 	doc.Find("#analy_BetfaStandard .odds").Each(func(j int, input *goquery.Selection) {
// 		bfindex := input.Children().Eq(input.Children().Length() - 3).Text()
// 		if "" == c.BfWin {
// 			c.BfWin = bfindex
// 		} else if "" == c.BfHe {
// 			c.BfHe = bfindex
// 		} else {
// 			c.BfLoss = bfindex
// 		}
// 	})
// }

func (c *Competition) GetBFData() {
	url := "http://vip.win0168.com/betfa/single.aspx?id=" + c.VSID
	cookie := `UM_distinctid=165c3525acc854-09c2457c51505e-34677908-13c680-165c3525acd67c; CNZZDATA1274585981=916050297-1536972537-http%253A%252F%252Fwww.win0168.com%252F%7C1536972537; Bet007EuropeIndex_Cookie=null; Bet007Odds_Cookie=0^0^0^1^1^0^0; Bet007live_hiddenID=_1552192_1552185_1583226_1583221_1583227_1583225_1560947_1560944_1585085_1585088_1585090_1585093_1585089_1585087_1549717_1549712_1549715_1511797_1511798_1600021_1600024_1600019_1562117_1562120_1562121_1584031_1584027_1584025_1584030_1584028_1584026_1555672_1555671_1555670_1551888_1551889_1551890_1551891_1551887_1559002_1558995_1558997_1559327_1559328_1559326_1600877_1600874_1600997_1600879_1489037_1489038_1489040_1489042_1489043_1489036_1489260_1489261_1489266_1496838_1500427_1500478_1500433_1500425_1500431_1500430_1500429_1500428_1584503_1584494_1584493_1584504_1584505_1510018_1510010_1510012_1510013_1510015_1510014_1553298_1582377_1488037_1488039_1500680_1500678_1500679_1489576_1489578_1489581_1489582_1489577_1489579_1489580_1503894_1503893_1503895_1550738_1550735_1550737_1550734_1565864_1565865_1565866_1565867_1565868_1565869_1559645_1517503_1551332_1551331_1560702_1560701_1550922_1550919_1562810_1562812_1562808_1557067_1557068_1557070_1557071_1555927_1555929_1567333_1567329_1503527_1503528_1503529_1503524_1570104_1570101_1613270_1613273_1613271_1613272_1613269_1510873_1525964_1525968_1525962_1572756_1572757_1572749_1572751_1572754_1563222_1574388_1574394_1574392_1574393_1574702_1574704_1574696_1574699_1618654_1555470_1557303_1557301_1617427_1617428_1617429_1575627_1575628_1575630_1575632_1575633_1588401_1588402_1621381_1572198_1572199_1577169_1577168_1577174_1577172_1577173_1577170_1541675_1541674_1621397_1621398_1621399_1601447_1601460_1601480_1591236_1591235_1554740_1554736_1554735_1554738_1554737_1554739_1581848_1581849_1581850_1593306_1593307_1567857_1567863_1567856_1567860_1567861_1567858_1567859_1605294_1605293_1605295_1605291_1605289_1575316_1575319_1500972_1500973_1500970_1500971_1576497_1621662_1621663_; win007BfCookie=0^0^1^1^1^1^1^0^0^0^0^0^1^2^1^1^0^1^1; FS007Filter=1^0^_36_31_8_34_11_60_40_9_33_5_16_10_3_32_22_26_25_21_2_4_17_138_122_15_123_284_7_127_128_61_766_118_6_119_137_129_124_166_159_242_358_230_250_135_30_130_133_131_505_240_308_354_391_460_466_469_504_593_613_140_423_596_700_847_1366_; bfWin007FirstMatchTime=2018,8,18,08,00,00; CNZZDATA1261430177=1073844010-1536981947-http%253A%252F%252Finfo.win0168.com%252F%7C1537275560`
	htmlCode := GET(url, cookie)
	htmlReader, _ := html.Parse(strings.NewReader(htmlCode))
	doc := goquery.NewDocumentFromNode(htmlReader)
	doc.Find("tr").Each(func(i int, tr *goquery.Selection) {
		switch i {
		case 3: //主队
			tr.Find("td").Each(func(j int, td *goquery.Selection) {
				switch j {
				case 10: //庄稼盈亏
					c.BfWin = td.Text()
				case 13: //购买倾向
					c.BfWinWantBuy = td.Text()
				case 14:
					c.BfWinWantSell = td.Text()
				}
			})
		case 4: //和
			tr.Find("td").Each(func(j int, td *goquery.Selection) {
				switch j {
				case 8: //庄稼盈亏
					c.BfHe = td.Text()
				case 11: //购买倾向
					c.BfDrawWantBuy = td.Text()
				case 12:
					c.BfDrawWantSell = td.Text()
				}
			})
		case 5: //客
			tr.Find("td").Each(func(j int, td *goquery.Selection) {
				switch j {
				case 8: //庄稼盈亏
					c.BfLoss = td.Text()
				case 11: //购买倾向
					c.BfLossWantBuy = td.Text()
				case 12:
					c.BfLossWantSell = td.Text()
				}
			})
		case 21: //客
			tr.Find("td").Each(func(j int, td *goquery.Selection) {
				switch j {
				case 2: //庄稼盈亏
					c.BFWinBuyCount = td.Text()
				case 4: //购买倾向
					c.BFDrawBuyCount = td.Text()
				case 6:
					c.BFLossBuyCount = td.Text()
				}
			})
		case 22: //客
			tr.Find("td").Each(func(j int, td *goquery.Selection) {
				switch j {
				case 2: //庄稼盈亏
					c.BFWinSellCount = td.Text()
				case 4: //购买倾向
					c.BFDrawSellCount = td.Text()
				case 6:
					c.BFLossSellCount = td.Text()
				}
			})
		}
	})

	//必发交易量
	// doc.Find(selector)
}

//获取主流公司凯利数据
// func (c *Competition) GetKaili() {
// 	jscode := GET(KailiHost+c.VSID+".js", "")

// 	vm.Run(jscode)
// 	games, _ := vm.Get("game")

// 	// gameCount := len(games)
// 	// fmt.Printf("一共家%d公司", gameCount)
// 	var top float64 = 31
// 	var win float64
// 	var he float64
// 	var loss float64

// 	for index := 0; top > 0; index++ {
// 		var winkaili, hekaili, losskaili float64
// 		game, _ := games.Object().Get(strconv.Itoa(index))
// 		gameStr, _ := game.ToString()
// 		kailiIndes := strings.Split(gameStr, "|")
// 		if len(kailiIndes) > 19 {
// 			winkaili, _ = strconv.ParseFloat(kailiIndes[17], 64)
// 			hekaili, _ = strconv.ParseFloat(kailiIndes[18], 64)
// 			losskaili, _ = strconv.ParseFloat(kailiIndes[19], 64)
// 		}
// 		// fmt.Println(strings.Split(gameStr, "|")[17], strings.Split(gameStr, "|")[18], strings.Split(gameStr, "|")[19], winbf, hebf, lossbf)

// 		win += winkaili
// 		he += hekaili
// 		loss += losskaili
// 		top--

// 	}

// 	c.KailiWin, c.KailHe, c.KailiLoss = strconv.FormatFloat(win/31, 'f', 10, 64), strconv.FormatFloat(he/31, 'f', 10, 64), strconv.FormatFloat(loss/31, 'f', 10, 64)
// }
func GET(url string, cookie string) string {
	fmt.Println("url:", url)
	client := &http.Client{Timeout: time.Second * 1}
	reqest, err := http.NewRequest("GET", url, nil) //建立一个请求
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	//Add 头协议
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	reqest.Header.Add("Accept-Language", "ja,zh-CN;q=0.8,zh;q=0.6")
	reqest.Header.Add("Connection", "keep-alive")
	reqest.Header.Add("Cookie", cookie)
	reqest.Header.Add("Referer", url)
	reqest.Header.Add("Host", "info.win0168.com")

	// reqest.Header.Add("Cookie", "UM_distinctid=16574bdb87b90d-09607a7cf6df6a-34677908-13c680-16574bdb87c53f; Bet007EuropeIndex_Cookie=0^1^1^1; win007BfCookie=null; bfWin007FirstMatchTime=2018,8,7,08,00,0")
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.92 Safari/537.36")
	response, err := client.Do(reqest) //提交

	if nil != err {
		fmt.Errorf("Get Error%v", err)
	}
	// defer response.Body.Close()
	// cookies := response.Cookies() //遍历cookies
	// for _, cookie := range cookies {
	// 	fmt.Println("cookie:", cookie)
	// }

	body, err1 := ioutil.ReadAll(response.Body)
	if err1 != nil {
		// handle error
		// fmt.Println("read body", err1)
	}
	return string(body) //网页源码

}
