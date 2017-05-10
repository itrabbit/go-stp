package float

import (
	"testing"
	"time"
)

func TestIntToFloat(t *testing.T) {
	res, err := To(int(2))
	if err != nil {
		t.Error(err)
		return
	}
	if res != 2.0 {
		t.Error(res, "!=", 2.0)
		return
	}
}

func TestInt8ToFloat(t *testing.T) {
	res, err := To(int8(8))
	if err != nil {
		t.Error(err)
		return
	}
	if res != 8.0 {
		t.Error(res, "!=", 8.0)
		return
	}
}

func TestInt16ToFloat(t *testing.T) {
	res, err := To(int16(16))
	if err != nil {
		t.Error(err)
		return
	}
	if res != 16.0 {
		t.Error(res, "!=", 16.0)
		return
	}
}

func TestInt32ToFloat(t *testing.T) {
	res, err := To(int32(32))
	if err != nil {
		t.Error(err)
		return
	}
	if res != 32.0 {
		t.Error(res, "!=", 32.0)
		return
	}
}

func TestInt64ToFloat(t *testing.T) {
	res, err := To(int64(64))
	if err != nil {
		t.Error(err)
		return
	}
	if res != 64.0 {
		t.Error(res, "!=", 64.0)
		return
	}
}

func TestUIntToFloat(t *testing.T) {
	res, err := To(uint(2))
	if err != nil {
		t.Error(err)
		return
	}
	if res != 2.0 {
		t.Error(res, "!=", 2.0)
		return
	}
}

func TestUInt8ToFloat(t *testing.T) {
	res, err := To(uint8(8))
	if err != nil {
		t.Error(err)
		return
	}
	if res != 8.0 {
		t.Error(res, "!=", 8.0)
		return
	}
}

func TestUInt16ToFloat(t *testing.T) {
	res, err := To(uint16(16))
	if err != nil {
		t.Error(err)
		return
	}
	if res != 16.0 {
		t.Error(res, "!=", 16.0)
		return
	}
}

func TestUInt32ToFloat(t *testing.T) {
	res, err := To(uint32(32))
	if err != nil {
		t.Error(err)
		return
	}
	if res != 32.0 {
		t.Error(res, "!=", 32.0)
		return
	}
}

func TestUInt64ToFloat(t *testing.T) {
	res, err := To(uint64(64))
	if err != nil {
		t.Error(err)
		return
	}
	if res != 64.0 {
		t.Error(res, "!=", 64.0)
		return
	}
}

func TestFloat32ToFloat(t *testing.T) {
	res, err := To(float32(32.25))
	if err != nil {
		t.Error(err)
		return
	}
	if res != 32.25 {
		t.Error(res, "!=", 32.25)
		return
	}
}

func TestFloat64ToFloat(t *testing.T) {
	res, err := To32(float64(64.1234))
	if err != nil {
		t.Error(err)
		return
	}
	if res != 64.1234 {
		t.Error(res, "!=", 64.1234)
		return
	}
}

func TestTimeToFloat(t *testing.T) {
	now := time.Now()
	res, err := To(now)
	if err != nil {
		t.Error(err)
		return
	}
	if res != float64(now.Unix()) {
		t.Error(res, "!=", float64(now.Unix()))
		return
	}
}
