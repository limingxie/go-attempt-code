package maps

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
	"unsafe"
)

func DeclareMap() {

	var mapTemp map[string]string
	fmt.Println(mapTemp)

	var mapInit = map[string]string{"xiaoli": "湖南", "xiaoliu": "天津"}
	fmt.Println(mapInit)

	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	scoreMap["王五"] = 60
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	//删除
	delete(scoreMap, "王五")
	fmt.Println(scoreMap)

	//判断是否有值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}

	//遍历
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	userInfo := map[string]string{
		"username": "xiaoming",
		"password": "123456",
	}
	fmt.Println(userInfo)
}

//map需要按照顺序遍历时处理
func OrderPrint() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

func MapValueSlice() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap, len(sliceMap))
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

func AA() {
	s := make([]int, 9, 20)
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(Len, len(s)) // 9 9

	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(Cap, cap(s)) // 20 20

	mp := make(map[string]int)
	mp["qcrao"] = 100
	mp["stefno"] = 18

	count := **(**int)(unsafe.Pointer(&mp))
	fmt.Println(count, len(mp)) // 2 2
}

func MainMaps() {
	AA()
}
