// Copyright (c) 2015 Square, Inc

// Package loadstat implements metrics collection related to loadavg
package loadstat

import (
	"bufio"
	"os"
	"strings"
	"time"

	"github.com/divan/inspect/metrics"
	"github.com/divan/inspect/os/misc"
)

// LoadStat represents load average metrics for 1/5/15 Minutes of
// current operating system.
// Caution: reflection is used to read this struct to discover names
// Do not add new types
type LoadStat struct {
	OneMinute     *metrics.Gauge
	FiveMinute    *metrics.Gauge
	FifteenMinute *metrics.Gauge
	m             *metrics.MetricContext
}

// New starts metrics collection every Step and registers with
// metricscontext
func New(m *metrics.MetricContext, Step time.Duration) *LoadStat {
	s := new(LoadStat)
	s.m = m
	// initialize all metrics and register them
	misc.InitializeMetrics(s, m, "loadstat", true)
	// collect once
	s.Collect()
	// collect metrics every Step
	ticker := time.NewTicker(Step)
	go func() {
		for _ = range ticker.C {
			s.Collect()
		}
	}()
	return s
}

// Collect populates Loadstat by reading /proc/loadavg
func (s *LoadStat) Collect() {
	file, err := os.Open("/proc/loadavg")
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		f := strings.Split(scanner.Text(), " ")
		if len(f) > 2 {
			s.OneMinute.Set(misc.ParseFloat(f[0]))
			s.FiveMinute.Set(misc.ParseFloat(f[1]))
			s.FifteenMinute.Set(misc.ParseFloat(f[2]))
		}
		break
	}
}
