package model

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func Result() ([][]float64, []float64) {
	tope := 3000
	output_result := []float64{}
	all_input := [][]float64{}
	csvFile, err := os.Open("data set.csv")
	defer csvFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfull")
	for i, line := range csvLines /*Algo...*/{
		tos, _ := strconv.ParseFloat(line[4], 64)
		cafelea, _ := strconv.ParseFloat(line[5], 64)
		congNasal, _ := strconv.ParseFloat(line[6], 64)
		dif_Respiratoria, _ := strconv.ParseFloat(line[7], 64)
		dolor_garganta, _ := strconv.ParseFloat(line[8], 64)
		fiebre, _ := strconv.ParseFloat(line[9], 64)
		diarrea, _ := strconv.ParseFloat(line[10], 64)
		nauseas, _ := strconv.ParseFloat(line[11], 64)
		anosmia_hiposmia, _  := strconv.ParseFloat(line[12], 64)
		dolor_abdominal, _ := strconv.ParseFloat(line[13], 64)
		dolor_articulaciones, _:= strconv.ParseFloat(line[14], 64)
		dolor_muscular, _ := strconv.ParseFloat(line[15], 64)
		dolor_pecho, _ := strconv.ParseFloat(line[16], 64)
		otros, _ := strconv.ParseFloat(line[17], 64)
		flag, _ := strconv.ParseFloat(line[2], 64)
		input := []float64{tos, cafelea, congNasal, dif_Respiratoria, dolor_garganta, fiebre, diarrea, nauseas, anosmia_hiposmia, dolor_abdominal, dolor_articulaciones, dolor_muscular, dolor_pecho, otros}

		if i != 0 && i < tope {
			all_input = append(all_input, input)
			output_result = append(output_result, flag)
		}else if i == tope {
			break
		}
	}
	return all_input, output_result
}
