#
# Makefile
#

NAME=dapp
CC=riscv64-linux-gnu-gcc-12
CARTESI_MACHINE=/home/renan/.local/outuni/bin/cartesi-machine
# CARTESI_MACHINE=cartesi-machine

all: clean go pack snapshot

go:
	CC=$(CC) CGO_ENABLED=1 GOOS=linux GOARCH=riscv64 go build -o ./$(NAME)/$(NAME)
	chmod +x ./$(NAME)/$(NAME)

pack:
	tar \
		--sort=name \
		--mtime="2022-01-01" \
		--owner=1000 \
		--group=1000 \
		--numeric-owner \
		-cf $(NAME).tar \
		--directory=$(NAME) .
	genext2fs \
		-f \
		-b 16384 \
		-a $(NAME).tar \
		$(NAME).ext2

snapshot:
	$(CARTESI_MACHINE) \
		--flash-drive=label:$(NAME),filename:$(NAME).ext2 \
	    --assert-rolling-template \
		--store="snapshot" \
		-- /mnt/$(NAME)/$(NAME)

clean:
	rm -rf ./snapshot
	rm -rf ./$(NAME)
	rm -f $(NAME).*
