package packet

import (
	"os"
	"testing"
	"time"

	"go.uber.org/zap"

	"github.com/alexbegoon/go-dcc/internal/driver/dummy"
)

func TestSend(t *testing.T) {
	if os.Getenv("TRAVIS") == "true" {
		// This facilitates that tests pass on travis :(
		dummy.ByteOneMax = 94 * time.Microsecond
	}
	d := &dummy.DCCDummy{Log: zap.NewExample()}
	p := NewBroadcastIdlePacket(d)
	d.TracksOn()
	p.Send()
	time.Sleep(1 * time.Second) // nolint:forbidigo
	packetStr := dummy.GuessBuffer.String()
	t.Log("Pckt: ", p.String())
	t.Log("Sent: ", packetStr)

	if packetStr != p.String() {
		t.Error("should have sent the encoded package")
	}
}

func TestNewPacket(t *testing.T) {
	p := NewPacket(&dummy.DCCDummy{}, 0xFF, []byte{0x01})
	if p.String() != "11111111111111110111111110000000010111111101" {
		t.Error("Bad packet: ", p.String())
	}
}

func TestNewBaselinePacket(t *testing.T) {
	p := NewBaselinePacket(&dummy.DCCDummy{}, 0xFF, []byte{0x01})
	if p.String() != "11111111111111110011111110000000010011111101" {
		t.Error("Bad packet: ", p.String())
	}
}

func TestIdlePacket(t *testing.T) {
	p := NewBroadcastIdlePacket(&dummy.DCCDummy{})
	if p.String() != "11111111111111110111111110000000000111111111" {
		t.Error("Bad idle packet")
	}
}

func TestNewSpeedAndDirectionPacket(t *testing.T) {
	p := NewSpeedAndDirectionPacket(&dummy.DCCDummy{}, 0xFF, 0xFF, byte(1))
	if p.String() != "11111111111111110011111110011111110000000001" {
		t.Error("Bad speed and direction packet: ", p.String())
	}
}

func TestNewFunctionGroupOnePacket(t *testing.T) {
	p := NewFunctionGroupOnePacket(&dummy.DCCDummy{}, 0xFF, true, true, true, true, true)
	if p.String() != "11111111111111110111111110100111110011000001" {
		t.Error("Bad Function Group One packet: ", p.String())
	}
}

func TestNewBroadcastResetPacket(t *testing.T) {
	p := NewBroadcastResetPacket(&dummy.DCCDummy{})
	if p.String() != "11111111111111110000000000000000000000000001" {
		t.Error("Bad reset packet")
	}
}

func TestNewBroadcastStopPacket(t *testing.T) {
	p := NewBroadcastStopPacket(&dummy.DCCDummy{}, byte(0), true, false)
	if p.String() != "11111111111111110000000000010000000010000001" {
		t.Error("Bad stop packet: ", p.String())
	}
}

func TestCorrectSpeedPacketBits(t *testing.T) {
	p := NewSpeedAndDirectionPacket(&dummy.DCCDummy{}, 0xFF, 0, byte(1))
	if p.String() != "11111111111111110011111110011000000000111111" {
		t.Error("Bad speed and direction packet: ", p.String())
	}

	p = NewSpeedAndDirectionPacket(&dummy.DCCDummy{}, 0xFF, 1, byte(1))
	if p.String() != "11111111111111110011111110011100000000011111" {
		t.Error("Bad speed and direction packet: ", p.String())
	}

	p = NewSpeedAndDirectionPacket(&dummy.DCCDummy{}, 0xFF, 2, byte(1))
	if p.String() != "11111111111111110011111110011000010000111101" {
		t.Error("Bad speed and direction packet: ", p.String())
	}

	p = NewSpeedAndDirectionPacket(&dummy.DCCDummy{}, 0xFF, 11, byte(1))
	if p.String() != "11111111111111110011111110011101010000010101" {
		t.Error("Bad speed and direction packet: ", p.String())
	}

	p = NewSpeedAndDirectionPacket(&dummy.DCCDummy{}, 0xFF, 13, byte(1))
	if p.String() != "11111111111111110011111110011101100000010011" {
		t.Error("Bad speed and direction packet: ", p.String())
	}

	p = NewSpeedAndDirectionPacket(&dummy.DCCDummy{}, 0xFF, 30, byte(1))
	if p.String() != "11111111111111110011111110011011110000100001" {
		t.Error("Bad speed and direction packet: ", p.String())
	}

	p = NewSpeedAndDirectionPacket(&dummy.DCCDummy{}, 0xFF, 31, byte(1))
	if p.String() != "11111111111111110011111110011111110000000001" {
		t.Error("Bad speed and direction packet: ", p.String())
	}

	p = NewSpeedAndDirectionPacket(&dummy.DCCDummy{}, 0xFF, 27, byte(1))
	if p.String() != "11111111111111110011111110011111010000000101" {
		t.Error("Bad speed and direction packet: ", p.String())
	}
}

func TestNewFunctionGroupTwoPacket(t *testing.T) {
	p0, p1 := NewFunctionGroupTwoPacket(&dummy.DCCDummy{}, 0xFF, true, false, true, false, true,
		false, true, false)
	if p0.String() != "11111111111111110111111110101101010010010101" {
		t.Error("Bad Function Group Two packet 0: ", p0.String())
	}
	if p1.String() != "11111111111111110111111110101001010010110101" {
		t.Error("Bad Function Group Two packet 1: ", p1.String())
	}
}

func TestNewFunctionExpansionPacket(t *testing.T) {
	p0, p1 := NewFunctionExpansionPacket(&dummy.DCCDummy{}, 0xFF, true, false, true, false,
		true, false, true, false, true,
		false, true, false, true, false, true, false)
	if p0.String() != "11111111111111110111111110110111100010101010011101001" {
		t.Error("Bad Function Expansion packet 0: ", p0.String())
	}
	if p1.String() != "11111111111111110111111110110111110010101010011101011" {
		t.Error("Bad Function Expansion packet 1: ", p1.String())
	}
}
