package main
import "fmt"

func main() {

animals := make(chan string, 2)

animals <- "bears"
animals <- "turtles"


critter1 := <-animals
critter2 := <-animals

fmt.Println(critter1)
fmt.Println(critter2)
}