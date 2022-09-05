
![](6_Rhino_Scraper_SS.gif)


# Rhino Scrapper

This project is a Web Scrapper that brings up facts about a Rhino from this [website](https://www.factretriever.com/rhino-facts) and stores in a JSON file.


## Little More
- Uses the given pacakges :- "encoding/json", "fmt", "io/ioutil", "log", "strconv", "github.com/gocolly/colly". 
- It picks up the id of the facts from the HTML page by traversing through the DOM
- Collects the data in a collector.
- Formats the data with indentaion
- Stores the data in a JSON file using ioutil packages.


