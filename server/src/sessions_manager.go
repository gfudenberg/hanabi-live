package main

type SessionsManager struct {
}

var (
	sessionsRequests chan *SessionsRequest
)

type SessionsRequest struct {
	Type    int
	Confirm chan bool
}

const (
	SessionsRequestTypeFoo = iota
	SessionsRequestTypeBar
)

func sessionsManagerInit() {
	for {
		request := <-sessionsRequests
		if request.Type == SessionsRequestTypeFoo {
			request.Confirm <- true
		} else {
			request.Confirm <- false
		}
	}
}
