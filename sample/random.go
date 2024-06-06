package sample

import (
	"github.com/google/uuid"
	"grpc4_test/pb"
	"math/rand"
	"time"
)

func init() {
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

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"i5-10400f",
			"i7-10700f",
			"i3-10300f",
			"Xeon E-2286M",
		)
	}

	return randomStringFromSet(
		"Ryzen 3 3200G",
		"Ryzen 5 3600X",
		"Ryzen 7 7700G",
	)
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return max + rand.Float64()*(max-min)
}

// GPU RANDOM

func randGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"RTX 4090",
			"RTX 3060",
			"GTX 1050",
			"GT 640",
		)
	}

	return randomStringFromSet(
		"RX 6600 XT",
		"RX 580",
		"RX 560",
	)
}

// SCREEN random
func randomFloat32(min float32, max float32) float32 {
	return max + rand.Float32()*(max-min)
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4230)
	width := height * 16 / 9
	resolution := &pb.Screen_Resolution{
		Width:  uint32(height),
		Height: uint32(width),
	}

	return resolution
}

// LAPTOP FIELD

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFromSet("Asus", "Acer", "Apple")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Asus":
		return randomStringFromSet("EnjoyBook", "TUF Gaming", "ROG Gaming")
	case "Acer":
		return randomStringFromSet("Aspire 7", "Predator", "Acer Nitro")
	default:
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	}
}
