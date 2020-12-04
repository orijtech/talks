package main
import "runtime/pprof"

func periodicallyCPUProfile(ctx context.Context, w io.Writer) error {
	for {
		if err := pprof.StartCPUProfile(w); err != nil {
			panic(err)
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(profilingPeriod):
			pprof.StopCPUProfile()
		}
		<-time.After(profilingPausePeriod)
	}
}
func main() {
	go periodicallyCPUProfile(ctx, w)
	defer cancel()
	// The rest of the logic goes down below...
}
