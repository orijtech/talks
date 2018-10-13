func properFsync(fd int) {
        syscall.Syscall(syscal.SYS_FCNTL, uintptr(fd), syscall.F_FULLFSYNC, 0)
}
