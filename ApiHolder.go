// ApiHolder
package ts3api

type ApiHolder struct {
	api *TS3Api
}

func (event *ApiHolder) setApi(api *TS3Api) {
	event.api = api
}

func (event *ApiHolder) Api() *TS3Api {
	return event.api
}
