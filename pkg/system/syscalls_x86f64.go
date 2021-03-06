package system

const (
	SyscallX86MaxNum64   = 322
	SyscallX86LastName64 = "execveat"
)

// line numbers are aligned with the syscall number (-10)
var syscallNumTableX86Family64 = [...]string{
	"read",
	"write",
	"open",
	"close",
	"stat",
	"fstat",
	"lstat",
	"poll",
	"lseek",
	"mmap",
	"mprotect",
	"munmap",
	"brk",
	"rt_sigaction",
	"rt_sigprocmask",
	"rt_sigreturn",
	"ioctl",
	"pread64",
	"pwrite64",
	"readv",
	"writev",
	"access",
	"pipe",
	"select",
	"sched_yield",
	"mremap",
	"msync",
	"mincore",
	"madvise",
	"shmget",
	"shmat",
	"shmctl",
	"dup",
	"dup2",
	"pause",
	"nanosleep",
	"getitimer",
	"alarm",
	"setitimer",
	"getpid",
	"sendfile",
	"socket",
	"connect",
	"accept",
	"sendto",
	"recvfrom",
	"sendmsg",
	"recvmsg",
	"shutdown",
	"bind",
	"listen",
	"getsockname",
	"getpeername",
	"socketpair",
	"setsockopt",
	"getsockopt",
	"clone",
	"fork",
	"vfork",
	"execve",
	"exit",
	"wait4",
	"kill",
	"uname",
	"semget",
	"semop",
	"semctl",
	"shmdt",
	"msgget",
	"msgsnd",
	"msgrcv",
	"msgctl",
	"fcntl",
	"flock",
	"fsync",
	"fdatasync",
	"truncate",
	"ftruncate",
	"getdents",
	"getcwd",
	"chdir",
	"fchdir",
	"rename",
	"mkdir",
	"rmdir",
	"creat",
	"link",
	"unlink",
	"symlink",
	"readlink",
	"chmod",
	"fchmod",
	"chown",
	"fchown",
	"lchown",
	"umask",
	"gettimeofday",
	"getrlimit",
	"getrusage",
	"sysinfo",
	"times",
	"ptrace",
	"getuid",
	"syslog",
	"getgid",
	"setuid",
	"setgid",
	"geteuid",
	"getegid",
	"setpgid",
	"getppid",
	"getpgrp",
	"setsid",
	"setreuid",
	"setregid",
	"getgroups",
	"setgroups",
	"setresuid",
	"getresuid",
	"setresgid",
	"getresgid",
	"getpgid",
	"setfsuid",
	"setfsgid",
	"getsid",
	"capget",
	"capset",
	"rt_sigpending",
	"rt_sigtimedwait",
	"rt_sigqueueinfo",
	"rt_sigsuspend",
	"sigaltstack",
	"utime",
	"mknod",
	"uselib",
	"personality",
	"ustat",
	"statfs",
	"fstatfs",
	"sysfs",
	"getpriority",
	"setpriority",
	"sched_setparam",
	"sched_getparam",
	"sched_setscheduler",
	"sched_getscheduler",
	"sched_get_priority_max",
	"sched_get_priority_min",
	"sched_rr_get_interval",
	"mlock",
	"munlock",
	"mlockall",
	"munlockall",
	"vhangup",
	"modify_ldt",
	"pivot_root",
	"_sysctl",
	"prctl",
	"arch_prctl",
	"adjtimex",
	"setrlimit",
	"chroot",
	"sync",
	"acct",
	"settimeofday",
	"mount",
	"umount2",
	"swapon",
	"swapoff",
	"reboot",
	"sethostname",
	"setdomainname",
	"iopl",
	"ioperm",
	"create_module",
	"init_module",
	"delete_module",
	"get_kernel_syms",
	"query_module",
	"quotactl",
	"nfsservctl",
	"getpmsg",
	"putpmsg",
	"afs_syscall",
	"tuxcall",
	"security",
	"gettid",
	"readahead",
	"setxattr",
	"lsetxattr",
	"fsetxattr",
	"getxattr",
	"lgetxattr",
	"fgetxattr",
	"listxattr",
	"llistxattr",
	"flistxattr",
	"removexattr",
	"lremovexattr",
	"fremovexattr",
	"tkill",
	"time",
	"futex",
	"sched_setaffinity",
	"sched_getaffinity",
	"set_thread_area",
	"io_setup",
	"io_destroy",
	"io_getevents",
	"io_submit",
	"io_cancel",
	"get_thread_area",
	"lookup_dcookie",
	"epoll_create",
	"epoll_ctl_old",
	"epoll_wait_old",
	"remap_file_pages",
	"getdents64",
	"set_tid_address",
	"restart_syscall",
	"semtimedop",
	"fadvise64",
	"timer_create",
	"timer_settime",
	"timer_gettime",
	"timer_getoverrun",
	"timer_delete",
	"clock_settime",
	"clock_gettime",
	"clock_getres",
	"clock_nanosleep",
	"exit_group",
	"epoll_wait",
	"epoll_ctl",
	"tgkill",
	"utimes",
	"vserver",
	"mbind",
	"set_mempolicy",
	"get_mempolicy",
	"mq_open",
	"mq_unlink",
	"mq_timedsend",
	"mq_timedreceive",
	"mq_notify",
	"mq_getsetattr",
	"kexec_load",
	"waitid",
	"add_key",
	"request_key",
	"keyctl",
	"ioprio_set",
	"ioprio_get",
	"inotify_init",
	"inotify_add_watch",
	"inotify_rm_watch",
	"migrate_pages",
	"openat",
	"mkdirat",
	"mknodat",
	"fchownat",
	"futimesat",
	"newfstatat",
	"unlinkat",
	"renameat",
	"linkat",
	"symlinkat",
	"readlinkat",
	"fchmodat",
	"faccessat",
	"pselect6",
	"ppoll",
	"unshare",
	"set_robust_list",
	"get_robust_list",
	"splice",
	"tee",
	"sync_file_range",
	"vmsplice",
	"move_pages",
	"utimensat",
	"epoll_pwait",
	"signalfd",
	"timerfd_create",
	"eventfd",
	"fallocate",
	"timerfd_settime",
	"timerfd_gettime",
	"accept4",
	"signalfd4",
	"eventfd2",
	"epoll_create1",
	"dup3",
	"pipe2",
	"inotify_init1",
	"preadv",
	"pwritev",
	"rt_tgsigqueueinfo",
	"perf_event_open",
	"recvmmsg",
	"fanotify_init",
	"fanotify_mark",
	"prlimit64",
	"name_to_handle_at",
	"open_by_handle_at",
	"clock_adjtime",
	"syncfs",
	"sendmmsg",
	"setns",
	"getcpu",
	"process_vm_readv",
	"process_vm_writev",
	"kcmp",
	"finit_module",
	"sched_setattr",
	"sched_getattr",
	"renameat2",
	"seccomp",
	"getrandom",
	"memfd_create",
	"kexec_file_load",
	"bpf",
	"execveat",
	"userfaultfd",
	"membarrier",
	"mlock2",
	"copy_file_range",
	"preadv2",
	"pwritev2",
	"pkey_mprotect",
	"pkey_alloc",
	"pkey_free",
	"statx",
	"io_pgetevents",
	"rseq",
	"reserved.335",
	"reserved.336",
	"reserved.337",
	"reserved.338",
	"reserved.339",
	"reserved.340",
	"reserved.341",
	"reserved.342",
	"reserved.343",
	"reserved.344",
	"reserved.345",
	"reserved.346",
	"reserved.347",
	"reserved.348",
	"reserved.349",
	"reserved.350",
	"reserved.351",
	"reserved.352",
	"reserved.353",
	"reserved.354",
	"reserved.355",
	"reserved.356",
	"reserved.357",
	"reserved.358",
	"reserved.359",
	"reserved.360",
	"reserved.361",
	"reserved.362",
	"reserved.363",
	"reserved.364",
	"reserved.365",
	"reserved.366",
	"reserved.367",
	"reserved.368",
	"reserved.369",
	"reserved.370",
	"reserved.371",
	"reserved.372",
	"reserved.373",
	"reserved.374",
	"reserved.375",
	"reserved.376",
	"reserved.377",
	"reserved.378",
	"reserved.379",
	"reserved.380",
	"reserved.381",
	"reserved.382",
	"reserved.383",
	"reserved.384",
	"reserved.385",
	"reserved.386",
	"reserved.387",
	"reserved.388",
	"reserved.389",
	"reserved.390",
	"reserved.391",
	"reserved.392",
	"reserved.393",
	"reserved.394",
	"reserved.395",
	"reserved.396",
	"reserved.397",
	"reserved.398",
	"reserved.399",
	"reserved.400",
	"reserved.401",
	"reserved.402",
	"reserved.403",
	"reserved.404",
	"reserved.405",
	"reserved.406",
	"reserved.407",
	"reserved.408",
	"reserved.409",
	"reserved.410",
	"reserved.411",
	"reserved.412",
	"reserved.413",
	"reserved.414",
	"reserved.415",
	"reserved.416",
	"reserved.417",
	"reserved.418",
	"reserved.419",
	"reserved.420",
	"reserved.421",
	"reserved.422",
	"reserved.423",
	"pidfd_send_signal",
	"io_uring_setup",
	"io_uring_enter",
	"io_uring_register",
	"open_tree",
	"move_mount",
	"fsopen",
	"fsconfig",
	"fsmount",
	"fspick",
	"pidfd_open",
	"clone3", //435
}

func callNameX86Family64(num uint32) string {
	if num > SyscallX86MaxNum64 {
		return SyscallX86UnknownName
	}

	return syscallNumTableX86Family64[num]
}

func callNumTableIsOkX86Family64() bool {
	if (len(syscallNumTableX86Family64) == SyscallX86MaxNum64+1) &&
		syscallNumTableX86Family64[SyscallX86MaxNum64] == SyscallX86LastName64 {
		return true
	}

	return false
}

func callNumberX86Family64(name string) (uint32, bool) {
	num, ok := syscallNameTableX86Family64[name]
	return num, ok
}

var syscallNameTableX86Family64 map[string]uint32

func init() {
	syscallNameTableX86Family64 = make(map[string]uint32, len(syscallNumTableX86Family64))

	for callNum, callName := range syscallNumTableX86Family64 {
		syscallNameTableX86Family64[callName] = uint32(callNum)
	}
}
