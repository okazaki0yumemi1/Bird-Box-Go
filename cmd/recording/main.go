package recording

import (
	"fmt"
)

func New() *[]string {
	devices, err := listAudioInputDevices()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil
	}

	fmt.Println("Available audio input devices:")
	for _, device := range devices {
		fmt.Println(device)
	}

	return &devices
}
