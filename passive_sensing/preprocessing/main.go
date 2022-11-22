package main

import (
	"fmt"
	"funcs/funcs"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		ids_loc, payloads_loc, dates_loc := funcs.ReadCsv("data/location_mt11.csv")
		funcs.WriteCsvLocation("output/location_new_mt11.csv", ids_loc, payloads_loc, dates_loc)
		wg.Done()
	}()
	wg.Add(1)

	go func() {
		ids_act, payloads_act, dates_act := funcs.ReadCsv("data/activity_mt11.csv")
		funcs.WriteCsvActivity("output/activity_new_mt11.csv", ids_act, payloads_act, dates_act)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		ids_bat, payloads_bat, dates_bat := funcs.ReadCsv("data/battery_mt11.csv")
		funcs.WriteCsvBattery("output/battery_new_mt11.csv", ids_bat, payloads_bat, dates_bat)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		ids_dist, payloads_dist, dates_dist := funcs.ReadCsv("data/distance_mt11.csv")
		funcs.WriteCsvDistance("output/distance_new_mt11.csv", ids_dist, payloads_dist, dates_dist)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		ids_st, payloads_st, dates_st := funcs.ReadCsv("data/steps_mt11.csv")
		funcs.WriteCsvSteps("output/steps_new_mt11.csv", ids_st, payloads_st, dates_st)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("End of the job")
}
