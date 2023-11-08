package subscribers

import (
	"github.com/14799678/echo_rest_api/pkg/hook"
)

type UserSubscriber struct {
}

func (s UserSubscriber) Created(payload hook.EventPayload) {

}

func (s UserSubscriber) Deleted(payload hook.EventPayload) {

}

func (s UserSubscriber) Updated(payload hook.EventPayload) {

}
