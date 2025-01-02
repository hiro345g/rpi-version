package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	var revcode string

	if len(args) == 0 {
		cmd := exec.Command("sh", "-c", "cat /proc/cpuinfo | awk '/Revision/ {print $3}'")
		out, err := cmd.Output()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		revcode = strings.TrimSpace(string(out))
	} else {
		revcode = args[0]
	}

	code, err := strconv.ParseInt(revcode, 16, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	new_value := (code >> 23) & 0x1
	mem_value := (code >> 20) & 0x7
	man_value := (code >> 16) & 0xf
	proc_value := (code >> 12) & 0xf
	model_value := (code >> 4) & 0xff
	rev_value := code & 0xf

	if new_value == 1 {
		memStrings := []string{"256MB", "512MB", "1GB", "2GB", "4GB", "8GB"}
		mem_svalue := memStrings[mem_value]

		manStrings := []string{"Sony UK", "Egoman", "Embest", "Sony Japan", "Embest", "Stadium"}
		man_svalue := manStrings[man_value]

		procStrings := []string{"BCM2835", "BCM2836", "BCM2837", "BCM2711"}
		proc_svalue := procStrings[proc_value]

		modelStrings := []string{"A", "B", "A+", "B+", "2B", "alpha", "CM1", "7", "3B", "Zero", "CM3", "b", "Zero W", "3B+", "3A+", "Internal use only", "CM3+", "4B", "Zero 2 W", "400", "CM4", "CM4S"}
		model_svalue := modelStrings[model_value]

		fmt.Printf("Revision: %s\n\tModel: %s\n\tRevision: 1.%d\n\tMemory Size: %s\n\tProcessor: %s\n\tManufacturer: %s\n", revcode, model_svalue, rev_value, mem_svalue, proc_svalue, man_svalue)
	} else if (code >= 0x2 && code <= 0x9) || (code >= 0xd && code <= 0x15) {
		mem_svalue := "MB"
		if (code >= 0x2 && code <= 0x9) || code == 0x12 {
			mem_svalue = "256MB"
		} else if code == 0x15 {
			mem_svalue = "256/512MB"
		} else {
			mem_svalue = "512MB"
		}

		model_svalue := ""
		if (code >= 0x2 && code <= 0x6) || (code >= 0xd && code <= 0xf) {
			model_svalue = "B"
		} else if code >= 0x7 && code <= 0x9 {
			model_svalue = "A"
		} else if code == 0x10 || code == 0x13 {
			model_svalue = "B+"
		} else if code == 0x12 || code == 0x15 {
			model_svalue = "A+"
		} else {
			model_svalue = "CM1"
		}

		rev_value := ""
		if (code >= 0x2 && code <= 0x3) || code == 0x11 || code == 0x14 {
			rev_value = "1.0"
		} else if code == 0x12 || code == 0x14 {
			rev_value = "1.1"
		} else if code == 0x10 || code == 0x13 {
			rev_value = "1.2"
		} else {
			rev_value = "2.0"
		}

		man_svalue := ""
		if (code >= 0x2 && code <= 0x3) || (code >= 0x6 && code <= 0x7) || code == 0xd || code == 0xf {
			man_svalue = "Egoman"
		} else if (code >= 0x10 && code <= 0x12) || code == 0x4 || code == 0x8 || code == 0xe {
			man_svalue = "Sony UK"
		} else if code == 0x5 || code == 0x9 {
			man_svalue = "Qisda"
		} else {
			man_svalue = "Embest"
		}
		fmt.Printf("Revision: %s\n\tModel: %s\n\tRevision: %s\n\tMemory Size: %s\n\tManufacturer: %s\n", revcode, model_svalue, rev_value, mem_svalue, man_svalue)
	} else {
		fmt.Printf("Revision: %s\n\tUnknown revision\n", revcode)
	}
}
