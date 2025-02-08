package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("Welcome to Go Programming language!")
	// Motivational message
	fmt.Println("Always keep high hopes for yourself when you learn something new -:)")
	fmt.Println("Finding system information:- ")
	cmd := exec.Command("whoami")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing whoami:", err)
		return
	}
	// Print the output of the 'whoami' command
	fmt.Println("Who logged into the system:", string(output))

	// Print the output of the 'hostname' command
	cmdHost := exec.Command("hostname")
	hostOutput, err := cmdHost.Output()
	if err != nil {
		fmt.Println("Error executing hostname:", err)
		return
	}
	// Print the output of the 'df -h' command
	fmt.Println("System hostname:", string(hostOutput))

	// Command to get the disk free output
	cmddiskfree := exec.Command("df", "-h")
	diskfreeOutput, err := cmddiskfree.Output()
	if err != nil {
		fmt.Println("Error executing diskfree command:", err)
		return
	}
	// Print the hostname
	fmt.Println("System disk free output:", string(diskfreeOutput))

	// Command to get the system detailed OS information
	cmdosinfo := exec.Command("system_profiler", "SPSoftwareDataType")
	osinfoOutput, err := cmdosinfo.Output()
	if err != nil {
		fmt.Println("Error executing 'system_profiler SPSoftwareDataType' command:", err)
		return
	}
	// Print the 'system_profiler SPSoftwareDataType' command output
	fmt.Println("System OS information output:", string(osinfoOutput))

	// Command to get the System Hardware Specification
	cmdhwinfo := exec.Command("system_profiler", "SPHardwareDataType")
	hwinfoOutput, err := cmdhwinfo.Output()
	if err != nil {
		fmt.Println("Error executing HWinfo command:'system_profiler SPHardwareDataType' in MacBook Pro:", err)
		return
	}
	// Print the 'system_profiler SPHardwareDataType' command output
	fmt.Println("System Hardware Spec output:", string(hwinfoOutput))
}
