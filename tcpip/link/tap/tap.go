package tap

import (
	"fmt"
	"os/exec"
	"syscall"
	"unsafe"
)

type Tap struct {
	Name string
	Fd   int
}

func NewTap(name string) (*Tap, error) {
	var (
		err error
	)

	tap := &Tap{
		Name: name,
	}

	//打开tuntap字符文件
	tap.Fd, err = syscall.Open("/dev/net/tun", syscall.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	//ioctl调用，将fd和虚拟网卡绑定在一起
	ifr := struct {
		name  [16]byte
		flags uint16
		_     [22]byte
	}{
		flags: syscall.IFF_TAP | syscall.IFF_NO_PI,
	}
	copy(ifr.name[:], name)

	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, uintptr(tap.Fd), syscall.TUNSETIFF, uintptr(unsafe.Pointer(&ifr)))
	if errno != 0 {
		syscall.Close(tap.Fd)
		return nil, fmt.Errorf("ioctrl err, errno:%d", errno)
	}

	return tap, nil
}

func (t *Tap) Linkup() error {
	// ip link set xx up
	if out, err := exec.Command("ip", "link", "set", t.Name, "up").CombinedOutput(); err != nil {
		return fmt.Errorf("%v: %v", err, string(out))
	}

	return nil
}

func (t *Tap) SetIp(ip string) error {
	//ip addr add  192.168.x.x dev xx
	if out, err := exec.Command("ip", "addr", "add", ip, "dev", t.Name).CombinedOutput(); err != nil {
		return fmt.Errorf("%v: %v", err, string(out))
	}

	return nil
}
