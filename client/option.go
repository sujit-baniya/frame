/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package client

import (
	"crypto/tls"
	"github.com/sujit-baniya/frame/client/retry"
	"time"

	"github.com/sujit-baniya/frame/pkg/common/config"
	"github.com/sujit-baniya/frame/pkg/network"
	"github.com/sujit-baniya/frame/pkg/network/standard"
	"github.com/sujit-baniya/frame/pkg/protocol/consts"
)

// WithDialTimeout sets dial timeout.
func WithDialTimeout(dialTimeout time.Duration) config.ClientOption {
	return config.ClientOption{F: func(o *config.ClientOptions) {
		o.DialTimeout = dialTimeout
	}}
}

// WithMaxConnsPerHost sets maximum number of connections per host which may be established.
func WithMaxConnsPerHost(mc int) config.ClientOption {
	return config.ClientOption{F: func(o *config.ClientOptions) {
		o.MaxConnsPerHost = mc
	}}
}

// WithMaxIdleConnDuration sets max idle connection duration, idle keep-alive connections are closed after this duration.
func WithMaxIdleConnDuration(t time.Duration) config.ClientOption {
	return config.ClientOption{F: func(o *config.ClientOptions) {
		o.MaxIdleConnDuration = t
	}}
}

// WithMaxConnDuration sets max connection duration, keep-alive connections are closed after this duration.
func WithMaxConnDuration(t time.Duration) config.ClientOption {
	return config.ClientOption{F: func(o *config.ClientOptions) {
		o.MaxConnDuration = t
	}}
}

// WithMaxConnWaitTimeout sets maximum duration for waiting for a free connection.
func WithMaxConnWaitTimeout(t time.Duration) config.ClientOption {
	return config.ClientOption{F: func(o *config.ClientOptions) {
		o.MaxConnWaitTimeout = t
	}}
}

// WithKeepAlive determines whether use keep-alive connection.
func WithKeepAlive(b bool) config.ClientOption {
	return config.ClientOption{F: func(o *config.ClientOptions) {
		o.KeepAlive = b
	}}
}

// WithClientReadTimeout sets maximum duration for full response reading (including body).
func WithClientReadTimeout(t time.Duration) config.ClientOption {
	return config.ClientOption{F: func(o *config.ClientOptions) {
		o.ReadTimeout = t
	}}
}

// WithTLSConfig sets tlsConfig to create a tls connection.
func WithTLSConfig(cfg *tls.Config) config.ClientOption {
	return config.ClientOption{F: func(o *config.ClientOptions) {
		o.TLSConfig = cfg
		o.Dialer = standard.NewDialer()
	}}
}

// WithDialer sets the specific dialer.
func WithDialer(d network.Dialer) config.ClientOption {
	return config.ClientOption{F: func(o *config.ClientOptions) {
		o.Dialer = d
	}}
}

// WithResponseBodyStream is used to determine whether read body in stream or not.
func WithResponseBodyStream(b bool) config.ClientOption {
	return config.ClientOption{F: func(o *config.ClientOptions) {
		o.ResponseBodyStream = b
	}}
}

// WithDisableHeaderNamesNormalizing is used to set whether disable header names normalizing.
func WithDisableHeaderNamesNormalizing(disable bool) config.ClientOption {
	return config.ClientOption{F: func(o *config.ClientOptions) {
		o.DisableHeaderNamesNormalizing = disable
	}}
}

// WithName sets client name which used in User-Agent Header.
func WithName(name string) config.ClientOption {
	return config.ClientOption{F: func(o *config.ClientOptions) {
		o.Name = name
	}}
}

// WithNoDefaultUserAgentHeader sets whether no default User-Agent header.
func WithNoDefaultUserAgentHeader(isNoDefaultUserAgentHeader bool) config.ClientOption {
	return config.ClientOption{F: func(o *config.ClientOptions) {
		o.NoDefaultUserAgentHeader = isNoDefaultUserAgentHeader
	}}
}

// WithDisablePathNormalizing sets whether disable path normalizing.
func WithDisablePathNormalizing(isDisablePathNormalizing bool) config.ClientOption {
	return config.ClientOption{F: func(o *config.ClientOptions) {
		o.DisablePathNormalizing = isDisablePathNormalizing
	}}
}

func WithRetryConfig(opts ...retry.Option) config.ClientOption {
	retryCfg := &retry.Config{
		MaxAttemptTimes: consts.DefaultMaxRetryTimes,
		Delay:           1 * time.Millisecond,
		MaxDelay:        100 * time.Millisecond,
		MaxJitter:       20 * time.Millisecond,
		DelayPolicy:     retry.CombineDelay(retry.DefaultDelayPolicy),
	}
	retryCfg.Apply(opts)

	return config.ClientOption{F: func(o *config.ClientOptions) {
		o.RetryConfig = retryCfg
	}}
}