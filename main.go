package main

import "fmt"

func main()  {
  var publisher, writer, artist, title string
  var year, pageNumber int
  var grade float32

  title = "Mr. GoToSleep"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year, pageNumber, grade = 1997, 14, 6.5


  fmt.Println("o livro", title, "written by", writer, "drawn by",   artist)
  fmt.Println("published by", publisher, "in", year,
 "has", pageNumber, "pages, and has a grade of", grade)

  title = "Epic Vol. 1"
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  year, pageNumber, grade = 2013, 160, 9.0


  fmt.Println("o livro", title, "written by", writer, "drawn by",   artist)
  fmt.Println("published by", publisher, "in", year,
 "has", pageNumber, "pages, and has a grade of", grade)
}