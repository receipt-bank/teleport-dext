// Copyright 2022 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helpers

import (
	"context"
	"time"

	"github.com/gravitational/teleport/lib/service"
	"github.com/gravitational/trace"
	authztypes "k8s.io/client-go/kubernetes/typed/authorization/v1"
)

func nullImpersonationCheck(context.Context, string, authztypes.SelfSubjectAccessReviewInterface) error {
	return nil
}

func StartAndWait(process *service.TeleportProcess, expectedEvents []string) ([]service.Event, error) {
	// register to listen for all ready events on the broadcast channel
	broadcastCh := make(chan service.Event)
	for _, eventName := range expectedEvents {
		process.WaitForEvent(context.TODO(), eventName, broadcastCh)
	}

	// start the process
	err := process.Start()
	if err != nil {
		return nil, trace.Wrap(err)
	}

	// wait for all events to arrive or a timeout. if all the expected events
	// from above are not received, this instance will not start
	receivedEvents := []service.Event{}
	timeoutCh := time.After(10 * time.Second)

	for idx := 0; idx < len(expectedEvents); idx++ {
		select {
		case e := <-broadcastCh:
			receivedEvents = append(receivedEvents, e)
		case <-timeoutCh:
			return nil, trace.BadParameter("timed out, only %v/%v events received. received: %v, expected: %v",
				len(receivedEvents), len(expectedEvents), receivedEvents, expectedEvents)
		}
	}

	// Not all services follow a non-blocking Start/Wait pattern. This means a
	// *Ready event may be emit slightly before the service actually starts for
	// blocking services. Long term those services should be re-factored, until
	// then sleep for 250ms to handle this situation.
	time.Sleep(250 * time.Millisecond)

	return receivedEvents, nil
}
