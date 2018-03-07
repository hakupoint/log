package log
import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"time"
)

type Log struct {
	cli  io.Writer
}
type info struct {
	Level  string          `json:"level"`
	Msg    string          `json:"msg"`
	Fields json.RawMessage `json:"fields"`
}

func NewLog() *Log {
	return &Log{
		cli: os.Stdout,
	}
}
func (l *Log) SetLean(lean *leancloud.LeanCloud) {
	l.lean = lean
}
func (l *Log) Info(msg string, f string) error {
	return l.write(&info{
		Level:  "info",
		Msg:    msg,
		Fields: json.RawMessage(f),
	})
}
func (l *Log) Warn(msg string, f string) error {
	return l.write(&info{
		Level:  "warn",
		Msg:    msg,
		Fields: json.RawMessage(f),
	})
}
func (l *Log) Error(msg string, f string) error {
	return l.write(&info{
		Level:  "error",
		Msg:    msg,
		Fields: json.RawMessage(f),
	})
}
func (l *Log) Fatal(msg string, f string) error {
	return l.write(&info{
		Level:  "fatal",
		Msg:    msg,
		Fields: json.RawMessage(f),
	})
}
func (l *Log) write(i *info) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	l.cli.Write([]byte(time.Now().Format("2006-01-02 15:04:05 ")))
	l.cli.Write(b)
	return nil
}