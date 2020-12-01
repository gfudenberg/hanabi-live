package sessions

type Manager struct {
}

func NewManager


var (
	sessionsRequests chan *Request
)

type Request struct {
	Type    int
	Confirm chan bool
}

const (
	SessionsRequestTypeFoo = iota
	SessionsRequestTypeBar
)

func (m *Manager) AddSession() {
	for {
		request := <-sessionsRequests
		if request.Type == SessionsRequestTypeFoo {
			request.Confirm <- true
		} else {
			request.Confirm <- false
		}
	}
}
