package mockgen

import (
	"fmt"
	"os"

	"github.com/sanposhiho/gomockhandler/internal/util"
)

type Runner interface {
	Run() error

	SetSource(new string)
	SetDestination(new string)
	GetDestination() string
	GetSource() string
}

// TestRun checks if mockgen can run with the given options.
// It does not actually update the mock.
func TestRun(r Runner) error {
	d := r.GetDestination()
	tmpFile := util.TmpFilePath(d)
	defer os.Remove(tmpFile)

	// use tmpfile to test generating mock
	r.SetDestination(tmpFile)
	defer r.SetDestination(d)

	if err := r.Run(); err != nil {
		return fmt.Errorf("run mockgen: %w", err)
	}
	return nil
}

func Checksum(r Runner) ([16]byte, error) {
	d := r.GetDestination()
	tmpFile := util.TmpFilePath(d)
	defer os.Remove(tmpFile)

	// use tmpfile to test generating mock
	r.SetDestination(tmpFile)
	defer r.SetDestination(d)

	if err := r.Run(); err != nil {
		return [16]byte{}, fmt.Errorf("run mockgen: %w", err)
	}

	checksum, err := util.MockCheckSum(tmpFile)
	if err != nil {
		return [16]byte{}, fmt.Errorf("calculate checksum of the mock: %v", err)
	}

	return checksum, nil
}