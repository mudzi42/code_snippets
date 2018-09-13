package codesnippets

import (
	"io/ioutil"
	"os"
)

// write a byte string to file

func writeByte() {
	lateBindVariableTestData := []byte(`{
	"buildPassword": "FAKEPassword",
	"production": false
}`)
	ioutil.WriteFile("runtimeVariables.json", lateBindVariableTestData, 0644)

	os.Remove("runtimeVariables.json")
}
