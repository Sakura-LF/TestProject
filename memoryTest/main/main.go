package main

import "sync"

var matrix3x3 = [][]int32{
	{1, 2, 3},
	{4, 5, 6},
	{7, 8, 9},
}

// 创建一个4x4的矩阵
var matrix4x4 = [][]int32{
	{11, 2, 4},
	{4, 5, 6},
	{10, 8, -12},
}

//func main() {
//	// 获取子数组的个数
//	cols := len(matrix4x4[0])
//
//	var leftNum int32 = 0
//	var rightNum int32 = 0
//	//leftNum, rightNum := 0, 0
//	left, right := 0, cols-1
//	// 遍历矩阵
//	for i, _ := range matrix4x4 {
//		fmt.Println(matrix4x4[i][left])
//		leftNum += matrix4x4[i][left]
//		rightNum += matrix4x4[i][right]
//
//		fmt.Println("right", matrix4x4[i][right])
//
//		left++
//		right--
//	}
//	fmt.Println(leftNum, rightNum)
//	i := int32(math.Abs(float64(leftNum - rightNum)))
//
//	fmt.Println(i)
//
//}

func main() {
	////var slice = []string{"$200.49", "$1,999.00", "$99", "50.00美元"}
	//text := "$200.49、$1,999.00、$99、50.00美元"
	//// 正则表达式用于匹配给定的格式
	//pattern := `\$[0-9]{1,3}(?:,[0-9]{3})*(?:\.[0-9]{2})?|[0-9]+(?:\.[0-9]{2})?美元`
	//re := regexp.MustCompile(pattern)
	//matches := re.FindAllString(text, -1)
	//
	//fmt.Println(matches)
	//// 遍历匹配结果
	//for _, match := range matches {
	//	// 去掉美元符号,汉字,逗号
	//	cleanMatch := strings.Replace(match, "$", "", -1)
	//	cleanMatch = strings.Replace(cleanMatch, "美元", "", -1)
	//	cleanMatch = strings.Replace(cleanMatch, ",", "", -1)
	//
	//	fmt.Printf("%s、", cleanMatch)
	//}

	s := sync.Map{}
	s.Store()
}
