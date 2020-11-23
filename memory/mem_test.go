package memory

import "testing"

func TestMemManager(t *testing.T) {
	mem, err := InitMemManager(0, 0, DefaultMinHeapMemSize, DefaultMaxHeapMemSize)
	if err != nil {
		t.Fatalf("init mem failed: %v", err)
	}
	p, err := mem.Malloc(100)
	if err != nil {
		t.Fatalf("Malloc failed: %v", err)
	}
	if err = mem.Free(p); err != nil {
		t.Fatalf("Free failed: %v", err)
	}
	t.Log("ok")
}
