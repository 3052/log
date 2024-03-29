package log

import (
   "io"
   "net/http"
   "testing"
)

const address = "https://go.dev/dl/go1.21.5.windows-amd64.zip"

func TestOne(t *testing.T) {
   var meter ProgressMeter
   meter.Set(1)
   res, err := http.Get(address)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   io.Copy(io.Discard, meter.Reader(res))
}

func TestTwo(t *testing.T) {
   var (
      meter ProgressMeter
      trip Transport
   )
   meter.Set(1)
   trip.Set()
   res, err := http.Get(address)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   io.Copy(io.Discard, meter.Reader(res))
}

func TestThree(t *testing.T) {
   var (
      log Level
      meter ProgressMeter
      trip Transport
   )
   log.Set()
   meter.Set(1)
   trip.Set()
   res, err := http.Get(address)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   io.Copy(io.Discard, meter.Reader(res))
}
