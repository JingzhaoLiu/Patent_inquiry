package read

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func Read() map[string]string {
	countryMap := make(map[string]string)
	xlsx, err := excelize.OpenFile("./ex.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	rows := xlsx.GetRows("Sheet1")

	// cell := xlsx.GetCellValue("Sheet1", "B2")
	// fmt.Println(cell)

	for k, row := range rows {
		countryMap[strconv.Itoa(k+1)] = row[0]
	}
	// fmt.Println(countryMap)
	return countryMap
}

func Write(obj map[string]string) {
	xlsx, err := excelize.OpenFile("./excopy.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	for s, val := range obj {

		fmt.Println("Sheet1", "D"+s, val)

		xlsx.SetCellValue("Sheet1", "D"+s, val)
		fmt.Println(s)
	}

	err = xlsx.SaveAs("./excopy.xlsx")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("ok")
}
