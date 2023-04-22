package sample

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"gitlab.com/techschool/pcbook/pb"
)

// 执行初始化 设置随机种子
func init() {
	// rand.Seed(1)
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-9400F",
			"COre i3-1005G1",
		)
	}
	// amd
	return randomStringFromSet(
		"Ryzen 7 PRO 2700U",
		"Ryzen 5 PRO 3500U",
		"Ryzen 3 PRO 3200GE",
	)
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDA" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"GTX 1070",
		)
	}

	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5700-XT",
		"RX Vega-56",
	)
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Lenovo", "Dell")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	case "Lenovo":
		return randomStringFromSet("Thinkpad X1", "Thinkpad P1", "Thinkpad P53")
	default:
		return randomStringFromSet("Latitude", "Vostro", "XPS", "Alienware")
	}
}

func randomResolution() *pb.Screen_Resolution {

	height := randomInt(1080, 4320)
	width := height * 16 / 9

	resolution := &pb.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}

	return resolution
}

func randomScreenPanel() pb.Screen_Panel {

	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomStringFromSet(arr ...string) string {
	n := len(arr)
	if n == 0 {
		return ""
	}
	return arr[rand.Intn(n)]
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomID() string {
	return uuid.New().String()
}
