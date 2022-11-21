package funcs

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type json_obj_battery struct {
	Level float64 `json:"battery_level"`
	State string  `json:"battery_state"`
}

func WriteCsvBattery(file_name string, id []string, payload []string, date []string) {

	csvFile, err := os.Create(file_name)
	if err != nil {
		log.Fatalf("failed creating the file: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)

	i := 0
	row := []string{"participant_id", "battery_level", "battery_state", "confidence", "date"}
	err = csvwriter.Write(row)
	if err != nil {
		log.Fatalf("Writing of the header into the csv went wrong: %s", err)
	}
	fmt.Println("Writing battery into the csv file")
	for i < len(id) {
		jsonRes := json_obj_battery{}
		err = json.Unmarshal([]byte(payload[i]), &jsonRes)
		if err != nil {
			log.Fatalf("Unmarshal went wrong: %s", err)
		}
		row := []string{id[i], fmt.Sprintf("%v", jsonRes.Level), fmt.Sprintf("%v", jsonRes.State), date[i]}
		err = csvwriter.Write(row)
		if err != nil {
			log.Fatalf("Writing into the csv went wrong: %s", err)
		}
		i++
	}
	csvwriter.Flush()
	fmt.Println("Writing the battery data into the csv file is finished")
	fmt.Println("---------------------")
}
