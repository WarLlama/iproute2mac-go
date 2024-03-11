package ifconfig

import (
	"slices"
	"strconv"
	"testing"
)

type testCase struct {
	testValue int
	flags     []string
}

var testCases = []testCase{
	{
		testValue: 0x8049,
		flags:     []string{"UP", "LOOPBACK", "RUNNING", "MULTICAST"},
	}, {
		testValue: 0x8010,
		flags:     []string{"POINTOPOINT", "MULTICAST"},
	}, {
		testValue: 0x8863,
		flags:     []string{"UP", "BROADCAST", "SMART", "RUNNING", "SIMPLEX", "MULTICAST"},
	}, {
		testValue: 0x8051,
		flags:     []string{"UP", "POINTOPOINT", "RUNNING", "MULTICAST"},
	},
}

func TestParseFlags(t *testing.T) {
	for _, tc := range testCases {
		t.Run(strconv.Itoa(tc.testValue), func(tt *testing.T) {
			flagList := ParseFlags(tc.testValue)
			if len(flagList) != len(tc.flags) {
				tt.Fatalf("incorrect flags returned, expected: %v got: %v", tc.flags, flagList)
			}
			for _, f := range tc.flags {
				if !slices.Contains(flagList, f) {
					tt.Fatalf("missing flag in list: %s\nTest List: %v\nGot List:  %v", f, tc.flags, flagList)
				}
			}
			for _, f := range flagList {
				if !slices.Contains(tc.flags, f) {
					tt.Fatalf("extra flag in list: %s\nTest List: %v\nGot List:  %v", f, tc.flags, flagList)
				}
			}
		})
	}
}
