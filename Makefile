.ONESHELL:
testacc:
	rm -rf work
	mkdir work && cd work
	git clone https://github.com/jerakia/go-jerakia
	cd go-jerakia
	make docker
	export JERAKIA_TOKEN=$$(docker exec jerakia-server jerakia token create terraform --quiet)
	export JERAKIA_URL="http://localhost:9843/v1"
	cd ../..
	TF_LOG=DEBUG TF_ACC=1 go test -v -run="$(TESTARGS)" ./jerakia/

build:
	go install

fmtcheck:
	echo "==> Checking that code complies with gofmt requirements..."
	files=$$(find . -name '*.go' | grep -v 'vendor' ) ; \
	gofmt_files=`gofmt -l $$files`; \
	if [ -n "$$gofmt_files" ]; then \
		echo 'gofmt needs running on the following files:'; \
		echo "$$gofmt_files"; \
		echo "You can use the command: \`make fmt\` to reformat code."; \
		exit 1; \
	fi

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

targets: $(TARGETS)

$(TARGETS):
	GOOS=$@ GOARCH=amd64 go build -o "dist/$@/terraform-provider-jerakia_${TRAVIS_TAG}_x4"
	zip -j dist/terraform-provider-jerakia_${TRAVIS_TAG}_$@_amd64.zip dist/$@/terraform-provider-jerakia_${TRAVIS_TAG}_x4
