package bloblang

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"github.com/benthosdev/benthos/v4/public/bloblang"
)

type Price struct {
	Date         string  `json:"date"`
	Open         float64 `json:"open"`
	High         float64 `json:"high"`
	Low          float64 `json:"low"`
	Close        float64 `json:"close"`
	Volume       int64   `json:"volume"`
	Dividends    float64 `json:"dividends"`
	StockSplits  float64 `json:"stocksplits"`
	CapitalGains float64 `json:"capitalgains"`
}

func init() {
	bleachSpec := bloblang.NewPluginSpec().
		Category("Bleach/Filter Data").
		Description("")

	if err := bloblang.RegisterMethodV2("bleach", bleachSpec,
		func(args *bloblang.ParsedParams) (bloblang.Method, error) {
			return bloblang.BytesMethod(func(s []byte) (interface{}, error) {
                start:=time.Now()
                defer func(){
                    log.Println("Bloblang Take Time:", time.Since(start))
                }()
				var day_price Price
				json.Unmarshal(s, &day_price)
				ss, err := json.Marshal(day_price)
                if err != nil {
                    return nil, err
                }
				return fmt.Sprintf("%s", string(ss)), nil
			}), nil
		},
	); err != nil {
		panic(err)
	}
}
