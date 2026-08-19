//go:build !productmetrics_testhook

package main

import "github.com/jonbaldie/gascity/internal/productmetrics"

func configuredPrivateProductMetricsRunner() privateProductMetricsRunFunc {
	return runProductionProductMetricsChild
}

func configuredProductMetricsControlService() (*productmetrics.Service, error) {
	return openProductionProductMetricsService()
}
