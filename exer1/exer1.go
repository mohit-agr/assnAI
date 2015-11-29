
package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

type node struct {
    name string
    parent *node
    end bool        // 0 for starting pair, 1 for ending pair
    cost int
}

type wordlists struct {
    char string
    word []string
}

/**
* Creates two lists using the beginning pair of literals and 
* ending pair of literals
*/
func create_lists(words []string) ([676]wordlists, [676]wordlists) {
    sp := [676]wordlists{}
    ep := [676]wordlists{}

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

func in_index_arr(ind uint16, _arr []uint16) (bool){
    for _, cur_ind := range _arr {
        if ind == cur_ind {
            return true
        }
    }
    return false
}

/**
* Checks if the given node is already in the explored set
*/
func in_queue(nd node, que []node) (bool) {
    for _, cur := range que {
        if nd.name == cur.name && nd.end == cur.end {
            return true
        }
    }
    return false
}

/**
* Search Breadth-first for the goal node
*/
func search(start *node, end *node , sp [676]wordlists, ep [676]wordlists) (*node){
    if start.name == end.name {
        return start
    }

    var st_index, en_index uint16
    var _arr_st_index, _arr_en_index []uint16
    var wlen uint8

    var cur [100000]node
    var child node
    var ndC uint64
    ndC = 0
    ftr := []node{}
    exp := []node{}
    depth := 0

    ftr = append(ftr, *start)

    for len(ftr) != 0 {
        cur[ndC], ftr = ftr[0], ftr[1:len(ftr)]
        if in_queue(cur[ndC], exp) {
            continue
        }

        child.parent = &cur[ndC]
        child.cost = cur[ndC].cost+1

        st_index = 26*(uint16(cur[ndC].name[0])-'a') + (uint16(cur[ndC].name[1])-'a')
        if !in_index_arr(st_index, _arr_st_index) {
            _arr_st_index = append(_arr_st_index, st_index)

            child.end = false
            for i:=0; i<len(ep[st_index].word); i++ {
                child.name = ep[st_index].word[i]
                if child.name == end.name {
                    return &child
                }
                ftr = append(ftr, child)
            }  
        }          

        wlen = uint8(len(cur[ndC].name))
        en_index = 26*(uint16(cur[ndC].name[wlen-2])-'a') + (uint16(cur[ndC].name[wlen-1])-'a')
        if !in_index_arr(en_index, _arr_en_index) {
            _arr_en_index = append(_arr_en_index, en_index)

            child.end = true
            for i:=0; i<len(sp[en_index].word); i++ {
                child.name = sp[en_index].word[i]
                if child.name == end.name {
                    return &child
                }
                ftr = append(ftr, child)
            }            
        }
        if cur[ndC].cost != depth {
            depth = cur[ndC].cost
        }
        exp = append(exp, cur[ndC])
        ndC++
    }

    child.name = ""
    return &child
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

    for scanner.Scan() {
        words = append(words, scanner.Text())
    }

    word1 := []string{"pen", "acorn", "autumn", "earthquake", "starship"}
    word2 := []string{"paper", "oak", "winter", "tsunami", "enterprise"}

    reader := bufio.NewReader(os.Stdin)

    file, err = os.Create("outputs.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    fmt.Print("Enter word1: ")
    word, _ := reader.ReadString('\n')
    word1 = append(word1, word[:len(word)-1])
    fmt.Print("Enter word2: ")
    word, _ = reader.ReadString('\n')
    word2 = append(word2, word[:len(word)-1])

    sp, ep := create_lists(words)

    for ix := 0; ix < len(word1); ix++ {
        start := new(node)
        start.name = word1[ix]
        start.cost = 0

        end := new(node)
        end.name = word2[ix]

        result := search(start, end, sp, ep)
        
        for result.cost !=0 {
            _, err = file.WriteString(result.name + "\n")
            check(err)
            if ix == len(word1)-1 {
                fmt.Println(word1[ix], len(word1[ix]))
                fmt.Println(word2[ix], len(word2[ix]))
                fmt.Println(result.name, result.cost, result.parent.name)
            }
            result = result.parent
        }
        _, err = file.WriteString(result.name + "\n")
        check(err)

        _, err = file.WriteString("\n")
        check(err)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
