package main

import (
	"fmt"
	"funcs/funcs"
)

func main() {
	ids, payloads, dates := funcs.ReadCsv("data/location_mt11.csv")
	funcs.WriteCsvLocation("output/location_new_mt11.csv", ids, payloads, dates)
	ids, payloads, dates = funcs.ReadCsv("data/activity_mt11.csv")
	funcs.WriteCsvActivity("output/activity_new_mt11.csv", ids, payloads, dates)
	ids, payloads, dates = funcs.ReadCsv("data/battery_mt11.csv")
	funcs.WriteCsvBattery("output/battery_new_mt11.csv", ids, payloads, dates)
	ids, payloads, dates = funcs.ReadCsv("data/distance_mt11.csv")
	funcs.WriteCsvDistance("output/distance_new_mt11.csv", ids, payloads, dates)
	ids, payloads, dates = funcs.ReadCsv("data/steps_mt11.csv")
	funcs.WriteCsvSteps("output/steps_new_mt11.csv", ids, payloads, dates)
	fmt.Println("End of the job")
}
