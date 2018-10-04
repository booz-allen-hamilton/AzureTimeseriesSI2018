// Script to traverse the Enron email dataset and extract it into a collection
// of relationship edges
package main

import (
"path/filepath"
"os"
"fmt"
"bufio"
"log"
"strings"
"github.com/gyuho/goraph"
)

//Container for scraped to/from communications
type commSumStruct struct {
    To string
    From string
}
var myArray = make([] commSumStruct, 0)
//Text flags for to/from in comms
var desiredStrings = [2] string {"X-From", "X-To"}

var fileLimit = 10
var fileCounter = 0


//Function passed to the walk command which operates over each file in a
//directory, in this case scraping out to/froms in communications
func visit (path string, info os.FileInfo, err error) error {

    //Crude filter operation to avoid trying to read folders
    fmt.Println(info.Name(), info.Size())
    stat, _ := os.Stat(path)
    if stat.IsDir() {
        return nil
    }
    
    if fileCounter > fileLimit {
        return nil
    }
    
    file, err := os.Open(path)
    
    if err != nil {
        log.Fatal(err)
    }
    
    //Holding variable for the to/from as we scroll through the file
    var edgeArray = [2] string {"", ""}
    
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        var text = scanner.Text()          
        for index, value := range desiredStrings {
            //If we see a to/from indicator, break the message apart after it then again before
            //the <> formatted string at the end, trim off spaces then store
            if strings.Contains(text, value) {
                split1 := strings.Split(text, ":")
                split2 := strings.Split(split1[1], "<")
                result := strings.Trim(split2[0], " ")
                edgeArray[index] = result
            }
        }
    }
    
    //Append the extracted to/from to the package level variable
    myArray = append(myArray, commSumStruct{edgeArray[0], edgeArray[1]})
    fileCounter += 1

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    
    
    return err
}

func main() {
        filepath.Walk("./maildir", visit)
        fmt.Println(myArray)
        
        
        graph := goraph.NewGraph()
        
        node1 := goraph.NewNode("whatever")
        node2 := goraph.NewNode("whatever2")
        
        graph.AddNode(node1)
        graph.AddNode(node2)
        node, _ := graph.GetNode(goraph.StringID("whatever"))
        //id := goraph.StringID("whatever")
        //fmt.Println(id);
        graph.AddEdge(node.ID(), goraph.StringID("whatever2"), 1)
        fmt.Println(node.ID())
        stuff, _ := graph.GetTargets(node.ID())
        fmt.Println(stuff)
        /*
        file, err := os.Open("./1.")
        if err != nil {
                log.Fatal(err)
        }
        defer file.Close()
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            var text = scanner.Text()            
                for _, value := range death {
                    if strings.Contains(text, value) {
                        print(text)
                    }
                }

                //fmt.Println(scanner.Text())

        }

        if err := scanner.Err(); err != nil {
                log.Fatal(err)
        }
        */
}
