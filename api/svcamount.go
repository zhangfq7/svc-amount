package api

import (
	"github.com/zonesan/clog"
)

type svcAmount struct {
	Name      string `json:"name"`
	Used      string `json:"used"`
	Size      string `json:"size,omitempty"`
	Available string `json:"available,omitempty"`
	Desc      string `json:"desc,omitempty"`
}

type svcAmountList struct {
	Items []svcAmount `json:"items"`
}

func DoomServiceInstance(bsi *BackingServiceInstance) (*svcAmountList, error) {
	bsname := bsi.Spec.BackingServiceName
	clog.Debug("bsi.ServiceName:", bsname)
	for k, v := range bsi.Spec.Creds {
		println("###", k, v)
	}

	agent, err := findDriver(bsname)
	if err != nil {
		clog.Error(err)
		return nil, err
	}
	amountlist := agent.GetAmount(bsname, nil)
	return amountlist, nil
}
