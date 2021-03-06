package integer

import (
	"testing"
	"time"
)

func TestFloat32ToInt(t *testing.T) {
	res, err := From(float32(32.25), 64)
	if err != nil {
		t.Error(err)
		return
	}
	if res != 32 {
		t.Error(res, "!=", 32)
		return
	}
}

func TestFloat32ToUInt(t *testing.T) {
	res := UnsignedFromByDef(float32(32.25), 64, 0)
	if res != 32 {
		t.Error(res, "!=", 32)
		return
	}
}

func TestFloat64ToInt(t *testing.T) {
	res, err := From(float64(64.1234), 64)
	if err != nil {
		t.Error(err)
		return
	}
	if res != 64 {
		t.Error(res, "!=", 64)
		return
	}
}

func TestFloat64ToUInt(t *testing.T) {
	res, err := UnsignedFrom(float64(64.1234), 64)
	if err != nil {
		t.Error(err)
		return
	}
	if res != 64 {
		t.Error(res, "!=", 64)
		return
	}
}

func TestTimeToInt(t *testing.T) {
	now := time.Now()
	res := FromByDef(now, 64, 0)
	if res != now.Unix() {
		t.Error(res, "!=", now.Unix())
		return
	}
}
