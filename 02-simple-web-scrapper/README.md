
![](02_Simple_Web_Scrapper_SS.webm)


# Simple Web Scrapper

This project is a Web Scrapper that brings up facts about a. internship posted from this [website](https://internshala.com/internships) and stores in a CSV file.


## Little More
- Uses the given pacakges :- "encoding/csv", "fmt", "io/ioutil", "log", "strconv", "github.com/gocolly/colly". 
- It picks up the "a" and "span" elements of the facts from the HTML page by traversing through the DOM using the colly package.
- Collects the data in a collector.
- Stores the data in a CSC file using csvWriter.


