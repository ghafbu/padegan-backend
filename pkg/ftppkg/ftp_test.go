package ftppkg

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

const (
	Host     = "194.147.142.197"
	Username = "pz15176"
	Password = "MCk6BApm"
	Port     = "21"
)

func TestFTP(t *testing.T) {

	newConnFTP := New()

	ConnFTP, err := newConnFTP.Dial(Host, Port)
	if err != nil {
		t.Errorf("%v", err)
	}

	err = ConnFTP.Login(Username, Password)
	if err != nil {
		t.Errorf("%v", err)
	}
	//
	err = newConnFTP.MakeDir("/public_html/1002")
	if err != nil {
		log.Println(err)

		return
	}

	err = newConnFTP.Store(
		"/public_html/1002/hello-world.txt",
		bytes.NewBufferString("Hello world!!!1111"),
	)
	if err != nil {
		t.Errorf("store: %v", err)
	}
	//
	log.Println("store is successful")

	log.Println(newConnFTP)
	if err = newConnFTP.Quit(); err != nil {
		fmt.Println(err)
	}

	// Read file
	//r, err := c.Retr(pathData)
	//if err != nil {
	//	panic(err)
	//}
	//defer r.Close()
	//
	//buf, err := io.ReadAll(r)
	//println(string(buf))
}
