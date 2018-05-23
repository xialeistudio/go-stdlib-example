// 打包示例
package main

import (
	"archive/tar"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 创建文件句柄
	fp, err := os.Create("pack.tar")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	// 创建打包句柄
	tw := tar.NewWriter(fp)
	defer tw.Close()
	// 需要打包的数据
	files := []string{"pack.go", "README.md"}
	for _, name := range files {
		info, err := os.Stat(name)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 写入文件头
		hdr, err := tar.FileInfoHeader(info, "")
		if err != nil {
			fmt.Println(err)
			return
		}
		tw.WriteHeader(hdr)
		// 写入文件数据
		fp, err := os.Open(name)
		if err != nil {
			fmt.Println(err)
			return
		}
		io.Copy(tw, fp)
		fp.Close()
	}
	log.Println("pack completed")
}
