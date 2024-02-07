package rsync_client

import (
	"github.com/superrxan/rsync/internal/receivermaincmd"
)

type NopIO struct {
}

func (nop NopIO) Read(p []byte) (n int, err error) {
	return len(p), nil
}

func (nop NopIO) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func Sync(address, destination string) error {
	IO := &NopIO{}
	args := []string{"gokr-rsync", "-v", "--archive"}
	args = append(args, address, destination)
	_, err := receivermaincmd.Main(args, IO, IO, IO)
	return err
}
