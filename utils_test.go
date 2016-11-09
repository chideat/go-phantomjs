package phantomjs

import "testing"

func TestRouteTemplateWithNotTemplate(t *testing.T) {
	route := "/session"

	val, err := RouteTemplate(route, nil)
	if err != nil {
		t.Fatal(err)
	}
	if val != route {
		t.Fatal("route changed")
	}
}

func TestRouteTemplateWithNotArg(t *testing.T) {
	route := "/session/:sessionId"

	_, err := RouteTemplate(route, nil)
	if err == nil {
		t.Fatal("missing args")
	}
}

func TestRouteTemplate(t *testing.T) {
	route := "/session/:sessionId"
	args := map[string]interface{}{
		"sessionId": "54321",
	}

	val, err := RouteTemplate(route, args)
	if err != nil {
		t.Fatal(err)
	}
	if val != "/session/54321" {
		t.Fatal("invalid route")
	}
}

func TestIsConnectAble(t *testing.T) {
	// start service at port 10000
	if !IsConnectAble(":10000") {
		t.Error("not connect able")
	}
}
