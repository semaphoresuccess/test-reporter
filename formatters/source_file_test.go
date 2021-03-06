package formatters

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_SourceFile_Merge_With_Numbers(t *testing.T) {
	r := require.New(t)
	a := SourceFile{
		BlobID:   "a",
		Coverage: Coverage{NewNullInt(0), NewNullInt(2), NewNullInt(3), NewNullInt(0)},
	}
	b := SourceFile{
		BlobID:   "b",
		Coverage: Coverage{NewNullInt(1), NewNullInt(0), NewNullInt(1), NewNullInt(0)},
	}

	c, err := a.Merge(b)
	r.NoError(err)
	r.Equal("a", c.BlobID)
	r.Equal(4, len(c.Coverage))
	r.Equal(Coverage{NewNullInt(1), NewNullInt(2), NewNullInt(4), NewNullInt(0)}, c.Coverage)
	r.InDelta(75.0, c.CoveredPercent, 1)
	r.InDelta(2.2, c.CoveredStrength, 1)
	r.Equal(LineCounts{Total: 4, Missed: 1, Covered: 3, Strength: 7}, c.LineCounts)
}

func Test_SourceFile_Merge_With_Nulls(t *testing.T) {
	r := require.New(t)
	a := SourceFile{
		BlobID:   "a",
		Coverage: Coverage{NullInt{}, NewNullInt(2), NewNullInt(3), NewNullInt(0)},
	}
	b := SourceFile{
		BlobID:   "b",
		Coverage: Coverage{NewNullInt(1), NullInt{}, NewNullInt(3), NewNullInt(3)},
	}

	c, err := a.Merge(b)
	r.NoError(err)
	r.Equal("a", c.BlobID)
	r.Equal(4, len(c.Coverage))
	r.Equal(Coverage{NullInt{}, NullInt{}, NewNullInt(6), NewNullInt(3)}, c.Coverage)
	r.InDelta(100.0, c.CoveredPercent, 1)
	r.InDelta(4.5, c.CoveredStrength, 1)
	r.Equal(LineCounts{Total: 2, Missed: 0, Covered: 2, Strength: 9}, c.LineCounts)
}

func Test_SourceFile_BlobID(t *testing.T) {
	r := require.New(t)
	sf, err := NewSourceFile("./coverage.go", nil)
	r.NoError(err)
	r.NotZero(sf.BlobID)
	r.NotContains(sf.BlobID, "blob")
}
