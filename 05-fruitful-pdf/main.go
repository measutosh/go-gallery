package main

import (
	"fmt"
  "os"
  "github.com/johnfercher/maroto/pkg/pdf"
  "github.com/johnfercher/maroto/pkg/consts"
  "github.com/johnfercher/maroto/pkg/props"
  "github.com/johnfercher/maroto/pkg/color"
  "main/data"
)

func main() {
  // using the package to create a pdf
  m := pdf.NewMaroto(consts.Portrait, consts.A4)
  // setting up the left, top and right margins
  m.SetPageMargins(20, 10, 20)

  // calling the add header to pdf function
  buildHeading(m)
  // calling the build fruit list to pdf function
  buildFruitList(m)

  // printing the results in a folder with naming convention
  // outputfileandclose returns an error so adding the error handling
  err := m.OutputFileAndClose("result/fruit.pdf")
  if err != nil {
    fmt.Println("Couldn't create the PDF, here is the error \n", err)
    os.Exit(1)
  }
	fmt.Println("PDF saved successfully, check in the result folder")
}


// Adding content to PDF, passing the Maroto instance as argument
func buildHeading(m pdf.Maroto) {
    // registerheader allows to create a global header for every page
  m.RegisterHeader(func() {
    // it accepts and anonymous call back function as its argument
    // this function can be thought as a container
    // adding a row with height of 50units
    m.Row(50, func() {
      // adding full-width column of 12 spaces
      m.Col(12, func(){
        // 2nd argument is there to configure the image a little
        err := m.FileImage("image/Banner.jpg", props.Rect {
          Center: true,
          Percent: 75,
        })

        if err != nil {
          fmt.Println("Image couldn't be loaded, more details here :\n", err)
        }
      })
    })
  })

  m.Row(10, func(){
    m.Col(12, func(){
      m.Text("Prepared for you by Asutosh Panda", props.Text {
        Top: 3,
        Style: consts.Bold,
        Align: consts.Center,
        Color: getDarkPurpleColor(),
      })
      
    })
  })
}

// Building table of products function that takes document instance as its argument
func buildFruitList(m pdf.Maroto){
  // building table headings of slice string
  tableHeadings := []string{"Fruit", "Description", "Price"}
  // tablelist expects 2d slice of string as content
  // to create randomly generated values in the pdf
  contents := data.FruitList(20)
  // below line can be used to hardcoded values to print in the pdf
  // contents := [][]string{{"Apple", "Read & Juicy", "120rs"},{"Orange", "Orange & Juicy", "70rs"}}
  // calling the lightpurplecolor creating function
  lightPurpleColor := getLightPurpleColor()
  
  m.SetBackgroundColor(getTealColor())

  m.Row(10, func() {
    m.Col(12, func() {
      m.Text("Products", props.Text{
        Top:    2,
        Size:   13,
        Color:  color.NewWhite(),
        Family: consts.Courier,
        Style:  consts.Bold,
        Align:  consts.Center,
      })
    })
  })

  // building the table
  m.SetBackgroundColor(color.NewWhite())

  // creating new table list with 3 arguments
  m.TableList(tableHeadings, contents, props.TableList{
    HeaderProp : props.TableListContent{
      Size: 9,
      // gridsize will be couple of unsigned integers that will add upto 12
      GridSizes: []uint{3, 7, 2},
    },
    ContentProp : props.TableListContent{
      Size: 8,
      GridSizes: []uint{3, 7, 2},
    },
    Align:               consts.Left,
    HeaderContentSpace : 1,
    // no lines will be there in between table rows
    Line :               false,
    AlternatedBackground: &lightPurpleColor,
  })
  
}

// creating the bgcolor for readability
func getLightPurpleColor() color.Color {
  return color.Color{
    Red:   210,
    Green: 200,
    Blue:  230,
  }
}
// adding the heading in purple color
func getDarkPurpleColor() color.Color {
  return color.Color{
    Red:   80,
    Green: 80,
    Blue:  99,
  }
}

// creating the teal color
func getTealColor() color.Color {
  return color.Color {
    Red:   3,
    Green: 166,
    Blue:  166,
  }
}
