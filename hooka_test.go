package hooka

import (
	"bytes"
	"fmt"
	"hooka/mock"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNew(t *testing.T) {
	// Test Case
	cases := []struct {
		uri     string
		service service
		none    bool
	}{
		{
			uri:     "https://azure.com:443/workflows",
			service: Teams,
			none:    false,
		},
		{
			uri:     "https://slack.com/workflow",
			service: Slack,
			none:    false,
		},
		{
			uri:     "https://example.com/webhook",
			service: Unknown,
			none:    false,
		},
		{
			uri:     "webhook_url",
			service: Unknown,
			none:    true,
		},
	}
	// Test
	for i, c := range cases {
		h := New(c.uri)
		if h.uri.String() != c.uri {
			t.Errorf(`[Case%d] URL Result: %v, Expected: %v`, i, h.uri, c.uri)
		}
		if h.service != c.service {
			t.Errorf(`[Case%d] Service Result: %v, Expected: %v`, i, h.service, c.service)
		}
		if c.none && h.uri.Hostname() != "" {
			t.Errorf(`[Case%d] URL Nil Expected: %v`, i, h.uri.Hostname())
		}
		if !c.none && fmt.Sprintf("%T", h) != "*hooka.Hooka" {
			t.Errorf(`[Case%d] *Hooka Expected: %v`, i, fmt.Sprintf("%T", h))
		}
	}
}

func TestHookaSend(t *testing.T) {
	// gomock controller作成とclose予約(呪文のようなもの)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock instance作成
	mockSample := mock.NewMockClient(ctrl)
	// 引数"hoge"でMethodが呼ばれることか確認する
	req, _ := http.NewRequest("POST", "https://example.com", bytes.NewReader([]byte("{}")))
	mockSample.EXPECT().Do(req).Return(&http.Response{StatusCode: 500})

}
