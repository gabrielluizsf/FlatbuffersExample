package main

import (
	"fmt"
	"log"
)

func main() {
   name, age := createPerson("Gabriel Luiz", 21)
   courseID, courseName, description := createCourse(1,"Análise e Desenvolvimento de Sistemas","Graduação")
   fmt.Println(name,age)
   printCourse(courseID,courseName,description)
   if err := startServer(":2500"); err != nil{
      log.Fatalf("Erro ao iniciar o servidor:",err)
   }
}
func printCourse(courseID int32, courseName, description string){
	fmt.Println(courseID,":"+courseName,"\nTipo de curso:"+description)
}