module github.com/regen-network/regen-ledger/api/v2

go 1.19

require (
	github.com/cosmos/cosmos-proto v1.0.0-beta.3
	github.com/cosmos/cosmos-sdk/api v0.1.0
	github.com/cosmos/cosmos-sdk/orm v1.0.0-alpha.12
	github.com/cosmos/gogoproto v1.4.10
	google.golang.org/genproto v0.0.0-20230410155749-daa745c078e1
	google.golang.org/grpc v1.56.2
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/DataDog/zstd v1.4.5 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cosmos/cosmos-sdk/errors v1.0.0-beta.5 // indirect
	github.com/cosmos/gorocksdb v1.2.0 // indirect
	github.com/dgraph-io/badger/v2 v2.2007.2 // indirect
	github.com/dgraph-io/ristretto v0.0.3 // indirect
	github.com/dgryski/go-farm v0.0.0-20200201041132-a6ae2369ad13 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/golang/snappy v0.0.3-0.20201103224600-674baa8c7fc3 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20200815110645-5c35d600f0ca // indirect
	github.com/tendermint/tm-db v0.6.7 // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	golang.org/x/exp v0.0.0-20230131160201-f062dba9d201 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
)

// https://github.com/regen-network/regen-ledger/security/dependabot/105
replace golang.org/x/net => golang.org/x/net v0.7.0
