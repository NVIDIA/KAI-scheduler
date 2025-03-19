/*
Copyright 2025 NVIDIA CORPORATION
SPDX-License-Identifier: Apache-2.0
*/
// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	http "net/http"

	scheme "github.com/NVIDIA/KAI-scheduler/pkg/apis/client/clientset/versioned/scheme"
	schedulingv1alpha2 "github.com/NVIDIA/KAI-scheduler/pkg/apis/scheduling/v1alpha2"
	rest "k8s.io/client-go/rest"
)

type SchedulingV1alpha2Interface interface {
	RESTClient() rest.Interface
	BindRequestsGetter
}

// SchedulingV1alpha2Client is used to interact with features provided by the scheduling.run.ai group.
type SchedulingV1alpha2Client struct {
	restClient rest.Interface
}

func (c *SchedulingV1alpha2Client) BindRequests(namespace string) BindRequestInterface {
	return newBindRequests(c, namespace)
}

// NewForConfig creates a new SchedulingV1alpha2Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*SchedulingV1alpha2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new SchedulingV1alpha2Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*SchedulingV1alpha2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &SchedulingV1alpha2Client{client}, nil
}

// NewForConfigOrDie creates a new SchedulingV1alpha2Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *SchedulingV1alpha2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new SchedulingV1alpha2Client for the given RESTClient.
func New(c rest.Interface) *SchedulingV1alpha2Client {
	return &SchedulingV1alpha2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := schedulingv1alpha2.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = rest.CodecFactoryForGeneratedClient(scheme.Scheme, scheme.Codecs).WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *SchedulingV1alpha2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
