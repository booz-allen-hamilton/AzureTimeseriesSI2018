// Script to traverse the Enron email dataset and extract it into a collection
// of relationship edges
package main

import (
	"path/filepath"
	"os"
	"fmt"
	"bufio"
	"io"
	"log"
	"strings"
	"encoding/csv"
)

//Container for scraped to/from communications
type commSumStruct struct {
    To string
    From string
	DateString string
}
var myArray = make([] commSumStruct, 0)
//Text flags for to/from in comms
var desiredStrings = [2] string {"X-From", "X-To"}

var fileLimit = 1000
var fileCounter = 0

func commonFieldExtractor (text string) string {
	split1 := strings.Split(text, ":")
	split2 := strings.Split(split1[1], "<")
	result := strings.Trim(split2[0], " ")
	return result
}


//Function passed to the walk command which operates over each file in a
//directory, in this case scraping out to/froms in communications
func visit (path string, info os.FileInfo, err error) error {

    //Crude filter operation to avoid trying to read folders
    fmt.Println(info.Name(), info.Size())
    stat, _ := os.Stat(path)
    if stat.IsDir() {
        return nil
    }
    
	//Throws error to exit the walk loop
    if fileCounter > fileLimit {
        return io.EOF
    }
    
    file, err := os.Open(path)
    
    if err != nil {
        log.Fatal(err)
    }

	holderStruct := commSumStruct{"", "", ""}
    
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        var text = scanner.Text()          
		if strings.Contains(text, "X-From") {
			holderStruct.From = commonFieldExtractor(text)
		}
		if strings.Contains(text, "X-To") {
			holderStruct.To = commonFieldExtractor(text)
		}
    }
    
    //For lines with TO/FROMs, append to the package level variable
    if (holderStruct.To != "" && holderStruct.From != "") {
		myArray = append(myArray, holderStruct)
    }

	fileCounter += 1
    
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    
    return err
}

func main() {
	filepath.Walk("/media/sf_GuestShared/maildir/", visit)
	fmt.Println(myArray)

	file, err := os.Create("result.csv")
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, value := range myArray {
		tempArray := [] string {value.From, value.To}
        err = writer.Write(tempArray)
    }
	
	if err != nil {
		fmt.Println("Devistation")	
	}
}