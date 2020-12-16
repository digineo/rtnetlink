package main

const (
	IFLA_UNSPEC = iota
	IFLA_ADDRESS
	IFLA_BROADCAST
	IFLA_IFNAME
	IFLA_MTU
	IFLA_LINK
	IFLA_QDISC
	IFLA_STATS
	IFLA_COST
	IFLA_PRIORITY
	IFLA_MASTER
	IFLA_WIRELESS
	IFLA_PROTINFO
	IFLA_TXQLEN
	IFLA_MAP
	IFLA_WEIGHT
	IFLA_OPERSTATE
	IFLA_LINKMODE
	IFLA_LINKINFO
	IFLA_NET_NS_PID
	IFLA_IFALIAS
	IFLA_NUM_VF
	IFLA_VFINFO_LIST
	IFLA_STATS64
	IFLA_VF_PORTS
	IFLA_PORT_SELF
	IFLA_AF_SPEC
	IFLA_GROUP
	IFLA_NET_NS_FD
	IFLA_EXT_MASK
	IFLA_PROMISCUITY
	IFLA_NUM_TX_QUEUES
	IFLA_NUM_RX_QUEUES
	IFLA_CARRIER
	IFLA_PHYS_PORT_ID
	IFLA_CARRIER_CHANGES
	IFLA_PHYS_SWITCH_ID
	IFLA_LINK_NETNSID
	IFLA_PHYS_PORT_NAME
)

const (
	// Subtype attributes for IFLA_PROTINFO
	IFLA_INET6_UNSPEC = iota
	IFLA_INET6_FLAGS
	IFLA_INET6_CONF
	IFLA_INET6_STATS
	IFLA_INET6_MCAST
	IFLA_INET6_CACHEINFO
	IFLA_INET6_ICMP6STATS
	IFLA_INET6_TOKEN
	IFLA_INET6_ADDR_GEN_MODE
)

const (
	// value for IFLA_INET6_ADDR_GEN_MODE
	IN6_ADDR_GEN_MODE_EUI64 = iota
	IN6_ADDR_GEN_MODE_NONE
	IN6_ADDR_GEN_MODE_STABLE_PRIVACY
)
