package partials

import (
	"github.com/maddalax/htmgo/framework/h"
	"htmgo-site/internal/datastructures"
	"htmgo-site/internal/dirwalk"
	"strings"
)

func formatPart(part string) string {
	if part[1] == '_' {
		part = part[2:]
	}
	part = strings.ReplaceAll(part, "-", " ")
	part = strings.ReplaceAll(part, "_", " ")
	part = strings.Title(part)
	return part
}

func CreateAnchor(parts []string) string {
	return strings.Join(h.Map(parts, func(part string) string {
		return strings.ReplaceAll(strings.ToLower(formatPart(part)), " ", "-")
	}), "-")
}

func partsToName(parts []string) string {
	builder := strings.Builder{}
	for i, part := range parts {
		if i == 0 {
			continue
		}
		part = formatPart(part)
		builder.WriteString(part)
		builder.WriteString(" ")
	}

	return builder.String()
}

func groupByFirstPart(pages []*dirwalk.Page) *datastructures.OrderedMap[string, []*dirwalk.Page] {
	grouped := datastructures.NewOrderedMap[string, []*dirwalk.Page]()
	for _, page := range pages {
		if len(page.Parts) > 0 {
			section := page.Parts[0]
			existing, has := grouped.Get(section)
			if !has {
				existing = []*dirwalk.Page{}
				grouped.Set(section, existing)
			}
			grouped.Set(section, append(existing, page))
		}
	}
	return grouped
}

func DocSidebar(pages []*dirwalk.Page) *h.Element {
	grouped := groupByFirstPart(pages)

	return h.Div(
		h.Class("px-3 py-2 pr-6 md:min-h-[(calc(100%))] md:min-h-screen bg-neutral-50 border-r border-r-slate-300"),
		h.Div(
			h.H4(h.Text("Contents"), h.Class("mt-4 text-slate-900 font-bold mb-3")),
			h.Div(
				h.Class("flex flex-col gap-4"),
				h.List(grouped.Entries(), func(entry datastructures.Entry[string, []*dirwalk.Page], index int) *h.Element {
					return h.Div(
						h.P(h.Text(formatPart(entry.Key)), h.Class("text-slate-800 font-bold")),
						h.Div(
							h.Class("pl-4 flex flex-col"),
							h.List(entry.Value, func(page *dirwalk.Page, index int) *h.Element {
								anchor := CreateAnchor(page.Parts)
								println(anchor)
								return h.A(
									h.Href("#"+anchor),
									h.Text(partsToName(page.Parts)),
									h.ClassX("text-slate-900", map[string]bool{}),
								)
							}),
						),
					)
				}),
			),
		),
	)
}
