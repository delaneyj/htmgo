package rendertest

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/valyala/bytebufferpool"
)

func BenchmarkMailToStatic(b *testing.B) {
	b.ReportAllocs()
	ctx := context.Background()
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	page := MailTo("myemail")
	for i := 0; i < b.N; i++ {
		buf.Reset()
		page.Render(ctx, buf)
	}
}

func BenchmarkMailToDynamic(b *testing.B) {
	b.ReportAllocs()
	ctx := context.Background()
	buf := bytebufferpool.Get()
	defer bytebufferpool.Put(buf)

	uuids := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		uuids[i] = uuid.NewString()
	}

	for i := 0; i < b.N; i++ {
		buf.Reset()
		MailTo(uuids[i]).Render(ctx, buf)
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
