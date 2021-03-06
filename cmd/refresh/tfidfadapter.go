package main

import (
	"context"

	"github.com/negapedia/negapedia/internal/exporter"
	"github.com/negapedia/wikitfidf"
	"github.com/pkg/errors"
)

func TFIDFExporter(ctx context.Context, fail func(error) error, tfidf wikitfidf.Exporter) (outs []<-chan exporter.ExtData) {
	if tfidf.Lang == "" { //invalid, TFIDF data is optional
		return
	}

	//Words Occurrences and TF-IDF
	{
		globalWords, err := tfidf.GlobalWords()
		if err != nil {
			fail(errors.Wrap(err, "Error while Retrieving GlobalWords"))
			outs = nil
			return
		}
		out := make(chan exporter.ExtData, 1)
		outs = append(outs, out)
		out <- exporter.ExtData{0, map[string]interface{}{"Word2Occur": globalWords.Words2Occur}} //If you add fields, remember to export them in data.html template
		close(out)
	}

	{
		in := tfidf.Topics(ctx, fail)
		out := make(chan exporter.ExtData, 1)
		outs = append(outs, out)
		go func() {
			defer close(out)
			for {
				select {
				case <-ctx.Done():
					return
				case p, ok := <-in:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case out <- exporter.ExtData{p.TopicID, map[string]interface{}{"Word2Occur": p.Words}}: //If you add fields, remember to export them in data.html template
						//Go on
					}
				}
			}
		}()
	}

	{
		in := tfidf.Pages(ctx, fail)
		out := make(chan exporter.ExtData, 1)
		outs = append(outs, out)
		go func() {
			defer close(out)
			for {
				select {
				case <-ctx.Done():
					return
				case p, ok := <-in:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case out <- exporter.ExtData{p.ID, map[string]interface{}{"Word2TFIDF": p.Word2TFIDF}}: //If you add fields, remember to export them in data.html template
						//Go on
					}
				}
			}
		}()
	}

	//BadWords Occurrences
	{
		in := tfidf.PageBadwords(ctx, fail)
		out := make(chan exporter.ExtData, 1)
		outs = append(outs, out)
		go func() {
			defer close(out)
			for {
				select {
				case <-ctx.Done():
					return
				case p, ok := <-in:
					if !ok {
						return
					}
					select {
					case <-ctx.Done():
						return
					case out <- exporter.ExtData{p.PageID, map[string]interface{}{"BWord2Occur": p.BadW}}: //If you add fields, remember to export them in data.html template
						//Go on
					}
				}
			}
		}()
	}

	return
}
