package main

// https://raw.githubusercontent.com/pierrerice/containers-from-scratch/master/main.go

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

// go run main.go run <cmd> <args>
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("help")
	}
}

func run() {
	fmt.Printf("Running %v \n", os.Args[2:])

	cg()

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Cloneflags is only available in Linux
	// CLONE_NEWUTS namespace isolates hostname
	// CLONE_NEWPID namespace isolates processes
	// CLONE_NEWNS namespace isolates mounts
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	// Run child using namespaces. The command provided will be executed inside that.
	must(cmd.Run())
}

func child() {
	fmt.Printf("Running %v \n", os.Args[2:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte("container")))

	// change root
	must(syscall.Chroot("/home/pierrez/workspace/demo/containers-vs-vms-demo/mychroot"))

	// Change directory after chroot
	must(os.Chdir("/"))

	// Mount /proc inside container so that `ps` command works
	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	// Mount a temporary filesystem
	must(syscall.Mount("tmp", "tmp", "tmpfs", 0, ""))

	must(cmd.Run())

	// Cleanup mount
	must(syscall.Unmount("proc", 0))
	must(syscall.Unmount("tmp", 0))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func cg() {
	cgroups := "/sys/fs/cgroup/"
	pids := filepath.Join(cgroups, "pids")
	os.Mkdir(filepath.Join(pids, "pierre"), 0755)
	must(ioutil.WriteFile(filepath.Join(pids, "pierre/pids.max"), []byte("20"), 0700))
	// Removes the new cgroup in place after the container exits
	must(ioutil.WriteFile(filepath.Join(pids, "pierre/notify_on_release"), []byte("1"), 0700))
	must(ioutil.WriteFile(filepath.Join(pids, "pierre/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}
