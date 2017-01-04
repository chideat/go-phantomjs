package phantomjs

import (
	"sync"
	"testing"
)

func TestGet(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	defer phantomJS.Quit()

	session, err := phantomJS.NewSession(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	err = session.Get("https://www.baidu.com")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRefresh(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	defer phantomJS.Quit()

	session, err := phantomJS.NewSession(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	err = session.Get("https://www.baidu.com")
	if err != nil {
		t.Fatal(err)
	}

	err = session.Refresh()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetCurrentUrl(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	defer phantomJS.Quit()

	session, err := phantomJS.NewSession(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	urlStr := "https://www.baidu.com/"
	err = session.Get(urlStr)
	if err != nil {
		t.Fatal(err)
	}

	currentUrlStr, err := session.GetCurrentUrl()
	if err != nil {
		t.Fatal(err)
	}
	if currentUrlStr != urlStr {
		t.Fatalf("current url %s not equal %s", currentUrlStr, urlStr)
	}
}

func TestGetTitle(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	defer phantomJS.Quit()

	session, err := phantomJS.NewSession(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	urlStr := "https://www.baidu.com/"
	err = session.Get(urlStr)
	if err != nil {
		t.Fatal(err)
	}

	title, err := session.GetTitle()
	if err != nil {
		t.Fatal(err)
	}
	if title != "百度一下，你就知道" {
		t.Fatalf("the title '%s' is not as excepted", title)
	}
}

func TestGetPageSource(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	defer phantomJS.Quit()

	session, err := phantomJS.NewSession(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	urlStr := "https://www.baidu.com/"
	err = session.Get(urlStr)
	if err != nil {
		t.Fatal(err)
	}

	source, err := session.GetPageSource()
	if err != nil {
		t.Fatal(err)
	}
	if source == "" {
		t.Fatal("get empty webpage source")
	}
}

func TestGetPageSources(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	defer phantomJS.Quit()

	urls := []string{
		"http://www.25pp.com/ios/wallpaper/64/13/",
		"http://www.25pp.com/ios/wallpaper/64/22/",
		"http://www.25pp.com/ios/wallpaper/64/19/",
		"http://www.25pp.com/ios/wallpaper/64/15/",
		"http://www.25pp.com/ios/wallpaper/64/12/",
		"http://www.25pp.com/ios/wallpaper/64/17/",
		"http://www.25pp.com/ios/wallpaper/64/20/",
		"http://www.25pp.com/ios/wallpaper/64/18/",
		"http://www.25pp.com/ios/wallpaper/64/16/",
		"http://www.25pp.com/ios/wallpaper/64/21/",
	}

	for _, urlStr := range urls {
		session, err := phantomJS.NewSession(nil)
		if err != nil {
			t.Fatal(err)
		}
		defer session.Close()

		err = session.Get(urlStr)
		if err != nil {
			t.Fatal(err)
		}

		source, err := session.GetPageSource()
		if err != nil {
			t.Fatal(err)
		}
		if source == "" {
			t.Fatal("get empty webpage source")
		}
	}
}

func TestGetPageSourcesAsync(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	defer phantomJS.Quit()

	urls := []string{
		"http://www.25pp.com/ios/wallpaper/64/13/",
		"http://www.25pp.com/ios/wallpaper/64/22/",
		"http://www.25pp.com/ios/wallpaper/64/19/",
		"http://www.25pp.com/ios/wallpaper/64/15/",
		"http://www.25pp.com/ios/wallpaper/64/12/",
		"http://www.25pp.com/ios/wallpaper/64/17/",
		"http://www.25pp.com/ios/wallpaper/64/20/",
		"http://www.25pp.com/ios/wallpaper/64/18/",
		"http://www.25pp.com/ios/wallpaper/64/16/",
		"http://www.25pp.com/ios/wallpaper/64/21/",
		"http://www.7kk.com/shoujibizhi/shouji_qiche/new-1---1.html",
		"http://www.7kk.com/shoujibizhi/shouji_meinv/new-1---2.html",
		"http://www.7kk.com/shoujibizhi/shouji_mingxing/new-1---1.html",
		"http://www.7kk.com/shoujibizhi/shouji_youxi/new-1---1.html",
		"http://www.7kk.com/shoujibizhi/shouji_tiyu/new----1.html",
		"http://www.7kk.com/shoujibizhi/shouji_yingshi/new-1---1.html",
		"http://www.7kk.com/shoujibizhi/shouji_chongwu/new----1.html",
		"http://www.7kk.com/tagyun/2029.html",
		"http://www.7kk.com/tagyun/2076_6.html",
		"http://theiphonewalls.com/category/city/",
		"http://theiphonewalls.com/category/plain/",
		"http://theiphonewalls.com/category/patterns/",
		"http://theiphonewalls.com/category/landscape/",
		"http://theiphonewalls.com/category/nature/",
		"http://theiphonewalls.com/category/technology/",
		"http://theiphonewalls.com/category/flowers/",
		"http://theiphonewalls.com/category/games/",
		"http://theiphonewalls.com/category/paintings/",
		"http://theiphonewalls.com/category/abstract/",
		"http://www.bz55.com/shoujibizhi/chuangyisjbz/",
		"http://www.bz55.com/shoujibizhi/fengjingsjbz/",
		"http://www.bz55.com/shoujibizhi/meinvsjbz/",
		"http://www.bz55.com/shoujibizhi/mingxingsjbz/",
		"http://www.bz55.com/shoujibizhi/xiaoqingxin/",
		"http://www.bz55.com/shoujibizhi/youxisjbz/",
		"http://www.bz55.com/shoujibizhi/zhiwusjbz/",
	}

	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(len(urls))
	for _, urlStr := range urls {
		go func(urlStr string) {
			defer waitGroup.Done()

			session, err := phantomJS.NewSession(nil)
			if err != nil {
				t.Fatal(err)
			}
			defer session.Close()

			err = session.Get(urlStr)
			if err != nil {
				t.Fatal(err)
			}

			source, err := session.GetPageSource()
			if err != nil {
				t.Fatal(err)
			}
			if source == "" {
				t.Fatal("get empty webpage source")
			}
		}(urlStr)
	}

	waitGroup.Wait()
}

func TestWindowSize(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	defer phantomJS.Quit()

	session, err := phantomJS.NewSession(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	width, height, err := session.GetWindowSize()
	if err != nil {
		t.Fatal(err)
	}
	err = session.SetWindowSize(1920, 1080)
	if err != nil {
		t.Fatal(err)
	}
	width, height, err = session.GetWindowSize()
	if width != 1920 && height != 1080 {
		t.Fatal("set window size or get window size failed.")
	}
}

func TestGetWindowPosition(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	defer phantomJS.Quit()

	session, err := phantomJS.NewSession(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	x, y, err := session.GetWindowPosition()
	if err != nil {
		t.Fatal(err)
	}
	if x != 0 && y != 0 {
		t.Fatal("get invalid window position")
	}
}

func TestSetWindowPosition(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	defer phantomJS.Quit()

	session, err := phantomJS.NewSession(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	err = session.SetWindowPosition(10, 10)
	if err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddCookie(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	defer phantomJS.Quit()

	session, err := phantomJS.NewSession(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	err = session.Get("https://www.baidu.com")
	if err != nil {
		t.Fatal(err)
	}

	cookie := Cookie{Name: "test", Path: "/", Value: "hello, world", Domain: ".baidu.com"}
	err = session.AddCookie(&cookie)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetAllCookies(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	defer phantomJS.Quit()

	session, err := phantomJS.NewSession(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	err = session.Get("https://www.baidu.com")
	if err != nil {
		t.Fatal(err)
	}
	cookies, err := session.GetAllCookies()
	if err != nil {
		t.Fatal(err)
	}
	if len(cookies) == 0 {
		t.Fatal("get cookies failed")
	}
}

func TestGetCookie(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	defer phantomJS.Quit()

	session, err := phantomJS.NewSession(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	err = session.Get("https://www.baidu.com")
	if err != nil {
		t.Fatal(err)
	}

	cookie := Cookie{Name: "test", Path: "/", Value: "hello, world", Domain: ".baidu.com"}
	err = session.AddCookie(&cookie)
	if err != nil {
		t.Fatal(err)
	}

	ck, err := session.GetCookie("test")
	if err != nil {
		t.Fatal(err)
	}
	if ck == nil || ck.Name != cookie.Name || ck.Value != cookie.Value {
		t.Fatal("get cookie test failed")
	}
}

func TestDeleteAllCookies(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	defer phantomJS.Quit()

	session, err := phantomJS.NewSession(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	err = session.Get("https://www.baidu.com")
	if err != nil {
		t.Fatal(err)
	}

	err = session.DeleteAllCookies()
	if err != nil {
		t.Fatal(err)
	}

	cookies, err := session.GetAllCookies()
	if err != nil {
		t.Fatal(err)
	}
	if len(cookies) != 0 {
		t.Fatal("delete cookies failed")
	}
}

func TestSetPageLoadTimeouts(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	defer phantomJS.Quit()

	session, err := phantomJS.NewSession(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	err = session.SetPageLoadTimeouts(1000)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetScriptTimeout(t *testing.T) {
	options := &Options{}
	phantomJS, err := NewPhantomJS(0, options)
	if err != nil {
		t.Fatal(err)
	}
	defer phantomJS.Quit()

	session, err := phantomJS.NewSession(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Close()

	err = session.SetScriptTimeout(0)
	if err != nil {
		t.Fatal(err)
	}
}
