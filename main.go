package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 打开输入文件
	inputFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("无法打开输入文件:", err)
		return
	}
	defer inputFile.Close()

	// 创建输出文件
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("无法创建输出文件:", err)
		return
	}
	defer outputFile.Close()

	// 正则表达式匹配

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	// 处理每一行
	for scanner.Scan() {
		line := scanner.Text()

		// 使用正则表达式匹配域名
		ln := strings.Split(line, ".")
		_, err := writer.WriteString(ln[len(ln)-2] + "." + ln[len(ln)-1] + "\n")
		if err != nil {
			fmt.Println("写入文件时发生错误:", err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("读取文件时发生错误:", err)
		return
	}

	// 刷新并关闭输出文件
	writer.Flush()
}
