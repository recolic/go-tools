package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Copy(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func main() {
	if len(os.Args) < 3 {
		panic("Usage: ./ThisProgram <SrcFile> <DstFile> \n Example: ./InfiniteCopy /usr/bin/shit ~/tmp/another.shit")
	}

	fmt.Println("Infinite copying ", os.Args[1], "to", os.Args[2], " on every 10ms...")
	for {
		err := Copy(os.Args[1], os.Args[2])
		if err != nil {
			fmt.Println("Error: failed to copy file, " + err.Error())
		}
		time.Sleep(10 * time.Millisecond)
	}
}
