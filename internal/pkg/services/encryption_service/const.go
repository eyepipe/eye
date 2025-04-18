package encryption_service

import (
	"fmt"
	"io"

	"encoding/binary"
)

const (
	ByteVer           = byte(10)
	ByteSalt          = byte(20)
	ByteIV            = byte(30)
	ByteBody          = byte(99)
	ByteReservedBegin = byte(100)
	ByteReservedEnd   = byte(255)

	// Ver1 buildinfo 1
	Ver1 = uint32(1)
)

var ByteOrder = binary.BigEndian

func PutByteWithData(w io.Writer, b byte, data []byte) (err error) {
	err = PutByte(w, b)
	if err != nil {
		return fmt.Errorf("failed to binary.Write: %w", err)
	}

	err = PutUint32(w, uint32(len(data)))
	if err != nil {
		return fmt.Errorf("failed to PutUint32: %w", err)
	}

	_, err = w.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write: %w", err)
	}

	return nil
}

func GetData(r io.Reader) (buf []byte, err error) {
	var size uint32
	err = GetUint32(r, &size)
	if err != nil {
		return nil, fmt.Errorf("failed to read size: %w", err)
	}

	buf = make([]byte, size)
	_, err = io.ReadAtLeast(r, buf, int(size))
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %w", err)
	}
	return buf, nil
}

func PutByte(w io.Writer, b byte) (err error) {
	return binary.Write(w, ByteOrder, b)
}

func PutUint32(w io.Writer, v uint32) (err error) {
	return binary.Write(w, ByteOrder, v)
}

func GetByte(r io.Reader, out *byte) error {
	return binary.Read(r, ByteOrder, out)
}

func GetUint32(r io.Reader, out *uint32) error {
	return binary.Read(r, ByteOrder, out)
}
