package strategy

import (
	"testing"
)

func TestStrategy(t *testing.T) {
	seller := &seller{}

	seller.changePaymentype(&wechat{})
	seller.swipe(100000)

	seller.changePaymentype(&alipay{})
	seller.swipe(999)

	seller.changePaymentype(&stripe{})
	seller.swipe(1234567)
}
