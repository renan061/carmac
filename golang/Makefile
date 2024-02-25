#
# Makefile
#

CARTESI_MACHINE=/home/renan/.local/outuni/bin/cartesi-machine
# CARTESI_MACHINE=cartesi-machine

all: clean go pack snapshot

go:
	CC=riscv64-linux-gnu-gcc-12 CGO_ENABLED=1 GOOS=linux GOARCH=riscv64 go build
	mkdir demo
	cp entry.sh ./demo
	cp  ./demo
	chmod +x ./demo/echo-go
	chmod +x ./demo/entry.sh

pack:
	tar \
		--sort=name \
		--mtime="2022-01-01" \
		--owner=1000 \
		--group=1000 \
		--numeric-owner \
		-cf demo.tar \
		--directory=demo .
	genext2fs \
		-f \
		-b 8192 \
		-a demo.tar \
		demo.ext2

snapshot:
	$(CARTESI_MACHINE) \
		--flash-drive=label:demo,filename:demo.ext2 \
	    --assert-rolling-template \
		--store="snapshot" \
		-- sh /mnt/demo/entry.sh

clean:
	rm -rf ./snapshot
	rm -rf ./demo
	rm -f demo.tar
	rm -f demo.ext2
	rm -f echo-go
