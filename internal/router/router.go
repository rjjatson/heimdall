package router

type Router struct {
	handlers map[string]func([]byte, ResponseWriter)
	// todo add hub outboun channel
}

// ResponseWriter
type ResponseWriter struct {
	outbound chan interface{}
}

func (resp *ResponseWriter) Write(userID string, message interface{}) {

}
