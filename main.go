package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("what?")
	}
}

func child() {
	fmt.Printf("runing %v as PID %d \n", os.Args[2], os.Getpid())
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	must(syscall.Chroot("/root/gorun/rootfs"))
	must(os.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	must(syscall.Setenv("PATH", "/bin"))
	must(cmd.Run())
}

func run() {
	fmt.Printf("Running %v \n", os.Args[2])

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	must(cmd.Run())
}

func must(inputs ...interface{}) {
	l := len(inputs)
	if l > 0 {
		if _, ok := inputs[l-1].(error); ok {
			if inputs[l-1] != nil {
				panic(inputs[l-1])
			}
		}
	}
}
