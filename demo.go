package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang/protobuf/proto"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	btdpb "github.com/bradfitz/grpc-go16-demo/proto/data_proto"
	btspb "github.com/bradfitz/grpc-go16-demo/proto/service_proto"
)

var (
	proj    = flag.String("project", "grpc-go16-dev", "GCE project")
	zone    = flag.String("zone", "us-central1-c", "GCE zone")
	cluster = flag.String("cluster", "dollar-95-per-hour", "Cloud Bigtable Cluster name")
	table   = flag.String("table", "mytable", "table name")
)

func main() {
	flag.Parse()
	const u = "https://bigtable.googleapis.com/google.bigtable.v1.BigtableService/MutateRow"

	mr := &btspb.MutateRowRequest{
		TableName: "projects/" + *proj + "/zones/" + *zone + "/clusters/" + *cluster + "/tables/" + *table,
		RowKey:    []byte(fmt.Sprintf("t-%d", time.Now().Unix())),
		Mutations: []*btdpb.Mutation{
			{Mutation: &btdpb.Mutation_SetCell_{&btdpb.Mutation_SetCell{
				FamilyName:      "fam",
				ColumnQualifier: []byte("col"),
				TimestampMicros: -1,
				Value:           []byte("blake says hi"),
			}}},
		},
	}

	payload, err := proto.Marshal(mr)
	if err != nil {
		log.Fatal("Marshal:", err)
	}

	header := make([]byte, 5)
	header[0] = 0 // uncompressed
	binary.BigEndian.PutUint32(header[1:5], uint32(len(payload)))

	req, err := http.NewRequest("POST", u, io.MultiReader(
		bytes.NewReader(header),
		bytes.NewReader(payload),
	))
	if err != nil {
		log.Fatal("NewRequest:", err)
	}
	req.Header.Set("Content-Type", "application/grpc")
	req.Header.Set("Authorization", "Bearer "+getToken())
	req.Header.Set("te", "trailers")
	req.Header.Set("User-Agent", "cbt-go/20150727 grpc-go/0.11")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Do:", err)
	}
	defer res.Body.Close()
	log.Printf("Res: %#v", res)
	if res.StatusCode != 200 {
		log.Fatalf("non-200 status code: %v", res.Status)
	}

	if err := foreachDelimitedMessage(res.Body, func(v []byte) error {
		log.Printf("delimited message of %d bytes: %q", len(v), v)
		return nil
	}); err != nil {
		log.Fatal(err)
	}

	// In the trailers by default, unless the server replies with
	// the grpc-status in the headers with no body.  (No observed,
	// but described in the grpc wire docs)
	statusHeader := res.Trailer
	if _, ok := statusHeader["Grpc-Status"]; !ok {
		statusHeader = res.Header
	}

	statusCode, statusMessage := statusHeader.Get("Grpc-Status"), statusHeader.Get("Grpc-Message")
	if statusCode != "0" {
		log.Fatalf("GPRC status %v: %s", statusCode, statusMessage)
	}
}

func foreachDelimitedMessage(rc io.ReadCloser, fn func([]byte) error) error {
	defer rc.Close()
	br := bufio.NewReader(rc) // TODO: recycle
	for {
		t, err := br.ReadByte()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if t != 0 {
			return errors.New("unsupported message compression type")
		}
		size, err := br.Peek(4)
		if err != nil {
			if err == io.EOF {
				err = io.ErrUnexpectedEOF
			}
			return err
		}
		br.Discard(4) // always succeeds
		usize := binary.BigEndian.Uint32(size)
		if int64(int(usize)) != int64(usize) {
			return errors.New("delimited message too large")
		}
		v, err := br.Peek(int(usize))
		if err != nil {
			if err == io.EOF {
				err = io.ErrUnexpectedEOF
			}
			return err
		}
		if err := fn(v); err != nil {
			return err
		}
	}
}

func getToken() string {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "grpc-go16-dev-608d72c9bce0.json")
	const scope = "https://www.googleapis.com/auth/bigtable.data"
	ts, err := google.DefaultTokenSource(context.Background(), scope)
	if err != nil {
		log.Fatal("DefaultTokenSource:", err)
	}
	tk, err := ts.Token()
	if err != nil {
		log.Fatal("Token:", err)
	}
	return tk.AccessToken

}
