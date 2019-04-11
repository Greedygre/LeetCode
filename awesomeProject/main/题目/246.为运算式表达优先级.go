package 题目

import (
	"fmt"
	"strconv"
)

// 给定一个含有数字和运算符的字符串，为表达式添加括号，改变其运算优先级以求出不同的结果。你需要给出所有可能的组合的结果。有效的运算符号包含 +, - 以及 * 。
//
//示例 1:
//
//输入: "2-1-1"
//输出: [0, 2]
//解释:
//((2-1)-1) = 0
//(2-(1-1)) = 2
//示例 2:
//
//输入: "2*3-4*5"
//输出: [-34, -14, -10, -10, 10]
//解释:
//(2*(3-(4*5))) = -34
//((2*3)-(4*5)) = -14
//((2*(3-4))*5) = -10
//(2*((3-4)*5)) = -10
//(((2*3)-4)*5) = 10
// 思路：感觉是递归
func diffWaysToCompute(input string) []int {

	result := []int{}
	if len(input) == 0 {
		return result
	}
	//a := input[0]
	flag := false
	for index := 0; index < len(input); index++ {
		//fmt.Println("index:", index, "value:", input[index])
		if isDigit(input[index]) {
			continue
		}
		flag = true
		// 这里不是很确定是是这么分
		lList := diffWaysToCompute(input[:index])

		rList := diffWaysToCompute(input[index+1:])
		for _, m := range lList {
			for _, n := range rList {
				cur := 0
				switch input[index] {
				case '+':
					cur = m + n
				case '-':
					cur = m - n
				case '*':
					cur = m * n
				}
				result = append(result, cur)
			}
		}

	}
	if flag == false {
		b, _ := strconv.Atoi(input)
		result = append(result, b)
	}
	return result
}

func main() {
	str := "2-1-1"
	fmt.Println(diffWaysToCompute(str))
}
func isDigit(b uint8) bool {
	if b >= 48 && b <= 57 {
		return true
	}
	return false
}

// class Solution {
//    public List<Integer> diffWaysToCompute(String input) {
//        //分治算法
//        List<Integer> list = new ArrayList<>();
//        if(input.length() == 0) return list;
//        boolean flag = false;
//        for(int i = 0; i < input.length(); ++i){
//            if(Character.isDigit(input.charAt(i))) continue;
//            flag = true;
//            List<Integer> lList = diffWaysToCompute(input.substring(0,i));
//            List<Integer> rList = diffWaysToCompute(input.substring(i+1,input.length()));
//            for(int m : lList){
//                for(int n : rList){
//                    char op = input.charAt(i);
//                    int result = 0;
//                    switch (op){
//                        case '+': result = m + n;break;
//                        case '-': result = m - n;break;
//                        case '*': result = m * n;break;
//                    }
//                    list.add(result);
//                }
//            }
//        }
//        if(flag == false) list.add(Integer.parseInt(input));
//
//        return list;
//    }
//}
