package main
import "runtime/pprof"

func periodicallyMemoryProfile(ctx context.Context, w io.Writer) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(profilingPeriod):
			if err := pprof.WriteHeapProfile(w); err != nil {
				panic(err)
			}
		}
		<-time.After(profilingPausePeriod)
	}
}
func main() {
	go periodicallyMemoryProfile(ctx, w)
	defer cancel()
	// The rest of the logic goes down below...
}
