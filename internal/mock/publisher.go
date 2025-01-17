/*******************************************************************************
 * Copyright 2023 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package mock

import (
	"github.com/project-alvarium/alvarium-sdk-go/pkg/config"
	"github.com/project-alvarium/alvarium-sdk-go/pkg/interfaces"
	"github.com/project-alvarium/alvarium-sdk-go/pkg/message"
)

type mockPublisher struct {
	cfg    config.MockStreamConfig
	logger interfaces.Logger
}

func NewMockPublisher(cfg config.MockStreamConfig, logger interfaces.Logger) (interfaces.StreamProvider, error) {
	return &mockPublisher{
		cfg:    cfg,
		logger: logger,
	}, nil
}

func (p *mockPublisher) Connect() error {
	return nil
}

func (p *mockPublisher) Publish(msg message.PublishWrapper) error {
	return nil
}

func (p *mockPublisher) Close() error {
	return nil
}
