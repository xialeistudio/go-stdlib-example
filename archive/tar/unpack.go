// 解包示例
package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开文件句柄
	fp, err := os.Open("pack.tar")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	// 创建读取器
	tr := tar.NewReader(fp)
	for hdr, err := tr.Next(); err != io.EOF; hdr, err = tr.Next() {
		if err != nil {
			fmt.Println(err)
			return
		}
		info := hdr.FileInfo()
		fp, err := os.Create(info.Name())
		if err != nil {
			fmt.Println(err)
			return
		}
		io.Copy(fp, tr)
		fp.Close()
	}
	fmt.Println("unpack completed")
}
