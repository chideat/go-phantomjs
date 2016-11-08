package phantomjs

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
)

var (
	route_param_reg *regexp.Regexp
	ip_addr_reg     *regexp.Regexp
)

func init() {
	route_param_reg = regexp.MustCompile(":[0-9a-zA-Z_]+")
	ip_addr_reg = regexp.MustCompile("^(.*):([0-9]+)$")
}

func RouteTemplate(route string, params map[string]string) (string, error) {
	paramKeys := route_param_reg.FindAllString(route, -1)

	for _, paramKey := range paramKeys {
		key := strings.TrimLeft(paramKey, ":")
		val, ok := params[key]
		if !ok {
			return "", fmt.Errorf("need param %s", key)
		}
		route = strings.Replace(route, paramKey, val, -1)
	}
	return route, nil
}

func FindFreePort() int {
	nl, err := net.Listen("TCP", ":")
	if err != nil {
		panic(err)
	}
	defer nl.Close()

	addrParts := ip_addr_reg.FindStringSubmatch(nl.Addr().String())
	if len(addrParts) != 2 {
		panic("invalid addr")
	}
	port, _ := strconv.Atoi(addrParts[1])

	return port
}
