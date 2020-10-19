package timeStamp

import "testing"

func TestGetUnixTime(t *testing.T) {
	start, end, err := GetUnixTimeStamp()
	if err != nil {
		t.Errorf("Error calculating start and end unix timestamps: %v\n", err)
	}
	if start < 0 && end < 0 {
		t.Errorf("An invalid timestamp was produced: %d, %d\n", start, end)
	}
}
