package simplejack

import (
	"testing"

	"bytes"

	"bufio"
)

func TestMinTrace(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	writer := bufio.NewWriter(buffer)
	tracelog := New(TRACE, writer)
	tracelog.Trace.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	tracelog.Debug.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	tracelog.Info.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	tracelog.Warning.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	tracelog.Error.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	tracelog.Fatal.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
}

func TestMinDebug(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	writer := bufio.NewWriter(buffer)
	tracelog := New(DEBUG, writer)
	tracelog.Trace.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	tracelog.Debug.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	tracelog.Info.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	tracelog.Warning.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	tracelog.Error.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	tracelog.Fatal.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
}

func TestMinInfo(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	writer := bufio.NewWriter(buffer)
	tracelog := New(INFO, writer)
	tracelog.Trace.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	tracelog.Debug.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	tracelog.Info.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	tracelog.Warning.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	tracelog.Error.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	tracelog.Fatal.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
}

func TestMinWarning(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	writer := bufio.NewWriter(buffer)
	tracelog := New(WARNING, writer)
	tracelog.Trace.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	tracelog.Debug.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	tracelog.Info.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	tracelog.Warning.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	tracelog.Error.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	tracelog.Fatal.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
}

func TestMinError(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	writer := bufio.NewWriter(buffer)
	tracelog := New(ERROR, writer)
	tracelog.Trace.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	tracelog.Debug.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	tracelog.Info.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	tracelog.Warning.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	tracelog.Error.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	tracelog.Fatal.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
}

func TestMinFatal(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	writer := bufio.NewWriter(buffer)
	tracelog := New(FATAL, writer)
	tracelog.Trace.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	tracelog.Debug.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	tracelog.Info.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	tracelog.Warning.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	tracelog.Error.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	tracelog.Fatal.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
}
