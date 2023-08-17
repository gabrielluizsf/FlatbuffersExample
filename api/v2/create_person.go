package main

import (
	schemas "github.com/gabrielluizsf/FlatbuffersExample/api/v2/flatbuffers"
	flatbuffers "github.com/google/flatbuffers/go"
)
func createPerson(name string, age int32)(string,int32){
	builder := flatbuffers.NewBuilder(1024)
	personName := builder.CreateString(name)
	
	schemas.PersonStart(builder)
	schemas.PersonAddName(builder,personName)
	schemas.PersonAddAge(builder,age)
 
	person := schemas.PersonEnd(builder)
 
	builder.Finish(person)
 
	buf := builder.FinishedBytes()
	p := schemas.GetRootAsPerson(buf,0)
 
	return string(p.Name()),p.Age()
 }