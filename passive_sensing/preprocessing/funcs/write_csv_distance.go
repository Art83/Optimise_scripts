package funcs

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type json_obj_distance struct {
	Distance float64 `json:"distance_metres"`
	Start    string  `json:"start_time"`
	Finish   string  `json:"end_time"`
}

func WriteCsvDistance(file_name string, id []string, payload []string, date []string) {

	csvFile, err := os.Create(file_name)
	if err != nil {
		log.Fatalf("failed creating the file: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)

	row := []string{"participant_id", "distance_metres", "start_time", "end_time", "date"}
	err = csvwriter.Write(row)
	if err != nil {
		log.Fatalf("Writing of the header into the csv went wrong: %s", err)
	}
	fmt.Println("Writing distance into the csv file")
	for i := 1; i < len(id); i++ {
		jsonRes := json_obj_distance{}
		err = json.Unmarshal([]byte(payload[i]), &jsonRes)
		if err != nil {
			log.Fatalf("Unmarshal went wrong: %s", err)
		}
		row := []string{id[i], fmt.Sprintf("%v", jsonRes.Distance), fmt.Sprintf("%v", jsonRes.Start), fmt.Sprintf("%v", jsonRes.Finish), date[i]}

		err = csvwriter.Write(row)
		if err != nil {
			log.Fatalf("Writing into the csv went wrong: %s", err)
		}
	}
	csvwriter.Flush()
	fmt.Println("Writing the distance data into the csv file is finished")
	fmt.Println("---------------------")
}
