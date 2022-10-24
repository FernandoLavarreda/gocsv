package main

import (
	"os"
	"encoding/csv"
	"fmt"
	"strconv"
	"io"
	"log"
)


var PARSING = []string{"INT", "FLOAT", "STR"};
var INIT_RECORDS = 42

type Record struct{
	Ints []int
	Floats []float64
	Strings []string
	Format []string
}

func count(arr []string, value string) int{
	mapd := make(map[string]int)
	for _, v := range arr{
		if _, in := mapd[v]; in{
			mapd[v] += 1
		}else{
			mapd[v] = 1;
		}
	}
	return mapd[value]
}

func (record *Record) readArr(str []string) Record{
	INITCAP_INT := count(record.Format, "INT")
	INITCAP_FLOAT := count(record.Format, "FLOAT")
	INITCAP_STRING := count(record.Format, "STR")
	
	ints := make([]int, INITCAP_INT)
	floats := make([]float64, INITCAP_FLOAT)
	strings := make([]string, INITCAP_STRING)
	curints := 0
	curfloats := 0
	curstrings := 0
	for i, v := range record.Format{
		switch ;v{
		case PARSING[0]:
			parse,_ := strconv.Atoi(str[i])
			ints[curints] = parse;
			curints+=1
		
		case PARSING[1]:
			parse, _ := strconv.ParseFloat(str[i], 64)
			floats[curfloats] = parse
			curfloats+=1
		
		case PARSING[2]:
			parse := string(str[i])
			strings[curstrings] = parse
			curstrings+=1
		}
	}

	return Record{Ints: ints, Floats:floats, Strings:strings, Format:record.Format}
}


func CheckColumns(fileloc string, delimiter rune) (int, []string, []string) {
	file, err := os.Open(fileloc)
	if err != nil{
		log.Fatal(err);
	}
	reader := csv.NewReader(file)
	reader.Comma = delimiter
	line, err := reader.Read()
	line2, err := reader.Read()
	return len(line), line, line2
}


func ReadCSV(fileloc string, delimiter rune, format []string, header bool) ([]Record, []int){
	template := Record{Format: format}
	records := make([]Record, INIT_RECORDS)
	cursor := 0;
	errors := make([]int, 10)
	file, err := os.Open(fileloc)
	if err != nil{
		log.Fatal(err);
	}
	reader := csv.NewReader(file)
	reader.Comma = delimiter
	if header{
		reader.Read()
	}
	for{
		line, err := reader.Read()
		if err == io.EOF{
			break
		}
		if err != nil{
			errors = append(errors, cursor)
		}
		reading := template.readArr(line);
		if cursor+1>INIT_RECORDS{
			records = append(records, reading);
		}else{
			records[cursor] = reading;
		}
		cursor+=1
	}
	return records, errors
}


func main(){
	rec := Record{Format: []string{"INT", "FLOAT", "FLOAT", "FLOAT", "FLOAT"}};
	fmt.Println(CheckColumns("cv.txt", '\t'))
	records, fails := ReadCSV("cv.txt", '\t', rec.Format, true)
	fmt.Println(records)
	fmt.Println(fails)

}
