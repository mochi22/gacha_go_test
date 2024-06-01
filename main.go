// STEP01: ファイルを分けよう

package main

import (
	"fmt"
	"os"

	"github.com/mochi22/gacha_go_test/gacha"
)

func main() {
	p := gacha.NewPlayer(10, 100)

	n := inputN(*p)
	results, summary := gacha.DrawN(p, n)
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
	f, err := os.Create("results.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	defer func() {
		if err := f.Close(); err != nil {
			//標準エラー出力にerrを出力する
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	for _, result := range results {
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
		os.Exit(1)
	}

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	}()

	for rarity, count := range summary {
		fmt.Fprintf(f, "%s %d\n", rarity.String(), count)
	}
}
