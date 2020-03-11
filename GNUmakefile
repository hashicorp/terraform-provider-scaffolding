default: testacc

# Run acceptance tests
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
