package main

type TablesManager struct {
}

var (
	tablesRequests chan *TablesRequest
)

type TablesRequest struct {
	Type    int
	Confirm chan bool
}

const (
	TablesRequestTypeFoo = iota
	TablesRequestTypeBar
)

func tablesManagerInit() {
	for {
		request := <-tablesRequests
		if request.Type == TablesRequestTypeFoo {
			request.Confirm <- true
		} else {
			request.Confirm <- false
		}
	}
}
