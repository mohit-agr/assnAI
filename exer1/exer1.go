
package main

import (
    "bufio"
    "fmt"
    // "io"
    "log"
    // "io/ioutil"
    "os"
)

type node struct {
    name string
    parent *node
    // end bool
    count int
}

type startpair struct {
    char string
    word []string
}

type endpair struct {
    char string
    word []string
}

func create_lists(words []string) ([676]startpair, [676]endpair){
    sp := [676]startpair{}
    ep := [676]endpair{}

    // initialize list with start and end pair of literals
    for i:='a'; i<='z'; i++ {
        for j:='a'; j<='z'; j++ {
            sp[26*(i-'a') + (j-'a')].char = string(i) + string(j)
            ep[26*(i-'a') + (j-'a')].char = string(i) + string(j)
        }
    }

    var index uint16
    for i:=0; i<len(words); i++ {
        wlen := len(words[i])
        if wlen >= 2 {
            index = 26*(uint16(words[i][0])-'a') + (uint16(words[i][1])-'a')
            sp[index].word = append(sp[index].word, words[i])

            index = 26*(uint16(words[i][wlen-2])-'a') + (uint16(words[i][wlen-1])-'a')
            ep[index].word = append(ep[index].word, words[i])
        }
    }

    return sp, ep
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    file, err := os.Open("wordsEn.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)
    var words []string
    // words := make([]string, 0)
    for scanner.Scan() {
        words = append(words, scanner.Text())
    }
    fmt.Println(words[2000], len(words))

    // reader := bufio.NewReader(os.Stdin)

    // fmt.Print("Enter word1: ")
    // word1, _ := reader.ReadString('\n')
    // fmt.Print("Enter word2: ")
    // word2, _ := reader.ReadString('\n')
    // word1 = word1[0 : len(word1) - 1]
    // word2 = word2[0 : len(word2) - 1]

    word1 := "pen"
    word2 := "paper"
    fmt.Println(word1, len(word1))
    fmt.Println(word2, len(word2))

    fmt.Println(word1[2])
    fmt.Println(word2[2])

    start := new(node)
    start.name = word1
    start.count = 0
    fmt.Println(start.name[2])

    frontier := start
    fmt.Println(frontier.name)

    stw, etw := create_lists(words)

    fmt.Println(stw[0], stw[10], etw[3], etw[15])

    // for frontier.name != word2 && frontier.count <= 10 {

    // }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
