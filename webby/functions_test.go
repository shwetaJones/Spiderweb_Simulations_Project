package main

import (
	"fmt"
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	type test struct {
		test1, test2 OrderedPair
		distance     float64
	}

	var Test test

	Test.test1.x = 4.0
	Test.test1.y = 6.0
	Test.test2.x = 1.0
	Test.test2.y = 2.0
	Test.distance = 5.0

	output := CalcDistance(Test.test1, Test.test2)

	if Test.distance == output {
		fmt.Println("TestDistance function works properly!")
	} else {
		t.Errorf("The TestDistance function is incorrect, the output was %f when it was supposed to be %f", output, Test.distance)
	}
	fmt.Println()
}

func TestCalcLengthSymmetric(t *testing.T) {
	type test struct {
		testX     float64
		testY     float64
		testAngle float64
		expOutput float64
	}
	var Test test
	Test.testX = 3.0
	Test.testY = 4.0
	Test.testAngle = 45.0
	Test.expOutput = 4.63
	testOutput := CalcLengthSymmetric(Test.testX, Test.testY, Test.testAngle)
	testOutput = math.Floor(testOutput*100) / 100
	if testOutput == Test.expOutput {
		fmt.Println("CalcLengthSymmetric function works properly!")
	} else {
		t.Errorf("The CalcLengthSymmetric function is incorrect, the output was %f when it was supposed to be %f", testOutput, Test.expOutput)
	}
	fmt.Println()
}

func TestCheckQuadrant(t *testing.T) {
	type test struct {
		testAngle float64
		expOutput int
	}
	var Test test
	Test.testAngle = 45.0
	Test.expOutput = 1
	testOutput := CheckQuadrant(Test.testAngle)
	if testOutput == Test.expOutput {
		fmt.Println("CheckQuadrant function works properly!")
	} else {
		t.Errorf("The CheckQuadrant function on the first test is incorrect, the output was %d when it was supposed to be %d", testOutput, Test.expOutput)
	}

	Test.testAngle = 100.0
	Test.expOutput = 2
	testOutput = CheckQuadrant(Test.testAngle)
	if testOutput == Test.expOutput {
		fmt.Println("CheckQuadrant function works properly!")
	} else {
		t.Errorf("The CheckQuadrant function on the second test is incorrect, the output was %d when it was supposed to be %d", testOutput, Test.expOutput)
	}

	Test.testAngle = 210.0
	Test.expOutput = 3
	testOutput = CheckQuadrant(Test.testAngle)
	if testOutput == Test.expOutput {
		fmt.Println("CheckQuadrant function works properly!")
	} else {
		t.Errorf("The CheckQuadrant function on the third test is incorrect, the output was %d when it was supposed to be %d", testOutput, Test.expOutput)
	}

	Test.testAngle = 350.0
	Test.expOutput = 4
	testOutput = CheckQuadrant(Test.testAngle)
	if testOutput == Test.expOutput {
		fmt.Println("CheckQuadrant function works properly!")
	} else {
		t.Errorf("The CheckQuadrant function on the fourth test is incorrect, the output was %d when it was supposed to be %d", testOutput, Test.expOutput)
	}

	Test.testAngle = 500.0
	Test.expOutput = 2
	testOutput = CheckQuadrant(Test.testAngle)
	if testOutput == Test.expOutput {
		fmt.Println("CheckQuadrant function works properly!")
	} else {
		t.Errorf("The CheckQuadrant function on the fifth test is incorrect, the output was %d when it was supposed to be %d", testOutput, Test.expOutput)
	}
	fmt.Println()
}

