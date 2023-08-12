package gotype

import (
	"testing"

	"github.com/mauricioromagnollo/go-packages/typevalidator"
)

func TestImport(t *testing.T) {
	if typevalidator.GetTypeOf(1) != "int" {
		t.Error("Import error")
	}
}
