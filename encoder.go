package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

// Nas5GSUpdateType represents the information element structure
type Nas5GSUpdateType struct {
	IEI          uint8 // Information Element Identifier
	Length       uint8 // Length of the Information Element
	EPS_PNBCIoT  uint8 // EPS-PNB-CIoT bit in the last octet
	G5S_PNBCIoT  uint8 // 5GS-PNB-CIoT bit in the last octet
	NGRAN_RCU    uint8 // NG-RAN-RCU bit in the last octet
	SMSRequested uint8 // SMS-requested bit in the last octet
}

// Encode encodes the Nas5GSUpdateType object into a byte stream
func (ie Nas5GSUpdateType) Encode(buffer *bytes.Buffer) {
	// Create a struct for the bitfield
	type structex struct {
		EPS_PNBCIoT  uint8 `bitfield:"1"`
		G5S_PNBCIoT  uint8 `bitfield:"1"`
		Reserved     uint8 `bitfield:"6"`
		NGRAN_RCU    uint8 `bitfield:"1"`
		SMSRequested uint8 `bitfield:"1"`
	}

	// Create a struct instance for the bitfield
	// initilising the bitfield with the values from Nas5GSUpdateType
	bitfield := structex{
		EPS_PNBCIoT:  ie.EPS_PNBCIoT,
		G5S_PNBCIoT:  ie.G5S_PNBCIoT,
		NGRAN_RCU:    ie.NGRAN_RCU,
		SMSRequested: ie.SMSRequested,
	}

	// Encode the bitfieldStruct into bytes
	byteValue := uint8(0)
	byteValue |= bitfield.EPS_PNBCIoT << 7
	byteValue |= bitfield.G5S_PNBCIoT << 6
	byteValue |= bitfield.Reserved << 0
	byteValue |= bitfield.NGRAN_RCU << 1
	byteValue |= bitfield.SMSRequested << 0

	// Write the encoded bytes to the buffer
	buffer.WriteByte(ie.IEI)
	buffer.WriteByte(ie.Length)
	buffer.WriteByte(byteValue)

	//return
}

func TestEncodeNas5GSUpdateType(t *testing.T) {
	input := Nas5GSUpdateType{
		IEI:          1,
		Length:       2,
		EPS_PNBCIoT:  0,
		G5S_PNBCIoT:  0,
		NGRAN_RCU:    1,
		SMSRequested: 1,
	}

	expectedOutput := []byte{0x01, 0x02, 0x03}
	//create a bufft
	var buffer bytes.Buffer
	//set input to buffer
	input.Encode(&buffer)
	//testing output
	output := buffer.Bytes()

	if !bytes.Equal(output, expectedOutput) {
		t.Errorf("Expected: %v, Got: %v", expectedOutput, output)
	}
}

func main() {
	// Run the unit test with coverage
	result := testing.MainStart(
		nil,
		[]testing.InternalTest{
			{"TestEncodeNas5GSUpdateType", TestEncodeNas5GSUpdateType},
		},
		nil,
		nil,
		nil,
	)

	// Output the unit test result
	fmt.Println(result)

	// Run tests with coverage and output coverage information

	// Set coverage mode to "count" for statement coverage
	coverMode := "count"
	// Specify the coverage profile file name
	coverProfile := "coverage.out"
	// Execute 'go test' with coverage and write the coverage profile to the specified file
	cmd := exec.Command("go", "test", "-covermode="+coverMode, "-coverprofile="+coverProfile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		// Handle errors if 'go test' command fails
		fmt.Println("Error running coverage test:", err)
	}

	// Output coverage report

	// Read the coverage profile file
	coverReport, err := ioutil.ReadFile(coverProfile)
	if err != nil {
		// Handle errors if reading the coverage profile fails
		fmt.Println("Error reading coverage profile:", err)
	} else {
		// Print the coverage report
		fmt.Println(string(coverReport))
	}
}
