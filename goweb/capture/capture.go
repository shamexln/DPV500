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
	user32, _              = syscall.LoadLibrary("user32.dll")
	findWindowW, _         = syscall.GetProcAddress(user32, "FindWindowW")
	getWindowRect, _       = syscall.GetProcAddress(user32, "GetWindowRect")
	setWindowPos, _        = syscall.GetProcAddress(user32, "SetWindowPos")
	setForegroundWindow, _ = syscall.GetProcAddress(user32, "SetForegroundWindow")
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

func SetWindowPos(hwnd HWND, cx int32, cy int32) HWND {
	//dd := 1

	ret, _, err := syscall.SyscallN(
		setWindowPos,
		uintptr(hwnd),
		//uintptr(unsafe.Pointer(&dd)),
		uintptr(0),
		uintptr(0),
		uintptr(0),
		uintptr(cx),
		uintptr(cy),
		uintptr(0x0040),
	)
	fmt.Printf("Return: %v\n", err)
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

func (c Capture) CaptureImage(imgFileName string) {
	defer syscall.FreeLibrary(user32)

	hwnd := FindWindowByTitle("DraegerEIT")
	//hwnd := FindWindowByTitle("计算器")
	fmt.Printf("Return: %d\n", hwnd)
	var rect *RECT
	if hwnd > 0 {
		rect = GetWindowDimensions(hwnd)

		rtnv := SetForegroundWindow(hwnd)
		fmt.Printf("Return: %v\n", rtnv)
		SetWindowPos(hwnd, rect.Right-rect.Left, rect.Bottom-rect.Top)

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
