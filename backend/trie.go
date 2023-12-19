// this is the most overcomplicated thing ever but its just a homegrown, grass fed, free range, DIY trie

package main

import (
    "fmt"
    "bufio"
    "os"
    "time"
    //"runtime"
    //"sync"
)

type trieNode struct {
    children []trieNode
    end bool 
}

func trieAppend(root *trieNode, word string) {
    for idx := 0; idx < len(word); idx++ {
        root = &root.children[word[idx] - 'A']
        if root.children == nil {
            root.children = make([]trieNode, 26)
        }
    }
    root.end = true
}

func validateWord(root *trieNode, word string) bool {
    for idx := 0; idx < len(word); idx++ {
        if root.children == nil {
            return false;
        }
        root = &root.children[word[idx] - 'A']
    }
    return root.end
}

func makeTrie() *trieNode {

    root := trieNode{
        children: make([]trieNode, 26),
        end:      false,
    }

    file, err := os.Open("./words.txt")

    if err != nil {
        panic(err)
    }

    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    startTime := time.Now() 

	for scanner.Scan() {
        line := scanner.Text() 
        trieAppend(&root, string(line))
	}
   
    fmt.Println("Successful Parsing:", time.Now().Sub(startTime))
    return &root

}