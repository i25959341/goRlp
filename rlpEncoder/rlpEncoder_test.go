package rlpEncoder

import (
	"bytes"
	"fmt"
	"math/big"
	"testing"

	rlpList "github.com/gorlp/rlpList"
	rlpString "github.com/gorlp/rlpString"
)

func TestNewRlpEncoderString(t *testing.T) {
	//  83 as offset is 80 and length of string is 3 thus, 83
	// the d is 64, o is 6f, g is 67
	rightAnswer := "83646f67"
	rString := rlpString.CreateRlpString("dog")

	result := EncodeAll(rString)
	var buffer bytes.Buffer

	for _, v := range result {
		buffer.WriteString(fmt.Sprintf("%02x", v&0xff))
	}

	if buffer.String() != rightAnswer {
		t.Errorf("EncodeAll dog incorrect, got: %s, want: %s.", buffer.String(), rightAnswer)
	}

	rightAnswer = "01"
	rString = rlpString.NewRlpString([]byte{byte(0x01)})

	result = EncodeAll(rString)
	buffer.Reset()

	for _, v := range result {
		buffer.WriteString(fmt.Sprintf("%02x", v&0xff))
	}

	if buffer.String() != rightAnswer {
		t.Errorf("EncodeAll one byte incorrect, got: %s, want: %s.", buffer.String(), rightAnswer)
	}

	rightAnswer = "80"
	rString = rlpString.CreateRlpString("")

	result = EncodeAll(rString)
	buffer.Reset()

	for _, v := range result {
		buffer.WriteString(fmt.Sprintf("%02x", v&0xff))
	}
	// fmt.Printf(buffer.String())

	if buffer.String() != rightAnswer {
		t.Errorf("EncodeAll empytp string incorrect, got: %s, want: %s.",
			buffer.String(),
			rightAnswer)
	}

	rightAnswer = "80"
	rString = rlpString.CreateRlpStringBigInt(big.NewInt(0))

	result = EncodeAll(rString)
	buffer.Reset()

	for _, v := range result {
		buffer.WriteString(fmt.Sprintf("%02x", v&0xff))
	}

	if buffer.String() != rightAnswer {
		t.Errorf("EncodeAll 0 string incorrect, got: %s, want: %s.",
			buffer.String(),
			rightAnswer)
	}

	rightAnswer = "820400"
	rString = rlpString.CreateRlpStringBigInt(big.NewInt(1024))

	result = EncodeAll(rString)
	buffer.Reset()

	for _, v := range result {
		buffer.WriteString(fmt.Sprintf("%02x", v&0xff))
	}

	if buffer.String() != rightAnswer {
		t.Errorf("EncodeAll 0 string incorrect, got: %s, want: %s.",
			buffer.String(),
			rightAnswer)
	}
}

func TestNewRlpEncoderList(t *testing.T) {

	rightAnswer := "c0"
	rList := rlpList.NewRlpListVariadic()

	result := EncodeAll(rList)
	var buffer bytes.Buffer

	for _, v := range result {
		buffer.WriteString(fmt.Sprintf("%02x", v&0xff))
	}

	if buffer.String() != rightAnswer {
		t.Errorf("EncodeAll 0 string incorrect, got: %s, want: %s.",
			buffer.String(),
			rightAnswer)
	}

	// [ [], [[]], [ [], [[]] ] ] = [ 0xc7, 0xc0, 0xc1, 0xc0, 0xc3, 0xc0, 0xc1, 0xc0 ]

	rightAnswer = "c7c0c1c0c3c0c1c0"
	rList = rlpList.NewRlpListVariadic(
		rlpList.NewRlpListVariadic(),                                                                                       //[]
		rlpList.NewRlpListVariadic(rlpList.NewRlpListVariadic()),                                                           //[[]]
		rlpList.NewRlpListVariadic(rlpList.NewRlpListVariadic(), rlpList.NewRlpListVariadic(rlpList.NewRlpListVariadic())), // [ [], [[]] ]
	)

	result = EncodeAll(rList)

	buffer.Reset()

	for _, v := range result {
		buffer.WriteString(fmt.Sprintf("%02x", v&0xff))
	}

	if buffer.String() != rightAnswer {
		t.Errorf("EncodeAll 0 string incorrect, got: %s, want: %s.",
			buffer.String(),
			rightAnswer)
	}
}
