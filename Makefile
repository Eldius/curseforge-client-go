
test:
	go test ./... -cover


release-build:
	$(MAKE) release-version bump=build

release-minor:
	$(MAKE) release-version bump=minor

release-major:
	$(MAKE) release-version bump=major

release-version:
	$(eval v := $(shell git describe --tags --abbrev=0 | sed -Ee 's/^v|-.*//'))
ifeq ($(bump), major)
	$(eval f := 1)
else ifeq ($(bump), minor)
	$(eval f := 2)
else
	$(eval f := 3)
endif
	$(eval next_version := $(shell echo $(v) | awk -F. -v OFS=. -v f=$(f) '{ $$f++ } 1'))
	@echo "last version: v$(v)"
	@echo "next version: v$(next_version)"
	git tag "v$(next_version)"
	git push origin "v$(next_version)"

search-create:
	go run ./cmd/cli mod search --term create --game-id 432 --debug
