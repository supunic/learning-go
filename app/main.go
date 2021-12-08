package main

import (
	"fmt"
	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func main() {
	spy, _ := quote.NewQuoteFromYahoo(
		"spy", "2018-01-01", "2019-01-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	// talib ... 株価の計算で使用するパッケージ
	// rsi ... テクニカルインディケータ（株が売られすぎか、買われすぎかを見る指標）
	// spy.Close ... 1日の株価の終値
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
	// Ema ... ムービングアベレージ（株価の平均）の1つ
	mva := talib.Ema(spy.Close, 14)
	fmt.Println(mva)
}
