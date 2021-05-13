package find

import (
	"reflect"

	"github.com/mniak/graphite"
)

func Methods(root interface{}) (result []graphite.Method, err error) {
	var interf graphite.Method
	items, err := Type(root, reflect.TypeOf(&interf).Elem())
	if err != nil {
		return
	}
	result = make([]graphite.Method, len(items))
	for i, item := range items {
		result[i] = item.(graphite.Method)
	}
	return
}
