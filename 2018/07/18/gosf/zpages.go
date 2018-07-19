import (
	"log"
	"net/http"

	"go.opencensus.io/zpages"
)

func main() {
	mux := http.NewServeMux()
	zpages.Handle(mux, "/debug")
	log.Fatal(http.ListenAndServe(":7788", mux))
}
