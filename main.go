package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Book struct {
	ID   int
	Name string
}

type Book2 struct {
	ID   int
	Name string
	Tags []Tag
}

type Tag struct {
	TagID   int
	TagName string
}

func main() {
	b := Book{
		ID:   1,
		Name: "book_name",
	}

	map1 := structToMap(&b)
	fmt.Println("structToMap: ", map1)

	b2 := Book2{
		ID:   1,
		Name: "book_name",
		Tags: []Tag{
			{
				TagID:   1,
				TagName: "tag_name1",
			},
			{
				TagID:   2,
				TagName: "tag_name2",
			},
		},
	}
	map2 := structToMap2(&b2)
	fmt.Println("structToMap2: ", map2)
}

func structToMap(data interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()

	for i := 0; i < size; i++ {
		field := strings.ToLower(elem.Type().Field(i).Name)
		value := elem.Field(i).Interface()

		result[field] = value
	}

	return result
}

func structToMap2(data interface{}) map[string]interface{} {
	result := map[string]interface{}{}

	elem := reflect.ValueOf(data).Elem()

	for i := 0; i < elem.NumField(); i++ {
		field := strings.ToLower(elem.Type().Field(i).Name)
		value := elem.Field(i).Interface()

		if reflect.TypeOf(value).Kind() == reflect.Slice {
			sliceValue := reflect.ValueOf(value)
			for j := 0; j < sliceValue.Len(); j++ {
				elem2 := sliceValue.Index(j)
				for k := 0; k < elem2.NumField(); k++ {
					field2 := strings.ToLower(elem2.Type().Field(k).Name)
					value2 := elem2.Field(k).Interface()

					result[field2] = value2
				}
			}
			continue
		}
		result[field] = value
	}
	return result
}
