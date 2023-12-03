package main

import (
    "fmt"
    "strings"
    "os"
)

func isNum(g byte) bool{
    return g >= '0' && g <= '9'
}

func getCalibrationValue(s string) int{
    i := 0
    var num1 int = 0
    var num2 int = 0
    for i < len(s){
        if(isNum(s[i])){
            num1 = int(byte(s[i]) - ('0'))
        }
        i++
    }
    i = len(s)-1
    for i >= 0{
        if(isNum(s[i])){
            num2 = int(byte(s[i]) - ('0'))
        }
        i--
    }
    return num2*10+num1
}

func main(){
    c,e := os.ReadFile(os.Args[1])
    if e != nil {
        fmt.Println("Error reading input.txt")
        os.Exit(1)
    }

    replacer := strings.NewReplacer("zero", "ze0o", "one", "o1e", "two", "t2o", "three", "thr3e", "four", "fo4r", "five", "fi5e", "six", "s6x", "seven", "sev7n", "eight", "eig8t", "nine", "ni9e")
    
    var temp string = string(c)
    var temp2 string

    for {
        temp2 = replacer.Replace(temp)
        if temp == temp2 {
            // when no replacements have been made
            break
        }
        temp = temp2
        
    }

    strs := strings.Split(temp,"\n")

    sum := 0
    for _,str := range strs{
        sum += getCalibrationValue( replacer.Replace(str))
    }
    fmt.Println(sum)
}
