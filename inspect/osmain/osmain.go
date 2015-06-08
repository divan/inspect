package osmain

import (
	"github.com/divan/inspect/os/cpustat"
	"github.com/divan/inspect/os/memstat"
	"github.com/divan/inspect/os/pidstat"
)

// OsIndependentStats must be implemented by all supported platforms
type OsIndependentStats struct {
	Cstat *cpustat.CPUStat
	Mstat *memstat.MemStat
	Procs *pidstat.ProcessStat
}
