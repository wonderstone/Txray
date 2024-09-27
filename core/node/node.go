package node

import (
	"Txray/core/protocols"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

type Node struct {
	protocols.Protocol `json:"-"`
	SubID              string  `json:"sub_id"`
	Data               string  `json:"data"`
	// TestResult         float64 `json:"-"`
	TestResult 		   float64 `json:"test_result"`
}

func (n *Node) TestResultStr() string {
	if n.TestResult == 0 {
		return ""
	} else if n.TestResult >= 99998 {
		return "-1ms"
	} else {
		return fmt.Sprintf("%.4vms", n.TestResult)
	}
}

// NewNode New一个节点
func NewNode(link, id string) *Node {
	if data := protocols.ParseLink(link); data != nil {
		return &Node{Protocol: data, SubID: id}
	}
	return nil
}

func NewNodeByData(protocol protocols.Protocol) *Node {
	return &Node{Protocol: protocol}
}

// ParseData 反序列化Data
func (n *Node) ParseData() {
	n.Protocol = protocols.ParseLink(n.Data)
}

// Serialize2Data 序列化数据-->Data
func (n *Node) Serialize2Data() {
	n.Data = n.GetLink()
}

var WG sync.WaitGroup

func (n *Node) Tcping() {
	count := 3
	var sum float64
	var timeout time.Duration = 3 * time.Second
	isTimeout := false
	for i := 0; i < count; i++ {
		start := time.Now()
		d := net.Dialer{Timeout: timeout}
		conn, err := d.Dial("tcp", fmt.Sprintf("%s:%d", n.GetAddr(), n.GetPort()))
		if err != nil {
			isTimeout = true
			break
		}
		conn.Close()
		elapsed := time.Since(start)
		sum += float64(elapsed.Nanoseconds()) / 1e6
	}
	if isTimeout {
		n.TestResult = 99999
	} else {
		n.TestResult = sum / float64(count)
	}
	WG.Done()
}

func (n *Node) Show() {
	ShowTopBottomSepLine('=', strings.Split(n.GetInfo(), "\n")...)
}
// ~ IsEqual 判断两个节点是否相等 
// todo Protocol 接口 Getlink()方法返回的是一个字符串，包含了节点的所有信息，包括地址、端口、协议等
// todo 通过Getlink()方法可以判断两个节点是否相等
// todo 但是这个方法可能不好，因为Getlink()方法返回的字符串可能会有很多不同的情况是由非关键信息导致的
// @ 目前的判断方式是通过判断节点的地址、端口、协议是否相等来判断两个节点是否相等
func (n *Node) IsEqual(node *Node) bool {
	return n.GetAddr() == node.GetAddr() && 
	n.GetPort() == node.GetPort() &&
	n.GetProtocolMode() == node.GetProtocolMode()
}