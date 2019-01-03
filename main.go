package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"golang.org/x/sync/errgroup"

	"github.com/olivere/bodleiandiff/diff"
	"github.com/olivere/bodleiandiff/diff/printer"
	"github.com/olivere/bodleiandiff/elastic"
	"github.com/olivere/bodleiandiff/elastic/config"
	"github.com/olivere/bodleiandiff/elastic/v6"
)

func main() {
	var (
		outputFormat = flag.String("o", "", "Output format, e.g. json")
		size         = flag.Int("size", 100, "Batch size")
		sortField    = flag.String("sort", "file.path", `Sort field, e.g. "_id", "file.path" or "-price" (descending)`)
		rawSrcQuery  = flag.String("src-query", "", `Raw query for filtering the source, e.g. {"term":{"user":"olivere"}}`)
		rawDstQuery  = flag.String("dst-query", "", `Raw query for filtering the destination, e.g. {"term":{"user":"olivere"}}`)
		unchanged    = flag.Bool("u", false, `Print unchanged docs`)
		updated      = flag.Bool("c", true, `Print changed docs`)
		changed      = flag.Bool("a", true, `Print added docs`)
		deleted      = flag.Bool("d", true, `Print deleted docs`)
	)

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 2 {
		usage()
		os.Exit(1)
	}

	options := []elastic.ClientOption{
		elastic.WithBatchSize(*size),
	}

	src, err := newClient(flag.Arg(0), options...)
	if err != nil {
		log.Fatal(err)
	}
	srcIterReq := &elastic.IterateRequest{
		RawQuery:  *rawSrcQuery,
		SortField: *sortField,
	}

	dst, err := newClient(flag.Arg(1), options...)
	if err != nil {
		log.Fatal(err)
	}
	dstIterReq := &elastic.IterateRequest{
		RawQuery:  *rawDstQuery,
		SortField: *sortField,
	}

	var p printer.Printer
	{
		switch *outputFormat {
		default:
			p = printer.NewStdPrinter(os.Stdout, *unchanged, *updated, *changed, *deleted)
		case "json":
			p = printer.NewJSONPrinter(os.Stdout, *unchanged, *updated, *changed, *deleted)
		}
	}

	g, ctx := errgroup.WithContext(context.Background())
	srcDocCh, srcErrCh := src.Iterate(ctx, srcIterReq)
	dstDocCh, dstErrCh := dst.Iterate(ctx, dstIterReq)
	diffCh, errCh := diff.Differ(ctx, srcDocCh, dstDocCh)
	g.Go(func() error {
		for {
			select {
			case d, ok := <-diffCh:
				if !ok {
					return nil
				}
				if err := p.Print(d); err != nil {
					return err
				}
			case err := <-srcErrCh:
				return err
			case err := <-dstErrCh:
				return err
			case err := <-errCh:
				return err
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	})
	if err = g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "General usage:\n\n")
	fmt.Fprintf(os.Stderr, "\t%s [flags] <source-url> <destination-url>\n\n", path.Base(os.Args[0]))
	fmt.Fprintf(os.Stderr, "General flags:\n")
	flag.PrintDefaults()
}

// newClient will create a new Elasticsearch client,
// matching the supported version.
func newClient(url string, opts ...elastic.ClientOption) (elastic.Client, error) {
	cfg, err := config.Parse(url)
	if err != nil {
		return nil, err
	}
	c, err := v6.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}
