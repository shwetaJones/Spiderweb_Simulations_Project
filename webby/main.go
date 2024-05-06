package main

import (
	"fmt"
	"gifhelper"
	"gowut/gwu"
	"image"
	"math/rand"
	"strconv"
	"time"
)

func WebPage(numberRadii int, initialAngle float64, minDistance float64, initialDistance float64, alphavalue int) {

	rand.Seed(time.Now().UTC().UnixNano())
	
	// Runs the Symmetric Web Function and creates the respective symmetric web
	canvasWidth := 1000.0
	alpha := float64(alphavalue) + rand.Float64()
	allLinesSymmetric := CreateWebSymmetric(initialAngle, canvasWidth, numberRadii, alpha, minDistance, initialDistance)
	var finalWebSym Web
	finalWebSym.lines = &allLinesSymmetric
	finalWebSym.width = int(canvasWidth)
	fmt.Println("Simuation Done, Generating Images!")
	imagesSym := make([]image.Image, 0)
	for i := 1; i < len(allLinesSymmetric)+1; i++ {
		imageSym := DrawWeb(allLinesSymmetric[:i], finalWebSym.width, numberRadii)
		imagesSym = append(imagesSym, imageSym)
	}
	fmt.Println("Symmetric Web Created, Making GIF!")
	gifhelper.ImagesToGIF(imagesSym, "symmetric")
	fmt.Println("GIF Made!")
	
	// Runs the Random Web Function and creates the respective random web
	allLinesRandom := CreateWebRandom(initialAngle, canvasWidth, numberRadii, alpha, minDistance, initialDistance)
	var finalWebRand Web
	finalWebRand.lines = &allLinesRandom
	finalWebRand.width = int(canvasWidth)
	fmt.Println("Random Web Created, Generating Images!")
	imagesRand := make([]image.Image, 0)
	for i := 1; i < len(allLinesRandom)+1; i++ {
		imageRand := DrawWeb(allLinesRandom[:i], finalWebRand.width, numberRadii)
		imagesRand = append(imagesRand, imageRand)
	}
	fmt.Println("Images made, Making GIF!")
	gifhelper.ImagesToGIF(imagesRand, "random")
	fmt.Println("GIF Made!")

}

func main() {

	// opening a new web browser called spider_web
	spider_web := gwu.NewWindow("spider", "spider_web")
	spider_web.SetCellPadding(10)

	// creating a new panel
	thread := gwu.NewVerticalPanel()

	// creating text boxe or drop down lists for user input parameters
	thread.Add(gwu.NewLabel("Enter Radii Number:"))
	radiitextbox := gwu.NewTextBox("")
	radiitextbox.Style().SetWidth("1000")
	radiitextbox.SetRows(2)
	thread.Add(radiitextbox)

	thread.Add(gwu.NewLabel("Choose Initial Angle:"))
	anglelist := gwu.NewTextBox("")
	anglelist.Style().SetWidth("1000")
	anglelist.SetRows(2)
	thread.Add(anglelist)

	thread.Add(gwu.NewLabel("Choose Minimum Distance:"))
	distancelist := gwu.NewListBox([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"})
	distancelist.Style().SetWidth("1000")
	distancelist.SetRows(9)
	thread.Add(distancelist)

	thread.Add(gwu.NewLabel("Choose Initial Distance:"))
	initialdistancelist := gwu.NewListBox([]string{"5", "6", "7", "8", "9", "10", "11", "12"})
	initialdistancelist.Style().SetWidth("1000")
	initialdistancelist.SetRows(8)
	thread.Add(initialdistancelist)

	thread.Add(gwu.NewLabel("Choose Alpha Value:"))
	alphavaluelist := gwu.NewListBox([]string{"4", "5", "6", "7", "8", "9", "10", "11", "12"})
	alphavaluelist.Style().SetWidth("1000")
	alphavaluelist.SetRows(9)
	thread.Add(alphavaluelist)

	// adding the panel to the web page

	spider_web.Add(thread)

	// create a button

	buttonmain := gwu.NewButton(fmt.Sprintf("Click for Symmetric Spider Web"))
	buttonmain.AddEHandlerFunc(func(e gwu.Event) {

		// convert the user inputs as inputs for the main function called WebPage

		radii := gwu.HasText.Text(radiitextbox)

		radii1, _ := strconv.Atoi(radii)

		angle := gwu.HasText.Text(anglelist)

		angle1, _ := strconv.ParseFloat(angle, 64)

		distance := distancelist.SelectedValue()

		distance1, _ := strconv.ParseFloat(distance, 64)

		initialdistance := initialdistancelist.SelectedValue()

		initialdistance1, _ := strconv.ParseFloat(initialdistance, 64)

		alpha := alphavaluelist.SelectedValue()

		alpha1, _ := strconv.Atoi(alpha)
		// insert the inputs in the main function called WebPage

		WebPage(radii1, angle1, distance1, initialdistance1, alpha1)

		// add the symmetric spider web GIF to local host 8080

		thread.Add(gwu.NewLabel("Spider Web Gif:"))

		spiderwebvisualisation := gwu.NewImage("See Your Web", "http://localhost:8080")

		thread.Add(spiderwebvisualisation)

		e.MarkDirty(thread)
	}, gwu.ETypeClick)
	thread.Add(buttonmain)

	//start of the second button for Random spider
	buttonmain2 := gwu.NewButton(fmt.Sprintf("Click for Random Spider Web"))
	buttonmain2.AddEHandlerFunc(func(e gwu.Event) {

		radii := gwu.HasText.Text(radiitextbox)

		radii1, _ := strconv.Atoi(radii)

		angle := gwu.HasText.Text(anglelist)

		angle1, _ := strconv.ParseFloat(angle, 64)

		distance := distancelist.SelectedValue()

		distance1, _ := strconv.ParseFloat(distance, 64)

		initialdistance := initialdistancelist.SelectedValue()

		initialdistance1, _ := strconv.ParseFloat(initialdistance, 64)

		alpha := alphavaluelist.SelectedValue()

		alpha1, _ := strconv.Atoi(alpha)

		WebPage(radii1, angle1, distance1, initialdistance1, alpha1)

		thread.Add(gwu.NewLabel("Spider Web Gif:"))

		// add the symmetric spider web GIF to local host 8081
		spiderwebvisualisation := gwu.NewImage("See Your Web", "http://localhost:8081")

		thread.Add(spiderwebvisualisation)

		e.MarkDirty(thread)
	}, gwu.ETypeClick)
	thread.Add(buttonmain2)

	// the main web page will run on server 8026

	server := gwu.NewServer("spider_web", "localhost:8026")
	server.AddWin(spider_web)
	server.Start("")

	if errorwebpage := server.Start(); errorwebpage != nil {
		fmt.Println("There is an error with your webpage! Sorry!", errorwebpage)
		return

	}

}
