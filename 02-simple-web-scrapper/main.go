package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"
)

func main() {
  // creating a file named data.csv, handling error, closing the file
  fName := "data.csv"
  file, err := os.Create(fName)
  if err != nil {
    log.Fatalf("Couldn't create the file, error was : %q", err)
    return
  }
  // anything after the keyword defer will get executed afterwards, not right away
  defer file.Close()

  // a csv writer is needed to write in the csv file after the data is fetched from the site
  writer := csv.NewWriter(file)
  // once done with writing the file, everything present in the buffer should be thrown into the writer which can be later passed on to the file
  defer writer.Flush()

  // now web scraping begins using Colly
  c := colly.NewCollector(
    colly.AllowedDomains("internshala.com"),
  )

  // the collector needs to be pointed to a tag of webpage from where it will collect info
  // the func will have a pointer to that html element
  c.OnHTML(".internship_meta", func (e *colly.HTMLElement){
    writer.Write( []string {
      // concatenated and striped text context from the tag provided 
      // here pointing to the text and salary part of the internship
      e.ChildText("a"),
      e.ChildText("span"),
    })
  })

  // using for loop for pagination of all pages of internshala
  for i := 0; i < 5; i ++ {
    fmt.Printf("Scraping page number : %d\n", i)
    // passing the integer i as string to the endpoint page using Itoa
    c.Visit("https://internshala.com/internships/page-" + strconv.Itoa(i))
  }

  log.Printf("Scraping is completed")
  
}
