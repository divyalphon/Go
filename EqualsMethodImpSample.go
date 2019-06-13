package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type ApplicationErrorDetails struct {
	ApplicationAppSk        string `json:"applicationAppSk"`
	ApplicationErrorType    string `json:"applicationErrorType"`
	ApplicationProcessName  string `json:"applicationProcessName"`
	ApplicationErrorMessage string `json:"applicationErrorMessage"`
}

type ApplnErrTestFileds struct {
	ApplicationAppSk     string `json:"applicationAppSk"`
	ApplicationErrorType string `json:"applicationErrorType"`
}

func main() {

	var applicationErrorDetailsList []ApplicationErrorDetails

	var aed1, aed2, aed3, aed4 ApplicationErrorDetails

	aed1 = ApplicationErrorDetails{"AppSk1", "ErrorType1", "ProcessName1", "ErrorMsg1"}
	aed2 = ApplicationErrorDetails{"AppSk1", "ErrorType1", "ProcessName2", "ErrorMsg2"}
	aed3 = ApplicationErrorDetails{"AppSk3", "ErrorType3", "ProcessName3", "ErrorMsg3"}
	aed4 = ApplicationErrorDetails{"AppSk4", "ErrorType4", "ProcessName4", "ErrorMsg4"}

	applicationErrorDetailsList = append(applicationErrorDetailsList, aed1)
	applicationErrorDetailsList = append(applicationErrorDetailsList, aed2)
	applicationErrorDetailsList = append(applicationErrorDetailsList, aed3)
	applicationErrorDetailsList = append(applicationErrorDetailsList, aed4)

	//fmt.Println(applicationErrorDetailsList)

	consolidateMap := make(map[ApplnErrTestFileds][]ApplicationErrorDetails)
	for _, value := range applicationErrorDetailsList {
		//fmt.Println("value", value)
		jsonappErrTestValues, err := json.Marshal(value)
		if err != nil {
			log.Fatal(err, "not able to marshal ")
		}

		//fmt.Println("marshel", jsonappErrTestValues)
		var appErrTestValues ApplnErrTestFileds
		json.Unmarshal(jsonappErrTestValues, &appErrTestValues)

		//fmt.Println("Unmarshal", appErrTestValues)
		if consolidateMap[appErrTestValues] != nil {

			for key, valueone := range consolidateMap {
				if strings.EqualFold(key.ApplicationAppSk, value.ApplicationAppSk) && strings.EqualFold(key.ApplicationErrorType, value.ApplicationErrorType) {
					var sliceone []ApplicationErrorDetails
					sliceone = append(sliceone, value)
					valueone = append(valueone, value)
					consolidateMap[key] = valueone

				}
			}

		} else {
			var slicetwo []ApplicationErrorDetails

			slicetwo = append(slicetwo, value)
			consolidateMap[appErrTestValues] = slicetwo
		}

	}

	fmt.Println("map:", consolidateMap)

}
