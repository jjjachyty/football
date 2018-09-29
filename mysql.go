package main

import (
	"database/sql"
	"fmt"

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
	fmt.Println("InserMatch-----------")
	stmt, err := DB.Prepare("insert into matchs values(?,?,?,?,?,?,?,?)")
	if nil == err {
		for _, match := range matchs {
			_, err := stmt.Exec(match.ID, match.HName, match.VName, match.Score, match.LeagueMatch, match.Rotation, match.Date)
			if nil != err {
				panic("insert into match error")
			}
		}

	} else {
		fmt.Println(err)
	}
}

func InserOuZhi(ouzhis []OuZhi) {
	stmt, err := DB.Prepare("insert into ouzhi values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	if nil == err {
		for _, ouzhi := range ouzhis {
			_, err := stmt.Exec(ouzhi.ID, ouzhi.Name,
				ouzhi.BeginWinRate,
				ouzhi.BeginWinPD,
				ouzhi.BeginDrawRate,
				ouzhi.BeginDrawPD,
				ouzhi.BeginLossRate,
				ouzhi.BeginLossPD,
				ouzhi.BeginReturn,
				ouzhi.BeginWinKaili,
				ouzhi.BeginDrawKaili,
				ouzhi.BeginLossKaili,
				ouzhi.EndWinRate,
				ouzhi.EndWinPD,
				ouzhi.EndDrawRate,
				ouzhi.EndDrawPD,
				ouzhi.EndLossRate,
				ouzhi.EndDrawPD,
				ouzhi.EndLossRate,
				ouzhi.EndLossPD,
				ouzhi.EndReturn,
				ouzhi.EndWinKaili,
				ouzhi.EndDrawKaili,
				ouzhi.EndLossKaili,
			)
			if nil != err {
				panic(err)
			}
		}

	} else {
		fmt.Println(err)
	}
}

func InserYaZhi(yazhis []YaZhi) {
	stmt, err := DB.Prepare("insert into yazhi values(?,?,?,?,?,?,?,?,?,?,?)")
	if nil == err {
		for _, yazhi := range yazhis {
			_, err := stmt.Exec(yazhi.ID, yazhi.Name, yazhi.BeginHWater, yazhi.BeginPan, yazhi.BeginVWater, yazhi.EndHWater, yazhi.EndHWaterDirection, yazhi.EndPan, yazhi.EndPanState, yazhi.EndVWater, yazhi.EndVWaterDirection)
			if nil != err {
				panic("insert into yazhi error")
			}
		}

	} else {
		fmt.Println(err)
	}
}
func InserBiFar(bifas []TouZhu) {
	stmt, err := DB.Prepare("insert into bifar values(?,?,?,?,?,?,?)")
	if nil == err {
		for _, bifa := range bifas {
			_, err := stmt.Exec(bifa.ID, bifa.DealerWin, bifa.DealerDraw, bifa.DealerLoss, bifa.WinIndex, bifa.DrawIndex, bifa.LossIndex)
			if nil != err {
				panic(err)
			}
		}

	} else {
		fmt.Println(err)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}
}
