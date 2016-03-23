package main

import (
       "fmt"
       "reflect"
)

var RegisteredTypes = make(map[string]reflect.Type)

type Inter interface {
     Hi() string
}

type Imp1 struct {
     Name string
}

func (i *Imp1) Hi() string {
     i.Name = "Implementation 1"
     return i.Name
}

type Imp2 struct {
     Alias string
}

func (i *Imp2) Hi() string {
     i.Alias = "Implementation 2"
     return i.Alias
}

func Factory(imp_type string) Inter {
     implementation_type := RegisteredTypes[imp_type]
     new_obj := reflect.New(implementation_type).Interface().(Inter)
     return new_obj
}

func main() {
     RegisteredTypes["1"] = reflect.TypeOf(Imp1{})
     RegisteredTypes["2"] = reflect.TypeOf(Imp2{})

     imp1 := Factory("1")
     imp2 := Factory("2")

     fmt.Printf("Hello from %s\n", imp1.Hi())
     fmt.Printf("Hello from %s\n", imp2.Hi())
}
