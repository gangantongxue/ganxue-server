package main

import (
	"fmt"
	"os"
	"syscall"

	seccomp "github.com/seccomp/libseccomp-golang"
)

func main() {
	filter, err := seccomp.NewFilter(seccomp.ActErrno.SetReturnCode(int16(syscall.EPERM)))
	if err != nil {
		fmt.Println("创建Seccomp失败:", err)
		os.Exit(1)
	}

	// 允许的系统调用列表
	allowedSyscalls := []string{
		// 原有条目保持不变
		"read", "write", "exit", "sigreturn",
		"futex", "mmap", "munmap", "brk",
		"rt_sigaction", "clock_gettime", "gettimeofday",
		"getpid", "gettid",

		// 新增关键系统调用
		"clone",       // 线程/进程创建
		"exit_group",  // 进程组退出（Go运行时使用）
		"wait4",       // 等待子进程
		"sigaltstack", // 信号栈管理
		"prlimit64",   // 资源限制设置
		"mprotect",    // 内存保护设置
		"close",       // 关闭文件描述符
		"ioctl",       // 设备控制（终端输入相关）
	}

	// 添加允许规则（基础系统调用）
	for _, syscallName := range allowedSyscalls {
		syscallNum, err := seccomp.GetSyscallFromName(syscallName)
		if err != nil {
			fmt.Printf("警告：系统调用 %s 不存在，跳过\n", syscallName)
			continue
		}
		filter.AddRule(syscallNum, seccomp.ActAllow)
	}

	// 加载过滤规则
	if err := filter.Load(); err != nil {
		fmt.Println("加载Seccomp策略失败:", err)
		os.Exit(1)
	}

	UserCode() // 用户代码
}
