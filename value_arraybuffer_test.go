package v8go_test

import (
	"bytes"
	"fmt"
	"testing"

	v8 "github.com/herb-go/v8go"
)

func TestValueArrayBuffer(t *testing.T) {
	t.Parallel()
	iso := v8.NewIsolate()
	defer iso.Dispose()
	ctx := v8.NewContext(iso)
	defer ctx.Close()
	tdata := []byte("test array buffer")
	v, err := ctx.RunScript(fmt.Sprintf("new ArrayBuffer(%d)", len(tdata)), "arraybuffer.js")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	v8.WriteToArrayBuffer(v, tdata)

	result := v8.ArrayBufferContent(v)
	if !bytes.Equal(tdata, result) {
		t.Fatal("value not equal")
	}
}
