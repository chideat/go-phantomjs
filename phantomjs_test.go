package phantomjs

import "testing"

func TestNewPhantomJSWithExecutePath(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJSWithExecutePath("/usr/local/bin/phantomjs", 10000, options)
	if err != nil {
		t.Fatal(err)
	}
	err = phantomJS.Quit()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}

func TestNewPhantomJS(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(10000, options)
	if err != nil {
		t.Fatal(err)
	}
	err = phantomJS.Quit()
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewSession(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(10000, options)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		err := phantomJS.Quit()
		if err != nil {
			t.Fatal(err)
		}
	}()

	session, err := phantomJS.NewSession(nil)
	if err != nil {
		t.Fatal(err)
	}
	err = session.Close()
	if err != nil {
		t.Fatal(err)
	}
}
