package endly

import (
	"github.com/viant/endly/common"
)

type ServiceResponse struct {
	Status   string
	Error    string
	Response interface{}
}

type Service interface {
	Id() string
	State() common.Map
	Run(context *Context, request interface{}) *ServiceResponse
	NewRequest(method string) (interface{}, error)
}

type AbstractService struct {
	Service
	id    string
	state common.Map
}

func (s *AbstractService) Id() string {
	return s.id
}

func (s *AbstractService) State() common.Map {
	return s.state
}

func NewAbstractService(id string) *AbstractService {
	return &AbstractService{
		id:    id,
		state: common.NewMap(),
	}
}