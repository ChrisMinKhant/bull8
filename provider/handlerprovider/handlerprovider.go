package handlerprovider

type handlerProvider struct {
}

func NewHandlerProvider() *handlerProvider {
	return &handlerProvider{}
}

/*
 * All the handlers that will be used by the application
 * must be pre-registered here. Each handler must be
 * mapped with their intended path and http method.
 * The path-method must be in " ( /path|METHOD ) " order.
 */

func (handlerProvider *handlerProvider) Register() {

}
