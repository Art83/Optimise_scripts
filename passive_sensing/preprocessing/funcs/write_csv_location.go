package funcs

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type json_obj_location struct {
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	Altitude      float64 `json:"altitude"`
	Hor_accuracy  float64 `json:"horizontal_accuracy"`
	Vert_accuracy float64 `json:"vertical_accuracy"`
}

func WriteCsvLocation(file_name string, id []string, payload []string, date []string) {

	csvFile, err := os.Create(file_name)
	if err != nil {
		log.Fatalf("failed creating the file: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)

	i := 0
	row := []string{"participant_id", "latitude", "longitude", "altitude", "horizontal_accuracy", "vertical_accuracy", "date"}
	err = csvwriter.Write(row)
	if err != nil {
		log.Fatalf("Writing of the header into the csv went wrong: %s", err)
	}
	fmt.Println("Writing location into the csv file")
	for i < len(id) {
		jsonRes := json_obj_location{}
		err = json.Unmarshal([]byte(payload[i]), &jsonRes)
		if err != nil {
			log.Fatalf("Unmarshal went wrong: %s", err)
		}
		row := []string{id[i], fmt.Sprintf("%v", jsonRes.Latitude), fmt.Sprintf("%v", jsonRes.Longitude), fmt.Sprintf("%v", jsonRes.Altitude),
			fmt.Sprintf("%v", jsonRes.Hor_accuracy), fmt.Sprintf("%v", jsonRes.Vert_accuracy), date[i]}
		err = csvwriter.Write(row)
		if err != nil {
			log.Fatalf("Writing into the csv went wrong: %s", err)
		}
		i++
	}
	csvwriter.Flush()
	fmt.Println("Writing the locations data into the csv file is finished")
	fmt.Println("---------------------")
}
