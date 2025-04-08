package resp

// Adapted from:
// https://www.redisgreen.net/blog/beginners-guide-to-redis-protocol/
// https://www.redisgreen.net/blog/reading-and-writing-redis-protocol/

// Maybe use this:
// https://godoc.org/github.com/fzzy/radix/redis/resp

// Either way this code will be refactored soon (20161229/thisisaaronland)

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	_ "log"
	"os"
	"strconv"
)

var (
	ErrInvalidSyntax = errors.New("resp: invalid syntax")
)

type RESPReader struct {
	*bufio.Reader
}

func NewRESPReader(reader io.Reader, sz int) *RESPReader {
	return &RESPReader{
		Reader: bufio.NewReaderSize(reader, sz),
	}
}

func NewRESPDebugReader(reader io.Reader, sz int) *RESPReader {

	tee := io.TeeReader(reader, os.Stdout)

	return &RESPReader{
		Reader: bufio.NewReaderSize(tee, sz),
	}
}

func (r *RESPReader) ReadObject() ([]byte, error) {

	line, err := r.readLine()

	if err != nil {
		return nil, err
	}

	switch string(line[0]) {
	case RESP_SIMPLE_STRING, RESP_INTEGER, RESP_ERROR:
		return line, nil
	case RESP_BULK_STRING:
		return r.readBulkString(line)
	case RESP_ARRAY:
		return r.readArray(line)
	default:
		return nil, ErrInvalidSyntax
	}
}

func (r *RESPReader) readLine() (line []byte, err error) {
	line, err = r.ReadSlice('\n')
	if err != nil {
		return nil, err
	}

	if len(line) > 1 && line[len(line)-2] == '\r' {
		return line, nil
	} else {
		// Line was too short or \n wasn't preceded by \r.
		return nil, ErrInvalidSyntax
	}
}

func (r *RESPReader) readBulkString(line []byte) ([]byte, error) {
	count, err := r.getCount(line)
	if err != nil {
		return nil, err
	}
	if count == -1 {
		return line, nil
	}

	buf := make([]byte, len(line)+count+2)
	copy(buf, line)
	_, err = r.Read(buf[len(line):])
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *RESPReader) getCount(line []byte) (int, error) {
	end := bytes.IndexByte(line, '\r')
	return strconv.Atoi(string(line[1:end]))
}

func (r *RESPReader) readArray(line []byte) ([]byte, error) {
	// Get number of array elements.
	count, err := r.getCount(line)
	if err != nil {
		return nil, err
	}

	// Read `count` number of RESP objects in the array.
	for i := 0; i < count; i++ {
		buf, err := r.ReadObject()
		if err != nil {
			return nil, err
		}
		line = append(line, buf...)
	}

	return line, nil
}
