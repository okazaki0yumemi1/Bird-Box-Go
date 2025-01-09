package recording

import (
	"fmt"
	"os/exec"
)

type InputDevice struct {
	Name    string
	Address string
}

func listAudioInputDevices() ([]string, error) {
	cmd := exec.Command("arecord", "-l")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to run arecord: %v, output: %s", err, output)
	}

	// Parse the output to extract device names
	lines := string(output)
	var devices []string
	for _, line := range lines {
		if len(string(line)) > 0 && string(line)[0] == '[' {
			devices = append(devices, string(line))
		}
	}

	return devices, nil
}

// func main() {
// 	devices, err := listAudioInputDevices()
// 	if err != nil {
// 		fmt.Printf("Error: %v\n", err)
// 		return
// 	}

// 	fmt.Println("Available audio input devices:")
// 	for _, device := range devices {
// 		fmt.Println(device)
// 	}
// }
