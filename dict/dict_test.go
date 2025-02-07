package dict

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDictKeys(t *testing.T) {
	dict := CreateDict()	
	dict.Add("name","Harris")
	dict.Add("addr","Prieuré 23")

	expectedKeys := []string{"name", "addr"}
	if !reflect.DeepEqual(dict.Keys(), expectedKeys) {
		t.Fatalf("Error: Dict keys don't match expected %v got %v", dict.Keys(), expectedKeys)
	}
}

func TestDictEntries(t *testing.T) {
	dict := CreateDict() 
	
	dict.Add("name","Harris")
	dict.Add("addr","Prieuré 23")

	expected := []Entry {
		{"name","Harris"},
		{"addr","Prieuré 23"},	
	}
	if !reflect.DeepEqual(expected, dict.Entries()) {
		t.Fatalf("Error: Dict entries don't match. Expected %v got %v", dict.Entries(), expected)
	}
}

func TestDictClear(t *testing.T) {
	dict := CreateDict()
	dict.Add("auteur","HB")
	dict.Add("siecle","XX")

	if dict.Size() !=2 {
		t.Fatalf("Error. Dict Size should be equal to 2")
	}

	dict.Clear()

	if dict.Size() !=0 {
		t.Fatalf("Error. Dict Size should equal to 0")
	}
}

func TestDictValues(t *testing.T) {
	dict := CreateDict()
	dict.Add("user","hb")
	dict.Add("password","pass")
	expectedValues :=[]string{"hb","pass"}
	fmt.Println(expectedValues)
	if !reflect.DeepEqual(dict.Values(), expectedValues) {
		t.Fatalf("Error: Dict values don't match expected %v got %v", dict.Keys(), expectedValues)
	}
}


func TestDictIterator(t *testing.T) {
	dict := CreateDict()
	dict.Add("strange","indeed")
	dict.Add("addr","Prieuré 23")
	expected := []string{"indeed","Prieuré 23"}
	result := []string{}
	for dict.HasNext() {
		term, err := dict.Next()
		if err != nil {
			fmt.Println(err)
		}
		result = append(result, term)
	}
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Error: Iterator entries don't match. Expected %v got %v", dict.Values(), expected)
	}
}
