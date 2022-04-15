package Gen_series

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"

	"github.com/Chain-Zhang/pinyin"
)

var Id_len = 1017

func Choose(start int, end int) int {
	choice := rand.Intn(end-start+1) + start
	return choice
}

func Gen_id() string {

	filepath := "Person\\Chinese_Hacker_ID.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file wrong!", err)
	}
	defer file.Close()

	choice := Choose(0, Id_len)

	bufReader := bufio.NewReader(file)
	choice_line := ""

	for i := 0; i <= choice; i++ {

		line, err := bufReader.ReadString('\n')
		if i == choice {
			choice_line = strings.TrimSpace(line)
		}
		if err != nil {
			fmt.Println("read error!", err)
		}
	}

	return choice_line
}

func Gen_first_name() string {
	first_name_list := []string{"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "褚", "卫", "蒋", "沈", "韩", "杨", "朱", "秦", "尤", "许",
		"何", "吕", "施", "张", "孔", "曹", "严", "华", "金", "魏", "陶", "姜", "戚", "谢", "邹", "喻", "柏", "水", "窦", "章",
		"云", "苏", "潘", "葛", "奚", "范", "彭", "郎", "鲁", "韦", "昌", "马", "苗", "凤", "花", "方", "俞", "任", "袁", "柳",
		"酆", "鲍", "史", "唐", "费", "廉", "岑", "薛", "雷", "贺", "倪", "汤", "滕", "殷", "罗", "毕", "郝", "邬", "安", "常",
		"乐", "于", "时", "傅", "皮", "卞", "齐", "康", "伍", "余", "元", "卜", "顾", "孟", "平", "黄", "和", "穆", "萧", "尹",
		"姚", "邵", "堪", "汪", "祁", "毛", "禹", "狄", "米", "贝", "明", "臧", "计", "伏", "成", "戴", "谈", "宋", "茅", "庞",
		"熊", "纪", "舒", "屈", "项", "祝", "董", "梁"}

	choice := Choose(0, len(first_name_list))
	f_name := first_name_list[choice]
	return f_name
}

func GBK_2312() string {
	head := Choose(0xb0, 0xba)
	body := Choose(0xa1, 0xf9)
	full := fmt.Sprintf("%X%X", head, body)
	b, _ := hex.DecodeString(full)
	var decodeBytes, _ = simplifiedchinese.GBK.NewDecoder().Bytes(b)
	goal := string(decodeBytes)
	return goal
}

func Gen_second_name() string {
	second_name_list := []string{GBK_2312(), ""}
	choice := Choose(0, 1)
	second_name := second_name_list[choice]
	return second_name
}

func Gen_third_name() string {
	third_name := GBK_2312()
	return third_name
}

func Gen_name() string {
	name := Gen_first_name() + Gen_second_name() + Gen_third_name()
	return name
}

func Gen_pinyin(name string) string {
	str, err := pinyin.New(name).Split("").Mode(pinyin.InitialsInCapitals).Convert()
	if err != nil {
		fmt.Println("pinyin error",err)
	} 
	return str
}

func Gen_id_card(age int, gender int) string {

	now := time.Now()
	choice := Choose(0, 3149)
	area_code := ""
	for k := range area_dict {
		if choice == 0 {
			area_code = k
		}
		choice--
	}

	//check
	id_code_list := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	check_code_list := []int{1, 0, 10, 9, 8, 7, 6, 5, 4, 3, 2} //use 10 as X
	year_str := fmt.Sprintf("%d", now.Year()-age)
	month_str := fmt.Sprintf("%02d", Choose(1, 12))
	day_str := ""
	switch month_str {
	case "2":
		day_str = fmt.Sprintf("%02d", Choose(1, 28))
	case "4", "6", "9", "11":
		day_str = fmt.Sprintf("%02d", Choose(1, 30))
	default:
		day_str = fmt.Sprintf("%02d", Choose(1, 31))
	}
	date_string := year_str + month_str + day_str

	rd := Choose(0, 999)
	var place_gender_number int
	if gender == 0 {
		if rd%2 == 0 {
			place_gender_number = rd
		} else {
			place_gender_number = rd + 1
		}
	} else {
		if rd%2 == 1 {
			place_gender_number = rd
		} else {
			place_gender_number = rd - 1
		}
	}
	place_gender_str := fmt.Sprintf("%03d", place_gender_number)
	result := area_code + date_string + place_gender_str
	var check_res int
	for i := 0; i < len(id_code_list); i++ {
		check_res += int(result[i]) * id_code_list[i]
	}
	check_res = check_res % 11
	last_code := fmt.Sprintf("%d", check_code_list[check_res])
	if last_code == "10" {
		last_code = "X"
	}
	last_res := result + last_code
	return last_res
}

