# LANG = go
GO_SRC_FILE := main.go

GO_DIR := $(ROOT_DIR)/go
GO_BIN := $(OUT_DIR)/out_go
GO_SRC := $(GO_DIR)/$(GO_SRC_FILE)

run_go: $(GO_BIN)
	$^

$(GO_BIN): $(GO_SRC)
	go build -o $@ $^

run_go_noopt: $(GO_BIN)_noopt
	$^

$(GO_BIN)_noopt: $(GO_SRC)
	go build -gcflags="-N -l" -o $@ $^

$(GO_BIN)_asm: $(GO_SRC)
	go tool compile -S $^ > $@.a

ALL_BINS += $(GO_BIN)
ALL_BINS += $(GO_BIN)_noopt
