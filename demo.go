package main

import (
	"bytes"
	"encoding/binary"
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

func main() {
	const u = "https://bigtable.googleapis.com/google.bigtable.v1.BigtableService/MutateRow"

	family := "fam"
	column := "col"
	ts := time.Now().Unix() / 1e3
	value := []byte("Blake says hi.")

	var ops []*btdpb.Mutation
	ops = append(ops, &btdpb.Mutation{Mutation: &btdpb.Mutation_SetCell_{&btdpb.Mutation_SetCell{
		FamilyName:      family,
		ColumnQualifier: []byte(column),
		TimestampMicros: int64(ts),
		Value:           value,
	}}})

	mr := &btspb.MutateRowRequest{
		TableName: "projects/grpc-go16-dev/zones/us-central1-c/clusters/dollar-95-per-hour/tables/mytable",
		RowKey:    []byte(fmt.Sprintf("t-%d", time.Now().Unix())),
		Mutations: ops,
	}

	pay, err := proto.Marshal(mr)
	if err != nil {
		log.Fatal("Marshal:", err)
	}

	log.Printf("payload: %q", pay)

	body := make([]byte, 5)
	binary.BigEndian.PutUint32(body[1:5], uint32(len(pay)))
	body = append(body, pay...)
	req, err := http.NewRequest("POST", u, bytes.NewReader(body))
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
		log.Fatalf("non-200 status code: %q", res.Status)
	}
	if _, err := io.Copy(debugWriter{os.Stdout}, res.Body); err != nil {
		log.Fatal("Copy:", err)
	}
	log.Printf("%#v", res.Trailer)
}

type debugWriter struct {
	w io.Writer
}

func (w debugWriter) Write(b []byte) (int, error) {
	_, err := fmt.Fprintf(w.w, "BODY: %q\n", b)
	return len(b), err
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