func TestChangeAxisCartToCanvas(t *testing.T) {
	type test struct {
		testPoint OrderedPair
		testWidth float64
		expOutput OrderedPair
	}

	var Test test
	Test.testPoint.x = 0.0
	Test.testPoint.y = 0.0
	Test.testWidth = 100.0
	Test.expOutput.x = 50.0
	Test.expOutput.y = 50.0
	testOutput := ChangeAxisCartToCanvas(Test.testPoint, Test.testWidth)
	if testOutput == Test.expOutput {
		fmt.Println("ChangeAxisCartToCanvas function works properly!")
	} else {
		t.Errorf("The ChangeAxisCartToCanvas function's first test is incorrect, the output was %f when it was supposed to be %f", testOutput, Test.expOutput)
	}

	Test.testPoint.x = 10.0
	Test.testPoint.y = 9.0
	Test.testWidth = 100.0
	Test.expOutput.x = 60.0
	Test.expOutput.y = 59.0
	testOutput = ChangeAxisCartToCanvas(Test.testPoint, Test.testWidth)
	if testOutput == Test.expOutput {
		fmt.Println("ChangeAxisCartToCanvas function works properly!")
	} else {
		t.Errorf("The ChangeAxisCartToCanvas function's second test is incorrect, the output was %f when it was supposed to be %f", testOutput, Test.expOutput)
	}

	Test.testPoint.x = -10.0
	Test.testPoint.y = 9.0
	Test.testWidth = 100.0
	Test.expOutput.x = 40.0
	Test.expOutput.y = 59.0
	testOutput = ChangeAxisCartToCanvas(Test.testPoint, Test.testWidth)
	if testOutput == Test.expOutput {
		fmt.Println("ChangeAxisCartToCanvas function works properly!")
	} else {
		t.Errorf("The ChangeAxisCartToCanvas function's third test is incorrect, the output was %f when it was supposed to be %f", testOutput, Test.expOutput)
	}

	Test.testPoint.x = 10.0
	Test.testPoint.y = -9.0
	Test.testWidth = 100.0
	Test.expOutput.x = 60.0
	Test.expOutput.y = 41.0
	testOutput = ChangeAxisCartToCanvas(Test.testPoint, Test.testWidth)
	if testOutput == Test.expOutput {
		fmt.Println("ChangeAxisCartToCanvas function works properly!")
	} else {
		t.Errorf("The ChangeAxisCartToCanvas function's fourth test is incorrect, the output was %f when it was supposed to be %f", testOutput, Test.expOutput)
	}

	Test.testPoint.x = -10.0
	Test.testPoint.y = -9.0
	Test.testWidth = 100.0
	Test.expOutput.x = 40.0
	Test.expOutput.y = 41.0
	testOutput = ChangeAxisCartToCanvas(Test.testPoint, Test.testWidth)
	if testOutput == Test.expOutput {
		fmt.Println("ChangeAxisCartToCanvas function works properly!")
	} else {
		t.Errorf("The ChangeAxisCartToCanvas function's fifth test is incorrect, the output was %f when it was supposed to be %f", testOutput, Test.expOutput)
	}

	fmt.Println()
}

func TestChangeAxisCanvasToCart(t *testing.T) {
	type test struct {
		testPoint OrderedPair
		testWidth float64
		expOutput OrderedPair
	}

	var Test test
	Test.testPoint.x = 50.0
	Test.testPoint.y = 50.0
	Test.testWidth = 100.0
	Test.expOutput.x = 0.0
	Test.expOutput.y = 0.0
	testOutput := ChangeAxisCanvastoCart(Test.testPoint, Test.testWidth)
	if testOutput == Test.expOutput {
		fmt.Println("ChangeAxisCanvastoCart function works properly!")
	} else {
		t.Errorf("The ChangeAxisCanvastoCart function's first test is incorrect, the output was %f when it was supposed to be %f", testOutput, Test.expOutput)
	}

	Test.testPoint.x = 60.0
	Test.testPoint.y = 59.0
	Test.testWidth = 100.0
	Test.expOutput.x = 10.00
	Test.expOutput.y = -9.00
	testOutput = ChangeAxisCanvastoCart(Test.testPoint, Test.testWidth)
	if testOutput == Test.expOutput {
		fmt.Println("ChangeAxisCanvastoCart function works properly!")
	} else {
		t.Errorf("The ChangeAxisCanvastoCart function's second test is incorrect, the output was %f when it was supposed to be %f", testOutput, Test.expOutput)
	}

	Test.testPoint.x = 40.0
	Test.testPoint.y = 59.0
	Test.testWidth = 100.0
	Test.expOutput.x = -10.00
	Test.expOutput.y = -9.00
	testOutput = ChangeAxisCanvastoCart(Test.testPoint, Test.testWidth)
	if testOutput == Test.expOutput {
		fmt.Println("ChangeAxisCanvastoCart function works properly!")
	} else {
		t.Errorf("The ChangeAxisCanvastoCart function's third test is incorrect, the output was %f when it was supposed to be %f", testOutput, Test.expOutput)
	}

	Test.testPoint.x = 60.0
	Test.testPoint.y = 41.0
	Test.testWidth = 100.0
	Test.expOutput.x = 10.00
	Test.expOutput.y = 9.00
	testOutput = ChangeAxisCanvastoCart(Test.testPoint, Test.testWidth)
	if testOutput == Test.expOutput {
		fmt.Println("ChangeAxisCanvastoCart function works properly!")
	} else {
		t.Errorf("The ChangeAxisCanvastoCart function's fourth test is incorrect, the output was %f when it was supposed to be %f", testOutput, Test.expOutput)
	}

	Test.testPoint.x = 40.0
	Test.testPoint.y = 41.0
	Test.testWidth = 100.0
	Test.expOutput.x = -10.00
	Test.expOutput.y = 9.00
	testOutput = ChangeAxisCanvastoCart(Test.testPoint, Test.testWidth)
	if testOutput == Test.expOutput {
		fmt.Println("ChangeAxisCanvastoCart function works properly!")
	} else {
		t.Errorf("The ChangeAxisCanvastoCart function's fifth test is incorrect, the output was %f when it was supposed to be %f", testOutput, Test.expOutput)
	}
	fmt.Println()
}

