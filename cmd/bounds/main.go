// bounds will print the bounding box (minx,miny,maxx,maxy) for one or more URIs describing GeoJSON files.
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/paulmach/orb/geojson"
	"io"
	"log"
	"os"
)

func FeatureFromURI(ctx context.Context, uri string) (*geojson.Feature, error) {

	fh, err := os.Open(uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to open URI '%s', %v", uri, err)
	}

	defer fh.Close()

	return FeatureFromReader(ctx, fh)
}

func FeatureFromReader(ctx context.Context, r io.Reader) (*geojson.Feature, error) {

	body, err := io.ReadAll(r)

	if err != nil {
		return nil, fmt.Errorf("Failed to read feature, %v", err)
	}

	return FeatureFromBytes(ctx, body)
}

func FeatureFromBytes(ctx context.Context, body []byte) (*geojson.Feature, error) {
	return geojson.UnmarshalFeature(body)
}

func main() {

	latlon := flag.Bool("latlon", false, "Print bounding box as miny,minx,maxy,maxx.")

	flag.Parse()

	uris := flag.Args()

	ctx := context.Background()

	wr := os.Stdout

	for _, path := range uris {

		f, err := FeatureFromURI(ctx, path)

		if err != nil {
			log.Fatalf("Failed to derive feature from URI, %v", err)
		}

		geom := f.Geometry
		bounds := geom.Bound()

		var coords []interface{}

		if *latlon {

			coords = []interface{}{
				bounds.Min.Y(),
				bounds.Min.X(),
				bounds.Max.Y(),
				bounds.Max.X(),
			}
		} else {

			coords = []interface{}{
				bounds.Min.X(),
				bounds.Min.Y(),
				bounds.Max.X(),
				bounds.Max.Y(),
			}
		}

		fmt.Fprintf(wr, "%f,%f,%f,%f\n", coords...)
	}
}
