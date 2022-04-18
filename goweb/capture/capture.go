package capture

import (
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"

	"github.com/kbinani/screenshot"
)

type Capture struct {
}
type HWND uintptr

type RECT struct {
	Left, Top, Right, Bottom int32
}

var (
	user32, _                   = syscall.LoadLibrary("user32.dll")
	findWindowW, _              = syscall.GetProcAddress(user32, "FindWindowW")
	getWindowRect, _            = syscall.GetProcAddress(user32, "GetWindowRect")
	setWindowPos, _             = syscall.GetProcAddress(user32, "SetWindowPos")
	setForegroundWindow, _      = syscall.GetProcAddress(user32, "SetForegroundWindow")
	getForegroundWindow, _      = syscall.GetProcAddress(user32, "GetForegroundWindow")
	getWindowThreadProcessId, _ = syscall.GetProcAddress(user32, "GetWindowThreadProcessId")
	attachThreadInput, _        = syscall.GetProcAddress(user32, "AttachThreadInput")
	showWindow, _               = syscall.GetProcAddress(user32, "ShowWindow")

	kernel32, _           = syscall.LoadLibrary("Kernel32.dll")
	getCurrentThreadId, _ = syscall.GetProcAddress(kernel32, "GetCurrentThreadId")
)

//HWND WINAPI FindWindow(
//  _In_opt_ LPCTSTR lpClassName,
//  _In_opt_ LPCTSTR lpWindowName
//);
func FindWindowByTitle(title string) HWND {
	ret, _, _ := syscall.SyscallN(
		findWindowW,
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		0,
	)
	return HWND(ret)
}

func GetWindowDimensions(hwnd HWND) *RECT {
	var rect RECT

	syscall.SyscallN(
		getWindowRect,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&rect)),
		0,
	)

	return &rect
}

func SetWindowPos(hwnd HWND, cx int32, cy int32, hWndInsertAfter HWND) HWND {

	ret, _, err := syscall.SyscallN(
		setWindowPos,
		uintptr(hwnd),
		uintptr(hWndInsertAfter),
		uintptr(0),
		uintptr(0),
		uintptr(cx),
		uintptr(cy),
		uintptr(0x0040),
	)
	fmt.Printf("SetWindowPos: %v\n", err)
	return HWND(ret)
}

func SetForegroundWindow(hwnd HWND) bool {
	ret, _, _ := syscall.SyscallN(
		setForegroundWindow,
		uintptr(hwnd),
		0,
		0,
	)
	val := (uint)(ret)
	return val != 0
}

func GetForegroundWindow() HWND {
	ret, _, err := syscall.SyscallN(
		getForegroundWindow,
		0,
		0,
		0,
	)
	fmt.Printf("GetForegroundWindow: %v\n", err)
	fmt.Printf("GetForegroundWindow: %v\n", ret)
	return HWND(ret)
}

func GetCurrentThreadId() int64 {
	ret, _, err := syscall.SyscallN(
		getCurrentThreadId,
		0,
	)
	fmt.Printf("GetCurrentThreadId: %v\n", err)
	fmt.Printf("GetCurrentThreadId: %v\n", ret)
	return int64(ret)
}

func GetWindowThreadProcessId(hwnd HWND) int64 {
	var processId int64
	ret, _, err := syscall.SyscallN(
		getWindowThreadProcessId,
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&processId)),
	)
	fmt.Printf("GetWindowThreadProcessId: %v\n", ret)
	fmt.Printf("GetWindowThreadProcessId: %v\n", err)

	return int64(ret)
}

func AttachThreadInput(curThreadId int64, curProcessId int64, attach int) bool {
	//var bVal bool = true
	ret, _, err := syscall.SyscallN(
		attachThreadInput,
		uintptr(curThreadId),
		uintptr(curProcessId),
		uintptr(attach),
	)

	fmt.Printf("AttachThreadInput: %v\n", err)
	val := (uint)(ret)
	return val != 0
}

func ShowWindow(hwnd HWND) bool {
	ret, _, err := syscall.SyscallN(
		showWindow,
		uintptr(hwnd),
		uintptr(5),
	)
	fmt.Printf("ShowWindow: %v\n", err)
	val := (uint)(ret)
	return val != 0
}

func (c Capture) CaptureImage(imgFileName string) {
	defer syscall.FreeLibrary(user32)

	hwnd := FindWindowByTitle("DraegerEIT")
	fmt.Printf("Return: %d\n", hwnd)

	curHwnd := GetForegroundWindow()
	curID := GetCurrentThreadId()
	curForeID := GetWindowThreadProcessId(curHwnd)
	AttachThreadInput(curID, curForeID, 1)
	ShowWindow(hwnd)

	var rect *RECT
	if hwnd > 0 {
		rect = GetWindowDimensions(hwnd)
		dd := -1
		SetWindowPos(hwnd, rect.Right-rect.Left, rect.Bottom-rect.Top, HWND(dd))
		rtnv := SetForegroundWindow(hwnd)
		fmt.Printf("SetForegroundWindow: %v\n", rtnv)
		//SetWindowPos(hwnd, rect.Right-rect.Left, rect.Bottom-rect.Top)

	}

	n := screenshot.NumActiveDisplays()

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)

		pt1 := image.Point{
			int(rect.Left),
			int(rect.Top),
		}
		pt2 := image.Point{
			int(rect.Right),
			int(rect.Bottom),
		}
		rg := image.Rectangle{
			pt1,
			pt2,
		}
		img, err := screenshot.CaptureRect(rg)
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf("%s.png", imgFileName)

		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(pwd)
		index := strings.Index(pwd, "goweb")
		pwd = pwd[:index]
		pwd = strings.TrimSuffix(pwd, "\\")

		newpath := filepath.Join(pwd, "src", "assets", "imgs", fileName)

		fmt.Println(newpath)
		_ = os.Remove(newpath)
		file, _ := os.Create(newpath)
		defer file.Close()
		png.Encode(file, img)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, newpath)
	}

	// reset previous active window
	if curHwnd > 0 {

		dd := -2
		SetWindowPos(hwnd, rect.Right-rect.Left, rect.Bottom-rect.Top, HWND(dd))
		AttachThreadInput(curID, curForeID, 0)
		SetForegroundWindow(curHwnd)
	}
}

func (c Capture) DeleteImage(imgFileName string) {

	fileName := fmt.Sprintf("%s.png", imgFileName)
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(pwd)
	index := strings.Index(pwd, "goweb")
	pwd = pwd[:index]
	pwd = strings.TrimSuffix(pwd, "\\")

	newpath := filepath.Join(pwd, "src", "assets", "imgs", fileName)
	fmt.Println(newpath)
	_ = os.Remove(newpath)
}
