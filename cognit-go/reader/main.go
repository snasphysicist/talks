package main

import "io"
import "log"
import "time"

type ProgressReporter struct{}

func (*ProgressReporter) BytesRead() uint64 { return 0 }

func (*ProgressReporter) ReportBytesRead(uint64) {}

type Source struct{}

func (*Source) Open() error { return nil }

func (*Source) Close() error { return nil }

func (*Source) Read([]byte) (int, error) { return 0, nil }

type Buffer struct{}

func (*Buffer) Offer([]byte) bool { return false }

type Reader struct {
	source     *Source
	b          *Buffer
	pr         *ProgressReporter
	toTransfer uint64
	done       chan struct{}
}

func (r *Reader) Start() {
	err := r.source.Open()
	if err != nil {
		panic(err)
	}
	defer r.source.Close()
	for {
		buf := make([]byte, 1000)
		n, err := r.source.Read(buf)
		if err == io.EOF {
			if r.pr.BytesRead() != r.toTransfer {
				log.Printf(
					"WARNING: transferred %d bytes, should have been %d",
					r.pr.BytesRead(), r.toTransfer,
				)
			}
			close(r.done)
			return
		}
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n > 0 {
			success := r.b.Offer(buf[:n])
			for !success {
				success = r.b.Offer(buf[:n])
				time.Sleep(1 * time.Millisecond)
			}
		}
		r.pr.ReportBytesRead(uint64(n))
	}
}
