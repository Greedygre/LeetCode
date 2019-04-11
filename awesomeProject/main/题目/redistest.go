//package main
//
//import (
//	"fmt"
//	"github.com/garyburd/redigo/redis"
//)
//
////
////client = redis.StrictRedis()
////for i in range(1000):
////   client.pfadd("codehole", "user%d" % i)
////   total = client.pfcount("codehole")
////   if total != i+1:
////       print total, i+1
////       break
//
//func main() {
//	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
//	conn.Do("del", "codehole")
//	if err != nil {
//		fmt.Println("connect redis error :", err)
//		return
//	}
//	defer conn.Close()
//	i := 0
//	for i <= 10000 {
//		i++
//
//		_, err := conn.Do("bf.add", "codehole", i)
//		if err != nil {
//			fmt.Println(err)
//		}
//		total, err := conn.Do("bf.exists", "codehole", i)
//		if err != nil {
//			fmt.Println(err)
//		}
//		if total == 0 {
//			fmt.Println("total:", i)
//		}
//
//	}
//
//}
//
////package main
////
////import (
////	"fmt"
////	"io/ioutil"
////	"net/http"
////	"unicode"
////)
////
////func main() {
////	resp, err := http.Get("https://www.jianshu.com/p/85f4624485b9")
////	defer resp.Body.Close()
////
////	if err != nil {
////		fmt.Println("error: ", err)
////	} else {
////		b, _ := ioutil.ReadAll(resp.Body)
////		str := string(b)
////		r := []rune(str)
////		//fmt.Println("rune=", r)
////		strSlice := []string{}
////		cnstr := ""
////		for i := 0; i < len(r); i++ {
////			if (r[i] <= 40869 && r[i] >= 19968) || r[i] == 0 || r[i] == 10 || r[i] == 63 {
////				cnstr = cnstr + string(r[i])
////				strSlice = append(strSlice, cnstr)
////
////			}
////			//fmt.Println("r[", i, "]=", r[i], "string=", string(r[i]))
////		}
////		if 0 == len(strSlice) {
////			//无中文，需要跳过，后面再找规律
////		}
////		fmt.Println("提取出的中文字符串:", cnstr)
////		fmt.Println(strSlice)
////		//fmt.Println(c)
////	}
////}
////func IsChineseChar(str string) bool {
////	for _, r := range str {
////		if unicode.Is(unicode.Scripts["Han"], r) {
////			fmt.Print(r)
////		}
////	}
////	return false
////}
package 题目

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:         []string{"127.0.0.1:6379"},
		ReadOnly:      true,
		RouteRandomly: true,
	})
	fmt.Println("lulalal")
	pipe := client.Pipeline()
	pipe.HGetAll("1371648200")
	pipe.HGetAll("1371648300")
	pipe.HGetAll("1371648400")
	cmders, err := pipe.Exec()
	if err != nil {
		fmt.Println("err1", err)
	}
	for _, cmder := range cmders {
		cmd := cmder.(*redis.StringStringMapCmd)
		strMap, err := cmd.Result()
		if err != nil {
			fmt.Println("err2", err)
		}
		fmt.Println("strMap", strMap)
	}
}
