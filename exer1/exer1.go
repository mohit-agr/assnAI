
package main

import (
    "bufio"
    "fmt"
    // "io"
    "log"
    // "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // dat, err := ioutil.ReadFile("wordsEn.txt")
    // check(err)
    // fmt.Print(string(dat))

    file, err := os.Open("wordsEn.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)

    var words [110000]string
    var word_idx uint32
    word_idx = 0
    for scanner.Scan() {
        words[word_idx] = scanner.Text()
        // fmt.Println(words[word_idx])
        word_idx += 1
    }
    fmt.Println(words[2000])

    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter word1: ")
    word1, _ := reader.ReadString('\n')
    fmt.Print("Enter word2: ")
    word2, _ := reader.ReadString('\n')
    word1 = word1[0 : len(word1) - 1]
    word2 = word2[0 : len(word2) - 1]
    fmt.Println(word1, " ", len(word1))
    fmt.Println(word2, " ", len(word2))


    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    // f, err := os.Open("/tmp/dat")
    // check(err)
    // b1 := make([]byte, 5)
    // n1, err := f.Read(b1)
    // check(err)
    // fmt.Printf("%d bytes: %s\n", n1, string(b1))

    // o2, err := f.Seek(6, 0)
    // check(err)
    // b2 := make([]byte, 2)
    // n2, err := f.Read(b2)
    // check(err)
    // fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

    // o3, err := f.Seek(6, 0)
    // check(err)
    // b3 := make([]byte, 2)
    // n3, err := io.ReadAtLeast(f, b3, 2)
    // check(err)
    // fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    // _, err = f.Seek(0, 0)
    // check(err)

    // r4 := bufio.NewReader(f)
    // b4, err := r4.Peek(5)
    // check(err)
    // fmt.Printf("5 bytes: %s\n", string(b4))

    // f.Close()

}
