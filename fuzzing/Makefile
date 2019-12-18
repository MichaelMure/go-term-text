fuzz:
	mkdir -p corpus
	echo "ADJUST THE FOLLOWING PATH!"
	cp ../../go-term-markdown/testdata_result/* corpus/
	go run github.com/dvyukov/go-fuzz/go-fuzz-build
	go run github.com/dvyukov/go-fuzz/go-fuzz