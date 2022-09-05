package main

import (
  "encoding/json"
	"fmt"
  "io/ioutil"
	"log"
  // "os"
	"strconv"

	"github.com/gocolly/colly"
  
)

// to present the structure of the rhino facts
type Fact struct{
  // the names are capitalised because we want them to be available outside of the package main
  ID int `json:"id"`
  Description string `json:"description"`
  
}

func main() {
  // initiating empty slice, facts will be appended to this later
  allFacts := make([]Fact, 0)

  // using collly creating a newcollctor and providing a cartain domain to it to allow
  collector := colly. NewCollector(
    colly.AllowedDomains("factretriever.com", "www.factretriever.com"),
  )

  // to traverse through the DOM
  // it takes 2 arguments : (1)go query selector,   (2)the callback function that will be executed every time our collector finds a fact list item
  // colly package uses go query to interact with DOM
  // go query is like jquery but for golang
  collector.OnHTML(".factsList li", func(element *colly.HTMLElement) {
    // creating a variable to hold the id of fact
    // factID := element.Attr("id")
    // id will be in string, we have to convert the type using the package strconv
    factID, err := strconv.Atoi(element.Attr("id"))
    // the Atoi(ascii to integer returns an error as it's second value)
    if err != nil {
      log.Println("Couldn't get an ID")
    }

    // creating variable to store the description of each fact
    // based the struct defined earlier, the desc will be of type string
    factDesc := element.Text

    // creating a new fact struct to store every list item we iterate over
    fact := Fact{
      ID:          factID,
      Description: factDesc,
    }

    // appending the the fact struct to the allFacts slice
    allFacts = append(allFacts, fact)
  })

  // FOR SANITY CHECK let's give a message of a url which the srapper is visiting
  collector.OnRequest(func (request *colly.Request) {
    // the onrequest method registers a callback function that will be executed everytime the collector makes a request
    fmt.Println("Visiting", request.URL.String())
  })

  // telling the collector where to start, crawl and start scrapping
  collector.Visit("https://www.factretriever.com/rhino-facts")

  // INSTEAD OF PRINTING IN TERMINAL, PRINTING IN A FILE
  writeJSON(allFacts)

  
  // TO PRINT THE FACTS IN THE TERMINAL=========
  // // using the json package to create a new JSON encoder that writes to the standard output
  // enc := json.NewEncoder(os.Stdout)
  // // using indent to format the JSON
  // enc.SetIndent("", " ")
  // // passing the all facts encoder to the encoder so that it can encoded into JSON
  // enc.Encode(allFacts)

}

func writeJSON(data []Fact) {
  // using marshalindeant to marshal the data in th body of the fact
  // it returns a json encoded data and also returns an error
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}
  // this part needs to ioutil package
  // we can use writefile method to write the json data into a file with permission code 0644

	_ = ioutil.WriteFile("rhinofacts.json", file, 0644)
}

