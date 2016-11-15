package bills

import (
	"fmt"

	"github.com/barnybug/gohome/config"
	"github.com/barnybug/gohome/lib/graphite"
	"github.com/barnybug/gohome/services"
)

func ExampleElectricityBill() {
	services.Config = config.ExampleConfig
	response := `
[{"datapoints":[[null,1479078900],[208.74974193762046,1479082500],[169.99837768305133,1479086100],[213.35026333908445,1479089700],[170.259225809913,1479093300],[215.70170907307784,1479096900],[178.71688191271733,1479100500],[222.24017699861088,1479104100],[392.65690806828025,1479107700],[409.2159394261653,1479111300],[411.04627848254495,1479114900],[357.10301162358974,1479118500],[366.9297715598914,1479122100],[336.8961499991692,1479125700],[969.1057201362746,1479129300],[658.341352387386,1479132900],[554.1869583208008,1479136500],[517.6474871332375,1479140100],[583.4296880857364,1479143700],[597.6046701503383,1479147300],[464.14261924889433,1479150900],[424.04454031389287,1479154500],[377.0063398102502,1479158100],[451.52215288043044,1479161700],[261.19681200471496,1479165300]],"target":"derivative(smartSummarize(sensor.power.total.avg, \"1h\", \"last\"))"}]
`
	gr = &graphite.MockGraphite{Response: response}
	s := electricityBill()
	fmt.Println(s)
	// Output:
	// Electricity: yesterday I used 9.51 kwh (5.78 day / 3.73 night), costing £1.00. Peak was around 1:15PM.
}
