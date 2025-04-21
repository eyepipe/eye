package uuidv7

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

const TimePrecision = time.Millisecond

type UUIDv7 struct {
	Time  time.Time
	Shard uint16
	data  []byte
}

// Decode inout "061cb26a-54b8-7a52-8000-2124e7041024"
func Decode(input string) (*UUIDv7, error) {
	cleanUUID := strings.ReplaceAll(input, "-", "")
	data, err := hex.DecodeString(cleanUUID)
	switch {
	case err != nil:
		return nil, fmt.Errorf("%w: invalid UUID format: %w", ErrDecodeFailed, err)
	case len(data) != 16:
		return nil, fmt.Errorf("%w: invalid UUID length: %d", ErrDecodeFailed, len(data))
	}

	// first 6 байт — is a timestamp
	timestampMs := uint64(data[0])<<40 |
		uint64(data[1])<<32 |
		uint64(data[2])<<24 |
		uint64(data[3])<<16 |
		uint64(data[4])<<8 |
		uint64(data[5])

	timestamp := time.UnixMilli(int64(timestampMs))
	// last 2 bytes is a shard number
	shard := binary.BigEndian.Uint16(data[14:16])

	return &UUIDv7{
		Time:  timestamp,
		Shard: shard,
		data:  data,
	}, nil
}

func (u *UUIDv7) String() string {
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		binary.BigEndian.Uint32(u.data[0:4]),
		binary.BigEndian.Uint16(u.data[4:6]),
		binary.BigEndian.Uint16(u.data[6:8]),
		binary.BigEndian.Uint16(u.data[8:10]),
		u.data[10:16],
	)
}

func NewWithShard(shard uint16) *UUIDv7 {
	return NewWithTimeShard(time.Now().Round(TimePrecision), shard)
}

func NewWithTimeShard(time time.Time, shard uint16) *UUIDv7 {
	data := make([]byte, 16)
	timestamp := uint64(time.UnixMilli())
	data[0] = byte(timestamp >> 40)
	data[1] = byte(timestamp >> 32)
	data[2] = byte(timestamp >> 24)
	data[3] = byte(timestamp >> 16)
	data[4] = byte(timestamp >> 8)
	data[5] = byte(timestamp)

	// entropy 8 bytes (6–13)
	_, err := rand.Read(data[6:14])
	if err != nil {
		panic(fmt.Errorf("failed to generate UUID rand: %v", err))
	}

	// setup buildinfo 7 (0111) — bits 48-51 → byte 6 (uuid[6])
	data[6] = (data[6] & 0x0F) | 0x70
	// setup variant (10xx) — bits 64-65 → byte 8 (uuid[8])
	data[8] = (data[8] & 0x3F) | 0x80

	// last 2 bytes — shard ID
	binary.BigEndian.PutUint16(data[14:], shard)
	return &UUIDv7{
		Time:  time,
		Shard: shard,
		data:  data,
	}
}
