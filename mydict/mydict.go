package mydict

import (
	"errors"
)

// Dictionary type
type Dictionary map[string]string


var errNotFound = errors.New("not Found");
var errWordExsist = errors.New("The word is exsist");

// Search for a wword
func (d Dictionary) Search(word string) (string,error){
	value, exists := d[word]
	if exists{
		return value, nil
	}
	return "", errNotFound
}

// Add for a word
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExsist
	}
	
	return nil
}

// Update for a word and def
func (d Dictionary) Update(word string, def string) (string, error) {
	err := d.Add(word,def)
	switch err{
	case errWordExsist:
		d[word] = def
	case nil:
		return "the word is add succesfully ",nil;
	}
	return "update is successful",nil
}

// Delete for a word
func (d Dictionary) Delete(word string) {
	
	delete(d,word)
}