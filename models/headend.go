package models

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type Headend struct {
	Addr string `json:"address"`
	Comm string `json:"community"`
}

func (h *Headend) GetModemOid(mac string, ch chan<- string) {
	fmt.Printf("%s <- %s\n", h.Addr, mac)
	time.Sleep(time.Millisecond * 100)

	index := time.Now().Nanosecond()
	oid, _ := h.macToDecimal(mac)

	ch <- fmt.Sprintf("%s -> %s.%d\n", h.Addr, oid, index)
}

func (h *Headend) macToDecimal(mac string) (string, error) {
	var macStr []string

	hwAddr, err := net.ParseMAC(mac)
	if err != nil {
		return "", err
	}

	for _, i := range hwAddr {
		s := strconv.Itoa(int(i))
		macStr = append(macStr, s)
	}

	return strings.Join(macStr, "."), nil
}
