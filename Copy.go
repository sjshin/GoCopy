package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"time"
)

func main() {
	/*
		fmt.Print("src text: ")
		var src string
		fmt.Scanln(&src)
		//fmt.Print(input)
		fmt.Print("dest text: ")
		var dest string
		fmt.Scanln(&dest)
		//Dir("D:/언세CD작업", "G:/test")
	*/

	now := time.Now()
	fmt.Println(now.Format("2006-01-02-15:04:05"))
	flag.Parse()

	src := flag.Arg(0)
	dest := flag.Arg(1)
	Dir(src, dest)
	end := time.Now()
	fmt.Println("")
	fmt.Print("Start Time : ")
	fmt.Println(now.Format("2006-01-02-15:04:05"))
	fmt.Print("End Time : ")
	fmt.Println(end.Format("2006-01-02-15:04:05"))
	t2 := now.Add(time.Second * 341)
	diff := t2.Sub(now)
	fmt.Print("Diff Time: ")
	fmt.Println(diff)
	println("복사가 완료 되었습니다")

}

func File(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	return os.Chmod(dst, srcinfo.Mode())
}

func Dir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())
		println(srcfp)
		if fd.IsDir() {
			if err = Dir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = File(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
