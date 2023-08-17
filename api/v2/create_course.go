package main

import (
	schemas "github.com/gabrielluizsf/FlatbuffersExample/api/v2/flatbuffers"
	flatbuffers "github.com/google/flatbuffers/go"
)

func createCourse(courseID int32,courseName, courseDescription string) (int32,string,string){
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
   return c.Id(),string(c.Name()),string(c.Description());
}