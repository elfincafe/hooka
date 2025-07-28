package hooka

import (
	"bytes"
	"hooka/mock"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestNew(t *testing.T) {
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
