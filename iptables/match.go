package iptables

import "fmt"

type matchType int

const (
	matchAddrType matchType = iota
	matchAH
	matchBPF
	matchCgroup
	matchCluster
	matchComment
	matchConnBytes
	matchConnLabel
	matchConnLimit
	matchConnMark
	matchConnTrack
	matchCPU
	matchDCCP
	matchDestination
	matchDevGroup
	matchDSCP
	matchDst
	matchECN
	matchESP
	matchEUI64
	matchFrag
	matchHashLimit
	matchHBH // Hop-by-Hop
	matchHelper
	matchHL // Hop Limit
	matchIcmp
	matchIcmp6
	matchIPRange
	matchIPv4
	matchIPv6
	matchIPv6Header
	matchIPVS
	matchLength
	matchLimit
	matchMAC
	matchMark
	matchMH
	matchMultiPort
	matchNFacct
	matchOSF
	matchOwner
	matchPhysDev
	matchPktType
	matchPolicy
	matchProtocol
	matchQuota
	matchRateEst
	matchRealm
	matchRecent
	matchRPFilter
	matchRT
	matchSCTP
	matchSet
	matchSocket
	matchSource
	matchState
	matchStatistic
	matchString
	matchTcp
)

type match interface {
	typ() matchType
	short() string
	long() string
}

type baseMatch struct {
	matchType matchType
	invert    bool
}

func (bm baseMatch) typ() matchType {
	return bm.matchType
}

func (bm baseMatch) short() string {
	return ""
}

func (bm baseMatch) long() string {
	return ""
}

type MatchIPv4 struct {
	baseMatch
}

func (mIPv4 *MatchIPv4) short() string {
	return "-4"
}

func (mIPv4 *MatchIPv4) long() string {
	return "--ipv4"
}

type MatchIPv6 struct {
	baseMatch
}

func (mIPv6 *MatchIPv6) short() string {
	return "-6"
}

func (mIPv6 *MatchIPv6) long() string {
	return "--ipv6"
}

type MatchProtocol struct {
	baseMatch
	protocol Protocol
}

func (mProtocol *MatchProtocol) short() string {
	if mProtocol.invert {
		return fmt.Sprintf("! -p %d", int(mProtocol.protocol))
	}
	return fmt.Sprintf("-p %d", int(mProtocol.protocol))
}

func (mProtocol *MatchProtocol) long() string {
	if mProtocol.invert {
		return fmt.Sprintf("! --protocol %d", int(mProtocol.protocol))
	}
	return fmt.Sprintf("--protocol %d", int(mProtocol.protocol))
}

type MatchSource struct {
	baseMatch
	ads *Address
}

func (mSrc *MatchSource) short() string {
	if mSrc.invert {
		return fmt.Sprintf("! -s %s", mSrc.ads.String())
	}
	return fmt.Sprintf("-s %s", mSrc.ads.String())
}

func (mSrc *MatchSource) long() string {
	if mSrc.invert {
		return fmt.Sprintf("! --source %s", mSrc.ads.String())
	}
	return fmt.Sprintf("--source %s", mSrc.ads.String())
}

type MatchDestination struct {
	baseMatch
	ads *Address
}

func (mDst *MatchDestination) short() string {
	if mDst.invert {
		return fmt.Sprintf("! -d %s", mDst.ads.String())
	}
	return fmt.Sprintf("-d %s", mDst.ads.String())
}

func (mDst *MatchDestination) long() string {
	if mDst.invert {
		return fmt.Sprintf("! --destination %s", mDst.ads.String())
	}
	return fmt.Sprintf("--destination %s", mDst.ads.String())
}
