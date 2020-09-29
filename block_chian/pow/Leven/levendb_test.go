package Leven

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_levendb(t *testing.T) {
	db, err := New("")
	check(err)
	err = db.Put([]byte("k1"), []byte("v1"))
	check(err)

	v, err := db.Get([]byte("k1"))
	fmt.Printf("%s\n", v)
	if !bytes.Equal(v, []byte("v1")) {
		t.Fatal()
	}
}

func check(err error) {
	if err != nil {
		panic(err)

	}
}
