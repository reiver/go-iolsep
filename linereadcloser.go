package iolsep

import (
	"errors"
	"io"

	"github.com/reiver/go-unicode"
	"sourcecode.social/reiver/go-utf8"
)

type internalLineReadCloser struct {
	reader io.Reader
	pipereader *io.PipeReader
	pipewriter *io.PipeWriter
}

func NewLineReadCloser(reader io.Reader) io.ReadCloser {

	if nil == reader {
		return nil
	}

	pipereader, pipewriter := io.Pipe()
	if nil == pipereader {
		return nil
	}
	if nil == pipewriter {
		return nil
	}

	linereader := internalLineReadCloser{
		reader:reader,
		pipereader:pipereader,
		pipewriter:pipewriter,
	}

	go linereader.pipewrite()

	return &linereader
}

func (receiver *internalLineReadCloser) Close() error {
	if nil == receiver {
		return errNilReceiver
	}

	var pipewritererror error
	{
		var pipewriter *io.PipeWriter = receiver.pipewriter
		if nil != pipewriter {
			err := pipewriter.Close()
			if nil != err {
				pipewritererror = err
			}
		}
	}

	{
		var pipereader *io.PipeReader = receiver.pipereader
		if nil != pipereader {
			err := pipereader.Close()
			if nil != err {
				return err
			}
		}
	}

	if nil != pipewritererror {
		return pipewritererror
	}

	return nil
}

func (receiver *internalLineReadCloser) Read(p []byte) (n int, err error) {
	if nil == receiver {
		return 0, errNilReceiver
	}

	var reader io.Reader = receiver.reader
	if nil == reader {
		return 0, errNilReader
	}

	var pipereader io.Reader = receiver.pipereader
	if nil == pipereader {
		return 0, errNilPipeReader
	}

	return pipereader.Read(p)
}

func (receiver *internalLineReadCloser) pipewrite() {
	if nil == receiver {
		panic(errNilReceiver)
	}

	var pipewriter *io.PipeWriter = receiver.pipewriter
	if nil == pipewriter {
		panic(errNilPipeWriter)
	}

	var reader io.Reader = receiver.reader
	if nil == reader {
		err := pipewriter.CloseWithError(errNilReader)
		if nil != err {
			panic(err)
		}
	}

	pipewrite(receiver.writerune, receiver.returneof, receiver.returnerror, reader)
}

func (receiver *internalLineReadCloser) returneof() {
	if nil == receiver {
		panic(errNilWriter)
	}

	var pipewriter *io.PipeWriter = receiver.pipewriter
	if nil == pipewriter {
		panic(errNilPipeWriter)
	}

	{
		err := pipewriter.CloseWithError(io.EOF)
		if nil != err {
			panic(err)
		}
	}
}

func (receiver *internalLineReadCloser) returnerror(err error) {
	if nil == receiver {
		panic(errNilWriter)
	}

	var pipewriter *io.PipeWriter = receiver.pipewriter
	if nil == pipewriter {
		panic(errNilPipeWriter)
	}

	{
		err := pipewriter.CloseWithError(err)
		if nil != err {
			panic(err)
		}
	}
}

func (receiver *internalLineReadCloser) writerune(r rune) (exit bool) {
	if nil == receiver {
		panic(errNilWriter)
	}

	var pipewriter *io.PipeWriter = receiver.pipewriter
	if nil == pipewriter {
		panic(errNilPipeWriter)
	}

	{
		_, err := utf8.WriteRune(pipewriter, r)
		if errors.Is(err, io.ErrClosedPipe) {
			pipewriter.Close()
			return true
		}
		if nil != err {
			e := pipewriter.CloseWithError(err)
			if nil != e {
				panic(e)
			}
		}
	}

	return false
}

func pipewrite(writerune func(rune)bool, returneof func(), returnerror func(error), reader io.Reader) {

	if nil == writerune {
		panic(errNilWriteRuneFunction)
	}

	if nil == returneof {
		panic(errNilReturnEOFFunction)
	}

	if nil == returnerror {
		panic(errNilReturnErrorFunction)
	}

	for {
		r, size, err := utf8.ReadRune(reader)
		if 0 < size {
			exit := writerune(r)
			if exit {
				return
			}

			switch r {
			case unicode.LF:
				returneof()
				return
			case unicode.NEL:
				returneof()
				return
			case unicode.LS:
				returneof()
				return
			case unicode.PS:
				returneof()
				return
			}
		}
		if io.EOF == err {
			returneof()
			return
		}
		if nil != err {
			returnerror(err)
			return
		}
	}
}
