package util

import "testing"

func Test_BoundaryIndices(t *testing.T) {
	tests := []struct {
		name                                string
		left, right, rng, minLeft, maxRight int
		expectedLeft, expectedRight         int
	}{
		{
			name: "(5,6) (rng3) => 2,9",
			left: 5, right: 6, rng: 3, minLeft: 0, maxRight: 10,
			expectedLeft: 2, expectedRight: 9,
		},
		{
			name: "(0,1) (rng3) => 0,4",
			left: 0, right: 1, rng: 3, minLeft: 0, maxRight: 7,
			expectedLeft: 0, expectedRight: 4,
		},
		{
			name: "(3,5) (rng2) => 1,7",
			left: 3, right: 5, rng: 2, minLeft: 0, maxRight: 50,
			expectedLeft: 1, expectedRight: 7,
		},
		{
			name: "(47,49) (rng2) => 45,50",
			left: 47, right: 49, rng: 2, minLeft: 0, maxRight: 50,
			expectedLeft: 45, expectedRight: 50,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			aL, aR := BoundaryIndices(tc.left, tc.right, tc.rng, tc.minLeft, tc.maxRight)
			if aL != tc.expectedLeft {
				t.Errorf("expected left: %d, got: %d", tc.expectedLeft, aL)
			}
			if aR != tc.expectedRight {
				t.Errorf("expected right: %d, got: %d", tc.expectedRight, aR)
			}
		})
	}
}
