#!/bin/sh

OUTPUT_FILE=benchmark.txt

cp -f /dev/null $OUTPUT_FILE
go test -bench=. -benchmem >>$OUTPUT_FILE
echo "Benchmark was completed."
