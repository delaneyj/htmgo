package rendertest

import (
	"context"
	"testing"

	"github.com/valyala/bytebufferpool"
)

func BenchmarkMailTo(b *testing.B) {
	b.ReportAllocs()

	ctx := context.Background()
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		MailTo().Render(ctx, buf)
	}
}

func BenchmarkComplexPage(b *testing.B) {
	b.ReportAllocs()
	ctx := context.Background()
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)
	for i := 0; i < b.N; i++ {
		buf.Reset()
		ComplexPage().Render(ctx, buf)
	}
}
