package main

import (
	"fmt"
	seccomp "github.com/seccomp/libseccomp-golang"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	//filter, err := seccomp.NewFilter(seccomp.ActErrno.SetReturnCode(int16(syscall.EPERM)))
	filter, err := seccomp.NewFilter(seccomp.ActAllow)

	if err != nil {
		fmt.Println("创建Seccomp失败:", err)
		os.Exit(1)
	}

	// 允许的系统调用列表
	//allowedSyscalls := []string{
	//	// 原有条目保持不变
	//	"read", "write", "exit", "sigreturn",
	//	"futex", "mmap", "munmap", "brk",
	//	"rt_sigaction", "clock_gettime", "gettimeofday",
	//	"getpid", "gettid",
	//
	//	// 新增关键系统调用
	//	"clone",       // 线程/进程创建
	//	"exit_group",  // 进程组退出（Go运行时使用）
	//	"wait4",       // 等待子进程
	//	"sigaltstack", // 信号栈管理
	//	"prlimit64",   // 资源限制设置
	//	"mprotect",    // 内存保护设置
	//	"close",       // 关闭文件描述符
	//	"ioctl",       // 设备控制（终端输入相关）
	//	"openat",      // 代替传统的open调用
	//	"newfstatat",  // 文件状态检查
	//	"faccessat",   // 文件访问权限检查
	//
	//	// 新增架构相关调用
	//	"arch_prctl",      // x86_64架构必需
	//	"execve",          // 执行新程序必需
	//	"mmap",            // 需要允许所有内存分配方式
	//	"set_tid_address", // 线程初始化必需
	//	"sched_yield",     // 进程调度必需
	//	"madvise",         // 内存管理
	//	"getrandom",       // 随机数生成
	//	"pread64",         // 文件读取
	//	"set_robust_list", // 线程锁管理
	//	"fchmodat",        // chmod操作实际使用的系统调用
	//	"connect",         // 动态链接库加载需要
	//	"socket",          // 动态链接库加载需要
	//	"access",          // 文件访问检查
	//	"execveat",        // 新版内核使用的执行方式
	//	"getdents64",      // 目录操作必需
	//	"readlinkat",      // 动态链接库解析
	//	"getrusage",       // 资源使用统计
	//	"rt_sigreturn",    // 信号处理
	//	"open", "unshare", "epoll_ctl",
	//	"epoll_create", // 创建 epoll 实例
	//	"epoll_wait",   // 等待事件
	//	"epoll_pwait",  // 带信号屏蔽的 epoll_wait
	//	"nanosleep",    // 时间相关操作
	//	"pipe2",        // 创建管道
	//	"pidfd_open",   // 进程文件描述符操作
	//	"rt_sigprocmask",
	//	"getrlimit",       // 获取资源限制
	//	"setrlimit",       // 设置资源限制
	//	"gettid",          // 获取线程ID
	//	"getuid",          // 获取用户ID
	//	"getgid",          // 获取组ID
	//	"geteuid",         // 获取有效用户ID
	//	"getegid",         // 获取有效组ID
	//	"set_tid_address", // 设置线程ID地址
	//	"waitid",          // 等待子进程状态
	//	"tgkill",          // 向指定线程发送信号
	//	"fcntl",           // 文件描述符操作
	//	"execve", "access", "openat",
	//}
	//
	//// 添加允许规则（基础系统调用）
	//for _, syscallName := range allowedSyscalls {
	//	syscallNum, err := seccomp.GetSyscallFromName(syscallName)
	//	if err != nil {
	//		fmt.Printf("警告：系统调用 %s 不存在，跳过\n", syscallName)
	//		continue
	//	}
	//	filter.AddRule(syscallNum, seccomp.ActAllow)
	//}

	// 加载过滤规则
	if err := filter.Load(); err != nil {
		fmt.Println("加载Seccomp策略失败:", err)
		os.Exit(1)
	}

	cmd := exec.Command("/home/ganxue-server/utils/run_code/c/user_code")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_VM, // 允许共享内存空间
		Ptrace:     false,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("执行用户代码失败:", err)
		os.Exit(1)
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("用户代码执行失败:", err)
	}
}
