package bigproc

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"io"
	"log"
	"os"
)

type Data struct {
	Field1  string `json:"field1"`
	Field2  string `json:"field2"`
	Field3  string `json:"field3"`
	Field4  string `json:"field4"`
	Field5  string `json:"field5"`
	Field6  string `json:"field6"`
	Field7  string `json:"field7"`
	Field8  string `json:"field8"`
	Field9  string `json:"field9"`
	Field10 string `json:"field10"`
	Field11 string `json:"field11"`
	Field12 string `json:"field12"`
	Field13 string `json:"field13"`
	Field14 string `json:"field14"`
	Field15 string `json:"field15"`
	Field16 string `json:"field16"`
	Field17 string `json:"field17"`
	Field18 string `json:"field18"`
	Field19 string `json:"field19"`
	Field20 string `json:"field20"`
	Field21 string `json:"field21"`
	Field22 string `json:"field22"`
	Field23 string `json:"field23"`
	Field24 string `json:"field24"`
	Field25 string `json:"field25"`
}

func Do(con *sql.DB) {
	var dataArr []Data
	csvFile, err := os.Open("./test.csv")
	if err != nil {
		log.Println(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		dataArr = append(dataArr, Data{
			Field1:  line[0][:20],
			Field2:  line[1],
			Field3:  line[2],
			Field4:  line[3],
			Field5:  line[4],
			Field6:  line[5],
			Field7:  line[6],
			Field8:  line[7],
			Field9:  line[8],
			Field10: line[9],
			Field11: line[10],
			Field12: line[11],
			Field13: line[12],
			Field14: line[13],
			Field15: line[14],
			Field16: line[15],
			Field17: line[16],
			Field18: line[17],
			Field19: line[18],
			Field20: line[19],
			Field21: line[20],
			Field22: line[21],
			Field23: line[22],
			Field24: line[23],
			Field25: line[24],
		})
	}

	savetoSQL(con, dataArr)
}
