package main

/*扫描单个端口*/

// import (
// 	"fmt"
// 	"net"
// )
//
// func main() {
// 		_, err := net.Dial("tcp", "scanme.nmap.org:80")
// 		if err == nil {
// 				fmt.Println("Connection successful")
// 		}
// }


/*扫描1024个端口*/

// import (
// 	"fmt"
// )
//
// func main() {
// 	for i := 1; i <= 1024; i++ {
// 		address := fmt.Sprintf("scanme.nmap.org:%d", i)
// 		fmt.Println(address)
// 	}
// }


/*基础端口扫描器*/

// import (
// 	"fmt"
// 	"net"
// )
//
// func main() {
// 	for i := 1; i <= 1024; i++ {
// 		address := fmt.Sprintf("scanme.nmap.org:%d", i)
// 		conn, err := net.Dial("tcp", address)
// 		if err != nil {
// 			//端口关闭
// 			continue
// 		}
// 		conn.Close()
// 		fmt.Printf("%d open\n", i)
// 	}
// }

import (
	"fmt"
	"net"
	"sort"
	"flag"
	"time"
	"os"
)


/*处理工作的worker函数*/

func worker(ports chan int, results chan int, ip string) {
	for p := range ports {
		// fmt.Printf("scan ports %d \n", p)
		address := fmt.Sprintf("%s:%d", ip, p)
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

/*主函数*/
func main() {
	banner := `
 ___  ___  __ _  _____
/ __|/ __|/ _  +/ ___ \
\__ \ (__| (_| || | | +
|___/\___|\__,_||_| |_| \|/ __| |/  /
										 scan version: 1.0.0
	`
	fmt.Println(banner)
	start := time.Now()
	// fmt.Println("开始扫描")
	// ip := "127.0.0.1"
	// ports := make(chan int, 1000)
	results := make(chan int)
	// p := 1024
	var ip string
	var openports []int
	var p int
	var t int

	flag.StringVar(&ip, "u", "", "目标ip")
	flag.IntVar(&p, "p", 1024, "目标端口数从1开始,默认为1-1024")
	flag.IntVar(&t, "t", 100, "线程并发数，默认100")
	flag.Parse()

	ports := make(chan int, t)

	if ip == "" {
		fmt.Println("ip is null")
		flag.Usage()
		os.Exit(0)
	}

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, ip)
	}

	go func() {
		for i := 1; i <= p; i++ {
			ports <- i
		}
	}()

 	for i := 0; i < p; i++ {
		port := <-results

		if port != 0 {
			// fmt.Printf(port)
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	fmt.Printf("IP: %s\n", ip)
	fmt.Printf("Port: 1-%d\n", p)
	for _, port := range openports {
		fmt.Printf("[+] %d open\n", port)
	}
	T := time.Now().Sub(start)
	fmt.Printf("[*] 扫描结束,耗时: %s", T)
}
