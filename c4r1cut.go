package main

import (
	"flag"
	"fmt"
	"github.com/seqyuan/annogene/io/fastq"
	"log"
	"compress/gzip"
	"os"
	)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func usage() {
	fmt.Printf("\nProgram: cut fastq length\n")
	fmt.Printf("Command:\n")
	fmt.Printf("    -inFQ          in.faseq\n")
	fmt.Printf("    -c             cut region\n")
	fmt.Printf("    -o             outfile.fastq\n")
	os.Exit(1)
}

func main() {
	infq := flag.String("inFQ", "", "test.fastq")
	extractRegion := flag.String("c", "5:15,21:31,36:46", "you can input multi-regons split.by ,")
	outfile := flag.String("o", "", "outfile.fastq")
	flag.Parse()
	if *infq == "" || *outfile == "" {
		usage()
	}

	file, err := os.Open(*infq)
	check(err)

	gz, err := gzip.NewReader(file)
	check(err)

	defer file.Close()
	defer gz.Close()

	r := fastq.NewReader(gz)
	sc := fastq.NewScanner(r)

	fo, err := os.OpenFile(*outfile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0660)
	check(err)

	defer fo.Close()

	w := fastq.NewWriter(fo)

	for sc.Next() {
		//fmt.Println(string(sc.Seq().Id1))
		CuTreads := fastq.ExtractRegion(sc.Seq(), *extractRegion)
		//fmt.Println(string(CuTreads.Id1))
		_, eer := w.Write(CuTreads)
		check(eer)
		}
	if err := sc.Error(); err != nil {
		log.Fatalf("failed to read fastq: %v", err)
	}

}

