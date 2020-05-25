package powerbox

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/DanNixon/go-2b/pkg/types"
)

type RestPowerbox struct {
	Address string
}

func NewRestPowerbox(address string) (*RestPowerbox, error) {
	pb := RestPowerbox{
		Address: address,
	}
	return &pb, nil
}

func (p *RestPowerbox) Reset() (types.Status, error) {
	u := bytes.NewReader([]byte("{}"))
	return receiveStatus(http.Post(p.Address+"/reset", "application/json", u))
}

func (p *RestPowerbox) Kill() (types.Status, error) {
	u := bytes.NewReader([]byte("{}"))
	return receiveStatus(http.Post(p.Address+"/kill", "application/json", u))
}

func (p *RestPowerbox) Set(s types.Settings) (types.Status, error) {
	jsonValue, _ := json.Marshal(s)
	u := bytes.NewReader(jsonValue)
	return receiveStatus(http.Post(p.Address+"/set", "application/json", u))
}

func (p *RestPowerbox) Get() (types.Status, error) {
	return receiveStatus(http.Get(p.Address + "/status"))
}

func receiveStatus(r *http.Response, err error) (types.Status, error) {
	if err != nil {
		return types.Status{}, err
	}
	defer r.Body.Close()

	var st types.Status
	err = json.NewDecoder(r.Body).Decode(&st)
	return st, err
}
