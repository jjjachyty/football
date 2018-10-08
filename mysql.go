package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:q1w2e3r4@tcp(106.12.10.77:3306)/football")
	check(err)
	DB = db
	fmt.Println("DB init")
}

func InserMatch(matchs []Match) {
	stmt, err := DB.Prepare("insert into matchs values(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if nil == err {
		for _, match := range matchs {
			scors := strings.Split(match.Score, ":")
			if len(scors) > 1 {
				_, err := stmt.Exec(match.Date, match.ID, match.HName, match.VName, match.LeagueMatch, match.Rotation, match.HRanking, match.HScoring, match.VRanking, match.VScoring, scors[0], scors[1], "")
				if nil != err {
					// panic(err)
					continue

					// panic(err)
				}
			}
		}

	} else {
		fmt.Println("InserMatch:", err)
	}
}

func InserOuZhi(ouzhis []OuZhi) {
	stmt, err := DB.Prepare("insert into ouzhi values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if nil == err {
		for _, ouzhi := range ouzhis {
			if "" == ouzhi.BeginWinRate {
				continue
			}
			_, err := stmt.Exec(ouzhi.ID, ouzhi.Name,
				ouzhi.BeginWinRate,
				strings.TrimRight(ouzhi.BeginWinPD, "%"),
				ouzhi.BeginDrawRate,
				strings.TrimRight(ouzhi.BeginDrawPD, "%"),
				ouzhi.BeginLossRate,
				strings.TrimRight(ouzhi.BeginLossPD, "%"),
				strings.TrimRight(ouzhi.BeginReturn, "%"),
				ouzhi.BeginWinKaili,
				ouzhi.BeginDrawKaili,
				ouzhi.BeginLossKaili,
				ouzhi.EndWinRate,
				strings.TrimRight(ouzhi.EndWinPD, "%"),
				ouzhi.EndDrawRate,
				strings.TrimRight(ouzhi.EndDrawPD, "%"),
				ouzhi.EndLossRate,
				strings.TrimRight(ouzhi.EndLossPD, "%"),
				strings.TrimRight(ouzhi.EndReturn, "%"),
				ouzhi.EndWinKaili,
				ouzhi.EndDrawKaili,
				ouzhi.EndLossKaili,
				"",
				"",
				"",
			)
			if nil != err {
				continue
				// fmt.Println(err)
				// fmt.Println("---", ouzhi.BeginWinRate)
				panic(err)

			}
		}

	} else {
		fmt.Println("InserOuZhi:", err)
	}
}

func InserYaZhi(yazhis []YaZhi) {
	stmt, err := DB.Prepare("insert into yazhi values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if nil == err {
		for _, yazhi := range yazhis {
			_, err := stmt.Exec(yazhi.ID, yazhi.Name, yazhi.BeginHWater, yazhi.BeginPan, yazhi.BeginVWater, yazhi.EndHWater, yazhi.EndHWaterDirection, yazhi.EndPan, yazhi.EndPanState, yazhi.EndVWater, yazhi.EndVWaterDirection, "", "", "")
			if nil != err {
				// fmt.Println()
				// panic(err)
				continue
			}
		}

	} else {
		fmt.Println("InserYaZhi:", err)
	}
}
func InserBiFar(bifas []TouZhu) {
	stmt, err := DB.Prepare("insert into bifar values(?,?,?,?,?,?,?,?)")
	if nil == err {
		for _, bifa := range bifas {
			_, err := stmt.Exec(bifa.ID,
				strings.Replace(bifa.DealerWin, ",", "", -1),
				strings.Replace(bifa.DealerDraw, ",", "", -1),
				strings.Replace(bifa.DealerLoss, ",", "", -1),
				bifa.WinIndex, bifa.DrawIndex, bifa.LossIndex, "")
			if nil != err {
				// panic(err)
				continue

			}
		}

	} else {
		fmt.Println("InserBiFar:", err)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}
}
