all:
	cd hexer && go build
	cd shifter && go build
	hexer/hexer input.txt | shifter/shifter 31 | perl -ne 'print scalar reverse'
	@echo
	@echo
	hexer/hexer input.txt | shifter/shifter 31 | perl -ne 'print scalar reverse' | ./prog 31
	@echo

fmt:
	go fmt hexer/*.go
	go fmt shifter/*.go
