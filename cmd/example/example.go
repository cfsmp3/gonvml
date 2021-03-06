/*
Copyright 2017 Google Inc.

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

package main

import (
	"fmt"
	"time"

	"github.com/cfsmp3/gonvml"
)

func main() {
	start := time.Now()
	err := gonvml.Initialize()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer gonvml.Shutdown()
	fmt.Printf("Initialize() took %v\n", time.Since(start))

	driverVersion, err := gonvml.SystemDriverVersion()
	if err != nil {
		fmt.Printf("SystemDriverVersion() error: %v\n", err)
		return
	}
	fmt.Printf("SystemDriverVersion(): %v\n", driverVersion)

	numDevices, err := gonvml.DeviceCount()
	if err != nil {
		fmt.Printf("DeviceCount() error: %v\n", err)
		return
	}
	fmt.Printf("DeviceCount(): %v\n", numDevices)

	for i := 0; i < int(numDevices); i++ {
		dev, err := gonvml.DeviceHandleByIndex(uint(i))
		if err != nil {
			fmt.Printf("\tDeviceHandleByIndex() error: %v\n", err)
			return
		}

		minorNumber, err := dev.MinorNumber()
		if err != nil {
			fmt.Printf("\tdev.MinorNumber() error: %v\n", err)
			return
		}
		fmt.Printf("\tminorNumber: %v\n", minorNumber)

		busID, err := dev.BusID()
		if err != nil {
			fmt.Printf("\tdev.BusID() error: %v\n", err)
			return
		}
		fmt.Printf("\tBusID: %v\n", busID)

		uuid, err := dev.UUID()
		if err != nil {
			fmt.Printf("\tdev.UUID() error: %v\n", err)
			return
		}
		fmt.Printf("\tuuid: %v\n", uuid)

		name, err := dev.Name()
		if err != nil {
			fmt.Printf("\tdev.Name() error: %v\n", err)
			return
		}
		fmt.Printf("\tname: %v\n", name)

		board, err := dev.BoardID()
		if err != nil {
			fmt.Printf("\tdev.BoardID() error: %v\n", err)
		} else {
		    fmt.Printf("\tBoardID: %v\n", board)
        }

		vbiosversion, err := dev.VBiosVersion()
		if err != nil {
			fmt.Printf("\tdev.VBiosVersion() error: %v\n", err)
		}
		fmt.Printf("\tVBiosVersion: %v\n", vbiosversion)

		serial, err := dev.Serial()
		if err != nil {
			fmt.Printf("\tdev.Serial() error: %v\n", err)
		} else {
		    fmt.Printf("\tSerial: %v\n", serial)
        }

		displaymode, err := dev.DisplayMode()
		if err != nil {
			fmt.Printf("\tdev.DisplayMode() error: %v\n", err)
			return
		}
		fmt.Printf("\tDisplayMode: %d (%v)\n", displaymode, displaymode)

		displayactive, err := dev.DisplayActive()
		if err != nil {
			fmt.Printf("\tdev.DisplayActive() error: %v\n", err)
			return
		}
		fmt.Printf("\tDisplayActive: %d (%v)\n", displayactive, displayactive)

		persistencemode, err := dev.PersistenceMode()
		if err != nil {
			fmt.Printf("\tdev.PersistenceMode() error: %v\n", err)
			return
		}
		fmt.Printf("\tPersistenceMode: %v\n", persistencemode)

		performancestate, err := dev.PerformanceState()
		if err != nil {
			fmt.Printf("\tdev.PerformanceState() error: %v\n", err)
			return
		}
		fmt.Printf("\tPerformanceState: %v\n", performancestate)

		totalMemory, usedMemory, err := dev.MemoryInfo()
		if err != nil {
			fmt.Printf("\tdev.MemoryInfo() error: %v\n", err)
			return
		}
		fmt.Printf("\tmemory.total: %v, memory.used: %v\n", totalMemory, usedMemory)

		totalBAR1Memory, usedBAR1Memory, err := dev.Bar1MemoryInfo()
		if err != nil {
			fmt.Printf("\tdev.Bar1MemoryInfo() error: %v\n", err)
			return
		}
		fmt.Printf("\tBAR1 memory.total: %v, memory.used: %v\n", totalBAR1Memory, usedBAR1Memory)

		gpuUtilization, memoryUtilization, err := dev.UtilizationRates()
		if err != nil {
			fmt.Printf("\tdev.UtilizationRates() error: %v\n", err)
			return
		}
		fmt.Printf("\tutilization.gpu: %v, utilization.memory: %v\n", gpuUtilization, memoryUtilization)

		PowerManagementDefaultLimit, err := dev.PowerManagementDefaultLimit()
		if err != nil {
			fmt.Printf("\tdev.PowerManagementDefaultLimit() error: %v\n", err)
		} else {
			fmt.Printf("\tpower.limit.default: %v\n", PowerManagementDefaultLimit)
		}

		minLimit, maxLimit, err := dev.PowerLimitConstraints()
		if err != nil {
			fmt.Printf("\tdev.PowerLimitConstraints() error: %v\n", err)
		} else {
			fmt.Printf("\tpower.min_limit: %v\n", minLimit / 1000)
			fmt.Printf("\tpower.max_limit: %v\n", maxLimit / 1000)
		}

		limit_mgmt, limit_enforced, err := dev.PowerLimits()
		if err != nil {
			fmt.Printf("\tdev.PowerLimits() error: %v\n", err)
		} else {
			fmt.Printf("\tpower.limit (management): %v\n", limit_mgmt / 1000)
			fmt.Printf("\tpower.limit (enforced): %v\n", limit_enforced / 1000)
		}

        /* Ignore errors here, if the library is not loaded something else
        must have fail before */
        /* CLOCKS */
		fmt.Printf("\n\tCLOCKS\n")
		GrClockCurrent, _ := dev.GrClock()
		GrClockMax, _ := dev.GrMaxClock()
		fmt.Printf("\tgraphics clock (current / max): %v / %v\n", GrClockCurrent, GrClockMax)
		SMClockCurrent, _ := dev.SMClock()
		SMClockMax, _ := dev.SMMaxClock()
		fmt.Printf("\tSM clock (current / max): %v / %v\n", SMClockCurrent, SMClockMax)
		MemClockCurrent, _ := dev.MemClock()
		MemClockMax, _ := dev.MemMaxClock()
		fmt.Printf("\tMemory clock (current / max): %v / %v\n", MemClockCurrent, MemClockMax)
		VideoClockCurrent, _ := dev.VideoClock()
		VideoClockMax, _ := dev.VideoMaxClock()
		fmt.Printf("\tVideo clock (current / max): %v / %v\n", VideoClockCurrent, VideoClockMax)

        /* THROTTLING */
		fmt.Printf("\n\tTHROTTLING\n")
		throttlereasons, err := dev.CurrentClocksThrottleReasons()
		if err != nil {
			fmt.Printf("\tdev.throttlereasons() error: %v\n", err)
			return
		}
		fmt.Printf("\tthrottlereasons (bitmap): %v\n", throttlereasons)
		throttleseriousreason, err := dev.MostSeriousClocksThrottleReason()
		if err != nil {
			fmt.Printf("\tdev.throttleseriousreason() error: %v\n", err)
			return
		}
		fmt.Printf("\tMost serious throttle reason: %v\n", throttleseriousreason)

        /* PCI */
		fmt.Printf("\n\tPCI\n")
        pci_tx, _ := dev.PcieTxThroughput()
        pci_rx, _ := dev.PcieRxThroughput()
		fmt.Printf("\tPCI throughput in KB/s: %v / %v\n", pci_tx, pci_rx)
        pci_link_gen_current, _ := dev.PcieGeneration()
        pci_link_gen_max, _ := dev.PcieMaxGeneration()
		fmt.Printf("\tCurrent PCIe link generation (current / max): %v / %v\n", pci_link_gen_current, pci_link_gen_max )
        pci_link_width_current, _ := dev.PcieWidth()
        pci_link_width_max, _ := dev.PcieMaxWidth()
		fmt.Printf("\tCurrent PCIe link width (current / max): %v / %v\n", pci_link_width_current, pci_link_width_max)


        /* POWER */
		fmt.Printf("\n\tPOWER\n")
		powerDraw, err := dev.PowerUsage()
		if err != nil {
			fmt.Printf("\tdev.PowerUsage() error: %v\n", err)
		} else {
		    fmt.Printf("\tpower.draw: %v\n", powerDraw)
        }

		energy, err := dev.TotalEnergyConsumption()
		if err != nil {
			fmt.Printf("\tdev.TotalEnergyConsumption() error: %v\n", err)
		} else {
            fmt.Printf("\tTotalEnergyConsumption: %v\n", energy)
        }

		averagePowerDraw, err := dev.AveragePowerUsage(10 * time.Second)
		if err != nil {
			fmt.Printf("\tdev.AveragePowerUsage() error: %v\n", err)
		} else {
		    fmt.Printf("\taverage power.draw for last 10s: %v\n", averagePowerDraw)
        }

		averageGPUUtilization, err := dev.AverageGPUUtilization(10 * time.Second)
		if err != nil {
			fmt.Printf("\tdev.AverageGPUUtilization() error: %v\n", err)
        } else {
		    fmt.Printf("\taverage utilization.gpu for last 10s: %v\n", averageGPUUtilization)
        }

        /* TEMPERATURE AND FANS */
		fmt.Printf("\n\tTEMPERATURE AND FANS\n")
		temperature, err := dev.Temperature()
		if err != nil {
			fmt.Printf("\tdev.Temperature() error: %v\n", err)
			return
		}
		fmt.Printf("\ttemperature.gpu: %v C\n", temperature)

		temp_threshold_shutdown, temp_threshold_slowdown, err := dev.TemperatureThresholds()
		if err != nil {
			fmt.Printf("\tdev.TemperatureThresholds() error: %v\n", err)
		} else {
			fmt.Printf("\temperature.threshold.shutdown: %v\n", temp_threshold_shutdown)
			fmt.Printf("\temperature.threshold.slowdown: %v\n", temp_threshold_slowdown)
		}

		fanSpeed, err := dev.FanSpeed()
		if err != nil {
			fmt.Printf("\tdev.FanSpeed() error: %v\n", err)
		} else {
		    fmt.Printf("\tfan.speed: %v%%\n", fanSpeed)
        }

        /* MEMORY */
		fmt.Printf("\n\tMEMORY\n")
		m1, m2, m3, m4, err := dev.TotalEccErrors()
		if err != nil {
			fmt.Printf("\tdev.dev.TotalEccErrors() error: %v\n", err)
		} else {
		    fmt.Printf("\tCorrected errors since last reboot: %v%%\n", m1)
		    fmt.Printf("\tCorrected errors ever: %v%%\n", m2)
		    fmt.Printf("\tUncorrected errors since last reboot: %v%%\n", m3)
		    fmt.Printf("\tUncorrected errors ever: %v%%\n", m4)
        }


        /* UTILIZATION */
		fmt.Printf("\n\tUTILIZATION\n")
		encoderUtilization, sampling_period, err := dev.EncoderUtilization()
		if err != nil {
			fmt.Printf("\tdev.EncoderUtilization() error: %v\n", err)
		} else {
		    fmt.Printf("\tutilization.encoder: %d over %d microseconds\n", encoderUtilization, sampling_period)
        }

		decoderUtilization, _, err := dev.DecoderUtilization()
		if err != nil {
			fmt.Printf("\tdev.DecoderUtilization() error: %v\n", err)
		} else {
		    fmt.Printf("\tutilization.decoder: %d\n", decoderUtilization)
        }


		caph264, caphevc, err := dev.EncoderCapacity()
		if err != nil {
			fmt.Printf("\tdev.EncoderCapacity() error: %v\n", err)
		} else {
            fmt.Printf("\tEncoder Capacity H264: %d  HEVC: %d\n", caph264, caphevc)
        }


		fmt.Printf("\n")
		modeStats, err := dev.AccountingMode()
		if err != nil {
			fmt.Printf("\tdev.DeviceGetAccountingMode() error: %v\n", err)
		} else {
		    fmt.Printf("\taccounting.mode enable: %v\n", modeStats)
        }

		computeMode, err := dev.ComputeMode()
		if err != nil {
			fmt.Printf("\tdev.ComputeMode() error: %v\n", err)
		} else {
		    fmt.Printf("\tcompute mode: %d (%v)\n", computeMode, computeMode)
        }

		bufferSize, err := dev.AccountingBufferSize()
		if err != nil {
			fmt.Printf("\tdev.DeviceGetAccountingBufferSize() error: %v\n", err)
		} else {
		    fmt.Printf("\taccounting.buffersize: %d\n", bufferSize)
        }

		pids, count, err := dev.AccountingPids(bufferSize)
		if err != nil {
			fmt.Printf("\tdev.DeviceGetAccountingPids() error: %v\n", err)
		} else {
			fmt.Printf("\taccounting.pids.count: %v\n", count)
			for _, pid := range pids[:count] {
				fmt.Printf("\t\tPid: %v", pid)
				stats, err := dev.AccountingStats(uint(pid))
				if err != nil {
					fmt.Printf("\tdev.DeviceGetAccountingStats() error: %v\n", err)
				} else {
					fmt.Printf(", GPUUtilization: %v", stats.GPUUtilization)
					fmt.Printf(", MemoryUtilization: %v", stats.MemoryUtilization)
					fmt.Printf(", MaxMemoryUsage: %v", stats.MaxMemoryUsage)
					fmt.Printf(", Time: %v", stats.Time)
					fmt.Printf(", StartTime: %v", stats.StartTime)
					fmt.Printf(", IsRunning: %v", stats.IsRunning)
					fmt.Println()
				}
			}
		}

		utilizations, err := dev.ProcessUtilization(10, 10*time.Second)
		if err != nil {
			fmt.Printf("\tdev.DeviceGetProcessUtilization() error: %v\n", err)
		} else {
			fmt.Printf("\tProcess count: %v\n", len(utilizations))

			utilizations = utilizations
			for _, sample := range utilizations {
				fmt.Printf("\t\tProcess: %v", sample.Pid)
				fmt.Printf(", SM  util: %v", sample.SMUtil)
				fmt.Printf(", Mem util: %v", sample.MemUtil)
				fmt.Printf(", Enc util: %v", sample.EncUtil)
				fmt.Printf(", Dec util: %v", sample.DecUtil)

				name, err := gonvml.SystemGetProcessName(sample.Pid, 64)
				if err != nil {
					fmt.Printf("\n\tdev.SystemGetProcessName() error: %v\n", err)
				} else {
					fmt.Printf(", Name: %s\n", name)
				}
			}
		}

		fmt.Println()
	}
}
