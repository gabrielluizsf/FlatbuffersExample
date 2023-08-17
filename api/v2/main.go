package main

import (
	"fmt"

	schemas "github.com/gabrielluizsf/FlatbuffersExample/api/v2/flatbuffers"
	flatbuffers "github.com/google/flatbuffers/go"
)
func main(){
   createPerson("Gabriel Luiz",21)
   createCourse(1,"Análise e Desenvolvimento de Sistemas","Graduação")
   
}
func createPerson(name string, age int32){
   builder := flatbuffers.NewBuilder(1024)
   personName := builder.CreateString(name)
   
   schemas.PersonStart(builder)
   schemas.PersonAddName(builder,personName)
   schemas.PersonAddAge(builder,age)

   person := schemas.PersonEnd(builder)

   builder.Finish(person)

   buf := builder.FinishedBytes()
   p := schemas.GetRootAsPerson(buf,0)

   fmt.Println(string(p.Name()),p.Age())
}
func createCourse(courseID int32,courseName, courseDescription string){
	builder := flatbuffers.NewBuilder(1024)
   
   courseN := builder.CreateString(courseName)
   courseDesc := builder.CreateString(courseDescription)

   schemas.CourseStart(builder)
   schemas.CourseAddId(builder,courseID)
   schemas.CourseAddName(builder,courseN)
   schemas.CourseAddDescription(builder,courseDesc)
   course := schemas.CourseEnd(builder)

   builder.Finish(course)

   buf := builder.FinishedBytes()
   c := schemas.GetRootAsCourse(buf,0)
   fmt.Println(c.Id(),string(c.Name()),string(c.Description()))
}