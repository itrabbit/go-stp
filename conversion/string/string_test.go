package string

import (
	"testing"
	"time"
)

func TestIntToString(t *testing.T) {
	res, err := To(int(2))
	if err != nil {
		t.Error(err)
		return
	}
	if res != "2" {
		t.Error(res, "!=", "2")
		return
	}
}

func TestInt8ToString(t *testing.T) {
	res, err := To(int8(8))
	if err != nil {
		t.Error(err)
		return
	}
	if res != "8" {
		t.Error(res, "!=", "8")
		return
	}
}

func TestInt16ToString(t *testing.T) {
	res, err := To(int16(16))
	if err != nil {
		t.Error(err)
		return
	}
	if res != "16" {
		t.Error(res, "!=", "16")
		return
	}
}

func TestInt32ToString(t *testing.T) {
	res, err := To(int32(32))
	if err != nil {
		t.Error(err)
		return
	}
	if res != "32" {
		t.Error(res, "!=", "32")
		return
	}
}

func TestInt64ToString(t *testing.T) {
	res, err := To(int64(64))
	if err != nil {
		t.Error(err)
		return
	}
	if res != "64" {
		t.Error(res, "!=", "64")
		return
	}
}

func TestUIntToString(t *testing.T) {
	res, err := To(uint(2))
	if err != nil {
		t.Error(err)
		return
	}
	if res != "2" {
		t.Error(res, "!=", "2")
		return
	}
}

func TestUInt8ToString(t *testing.T) {
	res, err := To(uint8(8))
	if err != nil {
		t.Error(err)
		return
	}
	if res != "8" {
		t.Error(res, "!=", "8")
		return
	}
}

func TestUInt16ToString(t *testing.T) {
	res, err := To(uint16(16))
	if err != nil {
		t.Error(err)
		return
	}
	if res != "16" {
		t.Error(res, "!=", "16")
		return
	}
}

func TestUInt32ToString(t *testing.T) {
	res, err := To(uint32(32))
	if err != nil {
		t.Error(err)
		return
	}
	if res != "32" {
		t.Error(res, "!=", "32")
		return
	}
}

func TestUInt64ToString(t *testing.T) {
	res, err := To(uint64(64))
	if err != nil {
		t.Error(err)
		return
	}
	if res != "64" {
		t.Error(res, "!=", "64")
		return
	}
}

func TestFloat32ToString(t *testing.T) {
	res, err := To(float32(32.25))
	if err != nil {
		t.Error(err)
		return
	}
	if res != "32.25" {
		t.Error(res, "!=", "32.25")
		return
	}
}

func TestFloat64ToString(t *testing.T) {
	res, err := To(float64(64.1234))
	if err != nil {
		t.Error(err)
		return
	}
	if res != "64.1234" {
		t.Error(res, "!=", "64.1234")
		return
	}
}

func TestTimeToString(t *testing.T) {
	now := time.Now()
	res, err := To(now)
	if err != nil {
		t.Error(err)
		return
	}
	if res != now.Format(time.RFC3339) {
		t.Error(res, "!=", now.Format(time.RFC3339))
		return
	}
}
