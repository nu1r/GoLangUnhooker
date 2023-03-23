// Code generated by 'go generate'; DO NOT EDIT.

package ThreadlessInject

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modCrypt32  = windows.NewLazySystemDLL("Crypt32.dll")
	modKernel32 = windows.NewLazySystemDLL("Kernel32.dll")
	modntdll    = windows.NewLazySystemDLL("ntdll.dll")

	procCertEnumSystemStore     = modCrypt32.NewProc("CertEnumSystemStore")
	procVirtualAlloc            = modKernel32.NewProc("VirtualAlloc")
	procNtAllocateVirtualMemory = modntdll.NewProc("NtAllocateVirtualMemory")
	procNtFreeVirtualMemory     = modntdll.NewProc("NtFreeVirtualMemory")
	procNtOpenProcess           = modntdll.NewProc("NtOpenProcess")
	procNtProtectVirtualMemory  = modntdll.NewProc("NtProtectVirtualMemory")
	procNtReadVirtualMemory     = modntdll.NewProc("NtReadVirtualMemory")
	procNtWriteVirtualMemory    = modntdll.NewProc("NtWriteVirtualMemory")
)

func CertEnumSystemStoreNu1r(p1 uintptr, p2 uintptr, p3 uintptr, p4 uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procCertEnumSystemStore.Addr(), 4, uintptr(p1), uintptr(p2), uintptr(p3), uintptr(p4), 0, 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func VirtualAllocNu1r(p1 uintptr, p2 uintptr, p3 uintptr, p4 uintptr) (p5 uintptr, err error) {
	r0, _, e1 := syscall.Syscall6(procVirtualAlloc.Addr(), 4, uintptr(p1), uintptr(p2), uintptr(p3), uintptr(p4), 0, 0)
	p5 = uintptr(r0)
	if p5 == 0 {
		err = errnoErr(e1)
	}
	return
}

func NtAllocateVirtualMemoryNu1r(p1 uintptr, p2 uintptr, p3 uintptr, p4 uintptr, p5 uintptr, p6 uintptr) (err2 error) {
	r0, _, _ := syscall.Syscall6(procNtAllocateVirtualMemory.Addr(), 6, uintptr(p1), uintptr(p2), uintptr(p3), uintptr(p4), uintptr(p5), uintptr(p6))
	if r0 != 0 {
		err2 = syscall.Errno(r0)
	}
	return
}

func NtFreeVirtualMemoryNu1r(p1 uintptr, p2 uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(procNtFreeVirtualMemory.Addr(), 2, uintptr(p1), uintptr(p2), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func OpenProcessNu1r(p1 uintptr, p2 uintptr, p3 uintptr, p4 uintptr) {
	syscall.Syscall6(procNtOpenProcess.Addr(), 4, uintptr(p1), uintptr(p2), uintptr(p3), uintptr(p4), 0, 0)
	return
}

func NtProtectVirtualMemoryNu1r(p1 uintptr, p2 uintptr, p3 uintptr, p4 uintptr, p5 uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procNtProtectVirtualMemory.Addr(), 5, uintptr(p1), uintptr(p2), uintptr(p3), uintptr(p4), uintptr(p5), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func NtReadVirtualMemoryNu1r(p1 uintptr, p2 uintptr, p3 uintptr, p4 uintptr, p5 uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procNtReadVirtualMemory.Addr(), 5, uintptr(p1), uintptr(p2), uintptr(p3), uintptr(p4), uintptr(p5), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}

func NtWriteVirtualMemoryNu1r(p1 uintptr, p2 uintptr, p3 uintptr, p4 uintptr, p5 uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procNtWriteVirtualMemory.Addr(), 5, uintptr(p1), uintptr(p2), uintptr(p3), uintptr(p4), uintptr(p5), 0)
	if r1 == 0 {
		err = errnoErr(e1)
	}
	return
}
