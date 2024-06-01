// STEP01: ファイルを分けよう

package main

import (
	"fmt"
	"os"

	"github.com/mochi22/gacha_go_test/gacha"
)

func main() {
	// TODO: ガチャ券10枚、コイン100枚を持ったプレイヤーを作る
	p := gacha.NewPlayer(10, 100)

	n := inputN(*p)
	// TODO: gacha.DrawN関数を呼び、変数resultsとsummaryに結果を代入する
	results, summary := gacha.DrawN(p, n)
	// fmt.Println(results)
	// fmt.Println(summary)
	saveResults(results)
	saveSummary(summary)
}

// TODO: 引数の型をgacha.Playerのポインタにする
func inputN(p gacha.Player) int {

	max := p.DrawableNum()
	fmt.Printf("ガチャを引く回数を入力してください（最大:%d回）\n", max)

	var n int
	for {
		fmt.Print("ガチャを引く回数>")
		fmt.Scanln(&n)
		if n > 0 && n <= max {
			break
		}
		fmt.Printf("1以上%d以下の数を入力してください\n", max)
	}

	return n
}

func saveResults(results []*gacha.Card) {
	// TODO: results.txtというファイルを作成する

	f, err := os.Create("results.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	for _, result := range results {
		// TODO: fmt.Fprintln関数を使ってresultをファイルに出力する
		_, err = fmt.Fprintln(f, result)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func saveSummary(summary map[gacha.Rarity]int) {
	f, err := os.Create("summary.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		// TODO: ファイルを閉じる
		// エラー発生した場合はfmt.Println関数で出力する
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}

	}()

	for rarity, count := range summary {
		fmt.Fprintf(f, "%s %d\n", rarity.String(), count)
	}
}
