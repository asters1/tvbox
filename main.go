package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"

	"github.com/go-yaml/yaml"
)

// 检查python环境
func Go_Init() string {
	M := make(map[string]interface{})
	cmd := exec.Command("python3", "--version")
	python_v, err := cmd.Output()
	if err != nil {
		M["code"] = 0
		M["message"] = "没有检测Python3,请自行安装!!!"
		M["python_v"] = ""
	} else {
		M["code"] = 1
		M["message"] = "您已安装Python3!"
		M["python_v"] = strings.TrimSpace(strings.Trim(string(python_v), "Python3"))
	}
	jstr, err := json.Marshal(M)
	if err != nil {
		return "{\"code\":0,\"message\":\"格式化json出错，请检查!!!函数名为Go_Init!!\",\"python_v\":\"\"}"
	} else {
		return string(jstr)
	}
}

func Go_HomeContent(filter bool, file_name string) string {
	M := make(map[string]interface{})
	flr := "False"
	if filter {
		flr = "True"
	}
	File_Name_Remove_py := strings.TrimSuffix(file_name, ".py")
	cmd := exec.Command("python3", "-c", "from "+File_Name_Remove_py+" import homeContent;homeContent("+flr+")")
	content, err := cmd.Output()
	if err != nil {
		M["code"] = 0
		M["message"] = "homeContent运行出错，请检查!!!"
		M["data"] = err.Error()

	} else {
		M["code"] = 0
		M["message"] = "success"
		M["data"] = strings.TrimSpace(string(content))
	}
	jstr, err := json.Marshal(M)
	if err != nil {
		return "{\"code\":0,\"message\":\"格式化json出错，请检查!!!函数名为Go_HomeContent!!\",\"data\":\"\"}"
	} else {
		return string(jstr)
	}
}

func Go_CategoryContent(tid string, pg string, filter bool, extend string, file_name string) string {
	M := make(map[string]interface{})
	flr := "False"
	if filter {
		flr = "True"
	}
	File_Name_Remove_py := strings.TrimSuffix(file_name, ".py")
	cmd := exec.Command("python3", "-c", "from "+File_Name_Remove_py+" import categoryContent;categoryContent(\""+tid+"\",\""+pg+"\","+flr+",\""+extend+"\")")
	content, err := cmd.Output()
	if err != nil {
		M["code"] = 0
		M["message"] = "categoryContent运行出错，请检查!!!"
		M["data"] = err.Error()

	} else {
		M["code"] = 0
		M["message"] = "success"
		M["data"] = strings.TrimSpace(string(content))
	}
	jstr, err := json.Marshal(M)
	if err != nil {
		return "{\"code\":0,\"message\":\"格式化json出错，请检查!!!函数名为Go_CategoryContent!!\",\"data\":\"\"}"
	} else {
		return string(jstr)
	}
}

func Go_DetailContent(ids string, file_name string) string {
	M := make(map[string]interface{})
	File_Name_Remove_py := strings.TrimSuffix(file_name, ".py")
	cmd := exec.Command("python3", "-c", "from "+File_Name_Remove_py+" import detailContent;detailContent(\""+ids+"\")")
	content, err := cmd.Output()
	if err != nil {
		M["code"] = 0
		M["message"] = "detailContent运行出错，请检查!!!"
		M["data"] = err.Error()

	} else {
		M["code"] = 0
		M["message"] = "success"
		M["data"] = strings.TrimSpace(string(content))
	}
	jstr, err := json.Marshal(M)
	if err != nil {
		return "{\"code\":0,\"message\":\"格式化json出错，请检查!!!函数名为Go_DetailContent!!\",\"data\":\"\"}"
	} else {
		return string(jstr)
	}
}

func Go_SearchContent(key string, file_name string) string {
	M := make(map[string]interface{})
	File_Name_Remove_py := strings.TrimSuffix(file_name, ".py")
	cmd := exec.Command("python3", "-c", "from "+File_Name_Remove_py+" import searchContent;searchContent(\""+key+"\")")
	content, err := cmd.Output()
	if err != nil {
		M["code"] = 0
		M["message"] = "searchContent运行出错，请检查!!!"
		M["data"] = err.Error()

	} else {
		M["code"] = 0
		M["message"] = "success"
		M["data"] = strings.TrimSpace(string(content))
	}
	jstr, err := json.Marshal(M)
	if err != nil {
		return "{\"code\":0,\"message\":\"格式化json出错，请检查!!!函数名为Go_SearchContent!!\",\"data\":\"\"}"
	} else {
		return string(jstr)
	}
}

func Go_PlayerContent(flag string, id string, file_name string) string {
	M := make(map[string]interface{})
	File_Name_Remove_py := strings.TrimSuffix(file_name, ".py")
	cmd := exec.Command("python3", "-c", "from "+File_Name_Remove_py+" import playerContent;playerContent(\""+flag+"\",\""+id+"\")")
	content, err := cmd.Output()
	if err != nil {
		M["code"] = 0
		M["message"] = "playerContent运行出错，请检查!!!"
		M["data"] = err.Error()

	} else {
		M["code"] = 0
		M["message"] = "success"
		M["data"] = strings.TrimSpace(string(content))
	}
	jstr, err := json.Marshal(M)
	if err != nil {
		return "{\"code\":0,\"message\":\"格式化json出错，请检查!!!函数名为Go_PlayerContent!!\",\"data\":\"\"}"
	} else {
		return string(jstr)
	}
}

// 读取yaml配置文件
func ReadYaml() {
	data, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Println("读取config.yaml文件失败!请检查!")
		return
	}
	var obj map[string]interface{}
	err = yaml.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("解析config.yaml文件失败!请检查!")
	}
}

func main() {
	ReadYaml()
}
