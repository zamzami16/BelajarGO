package test

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string;

func TestEmbedFile(t *testing.T)  {
	fmt.Println(version)
}

//go:embed logo.png
var logo []byte

func TestEmbedBinary(t *testing.T)  {
	err:= os.WriteFile("logo_tulis.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestEmbedMultiFiles(t *testing.T)  {
	a, err := files.ReadFile("files/a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(a))

	b, err := files.ReadFile("files/b.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	c, err := files.ReadFile("files/c.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(c))
}
