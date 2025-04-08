package resp

// Adapted from:
// https://www.redisgreen.net/blog/beginners-guide-to-redis-protocol/
// https://www.redisgreen.net/blog/reading-and-writing-redis-protocol/

// Maybe use this:
// https://godoc.org/github.com/fzzy/radix/redis/resp

// Either way this code will be refactored soon (20161229/thisisaaronland)

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

type RESPWriter struct {
	*bufio.Writer
}

func NewRESPWriter(writer io.Writer) *RESPWriter {

	return &RESPWriter{
		Writer: bufio.NewWriter(writer),
	}
}

func NewRESPDebugWriter(writer io.Writer) *RESPWriter {

	writers := []io.Writer{
		writer,
		os.Stdout,
	}

	multi := io.MultiWriter(writers...)

	return &RESPWriter{
		Writer: bufio.NewWriter(multi),
	}
}

func (w *RESPWriter) WriteCountString(count int) error {

	w.WriteString(string(RESP_ARRAY))
	w.WriteString(strconv.Itoa(count))
	w.WriteString(string(RESP_NEWLINE))

	return w.Flush()
}

func (w *RESPWriter) WriteNumberString(count int) error {

	w.WriteString(string(RESP_INTEGER))
	w.WriteString(strconv.Itoa(count))
	w.WriteString(string(RESP_NEWLINE))

	return w.Flush()
}

func (w *RESPWriter) WriteBulkStringMessage(str string) error {

	str_len := len(str)

	w.WriteString(string(RESP_BULK_STRING))
	w.WriteString(strconv.Itoa(str_len))
	w.WriteString(string(RESP_NEWLINE))

	w.WriteString(str)
	w.WriteString(string(RESP_NEWLINE))

	return w.Flush()
}

func (w *RESPWriter) WriteStringMessage(str ...string) error {

	for _, s := range str {
		w.WriteString(string(RESP_SIMPLE_STRING))
		w.WriteString(s)
		w.WriteString(string(RESP_NEWLINE))
	}

	return w.Flush()
}

func (w *RESPWriter) WriteNullMessage() error {

	w.WriteString(string(RESP_BULK_STRING))
	w.WriteString("-1")
	w.WriteString(string(RESP_NEWLINE))

	return w.Flush()
}

func (w *RESPWriter) WriteSubscribeMessage(channels []string) error {

	for i, ch := range channels {

		w.WriteCountString(3)
		w.WriteBulkStringMessage("subscribe")
		w.WriteBulkStringMessage(ch)
		w.WriteNumberString(i + 1)
	}

	return w.Flush()
}

func (w *RESPWriter) WriteUnsubscribeMessage(channels []string) error {

	i := len(channels) - 1

	for _, ch := range channels {

		w.WriteCountString(3)
		w.WriteBulkStringMessage("unsubscribe")
		w.WriteBulkStringMessage(ch)
		w.WriteNumberString(i - 1)
	}

	return w.Flush()
}

func (w *RESPWriter) WritePublishMessage(channel string, msg string) error {

	w.WriteCountString(3)
	w.WriteBulkStringMessage("message")
	w.WriteBulkStringMessage(channel)
	w.WriteBulkStringMessage(msg)

	return w.Flush()
}

func (w *RESPWriter) WriteErrorMessage(err error) error {

	w.WriteString(string(RESP_SIMPLE_STRING))
	w.WriteString(fmt.Sprintf("%s", err))
	w.WriteString(string(RESP_NEWLINE))

	return w.Flush()
}
