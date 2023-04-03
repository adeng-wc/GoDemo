package main

import (
	"GoDemo/generics"
	"fmt"
)

//var visits = expvar.NewInt("visits")

//func handler(w http.ResponseWriter, r *http.Request) {
//	visits.Add(1)
//	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
//}

func main() {

	ints := map[string]int{
		"first":  34,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("非泛型计算结果，SumInts: %v, SumFloats: %v\n", generics.SumInts(ints), generics.SumFloats(floats))
	fmt.Printf("泛型计算结果，Ints 结果: Floats 结果: %v\n", generics.SumIntsOrFloats[string, int](ints), generics.SumIntsOrFloats[string, float64](floats))
	fmt.Printf("泛型计算结果，Ints 结果: Floats 结果: %v\n", generics.SumIntsOrFloats(ints), generics.SumIntsOrFloats(floats))

	fmt.Printf("泛型计算结果（带 Constraint），Ints 结果: %v, Floats 结果: %v\n",
		generics.SumNumbers(ints),
		generics.SumNumbers(floats))

	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", nil)

	//t, _ := time.Parse("2006-01-02 15:04:05", "2016-06-13 09:14:00")
	//fmt.Println(time.Now().Sub(t).Hours())
	//fmt.Println(t.Format("2006-01-02 15:04:05"))

	//
	//t, _ := time.ParseInLocation("2006-01-02 15:04:05", "2016-06-13 15:34:39", time.Local)
	//// 整点（向下取整）
	//fmt.Println(t.Truncate(1 * time.Hour))
	//// 整点（最接近）
	//fmt.Println(t.Round(1 * time.Hour))
	//
	//// 整分（向下取整）
	//fmt.Println(t.Truncate(1 * time.Minute))
	//// 整分（最接近）
	//fmt.Println(t.Round(1 * time.Minute))
	//
	//t2, _ := time.ParseInLocation("2006-01-02 15:04:05", t.Format("2006-01-02 15:00:00"), time.Local)
	//fmt.Println(t2)

	//str := "{\"appId\":13,\"instance\":\"10.221.54.102\",\"timestamp\":1679487043616,\"collectName\":\"go_vm\",\"data\":{\"CPU\":[{\"cpu.busy\":5.709487830602782,\"cpu.guest\":0,\"cpu.idle\":94.29051216939722,\"cpu.iowait\":0,\"cpu.irq\":0,\"cpu.nice\":0,\"cpu.softirq\":0,\"cpu.steal\":0,\"cpu.system\":1.8471872376021128,\"cpu.usageRate\":7.238381812045507,\"cpu.user\":3.8623005873802843}]}}"
	//fmt.Printf("监控数据开始上报 %s\n", str)
	//log.Printf("监控数据开始上报 %s\n", str)
	//fmt.Println("Hello, World!")

	//for {
	//DONE:
	//	select {
	//	case <-time.After(time.Second):
	//		fmt.Println("Hello, World!")
	//		continue
	//	default:
	//		fmt.Println("default!")
	//		break DONE
	//	}
	//
	//	fmt.Println("break 2!")
	//	break
	//}

	//t, _ := tail.TailFile("/Users/wengcheng/CodeWorkspces/Netease/gameserver/logs/spring-web/AccessLog.log", tail.Config{Follow: true})
	//for line := range t.Lines {
	//	fmt.Println(line.Text)
	//}

	//line := "[2023-03-13 17:45:36] [http-nio-8080-exec-9] [userId=] {\"sender\":\"0:0:0:0:0:0:0:1\",\"method\":\"GET\",\"path\":\"/test/treasure/hello\",\"params\":{\"userId\":\"222\"},\"headers\":{\"content-length\":\"32\",\"postman-token\":\"f5e98b4a-6b50-41b4-92b4-6456241ea66c\",\"host\":\"localhost:8080\",\"content-type\":\"application/json\",\"connection\":\"keep-alive\",\"cache-control\":\"no-cache\",\"accept-encoding\":\"gzip, deflate, br\",\"user-agent\":\"PostmanRuntime/7.31.1\",\"accept\":\"*/*\"}}\n"
	//regTag, _ := regexp.Compile("(test/.*)")
	//t := regTag.FindStringSubmatch(line)
	//fmt.Println(t)
	//fmt.Println(len(t))
	//re := regexp.MustCompile( `a(x*)b(y|z)c` )
	//fmt.Printf( "%q \n " , re.FindStringSubmatch( "-axxxbyc-" ))
	//fmt.Printf( "%q \n " , re.FindStringSubmatch( "-abzc-" ))

}
