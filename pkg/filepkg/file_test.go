package filepkg

import (
	"fmt"
	"testing"
)

var file = New()

func TestCheckValidExtension(t *testing.T) {
	fmt.Println(file.IsValidExtension(".MO"))
}

func TestGetFilenameHashed(t *testing.T) {
	//fmt.Println(file.GetInfoFile("hello.world.exe.txt"))
	//fmt.Println(file.GetInfoFile("hello-world.exe.txt"))
	//fmt.Println(file.GetInfoFile("hello-world..txt"))
	//fmt.Println(file.GetInfoFile("hello-world.txt"))
}

func TestCheck(t *testing.T) {
	f := file.SetFile("ascii-hello..text.png").
		SetContentType("images/png").
		SetSize(2763578256)

	fmt.Printf("%#v", f.GetInfoFile())
}
