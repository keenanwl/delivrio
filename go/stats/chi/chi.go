package chi

import (
	"fmt"
)

func TestStatistic(controlSuccess float64, controlFail float64, testSuccess float64, testFail float64) float64 {

	controlCount := controlSuccess + controlFail
	testCount := testSuccess + testFail

	totalSuccess := controlSuccess + testSuccess
	totalFail := controlFail + testFail
	totalObservations := float64(controlCount + testCount)

	expectedControlSuccess := (totalSuccess * controlCount) / totalObservations
	expectedControlFail := (totalFail * controlCount) / totalObservations

	expectedTestSuccess := (totalSuccess * testCount) / totalObservations
	expectedTestFail := (totalFail * testCount) / totalObservations

	fmt.Println(controlSuccess, expectedControlSuccess, controlFail, expectedControlFail, testSuccess, expectedTestSuccess, testFail, expectedTestFail)

	controlSuccessStat := ((controlSuccess - expectedControlSuccess) * (controlSuccess - expectedControlSuccess)) / expectedControlSuccess
	controlFailStat := ((controlFail - expectedControlFail) * (controlFail - expectedControlFail)) / expectedControlFail
	testSuccessStat := ((testSuccess - expectedTestSuccess) * (testSuccess - expectedTestSuccess)) / expectedTestSuccess
	testFailStat := ((testFail - expectedTestFail) * (testFail - expectedTestFail)) / expectedControlFail

	fmt.Println(controlSuccessStat, controlFailStat, testSuccessStat, testFailStat)

	testStatistic := controlSuccessStat + controlFailStat + testSuccessStat + testFailStat

	return testStatistic

}
