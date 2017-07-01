package simplejack

import (
	"testing"

	"bytes"

	"bufio"
)

func TestMinTrace(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	writer := bufio.NewWriter(buffer)
	logger := New(TRACE, writer)
	logger.Trace.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	logger.Debug.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	logger.Info.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	logger.Warning.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	logger.Error.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	logger.Fatal.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
}

func TestMinDebug(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	writer := bufio.NewWriter(buffer)
	logger := New(DEBUG, writer)
	logger.Trace.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	logger.Debug.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	logger.Info.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	logger.Warning.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	logger.Error.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	logger.Fatal.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
}

func TestMinInfo(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	writer := bufio.NewWriter(buffer)
	logger := New(INFO, writer)
	logger.Trace.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	logger.Debug.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	logger.Info.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	logger.Warning.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	logger.Error.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	logger.Fatal.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
}

func TestMinWarning(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	writer := bufio.NewWriter(buffer)
	logger := New(WARNING, writer)
	logger.Trace.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	logger.Debug.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	logger.Info.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	logger.Warning.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	logger.Error.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	logger.Fatal.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
}

func TestMinError(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	writer := bufio.NewWriter(buffer)
	logger := New(ERROR, writer)
	logger.Trace.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	logger.Debug.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	logger.Info.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	logger.Warning.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	logger.Error.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
	buffer.Reset()
	logger.Fatal.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
}

func TestMinFatal(t *testing.T) {
	buffer := bytes.NewBuffer(nil)
	writer := bufio.NewWriter(buffer)
	logger := New(FATAL, writer)
	logger.Trace.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	logger.Debug.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	logger.Info.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	logger.Warning.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	logger.Error.Print("TEST")
	writer.Flush()
	if buffer.Len() != 0 {
		t.Error("Buffer should be empty here")
	}
	buffer.Reset()
	logger.Fatal.Print("TEST")
	writer.Flush()
	if buffer.Len() == 0 {
		t.Error("Buffer shouldnt be empty here")
	}
}