func Gen_creditcode() map[string]string {
	org_code := map[string]string{"1": "机构编制", "2": "外交", "3": "教育", "4": "公安", "5": "民政", "6": "司法", "7": "交通运输", "8": "文化", "9": "工商", "A": "中央军委改革和编制办公室", "N": "农业", "Y": "其他"}

	ic_code := map[string]map[string]string{"1": {"1": "机关", "2": "事业单位", "3": "中央编办直接管理机构编制的群众团体", "9": "其他"}, "2": {"1": "外国常驻新闻机构", "9": "其他"}, "3": {"1": "律师执业机构", "2": "公证处", "3": "基层法律服务所", "4": "司法鉴定机构", "5": "仲裁委员会", "9": "其他"}, "4": {"1": "外国在华文化中心", "9": "其他"}, "5": {"1": "社会团体", "2": "民办非企业单位", "3": "基金会", "9": "其他"}, "6": {"1": "外国旅游部门常驻机构代表机构", "2": "港澳台地区旅游部门常驻内地（大陆）代表机构", "9": "其他"}, "7": {"1": "宗教活动场所", "2": "宗教院校", "9": "其他"}, "8": {"1": "基层工会", "9": "其他"}, "9": {"1": "企业", "2": "个体工商户", "3": "农民专业社"}, "A": {"1": "军队事业单位", "9": "其他"}, "N": {"1": "组级集体经济组织", "2": "村级集体经济组织", "3": "乡镇集体经济组织", "9": "其他"}, "Y": {"1": "其他"}}
	areas := []string{"110000", "110101", "110102", "110103", "110104", "110105", "110106", "110107", "110108", "110109", "110111", "110112", "110113", "110114", "110115", "110116", "110117", "110228", "110229", "120000", "120101", "120102", "120103", "120104", "120105", "120106", "120107", "120108", "120109", "120110", "120111", "120112", "120113", "120114", "120115", "120221", "120223", "120225", "130000", "130100", "130102", "130103", "130104", "130105", "130107", "130108", "130121", "130123", "130124", "130125", "130126", "130127", "130128", "130129", "130130", "130131", "130132", "130133", "130181", "130182", "130183", "130184", "130185", "130200", "130202", "130203", "130204", "130205", "130207", "130208", "130223", "130224", "130225", "130227", "130229", "130230", "130281", "130283", "130300", "130302", "130303", "130304", "130321", "130322", "130323", "130324", "130400", "130402", "130403", "130404", "130406", "130421", "130423", "130424", "130425", "130426", "130427", "130428", "130429", "130430"}
	choice1 := Choose(0, len(org_code)-1)

	regorg := ""
	for k := range org_code {
		if choice1 == 0 {
			regorg = k
		}
		choice1--
	}
	cho_org := ic_code[regorg]
	choice2 := Choose(0, len(cho_org)-1)
	orgtype := ""
	for k := range cho_org {
		if choice2 == 0 {
			orgtype = k
		}
		choice2--
	}

	ot := ic_code[regorg][orgtype]
	choice3 := Choose(0, len(areas)-1)
	area := areas[choice3]
	num := Choose(10000000, 99999999)
	num_str := fmt.Sprintf("%d", num)
	ws := []int{3, 7, 9, 10, 5, 8, 4, 2}
	sum := 0
	for i := 0; i < len(ws); i++ {
		sum += int(num_str[i]) * ws[i]
	}
	c9 := 11 - (sum % 11)
	c9_str := ""
	if c9 == 11 {
		c9_str = "0"
	} else if c9 == 10 {
		c9_str = "X"
	} else {
		c9_str = fmt.Sprintf("%d", c9)
	}
	orgcode := num_str + c9_str
	code := regorg + orgtype + area + orgcode

	s := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	wi := []int{1, 3, 9, 27, 19, 26, 16, 17, 20, 29, 25, 13, 8, 24, 10, 30, 28}
	sum = 0
	for i := 0; i < len(wi); i++ {
		if c9_str == "X" {
			sum += 33 * wi[i]
			continue
		}
		sum += int(code[i]) * wi[i]
	}
	c18 := 31 - (sum % 31)
	c18_str := ""
	if c18 == 31 {
		c18_str = "0"
	} else {
		c18_str = string(s[c18])
	}
	res := map[string]string{code + c18_str: ot}
	return res
}

func Gen_orgcode() string {
	num := Choose(10000000, 99999999)
	num_str := fmt.Sprintf("%d", num)
	ws := []int{3, 7, 9, 10, 5, 8, 4, 2}
	sum := 0
	for i := 0; i < len(ws); i++ {
		sum += int(num_str[i]) * ws[i]
	}

	c9 := 11 - (sum % 11)
	c9_str := ""
	if c9 == 11 {
		c9_str = "0"
	} else if c9 == 10 {
		c9_str = "X"
	} else {
		c9_str = fmt.Sprintf("%d", c9)
	}
	return num_str + "-" + c9_str
}

func Gen_mobile() map[string]string {
	prelist := map[string]string{"133": "电信", "149": "电信", "153": "电信", "173": "电信", "177": "电信", "180": "电信", "181": "电信", "189": "电信", "199": "电信", "130": "联通", "131": "联通", "132": "联通", "145": "联通", "155": "联通", "156": "联通",  "171": "联通", "175": "联通", "176": "联通", "185": "联通", "186": "联通", "166": "联通", "134": "移动", "135": "移动", "136": "移动", "137": "移动", "138": "移动", "139": "移动", "147": "移动", "150": "移动", "151": "移动", "152": "移动", "157": "移动", "158": "移动", "159": "移动", "172": "移动", "178": "移动", "182": "移动", "183": "移动", "184": "移动", "187": "移动", "188": "移动", "198": "移动"}
	choice := Choose(0, len(prelist)-1)
	three := ""
	mobile := ""
	for k := range prelist {
		if choice == 0 {
			three = k
		}
		choice--
	}

	mobile = three
	for i := 0; i < 8; i++ {
		mobile += strconv.FormatInt(int64(Choose(0, 9)),10)
	}
	op := prelist[three]
	return map[string]string{mobile: op}
}
