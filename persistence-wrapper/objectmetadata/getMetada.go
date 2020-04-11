package pq_metadata

import (
	"fmt"
	"reflect"
	"strings"

)

func BuildSelectQuery(q interface{}) (query string) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		nameTable := strings.ToLower(reflect.TypeOf(q).Name())
		fields := GetFields(q)
		query = fmt.Sprintf(strings.ToLower(" select %s from  %s "), fields, nameTable)
	}
	return
}

func BuildInsertQuery(q interface{}) (query string) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		nameTable := strings.ToLower(reflect.TypeOf(q).Name())
		values, _ := GetValues(q)
		fields := strings.ToLower(GetFields(q))

		query = fmt.Sprintf(strings.ToLower(" insert into %s   (%s) values (%s) "), nameTable, fields, values)
	}
	return
}

func GetValues(q interface{}) (values string, prepareFields string) {

	v := reflect.ValueOf(q)
	for i := 0; i < v.NumField(); i++ {

		fmt.Println( "dados do objeto field ",v.Field(i))
		switch v.Field(i).Kind() {
		case reflect.Int:
			values = fmt.Sprintf("%s, %d", values, v.Field(i).Int())
			prepareFields = fmt.Sprintf(" %s ", prepareFields, "?")
		case reflect.Float64, reflect.Float32:
			values = fmt.Sprintf("%s, %d", values, v.Field(i).Float())
			prepareFields = fmt.Sprintf(" %s ", prepareFields, "?")
		case reflect.Bool:
			values = fmt.Sprintf("%s, %t", values, v.Field(i).Bool())
			prepareFields = fmt.Sprintf(" %s ", prepareFields, "?")
		case reflect.String:
			values = fmt.Sprintf("%s, '%s'", values, v.Field(i).String())
			prepareFields = fmt.Sprintf(" %s ", prepareFields, "?")
		default:
			fmt.Println("Unsupported type")
			return
		}
	}
	values = strings.TrimSpace(values)
	values = values[1:len(values)] //remove primeira virgula
	return

}

func GetFields(q interface{}) (fields string) {

	v := reflect.TypeOf(q)
	for i := 0; i < v.NumField(); i++ {

		fields = fmt.Sprintf("%s,%s", fields, v.Field(i).Name)
	}
	fields = fields[1:len(fields)] //remove primeira virgula
	return

}

func GetFieldsNames(q interface{}) (fields[] string) {
	v := reflect.TypeOf(q)
	fields =make([]string, v.NumField() )
	for i := 0; i < v.NumField(); i++ {
		fields[i]=v.Field(i).Name
	}
	return
}

//Retorna um map com o nome e tipo dos campos de uma struct
func GetFieldsMap(a interface{}) (fields map[string]string) {
	v := reflect.ValueOf(a).Elem()
	fields = make(map[string]string,v.NumField())
	for j := 0; j < v.NumField(); j++ {
		fields[v.Type().Field(j).Name] =v.Field(j).Type().Name()
	}
	return
}
