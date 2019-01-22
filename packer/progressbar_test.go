package packer

import (
	"bytes"
	"io/ioutil"
	"testing"
	"time"

	"github.com/cheggaaa/pb"
	"golang.org/x/sync/errgroup"
)

func speedyProgressBar(bar *pb.ProgressBar) {
	bar.SetUnits(pb.U_BYTES)
	bar.SetRefreshRate(1 * time.Millisecond)
	bar.NotPrint = true
	bar.Format("[\x00=\x00>\x00-\x00]")
}
func TestProgressTracking_races(t *testing.T) {
	var bar *uiProgressBar
	g := errgroup.Group{}
	txt := []byte("foobarbaz dolores")
	b := bytes.NewReader(txt)

	for i := 0; i < 100; i++ {
		g.Go(func() error {
			tracker := bar.TrackProgress("file,", 1, 42, ioutil.NopCloser(b))

			b := []byte("i")
			for i := 0; i < 42; i++ {
				tracker.Read(b)
			}
			return nil
		})
	}

	g.Wait()
}
