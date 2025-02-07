package dict

import "errors"

// Iterator
type Iterator interface {
	Next() interface{}
	HasNext() bool
}

type Iteratable struct {
	index int
	Iterator
}

// Dict
type Dict struct {
	terms map[string]string
	keys []string // Golang doesn't keep keys
	Iteratable // extends [next(), hasNext(), done()]
}


type Entry [2]string

/* API: keys, entries, Add, Remove */
func (dict *Dict) Add(key string, value string) {
	dict.terms[key] = value
	dict.keys = append(dict.keys, key)
}

func (dict *Dict) Remove(key string) {
	delete(dict.terms, key)
}

func (dict *Dict) Keys() []string {
	return dict.keys
}

func (dict *Dict) Values() []string {
	values := make([]string, 0)
	for _,key := range dict.keys {
		value := dict.terms[key]
		values = append(values, value)
	}
	return values
}


func (dict *Dict) Entries() []Entry {
	var entries []Entry
	for _, term := range dict.keys {
		var entry Entry
		entry[0] = term
		entry[1] = dict.terms[term]
		entries = append(entries, entry)
	}
	return entries
}

// clear function
func (dict *Dict) Clear() {
	dict.terms = make(map[string]string)
	dict.keys = make([]string, 0)
}

func (dict *Dict) Size() int {
	return len(dict.keys)
}

// > Iterator
func(dict *Dict) Next() (string, error) {
	if dict.HasNext() {
		key := dict.keys[dict.index]
		dict.index = dict.index + 1
		return dict.terms[key], nil
	}
	return "", errors.New("")
}

func(dict *Dict) HasNext() bool {
	return dict.index < dict.Size()
}

// Create Dict()
func CreateDict() *Dict {
	dict := new(Dict)
	dict.terms = make(map[string]string)
	dict.index = 0
	return dict
}

