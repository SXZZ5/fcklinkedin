package linkedin

import (
	"fmt"
	"time"

	"github.com/go-rod/rod/lib/proto"
)

var (
	selector_signinButton = `#main-content > div.flip-card.flex-grow.min-w-\[300px\].max-w-\[416px\] > form > p > button`

	signin_link = `https://www.linkedin.com/authwall?trk=gf&trkInfo=AQF5-wnB71mxXwAAAZVm5lkYW7EfQG_lqsi8v5jYizvLLLKpfLCX9GatpeKCHkMM_TUhIq5_Pur5AVtCRWiWFlKBBY6exknPnoU8xCQhLvEQsWG5wLLERTYkRkq4YmqYYWx3Yzo=&original_referer=&sessionRedirect=https%3A%2F%2Fwww.linkedin.com%2Fin%2Fsushantsxzz5%2F#main-content`

	selector_UsernameEmailDalo = "#session_key"
	selector_PasswordDalo      = "#session_password"
)

func (z *Navigator) Login() {
	z.Page = z.Browser.MustPage(signin_link)
	z.Page.MustWaitStable()

	el := z.Page.MustElement(selector_signinButton)
	time.Sleep(time.Second * 2)
	if err := el.Click(proto.InputMouseButtonLeft, 1); err != nil {
		fmt.Println(err)
	}

	z.Page.MustActivate()
	//selecting the username/email field.
	el = z.Page.MustElement(selector_UsernameEmailDalo)
	el.MustInput("jadaTezNaBano")
	el = z.Page.MustElement(selector_PasswordDalo)
	el.MustInput("jadaTezNaBano")
	el = z.Page.MustElement(`#main-content > div.flip-card.flex-grow.min-w-\[300px\].max-w-\[416px\].show-login > div.authwall-sign-in-form > form > div.flex.justify-between.sign-in-form__footer--full-width > button`)
	el.MustClick()
}