func TestBasePointsSymmetric(t *testing.T) {
	type test struct {
		testWidth float64
		outputBP1 OrderedPair
		outputBP2 OrderedPair
		outputBP3 OrderedPair
		outputBP4 OrderedPair
	}

	var Test test
	Test.testWidth = 1000.0
	Test.outputBP1.y = 0.0
	Test.outputBP1.x = 250.0
	Test.outputBP2.x = 0.0
	Test.outputBP2.y = 250.0
	Test.outputBP3.x = -250.0
	Test.outputBP3.y = 0.0
	Test.outputBP4.x = 0.0
	Test.outputBP4.y = -250.0

	testOutput1, testOutput2, testOutput3, testOutput4 := BasePointsSymmetric(Test.testWidth)

	if testOutput1 == Test.outputBP1 {
		fmt.Println("BasePointsSymmetric function works properly!")
	} else {
		t.Errorf("The BasePointsSymmetric function is incorrect, the output was %f when it was supposed to be %f", testOutput1, Test.outputBP1)
	}

	if testOutput2 == Test.outputBP2 {
		fmt.Println("BasePointsSymmetric function works properly!")
	} else {
		t.Errorf("The BasePointsSymmetric function is incorrect, the output was %f when it was supposed to be %f", testOutput2, Test.outputBP2)
	}

	if testOutput3 == Test.outputBP3 {
		fmt.Println("BasePointsSymmetric function works properly!")
	} else {
		t.Errorf("The BasePointsSymmetric function is incorrect, the output was %f when it was supposed to be %f", testOutput3, Test.outputBP3)
	}

	if testOutput4 == Test.outputBP4 {
		fmt.Println("BasePointsSymmetric function works properly!")
	} else {
		t.Errorf("The BasePointsSymmetric function is incorrect, the output was %f when it was supposed to be %f", testOutput4, Test.outputBP4)
	}
	fmt.Println()
}

