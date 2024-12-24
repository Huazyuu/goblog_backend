package models

import (
	"fmt"
	"os"
	"testing"
)

func TestFullTextModel_Mapping(t *testing.T) {
	currentDir, _ := os.Getwd()
	fmt.Println("current dir:", currentDir)
	txt := FullTextModel{}.Mapping()
	fmt.Println(txt)
}
