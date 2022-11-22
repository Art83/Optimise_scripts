package funcs

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type json_obj_steps struct {
	Steps  float64 `json:"step_count"`
	Start  string  `json:"start_time"`
	Finish string  `json:"end_time"`
}

func WriteCsvSteps(file_name string, id []string, payload []string, date []string) {

	csvFile, err := os.Create(file_name)
	if err != nil {
		log.Fatalf("failed creating the file: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)

	row := []string{"participant_id", "step_count", "start_time", "end_time", "date"}
	err = csvwriter.Write(row)
	if err != nil {
		log.Fatalf("Writing of the header into the csv went wrong: %s", err)
	}
	fmt.Println("Writing steps into the csv file")
	for i := 1; i < len(id); i++ {
		jsonRes := json_obj_steps{}
		err = json.Unmarshal([]byte(payload[i]), &jsonRes)
		if err != nil {
			log.Fatalf("Unmarshal went wrong: %s", err)
		}
		row := []string{id[i], fmt.Sprintf("%v", jsonRes.Steps), fmt.Sprintf("%v", jsonRes.Start), fmt.Sprintf("%v", jsonRes.Finish), date[i]}
		err = csvwriter.Write(row)
		if err != nil {
			log.Fatalf("Writing into the csv went wrong: %s", err)
		}
	}
	csvwriter.Flush()
	fmt.Println("Writing the steps data into the csv file is finished")
	fmt.Println("---------------------")
}
