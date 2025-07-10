package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	// Each key gonna be transformed to a certain path on disk
	// 68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff
	key := "momsbestpicture"
	pathname := CASPathTransformFunc(key)

	expectedPathName := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"

	if pathname != expectedPathName {
		t.Errorf("have %s want %s", pathname, expectedPathName)
	}

	fmt.Println(pathname)
}

func TestStore(t *testing.T) {

	opts := StoreOpts{
		PathTransformFunc: DefaultPathTransformFunc,
	}

	s := NewStore(opts)

	data := bytes.NewReader([]byte("some jpeg bytes"))

	if err := s.writeStream("myspecialpicture", data); err != nil {
		t.Error(err)
	}
}
