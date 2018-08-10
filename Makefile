all:
	cd hexer && go build
	cd shifter && go build
	hexer/hexer input.txt | shifter/shifter 33 | perl -ne 'print scalar reverse'
	@echo