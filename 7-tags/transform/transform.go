package transform

import (
	"errors"
	"reflect"
	"strings"
)

func Tag(value any) error {
	fieldType := reflect.TypeOf(value).Elem()   // obtém o tipo da variável
	fieldValue := reflect.ValueOf(value).Elem() // obtém o valor da variável

	// quero saber se o tipo da variável é struct
	if fieldType.Kind() != reflect.Struct {
		return errors.New("Expected a struct")
	}

	// percorrendo todos os campos que temos na struct
	for i := 0; i < fieldType.NumField(); i++ {
		field := fieldType.Field(i)  // obtém o tipo do campo
		value := fieldValue.Field(i) // obtém o valor do campo

		if !value.CanSet() { // verifica se o campo pode ser modificado, pois se não for, não faz sentido utilizar a tag
			continue
		}

		transform := field.Tag.Get("transform")  // obtém o valor da tag transform
		if transform == "" || transform == "-" { // verifica se a tag está vazia ou se o valor é "-"
			continue
		}

		switch value.Kind() { // verifica o tipo do campo
		case reflect.String: // se for string podemos verificar qual a formatação desejada
			stringValue := value.String()
			if strings.Contains(transform, "upper") { // transformar para maiúsculo
				value.SetString(strings.ToUpper(stringValue))
			}
			if strings.Contains(transform, "lower") { // transformar para minúsculo
				value.SetString(strings.ToLower(stringValue))
			}
		default:
			return errors.New("Unsupported type") // se o tipo não for string, retornamos um erro
		}
	}

	return nil
}
