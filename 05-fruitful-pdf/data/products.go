package data

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
  
)

type Fruit struct {
  // creating the fruit struct using fake function from the fake package
  Name string `fake:"{fruit}"`
  // Description will be around 10 words
  Description string `fake:"{loremipsumsentence:10}"`
  // price will be a float64 between 1 to 10
  Price float64 `fake:"{price:1, 10}"`
}

func generateFruit() []string {
  var f Fruit
  gofakeit.Struct(&f)

  froot := []string{}
  // to this froot every random data will be appended
  froot = append(froot, f.Name)
  froot = append(froot, f.Description)
  // the float price should be converted into string to print
  froot = append(froot, fmt.Sprint("%.2f", f.Price))

  return froot
  
}

// tablelist expects the fruit in 2d slice format of a string
// length variable decides how long the fruit list will be
func FruitList(length int) [][]string{
  var fruits [][]string

  // for loop to generate no of random FruitList
  for i := 0; i < length; i ++ {
    // creating a new fruit and appending it to the fruits variable
    onefruit := generateFruit()
    fruits = append(fruits, onefruit)
    
  }
  return fruits
  
}
