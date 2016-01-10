package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/cloud/bigtable"
)

var (
	create = flag.Bool("create", false, "create table")
	insert = flag.Bool("insert", true, "insert a row")
)

func main() {
	flag.Parse()
	ctx := context.Background()
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "grpc-go16-dev-608d72c9bce0.json")
	const tableName = "mytable"
	if *create {
		ac, err := bigtable.NewAdminClient(ctx, "grpc-go16-dev", "us-central1-c", "dollar-95-per-hour")
		if err != nil {
			log.Fatalf("AdminClient: %v", err)
		}
		if err := ac.CreateTable(ctx, tableName); err != nil {
			log.Printf("Create: %v", err)
		}
		if err := ac.CreateColumnFamily(ctx, tableName, "fam"); err != nil {
			log.Printf("CreateColumnFamily: %v", err)
		}
	}

	bc, err := bigtable.NewClient(ctx, "grpc-go16-dev", "us-central1-c", "dollar-95-per-hour")
	if err != nil {
		log.Fatalf("NewClient: %v", err)
	}
	t := bc.Open(tableName)
	m := bigtable.NewMutation()
	m.Set("fam", "col", bigtable.Now(), []byte("hello"))

	if *insert {
		log.Printf("Apply...")
		err = t.Apply(ctx, fmt.Sprintf("t:%v", time.Now().Unix()), m)
		if err != nil {
			log.Fatalf("Apply: %v", err)
		}
	}

	log.Printf("ReadRows...")

	t.ReadRows(ctx, bigtable.InfiniteRange(""), func(row bigtable.Row) bool {
		log.Printf("Row: %v", row.Key())
		for fam, ris := range row {
			log.Printf("  col fam %q", fam)
			for _, ri := range ris {
				log.Printf("    row %q, col %q, ts %v, val %q", ri.Row, ri.Column, ri.Timestamp, ri.Value)
			}
		}
		return true
	})
	log.Printf("done.")
}