func TestFirstRadiiSymmetric(t *testing.T) {
	type test struct {
		inputAngle             float64
		inputWidth             float64
		x1, y1, x2, y2         OrderedPair
		outputStart, outputEnd OrderedPair
	}

	var Test test
	var outputRadius line

	Test.inputAngle = 30.0
	Test.inputWidth = 1000.0
	outputRadius.startPoint.x = 0.0
	outputRadius.startPoint.y = 0.0
	outputRadius.endPoint.x = 433.079
	outputRadius.endPoint.y = 249.885

	Test.x1.x, Test.x1.y = 0.0, 0.0
	Test.x2.x, Test.x2.y = 0.0, 0.0
	Test.y1.x, Test.y1.y = 0.0, 0.0
	Test.y2.x, Test.y2.y = 0.0, 0.0

	testOutput := FirstRadiiSymmetric(Test.inputAngle, Test.inputWidth, Test.x1, Test.y1, Test.x2, Test.y2)

	if roundFloat(testOutput.endPoint.x, 3) == outputRadius.endPoint.x {
		fmt.Println("FirstRadiiSymmetric function works properly!")
	} else {
		t.Errorf("The FirstRadiiSymmetric function is incorrect, the output was %f when it was supposed to be %f", testOutput.endPoint.x, outputRadius.endPoint.x)
	}

	if roundFloat(testOutput.endPoint.y, 3) == outputRadius.endPoint.y {
		fmt.Println("FirstRadiiSymmetric function works properly!")
	} else {
		t.Errorf("The FirstRadiiSymmetric function is incorrect, the output was %f when it was supposed to be %f", testOutput.endPoint.y, outputRadius.endPoint.y)
	}
	fmt.Println()
}

func TestCreateFrameSymmetric(t *testing.T) {
	type test struct {
		inputRadii  []line
		outputFrame []line
	}

	var Test test
	Test.inputRadii = make([]line, 3)

	Test.inputRadii[0].startPoint.x = 0.0
	Test.inputRadii[0].startPoint.y = 0.0
	Test.inputRadii[1].startPoint.x = 0.0
	Test.inputRadii[1].startPoint.y = 0.0
	Test.inputRadii[2].startPoint.x = 0.0
	Test.inputRadii[2].startPoint.y = 0.0

	Test.inputRadii[0].endPoint.x = 100.0
	Test.inputRadii[0].endPoint.y = 100.0
	Test.inputRadii[1].endPoint.x = 100.0
	Test.inputRadii[1].endPoint.y = 900.0
	Test.inputRadii[2].endPoint.x = 500.0
	Test.inputRadii[2].endPoint.y = 800.0

	Test.outputFrame = make([]line, 3)

	Test.outputFrame[0].startPoint.x = 100.0
	Test.outputFrame[0].startPoint.y = 100.0
	Test.outputFrame[1].startPoint.x = 100.0
	Test.outputFrame[1].startPoint.y = 900.0
	Test.outputFrame[2].startPoint.x = 500.0
	Test.outputFrame[2].startPoint.y = 800.0

	Test.outputFrame[0].endPoint.x = 100.0
	Test.outputFrame[0].endPoint.y = 900.0
	Test.outputFrame[1].endPoint.x = 500.0
	Test.outputFrame[1].endPoint.y = 800.0
	Test.outputFrame[2].endPoint.x = 100.0
	Test.outputFrame[2].endPoint.y = 100.0

	testOutput := CreateFrame(Test.inputRadii)

	if testOutput[0].startPoint.x == Test.outputFrame[0].startPoint.x {
		fmt.Println("CreateFrameSymmetric function works properly!")
	} else {
		t.Errorf("The CreateFrameSymmetric function is incorrect, the output was %f when it was supposed to be %f", testOutput[0].startPoint.x, Test.outputFrame[0].startPoint.x)
	}

	if testOutput[0].startPoint.x == Test.outputFrame[0].startPoint.x {
		fmt.Println("CreateFrameSymmetric function works properly!")
	} else {
		t.Errorf("The CreateFrameSymmetric function is incorrect, the output was %f when it was supposed to be %f", testOutput[0].startPoint.x, Test.outputFrame[0].startPoint.x)
	}

	if testOutput[2].endPoint.y == Test.outputFrame[0].startPoint.y {
		fmt.Println("CreateFrameSymmetric function works properly!")
	} else {
		t.Errorf("The CreateFrameSymmetric function is incorrect, the output was %f when it was supposed to be %f", testOutput[0].startPoint.y, Test.outputFrame[0].startPoint.y)
	}

	if testOutput[0].startPoint.y == Test.outputFrame[0].startPoint.y {
		fmt.Println("CreateFrameSymmetric function works properly!")
	} else {
		t.Errorf("The CreateFrameSymmetric function is incorrect, the output was %f when it was supposed to be %f", testOutput[0].startPoint.y, Test.outputFrame[0].startPoint.y)
	}

}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
