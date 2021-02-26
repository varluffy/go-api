/**
 * @Time: 2021/2/25 9:45 上午
 * @Author: varluffy
 * @Description: //TODO
 */

package app

import (
	"github.com/varluffy/ginx/transport/http"
	"testing"
	"time"
)

func TestApp(t *testing.T) {
	hs := http.NewServer()
	app := New(Name("test"), Version("1.0.0"), Server(hs))
	time.AfterFunc(time.Second, func() {
		_ = app.Stop()
	})
	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
}
