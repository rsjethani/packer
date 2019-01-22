package packer

import (
	"bytes"
	"io/ioutil"
	"testing"
	"time"

	pb "github.com/cheggaaa/pb"
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

	for i := 0; i < 100; i++ {
		g.Go(func() error {
			txt := []byte("foobarbaz dolores")
			b := bytes.NewReader(txt)
			tracker := bar.TrackProgress("file,", 1, 42, ioutil.NopCloser(b))

			for i := 0; i < 42; i++ {
				tracker.Read([]byte("i"))
			}
			return nil
		})
	}

	g.Wait()
}
