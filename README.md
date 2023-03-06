# studying_go

## 목적 : Golang을 좀 더 친숙하게 공부하고 활용하기 위해 만든 Repository

## 알고리즘, 문법, 코어기술에 관련된 코드를 이 repository에 모아 둘 생각임

## 디렉토리 구조 : 

.
├── algorithm
├── baek_joon
├── blockchain
│   ├── __pycache__
│   ├── go-libp2p
│   │   ├── config
│   │   ├── core
│   │   │   ├── canonicallog
│   │   │   ├── connmgr
│   │   │   ├── control
│   │   │   ├── crypto
│   │   │   │   ├── pb
│   │   │   │   └── test_data
│   │   │   ├── discovery
│   │   │   ├── event
│   │   │   ├── host
│   │   │   ├── internal
│   │   │   │   └── catch
│   │   │   ├── metrics
│   │   │   ├── network
│   │   │   │   └── mocks
│   │   │   ├── peer
│   │   │   │   └── pb
│   │   │   ├── peerstore
│   │   │   ├── pnet
│   │   │   ├── protocol
│   │   │   ├── record
│   │   │   │   └── pb
│   │   │   ├── routing
│   │   │   ├── sec
│   │   │   │   └── insecure
│   │   │   │       └── pb
│   │   │   ├── test
│   │   │   └── transport
│   │   ├── dashboards
│   │   │   ├── autonat
│   │   │   ├── eventbus
│   │   │   ├── resource-manager
│   │   │   └── swarm
│   │   ├── examples
│   │   │   ├── chat
│   │   │   ├── chat-with-mdns
│   │   │   ├── chat-with-rendezvous
│   │   │   ├── echo
│   │   │   ├── http-proxy
│   │   │   ├── ipfs-camp-2019
│   │   │   │   ├── 01-Transports
│   │   │   │   ├── 02-Multiaddrs
│   │   │   │   ├── 03-Muxing-Encryption
│   │   │   │   ├── 05-Discovery
│   │   │   │   ├── 06-Pubsub
│   │   │   │   ├── 07-Messaging
│   │   │   │   └── 08-End
│   │   │   ├── libp2p-host
│   │   │   ├── multipro
│   │   │   │   └── pb
│   │   │   ├── p2p_blockchain
│   │   │   ├── pubsub
│   │   │   │   ├── basic-chat-with-rendezvous
│   │   │   │   └── chat
│   │   │   ├── relay
│   │   │   ├── routed-echo
│   │   │   └── testutils
│   │   ├── p2p
│   │   │   ├── discovery
│   │   │   │   ├── backoff
│   │   │   │   ├── mdns
│   │   │   │   ├── mocks
│   │   │   │   ├── routing
│   │   │   │   └── util
│   │   │   ├── host
│   │   │   │   ├── autonat
│   │   │   │   │   ├── pb
│   │   │   │   │   └── test
│   │   │   │   ├── autorelay
│   │   │   │   ├── basic
│   │   │   │   ├── blank
│   │   │   │   ├── eventbus
│   │   │   │   ├── peerstore
│   │   │   │   │   ├── pstoreds
│   │   │   │   │   │   └── pb
│   │   │   │   │   ├── pstoremem
│   │   │   │   │   └── test
│   │   │   │   ├── pstoremanager
│   │   │   │   ├── relaysvc
│   │   │   │   ├── resource-manager
│   │   │   │   │   ├── docs
│   │   │   │   │   └── obs
│   │   │   │   └── routed
│   │   │   ├── metricshelper
│   │   │   ├── muxer
│   │   │   │   ├── mplex
│   │   │   │   ├── testsuite
│   │   │   │   └── yamux
│   │   │   ├── net
│   │   │   │   ├── conngater
│   │   │   │   ├── connmgr
│   │   │   │   ├── mock
│   │   │   │   ├── nat
│   │   │   │   ├── pnet
│   │   │   │   ├── reuseport
│   │   │   │   ├── swarm
│   │   │   │   │   └── testing
│   │   │   │   └── upgrader
│   │   │   ├── protocol
│   │   │   │   ├── circuitv2
│   │   │   │   │   ├── client
│   │   │   │   │   ├── pb
│   │   │   │   │   ├── proto
│   │   │   │   │   ├── relay
│   │   │   │   │   └── util
│   │   │   │   ├── holepunch
│   │   │   │   │   └── pb
│   │   │   │   ├── identify
│   │   │   │   │   ├── grafana-dashboards
│   │   │   │   │   └── pb
│   │   │   │   └── ping
│   │   │   ├── security
│   │   │   │   ├── noise
│   │   │   │   │   └── pb
│   │   │   │   └── tls
│   │   │   │       └── cmd
│   │   │   │           └── tlsdiag
│   │   │   ├── test
│   │   │   │   ├── backpressure
│   │   │   │   ├── negotiation
│   │   │   │   ├── notifications
│   │   │   │   ├── quic
│   │   │   │   ├── reconnects
│   │   │   │   ├── resource-manager
│   │   │   │   └── webtransport
│   │   │   └── transport
│   │   │       ├── quic
│   │   │       │   └── cmd
│   │   │       │       ├── client
│   │   │       │       └── server
│   │   │       ├── quicreuse
│   │   │       ├── tcp
│   │   │       ├── testsuite
│   │   │       ├── websocket
│   │   │       └── webtransport
│   │   ├── scripts
│   │   └── test-plans
│   │       └── cmd
│   │           └── ping
│   ├── go-spew
│   │   └── spew
│   │       └── testdata
│   ├── godotenv
│   │   ├── autoload
│   │   ├── cmd
│   │   │   └── godotenv
│   │   └── fixtures
│   ├── mux
│   ├── pyrlp
│   │   ├── docs
│   │   ├── rlp
│   │   │   ├── __pycache__
│   │   │   └── sedes
│   │   │       └── __pycache__
│   │   └── tests
│   └── web3_api
├── crypto
├── geth_study
│   └── p2p
│       ├── v4wire
│       └── v5wire
├── grammer
│   ├── array
│   ├── bufio
│   ├── channel
│   ├── context
│   ├── defer
│   ├── etc
│   ├── flag
│   ├── fmt
│   ├── goRoutine
│   ├── log
│   ├── marshal
│   ├── mutex
│   ├── operation
│   ├── panic
│   ├── pointer
│   ├── rand
│   ├── slice
│   ├── socket
│   ├── structure
│   └── test
├── programmers
└── reverse_engineering
    └── reverse_study.rep
        ├── idata
        │   └── 00
        │       └── ~00000000.db
        ├── user
        │   └── 00
        │       └── ~00000000.db
        └── versioned

198 directories

