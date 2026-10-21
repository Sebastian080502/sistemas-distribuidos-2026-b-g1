package observability

import "net/http"

func TraceMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := NewContext(r.Header.Get(HeaderTraceparent), r.Header.Get(HeaderCorrelation))
		w.Header().Set(HeaderTraceparent, FormatTraceparent(ctx))
		w.Header().Set(HeaderCorrelation, ctx.CorrelationID)
		r = r.WithContext(WithValue(r.Context(), ctx))
		next.ServeHTTP(w, r)
	})
}
