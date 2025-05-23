/*
Copyright 2024 The Aibrix Team.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package routingalgorithms

import (
	"context"
	"math"

	"github.com/aibrix/aibrix/pkg/cache"
	ratelimiter "github.com/aibrix/aibrix/pkg/plugins/gateway/ratelimiter"
	v1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
)

type leastRequestRouter struct {
	ratelimiter ratelimiter.RateLimiter
	cache       *cache.Cache
}

func NewLeastRequestRouter(ratelimiter ratelimiter.RateLimiter) Router {
	cache, err := cache.GetCache()
	if err != nil {
		panic(err)
	}

	return leastRequestRouter{
		ratelimiter: ratelimiter,
		cache:       cache,
	}
}

func (r leastRequestRouter) Route(ctx context.Context, pods map[string]*v1.Pod) (string, error) {
	var targetPodIP string
	minCount := math.MaxFloat64

	for _, pod := range pods {
		if pod.Status.PodIP == "" {
			continue
		}

		runningReq, err := r.cache.GetPodMetric(pod.Name, num_requests_running)
		if err != nil {
			klog.Error(err)
			continue
		}
		waitingReq, err := r.cache.GetPodMetric(pod.Name, num_requests_waiting)
		if err != nil {
			klog.Error(err)
			continue
		}
		swappedReq, err := r.cache.GetPodMetric(pod.Name, num_requests_swapped)
		if err != nil {
			klog.Error(err)
			continue
		}
		totalReq := runningReq + waitingReq + swappedReq
		klog.V(4).Infof("pod: %v, podIP: %v, runningReq: %v, waitingReq: %v, swappedReq: %v, totalReq: %v",
			pod.Name, pod.Status.PodIP, runningReq, waitingReq, swappedReq, totalReq)

		if totalReq <= minCount {
			minCount = totalReq
			targetPodIP = pod.Status.PodIP
		}
	}

	return targetPodIP + ":" + podPort, nil
}
