package main

import  (
    "fmt"
)

const englishHelloPrefix = "Hello "
const portugueseHelloPrefix  = "Oi " 
const koreanHelloPrefix  = "안녕 "
func Hello(name string, language string)string{
    if name == ""{
        name = "World"
    }
    return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

    switch language{
    case "korean":
        prefix = koreanHelloPrefix
    case "portuguese" :
        prefix = portugueseHelloPrefix
    default :
        prefix = englishHelloPrefix
}
  return
}
func main(){
    fmt.Println(Hello("Joao", ""))
}
