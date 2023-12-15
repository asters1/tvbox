package main

import (
	"C"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/tidwall/gjson"
)

import (
	"strconv"
)

var (
	extend              string
	filter_switch       bool
	filter_type_index   int
	filter_num_index    int
	search_switch       bool
	search_keyword      string
	test_type_index     int
	test_category_page  string
	test_vod_index      int
	test_vod_from_index int
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

// 解析homeContent
//
//export Go_HomeContent
func Go_HomeContent(etd *C.char, filter bool, file_name *C.char) *C.char {
	c_etd := C.GoString(etd)
	c_filter := filter
	c_file_name := C.GoString(file_name)

	M := make(map[string]interface{})
	flr := "False"
	if c_filter {
		flr = "True"
	}
	head, c_file_name := filepath.Split(c_file_name)
	File_Name_Remove_py := strings.TrimSuffix(c_file_name, ".py")
	// fmt.Println("python3", "-c", "from "+File_Name_Remove_py+" import homeContent,init;init(\""+etd+"\");homeContent("+flr+")")
	// fmt.Println("python3", "-c", "import sys;sys.path.append(\""+head+"\");from "+File_Name_Remove_py+" import homeContent,init;init(\""+etd+"\");homeContent("+flr+")")
	cmd := exec.Command("python3", "-c", "import sys;sys.path.append(\""+head+"\");from "+File_Name_Remove_py+" import homeContent,init;init(\""+c_etd+"\");homeContent("+flr+")")
	content, err := cmd.Output()
	if err != nil {
		M["code"] = 0
		M["message"] = "homeContent运行出错，请检查!!!"
		M["data"] = err.Error()
		fmt.Printf("%v\n", cmd)

	} else {
		M["code"] = 1
		M["message"] = "success"
		M["data"] = strings.TrimSpace(string(content))
	}
	jstr, err := json.Marshal(M)
	if err != nil {
		return C.CString("{\"code\":0,\"message\":\"格式化json出错，请检查!!!函数名为Go_HomeContent!!\",\"data\":\"\"}")
	} else {
		return C.CString(string(jstr))
	}
}

func Go_CategoryContent(etd *C.char, tid *C.char, pg *C.char, filter bool, extend *C.char, file_name *C.char) *C.char {
	c_etd := C.GoString(etd)
	c_tid := C.GoString(tid)
	c_pg := C.GoString(pg)
	c_extend := C.GoString(extend)
	c_filter := filter
	c_file_name := C.GoString(file_name)
	M := make(map[string]interface{})
	flr := "False"
	if c_filter {
		flr = "True"
	}
	head, c_file_name := filepath.Split(c_file_name)
	File_Name_Remove_py := strings.TrimSuffix(c_file_name, ".py")
	cmd := exec.Command("python3", "-c", "import sys;sys.path.append(\""+head+"\");from "+File_Name_Remove_py+" import categoryContent,init;init(\""+c_etd+"\");categoryContent(\""+c_tid+"\",\""+c_pg+"\","+flr+",\""+c_extend+"\")")
	content, err := cmd.Output()
	if err != nil {
		M["code"] = 0
		M["message"] = "categoryContent运行出错，请检查!!!"
		M["data"] = err.Error()
		fmt.Printf("%v\n", cmd)

	} else {
		M["code"] = 0
		M["message"] = "success"
		M["data"] = strings.TrimSpace(string(content))
	}
	jstr, err := json.Marshal(M)
	if err != nil {
		return C.CString("{\"code\":0,\"message\":\"格式化json出错，请检查!!!函数名为Go_CategoryContent!!\",\"data\":\"\"}")
	} else {
		return C.CString(string(jstr))
	}
}

func Go_DetailContent(etd *C.char, ids *C.char, file_name *C.char) *C.char {
	c_etd := C.GoString(etd)
	c_ids := C.GoString(ids)
	c_file_name := C.GoString(file_name)
	M := make(map[string]interface{})
	head, c_file_name := filepath.Split(c_file_name)
	File_Name_Remove_py := strings.TrimSuffix(c_file_name, ".py")
	cmd := exec.Command("python3", "-c", "import sys;sys.path.append(\""+head+"\");from "+File_Name_Remove_py+" import detailContent,init;init(\""+c_etd+"\");detailContent(\""+c_ids+"\")")
	content, err := cmd.Output()
	if err != nil {
		M["code"] = 0
		M["message"] = "detailContent运行出错，请检查!!!"
		M["data"] = err.Error()
		fmt.Printf("%v\n", cmd)

	} else {
		M["code"] = 0
		M["message"] = "success"
		M["data"] = strings.TrimSpace(string(content))
	}
	jstr, err := json.Marshal(M)
	if err != nil {
		return C.CString("{\"code\":0,\"message\":\"格式化json出错，请检查!!!函数名为Go_DetailContent!!\",\"data\":\"\"}")
	} else {
		return C.CString(string(jstr))
	}
}

func Go_SearchContent(etd *C.char, key *C.char, file_name *C.char) *C.char {
	c_etd := C.GoString(etd)
	c_key := C.GoString(key)
	c_file_name := C.GoString(file_name)
	M := make(map[string]interface{})
	head, c_file_name := filepath.Split(c_file_name)
	File_Name_Remove_py := strings.TrimSuffix(c_file_name, ".py")
	cmd := exec.Command("python3", "-c", "import sys;sys.path.append(\""+head+"\");from "+File_Name_Remove_py+" import searchContent,init;init(\""+c_etd+"\");searchContent(\""+c_key+"\")")
	content, err := cmd.Output()
	if err != nil {
		M["code"] = 0
		M["message"] = "searchContent运行出错，请检查!!!"
		M["data"] = err.Error()
		fmt.Printf("%v\n", cmd)

	} else {
		M["code"] = 0
		M["message"] = "success"
		M["data"] = strings.TrimSpace(string(content))
	}
	jstr, err := json.Marshal(M)
	if err != nil {
		return C.CString("{\"code\":0,\"message\":\"格式化json出错，请检查!!!函数名为Go_SearchContent!!\",\"data\":\"\"}")
	} else {
		return C.CString(string(jstr))
	}
}

func Go_PlayerContent(etd string, flag string, id string, file_name string) string {
	M := make(map[string]interface{})
	head, file_name := filepath.Split(file_name)
	File_Name_Remove_py := strings.TrimSuffix(file_name, ".py")
	cmd := exec.Command("python3", "-c", "import sys;sys.path.append(\""+head+"\");from "+File_Name_Remove_py+" import playerContent,init;init(\""+etd+"\");playerContent(\""+flag+"\",\""+id+"\")")
	content, err := cmd.Output()
	if err != nil {
		M["code"] = 0
		M["message"] = "playerContent运行出错，请检查!!!"
		M["data"] = err.Error()
		fmt.Printf("%v\n", cmd)

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
func InitConfig() {
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
	for k, v := range obj {
		switch {
		case k == "extend":
			if str, ok := v.(string); ok {
				extend = str
			} else {
				fmt.Println("extend不是字符串类型!请检查!")
			}
		case k == "filter_switch":
			if bl, ok := v.(bool); ok {
				filter_switch = bl
			} else {
				fmt.Println("filter_switch不是bool类型!请检查!")
			}
		case k == "filter_type_index":
			if num, ok := v.(int); ok {
				filter_type_index = num
			} else {
				fmt.Println("filter_type_index不是int类型!请检查!")
			}
		case k == "filter_num_index":
			if num, ok := v.(int); ok {
				filter_num_index = num
			} else {
				fmt.Println("filter_num_index不是int类型!请检查!")
			}
		case k == "search_switch":
			if bl, ok := v.(bool); ok {
				search_switch = bl
			} else {
				fmt.Println("search_switch不是bool类型!请检查!")
			}
		case k == "search_keyword":
			if str, ok := v.(string); ok {
				search_keyword = str
			} else {
				fmt.Println("search_keyword不是字符串类型!请检查!")
			}
		case k == "test_type_index":
			if num, ok := v.(int); ok {
				test_type_index = num
			} else {
				fmt.Println("test_type_index不是int类型!请检查!")
			}
		case k == "test_category_page":
			if str, ok := v.(string); ok {
				test_category_page = str
			} else {
				fmt.Println("test_category_page不是字符串类型!请检查!")
			}
		case k == "test_vod_index":
			if num, ok := v.(int); ok {
				test_vod_index = num
			} else {
				fmt.Println("test_vod_index不是int类型!请检查!")
			}
		case k == "test_vod_from_index":
			if num, ok := v.(int); ok {
				test_vod_from_index = num
			} else {
				fmt.Println("test_vod_from_index不是int类型!请检查!")
			}

		}
	}
}

func debug_homeContent(file_path string) (string, string) {
	resurl_filter := ""
	f_name := ""
	res_homeContent := C.GoString(Go_HomeContent(C.CString(extend), filter_switch, C.CString(file_path)))
	if gjson.Get(res_homeContent, "code").Int() == 1 {
		tid := ""
		type_name := ""
		fmt.Println("\n=======homeContent=======\n")
		res_data := gjson.Get(res_homeContent, "data").String()
		if res_data == "" {
			fmt.Println("homeContent返回为空")
			os.Exit(0)
		}
		for i := 0; i < len(gjson.Get(res_data, "class").Array()); i++ {
			// fmt.Println(res_data)
			tn := gjson.Get(res_data, "class."+strconv.Itoa(i)+".type_name").String()
			ti := gjson.Get(res_data, "class."+strconv.Itoa(i)+".type_id").String()
			if tn == "" || ti == "" {
				fmt.Println("没有检测到type_id或者type_name!!请检查")
				os.Exit(0)
			}
			fmt.Print(tn + "[" + ti + "] ")
			if i == test_type_index {
				tid = gjson.Get(res_data, "class."+strconv.Itoa(i)+".type_id").String()
				type_name = gjson.Get(res_data, "class."+strconv.Itoa(i)+".type_name").String()
			}
		}
		fmt.Println("\n")
		if filter_switch {
			fmt.Println("筛选开关已开启!\n")
			// fmt.Println(gjson.Get(res_data, "class").Raw)
			for i := 0; i < len(gjson.Get(res_data, "class").Array()); i++ {
				// fmt.Println(i)
				ti := gjson.Get(res_data, "class."+strconv.Itoa(i)+".type_id").String()
				res_data_filter := gjson.Get(res_data, "filter."+ti).Array()

				for j := 0; j < len(res_data_filter); j++ {

					// fmt.Println(j)
					fmt.Print(gjson.Get(res_data_filter[j].String(), "name").String() + " ")
					f_value := gjson.Get(res_data_filter[j].String(), "value").Array()
					for k := 0; k < len(f_value); k++ {

						n := gjson.Get(f_value[k].String(), "n").String()
						v := gjson.Get(f_value[k].String(), "v").String()
						fmt.Print(n + " ")
						if j == filter_type_index && k == filter_num_index {

							resurl_filter = "{\\\"" + gjson.Get(res_data_filter[j].String(), "key").String() + "\\\":\\\"" + v + "\\\"}"
							f_name = "[" + gjson.Get(res_data_filter[j].String(), "name").String() + ":" + n + "]"
						}
					}
					fmt.Println()
				}
			}

		}
		fmt.Println("\n你测试的类型是->" + type_name + "[" + tid + "]")
		fmt.Println("\n你测试的筛选是:" + f_name + "->" + strings.ReplaceAll(resurl_filter, "\\", ""))

		return tid, resurl_filter
	} else {
		os.Exit(0)
	}
	return "", ""
}

func debug_categoryContent(tid string, pg string, json_filter string, file_path string) string {
	ids := ""
	ids_name := ""

	fmt.Println("\n=======categoryContent=======\n")
	res_categoryContent := C.GoString(Go_CategoryContent(C.CString(extend), C.CString(tid), C.CString(pg), filter_switch, C.CString(json_filter), C.CString(file_path)))
	res_data := gjson.Get(res_categoryContent, "data").String()
	if res_data == "" {
		fmt.Println("categoryContent返回为空")
		os.Exit(0)
	}
	// fmt.Println(res_data)
	res_pg := gjson.Get(res_data, "page").String()
	res_pagecount := gjson.Get(res_data, "pagecount").String()
	res_limit := gjson.Get(res_data, "limit").String()
	res_total := gjson.Get(res_data, "total").String()
	res_list := gjson.Get(res_data, "list").Array()
	fmt.Println("//当前页")
	if res_pg == "" {
		fmt.Println("没有解析到page,请检查!")
		os.Exit(0)
	} else {
		fmt.Println("page:" + res_pg + "\n")
	}
	fmt.Println("//总共几页")
	if res_pagecount == "" {
		fmt.Println("没有解析到pagecount,请检查!")
		os.Exit(0)
	} else {
		fmt.Println("page:" + res_pagecount + "\n")
	}
	fmt.Println("//每页几条数据")
	if res_limit == "" {
		fmt.Println("没有解析到limit,请检查!")
		os.Exit(0)
	} else {
		fmt.Println("page:" + res_limit + "\n")
	}
	fmt.Println("//总共多少条数据")
	if res_total == "" {
		fmt.Println("没有解析到total,请检查!")
		os.Exit(0)
	} else {
		fmt.Println("page:" + res_total + "\n")
	}
	fmt.Println("//视频列表")
	for i := 0; i < len(res_list); i++ {
		vod_id := gjson.Get(res_list[i].String(), "vod_id").String()
		vod_name := gjson.Get(res_list[i].String(), "vod_name").String()
		fmt.Println(vod_name + "[" + vod_id + "]")
		if test_vod_index == i {
			ids = vod_id
			ids_name = vod_name
		}

	}
	fmt.Println("\n你测试的视频是->" + ids + "[" + ids_name + "]")
	return ids
}

func debug_searchContent(key string, file_path string) string {
	ids := ""
	ids_name := ""
	fmt.Println("\n你测试的搜索关键字是->" + key)
	fmt.Println("\n=======searchContent=======\n")
	res_searchContent := C.GoString(Go_SearchContent(C.CString(extend), C.CString(key), C.CString(file_path)))
	res_data := gjson.Get(res_searchContent, "data").String()
	if res_data == "" {
		fmt.Println("searchContent返回为空")
		os.Exit(0)
	}
	// fmt.Println(res_data)
	list := gjson.Get(res_data, "list").Array()
	for i := 0; i < len(list); i++ {
		vod_name := gjson.Get(list[i].String(), "vod_name").String()
		vod_id := gjson.Get(list[i].String(), "vod_id").String()
		fmt.Println(vod_name + "[" + vod_id + "]")

		if test_vod_index == i {
			ids = vod_id
			ids_name = vod_name
		}
	}
	fmt.Println("\n你测试的视频是->" + ids + "[" + ids_name + "]")
	return ids
}

func debug_detailContent(ids string, file_path string) {
	fmt.Println(ids)
	res_detailContent := C.GoString(Go_DetailContent(C.CString(extend), C.CString(ids), C.CString(file_path)))
	fmt.Println(res_detailContent)
}

func main() {
	InitConfig()
	if len(os.Args) == 1 {
		fmt.Println("请输入测试的py文件!!!")
		return
	}
	file_path := os.Args[1]
	ids := ""
	fmt.Println("你测试的文件是===>" + file_path)
	if search_switch {
		ids = debug_searchContent(search_keyword, file_path)
	} else {

		tid, json_filter := debug_homeContent(file_path)
		// fmt.Println(tid, filter)
		ids = debug_categoryContent(tid, test_category_page, json_filter, file_path)

	}
	debug_detailContent(ids, file_path)
	fmt.Println("\n")
}
