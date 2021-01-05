package huawei_hilink

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var fieldsA = []interface{}{
	"4", "-13dB", "-92dBm", "-65dBm", "None", "",
}

var fieldsAParsed = []interface{}{
	4, -13, -92, -65,
}

func TestParser(t *testing.T) {
	m := Module{}
	parse := make([]interface{}, 0)

	for _, field := range fieldsA {
		parseField, err := m.Parse(field)
		if err == nil {
			parse = append(parse, parseField)
		}
	}
	eq := reflect.DeepEqual(fieldsAParsed, parse)
	assert.True(t, eq, "Wrong parse elements")
}
