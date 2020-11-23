

.PHONY: clean
clean:
	chmod +w -R $$GOPATH/ && rm -rf $$GOPATH/*
