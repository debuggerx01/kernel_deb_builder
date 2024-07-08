package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	origin_config := readByLine("config")
	new_config := []string{}
	for _, line := range origin_config {
		if strings.HasPrefix(line, "CONFIG_R8169") {
			log.Printf("找到不需要的配置%s", line)
			continue
		}
		if strings.HasPrefix(line, "CONFIG_DRM_AMDGPU") {
			log.Printf("找到不需要的配置%s", line)
			continue
		}
		if strings.HasPrefix(line, "CONFIG_MLX4") || strings.HasPrefix(line, "CONFIG_MLX5") {
			log.Printf("找到不需要的配置%s", line)
			continue
		}
		if strings.HasPrefix(line, "CONFIG_ATH") {
			log.Printf("找到不需要的配置%s", line)
			continue
		}
		//log.Println(new_config)
		//log.Printf("第%d行%s\n", i, line)
		new_config = append(new_config, line)
	}
	if _, err := os.Stat("smartConfig"); err == nil {
		// 文件存在，删除文件
		err := os.Remove("smartConfig")
		if err != nil {
			panic(err)
		}
		fmt.Println("File deleted successfully.")
	} else if os.IsNotExist(err) {
		fmt.Println("File does not exist.")
	} else {
		panic(err)
	}
	WriteByLine("smartConfig", new_config)
}
func readByLine(fp string) []string {
	lines := []string{}
	fi, err := os.Open(fp)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		log.Println("按行读文件出错")
		return []string{}
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		lines = append(lines, string(a))
	}
	return lines
}
func WriteByLine(fp string, s []string) {
	file, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, v := range s {
		writer.WriteString(v)
		writer.WriteString("\n")
	}
	writer.Flush()
	return
}
func removeRealtek(s string) string {
	var disable string
	if strings.HasPrefix(s, "CONFIG_R8169") {
		aim := strings.Split(s, "=")
		k := aim[0]
		v := aim[1]
		log.Printf("原始值%v is %v\n", k, v)
		disable = strings.Join([]string{k, "n"}, "=")
		log.Printf("修改后:%v\n", disable)
	}
	return disable
}
func removeAMDGPU(s string) string {
	var disable string
	if strings.HasPrefix(s, "CONFIG_DRM_AMDGPU") {
		//CONFIG_R8169=m
		aim := strings.Split(s, "=")
		k := aim[0]
		v := aim[1]
		log.Printf("原始值%v=%v\n", k, v)
		disable = strings.Join([]string{k, "n"}, "=")
	}
	return disable
}
func removeMellanox(s string) string {
	var disable string
	if strings.HasPrefix(s, "CONFIG_MLX4_EN") || strings.HasPrefix(s, "CONFIG_MLX5_CORE") {
		//CONFIG_R8169=m
		aim := strings.Split(s, "=")
		k := aim[0]
		v := aim[1]
		log.Printf("原始值%v=%v\n", k, v)
		disable = strings.Join([]string{k, "n"}, "=")
	}
	return disable
}
func removeAtheros(s string) string {
	var disable string
	if strings.HasPrefix(s, "CONFIG_ATH") {
		//CONFIG_R8169=m
		aim := strings.Split(s, "=")
		k := aim[0]
		v := aim[1]
		log.Printf("原始值%v=%v\n", k, v)
		disable = strings.Join([]string{k, "n"}, "=")
	}
	return disable
}
