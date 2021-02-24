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

package gonvml

// #cgo LDFLAGS: -ldl
/*
#include <stddef.h>
#include <dlfcn.h>
#include <stdlib.h>
#include <stdio.h>

#include "nvml.h"

// nvmlHandle is the handle for dynamically loaded libnvidia-ml.so
void *nvmlHandle;

nvmlReturn_t (*nvmlInitFunc)(void);

nvmlReturn_t (*nvmlShutdownFunc)(void);

const char* (*nvmlErrorStringFunc)(nvmlReturn_t result);
const char* nvmlErrorString(nvmlReturn_t result) {
  if (nvmlErrorStringFunc == NULL) {
    return "nvmlErrorString Function Not Found";
  }
  return nvmlErrorStringFunc(result);
}

nvmlReturn_t (*nvmlSystemGetDriverVersionFunc)(char *version, unsigned int length);
nvmlReturn_t nvmlSystemGetDriverVersion(char *version, unsigned int length) {
  if (nvmlSystemGetDriverVersionFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlSystemGetDriverVersionFunc(version, length);
}

nvmlReturn_t (*nvmlSystemGetNVMLVersionFunc)(char *version, unsigned int length);
nvmlReturn_t nvmlSystemGetNVMLVersion(char *version, unsigned int length) {
  if (nvmlSystemGetNVMLVersionFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlSystemGetDriverVersionFunc(version, length);
}

nvmlReturn_t (*nvmlDeviceGetCountFunc)(unsigned int *deviceCount);
nvmlReturn_t nvmlDeviceGetCount(unsigned int *deviceCount) {
  if (nvmlDeviceGetCountFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetCountFunc(deviceCount);
}

nvmlReturn_t (*nvmlDeviceGetHandleByIndexFunc)(unsigned int index, nvmlDevice_t *device);
nvmlReturn_t nvmlDeviceGetHandleByIndex(unsigned int index, nvmlDevice_t *device) {
  if (nvmlDeviceGetHandleByIndexFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetHandleByIndexFunc(index, device);
}

nvmlReturn_t (*nvmlDeviceGetIndexFunc)(nvmlDevice_t device, unsigned int* index);
nvmlReturn_t nvmlDeviceGetIndex(nvmlDevice_t device, unsigned int* index){
	if(nvmlDeviceGetIndexFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	return nvmlDeviceGetIndexFunc(device, index);
}

nvmlReturn_t (*nvmlDeviceGetBoardIdFunc)(nvmlDevice_t device, unsigned int* boardId);
nvmlReturn_t nvmlDeviceGetBoardId(nvmlDevice_t device, unsigned int* boardId) {
	if(nvmlDeviceGetBoardIdFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	return nvmlDeviceGetBoardIdFunc(device, boardId);
}

nvmlReturn_t (*nvmlDeviceGetDisplayModeFunc)(nvmlDevice_t device, nvmlEnableState_t* display);
nvmlReturn_t nvmlDeviceGetDisplayMode(nvmlDevice_t device, nvmlEnableState_t* display) {
	if(*nvmlDeviceGetDisplayModeFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	return nvmlDeviceGetDisplayModeFunc(device, display);
}

nvmlReturn_t (*nvmlDeviceGetDisplayActiveFunc)(nvmlDevice_t device, nvmlEnableState_t* display);
nvmlReturn_t nvmlDeviceGetDisplayActive(nvmlDevice_t device, nvmlEnableState_t* display) {
	if(nvmlDeviceGetDisplayActiveFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	return nvmlDeviceGetDisplayActiveFunc(device, display);
}

nvmlReturn_t (*nvmlDeviceGetVbiosVersionFunc)(nvmlDevice_t device, char* version, unsigned int  length);
nvmlReturn_t nvmlDeviceGetVbiosVersion (nvmlDevice_t device, char* version, unsigned int  length) {
	if(nvmlDeviceGetVbiosVersionFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	return nvmlDeviceGetVbiosVersionFunc(device, version, length);
}

nvmlReturn_t (*nvmlDeviceGetCurrentClocksThrottleReasonsFunc)(nvmlDevice_t device, unsigned long long *clocksThrottleReasons);
nvmlReturn_t nvmlDeviceGetCurrentClocksThrottleReasons(nvmlDevice_t device, unsigned long long *clocksThrottleReasons)
{
	if(*nvmlDeviceGetCurrentClocksThrottleReasonsFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	return nvmlDeviceGetCurrentClocksThrottleReasonsFunc(device, clocksThrottleReasons);
}

nvmlReturn_t (*nvmlDeviceGetBAR1MemoryInfoFunc)(nvmlDevice_t device, nvmlBAR1Memory_t* bar1Memory);
nvmlReturn_t nvmlDeviceGetBAR1MemoryInfo(nvmlDevice_t device, nvmlBAR1Memory_t* bar1Memory) {
	if(nvmlDeviceGetVbiosVersionFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	return nvmlDeviceGetBAR1MemoryInfoFunc(device, bar1Memory);
}

nvmlReturn_t (*nvmlDeviceGetMinorNumberFunc)(nvmlDevice_t device, unsigned int *minorNumber);
nvmlReturn_t nvmlDeviceGetMinorNumber(nvmlDevice_t device, unsigned int *minorNumber) {
  if (nvmlDeviceGetMinorNumberFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetMinorNumberFunc(device, minorNumber);
}

nvmlReturn_t (*nvmlDeviceGetUUIDFunc)(nvmlDevice_t device, char *uuid, unsigned int length);
nvmlReturn_t nvmlDeviceGetUUID(nvmlDevice_t device, char *uuid, unsigned int length) {
  if (nvmlDeviceGetUUIDFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetUUIDFunc(device, uuid, length);
}

nvmlReturn_t (*nvmlDeviceGetSerialFunc)(nvmlDevice_t device, char* serial, unsigned int  length);
nvmlReturn_t nvmlDeviceGetSerial(nvmlDevice_t device, char* serial, unsigned int  length){
	if (nvmlDeviceGetSerialFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	return nvmlDeviceGetSerialFunc(device, serial, length);
}

nvmlReturn_t (*nvmlDeviceGetNameFunc)(nvmlDevice_t device, char *name, unsigned int length);
nvmlReturn_t nvmlDeviceGetName(nvmlDevice_t device, char *name, unsigned int length) {
  if (nvmlDeviceGetNameFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetNameFunc(device, name, length);
}

nvmlReturn_t (*nvmlDeviceGetBrandFunc)(nvmlDevice_t device, nvmlBrandType_t *type);
nvmlReturn_t nvmlDeviceGetBrand(nvmlDevice_t device, nvmlBrandType_t *type) {
  if (nvmlDeviceGetBrandFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetBrandFunc(device, type);
}

nvmlReturn_t (*nvmlDeviceGetPersistenceModeFunc)(nvmlDevice_t device, nvmlEnableState_t *state);
nvmlReturn_t nvmlDeviceGetPersistenceMode(nvmlDevice_t device, nvmlEnableState_t *state) {
  if (nvmlDeviceGetPersistenceModeFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetPersistenceModeFunc(device, state);
}

nvmlReturn_t (*nvmlDeviceSetPersistenceModeFunc)(nvmlDevice_t device, nvmlEnableState_t state);
nvmlReturn_t nvmlDeviceSetPersistenceMode(nvmlDevice_t device, nvmlEnableState_t state) {
  if (nvmlDeviceSetPersistenceModeFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceSetPersistenceModeFunc(device, state);
}

nvmlReturn_t (*nvmlDeviceGetComputeModeFunc)(nvmlDevice_t device, nvmlComputeMode_t *mode);
nvmlReturn_t nvmlDeviceGetComputeMode(nvmlDevice_t device, nvmlComputeMode_t *mode) {
  if (nvmlDeviceGetComputeModeFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetComputeModeFunc(device, mode);
}

nvmlReturn_t (*nvmlDeviceSetComputeModeFunc)(nvmlDevice_t device, nvmlComputeMode_t mode);
nvmlReturn_t nvmlDeviceSetComputeMode(nvmlDevice_t device, nvmlComputeMode_t mode) {
  if (nvmlDeviceSetComputeModeFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceSetComputeModeFunc(device, mode);
}

nvmlReturn_t (*nvmlDeviceGetPerformanceStateFunc)(nvmlDevice_t device, nvmlPstates_t *pstate);
nvmlReturn_t nvmlDeviceGetPerformanceState(nvmlDevice_t device, nvmlPstates_t *pstate) {
  if (nvmlDeviceGetPerformanceStateFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetPerformanceStateFunc(device, pstate);
}

nvmlReturn_t (*nvmlDeviceGetClockInfoFunc)(nvmlDevice_t device, nvmlClockType_t clockType, unsigned int* clockMHz);
nvmlReturn_t nvmlDeviceGetClockInfo(nvmlDevice_t device, nvmlClockType_t clockType, unsigned int* clockMHz) {
  if (nvmlDeviceGetClockInfoFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetClockInfoFunc(device, clockType, clockMHz);
}

nvmlReturn_t (*nvmlDeviceGetMaxClockInfoFunc)(nvmlDevice_t device, nvmlClockType_t clockType, unsigned int* clockMHz);
nvmlReturn_t nvmlDeviceGetMaxClockInfo(nvmlDevice_t device, nvmlClockType_t clockType, unsigned int* clockMHz) {
  if (nvmlDeviceGetMaxClockInfoFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetMaxClockInfoFunc(device, clockType, clockMHz);
}

nvmlReturn_t (*nvmlDeviceGetMemoryInfoFunc)(nvmlDevice_t device, nvmlMemory_t *memory);
nvmlReturn_t nvmlDeviceGetMemoryInfo(nvmlDevice_t device, nvmlMemory_t *memory) {
  if (nvmlDeviceGetMemoryInfoFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetMemoryInfoFunc(device, memory);
}

nvmlReturn_t (*nvmlDeviceGetUtilizationRatesFunc)(nvmlDevice_t device, nvmlUtilization_t *utilization);
nvmlReturn_t nvmlDeviceGetUtilizationRates(nvmlDevice_t device, nvmlUtilization_t *utilization) {
  if (nvmlDeviceGetUtilizationRatesFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetUtilizationRatesFunc(device, utilization);
}

nvmlReturn_t (*nvmlDeviceGetPowerUsageFunc)(nvmlDevice_t device, unsigned int *power);
nvmlReturn_t nvmlDeviceGetPowerUsage(nvmlDevice_t device, unsigned int *power) {
  if (nvmlDeviceGetPowerUsageFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetPowerUsageFunc(device, power);
}

nvmlReturn_t (*nvmlDeviceGetPowerManagementLimitConstraintsFunc)(nvmlDevice_t device, unsigned int *minLimit, unsigned int *maxLimit);
nvmlReturn_t nvmlDeviceGetPowerManagementLimitConstraints(nvmlDevice_t device, unsigned int *minLimit, unsigned int *maxLimit) {
  if (nvmlDeviceGetPowerManagementLimitConstraintsFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetPowerManagementLimitConstraintsFunc(device, minLimit, maxLimit);
}

nvmlReturn_t (*nvmlDeviceGetPowerManagementLimitFunc)(nvmlDevice_t device, unsigned int *limit);
nvmlReturn_t nvmlDeviceGetPowerManagementLimit(nvmlDevice_t device, unsigned int *limit) {
  if (nvmlDeviceGetPowerManagementLimitFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetPowerManagementLimitFunc(device, limit);
}

nvmlReturn_t (*nvmlDeviceGetPowerManagementDefaultLimitFunc)(nvmlDevice_t device, unsigned int *defaultLimit);
nvmlReturn_t nvmlDeviceGetPowerManagementDefaultLimit(nvmlDevice_t device, unsigned int *defaultLimit) {
  if (nvmlDeviceGetPowerManagementDefaultLimitFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetPowerManagementDefaultLimitFunc(device, defaultLimit);
}

nvmlReturn_t (*nvmlDeviceGetPcieThroughputFunc)(nvmlDevice_t device, nvmlPcieUtilCounter_t counter, unsigned int *value);
nvmlReturn_t nvmlDeviceGetPcieThroughput(nvmlDevice_t device, nvmlPcieUtilCounter_t counter, unsigned int *value) {
  if (nvmlDeviceGetPcieThroughputFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetPcieThroughputFunc(device, counter, value);
}

nvmlReturn_t (*nvmlDeviceGetCurrPcieLinkGenerationFunc)(nvmlDevice_t device, unsigned int *currLinkGen);
nvmlReturn_t nvmlDeviceGetCurrPcieLinkGeneration(nvmlDevice_t device, unsigned int *currLinkGen) {
  if (nvmlDeviceGetCurrPcieLinkGenerationFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetCurrPcieLinkGenerationFunc(device, currLinkGen);
}

nvmlReturn_t (*nvmlDeviceGetCurrPcieLinkWidthFunc)(nvmlDevice_t device, unsigned int *currLinkWidth);
nvmlReturn_t nvmlDeviceGetCurrPcieLinkWidth(nvmlDevice_t device, unsigned int *currLinkWidth) {
  if (nvmlDeviceGetCurrPcieLinkWidthFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetCurrPcieLinkWidthFunc(device, currLinkWidth);
}

nvmlReturn_t (*nvmlDeviceGetMaxPcieLinkGenerationFunc)(nvmlDevice_t device, unsigned int *maxLinkGen);
nvmlReturn_t nvmlDeviceGetMaxPcieLinkGeneration(nvmlDevice_t device, unsigned int *maxLinkGen) {
  if (nvmlDeviceGetMaxPcieLinkGenerationFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetMaxPcieLinkGenerationFunc(device, maxLinkGen);
}

nvmlReturn_t (*nvmlDeviceGetMaxPcieLinkWidthFunc)(nvmlDevice_t device, unsigned int *maxLinkWidth);
nvmlReturn_t nvmlDeviceGetMaxPcieLinkWidth(nvmlDevice_t device, unsigned int *maxLinkWidth) {
  if (nvmlDeviceGetMaxPcieLinkWidthFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetMaxPcieLinkWidthFunc(device, maxLinkWidth);
}

nvmlReturn_t (*nvmlDeviceGetEnforcedPowerLimitFunc)(nvmlDevice_t device, unsigned int* limit);
nvmlReturn_t nvmlDeviceGetEnforcedPowerLimit(nvmlDevice_t device, unsigned int* limit){
	if (nvmlDeviceGetEnforcedPowerLimitFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	return nvmlDeviceGetEnforcedPowerLimitFunc(device, limit);
}

nvmlReturn_t (*nvmlDeviceGetTemperatureFunc)(nvmlDevice_t device, nvmlTemperatureSensors_t sensorType, unsigned int *temp);
nvmlReturn_t nvmlDeviceGetTemperature(nvmlDevice_t device, nvmlTemperatureSensors_t sensorType, unsigned int *temp) {
  if (nvmlDeviceGetTemperatureFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetTemperatureFunc(device, sensorType, temp);
}

nvmlReturn_t (*nvmlDeviceGetTemperatureThresholdFunc)(nvmlDevice_t device, nvmlTemperatureThresholds_t thresholdType, unsigned int* temp);
nvmlReturn_t nvmlDeviceGetTemperatureThreshold(nvmlDevice_t device, nvmlTemperatureThresholds_t thresholdType, unsigned int* temp) {
  if (nvmlDeviceGetTemperatureThresholdFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetTemperatureThresholdFunc(device, thresholdType, temp);
}

nvmlReturn_t (*nvmlDeviceGetFanSpeedFunc)(nvmlDevice_t device, unsigned int *speed);
nvmlReturn_t nvmlDeviceGetFanSpeed(nvmlDevice_t device, unsigned int *speed) {
  if (nvmlDeviceGetFanSpeedFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetFanSpeedFunc(device, speed);
}

nvmlReturn_t (*nvmlDeviceGetEncoderUtilizationFunc)(nvmlDevice_t device, unsigned int* utilization, unsigned int* samplingPeriodUs);
nvmlReturn_t nvmlDeviceGetEncoderUtilization(nvmlDevice_t device, unsigned int* utilization, unsigned int* samplingPeriodUs) {
  if (nvmlDeviceGetEncoderUtilizationFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetEncoderUtilizationFunc(device, utilization, samplingPeriodUs);
}

nvmlReturn_t (*nvmlDeviceGetDecoderUtilizationFunc)(nvmlDevice_t device, unsigned int* utilization, unsigned int* samplingPeriodUs);
nvmlReturn_t nvmlDeviceGetDecoderUtilization(nvmlDevice_t device, unsigned int* utilization, unsigned int* samplingPeriodUs) {
  if (nvmlDeviceGetDecoderUtilizationFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetDecoderUtilizationFunc(device, utilization, samplingPeriodUs);
}

nvmlReturn_t (*nvmlSystemGetProcessNameFunc)(unsigned int pid, char *name, unsigned int length);
nvmlReturn_t nvmlSystemGetProcessName(unsigned int pid, char *name, unsigned int length) {
  if (nvmlSystemGetProcessNameFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlSystemGetProcessNameFunc(pid, name, length);
}

nvmlReturn_t (*nvmlDeviceGetAccountingModeFunc)(nvmlDevice_t device, nvmlEnableState_t *mode);
nvmlReturn_t nvmlDeviceGetAccountingMode(nvmlDevice_t device, nvmlEnableState_t *mode) {
  if (nvmlDeviceGetAccountingModeFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetAccountingModeFunc(device, mode);
}

nvmlReturn_t (*nvmlDeviceGetAccountingStatsFunc)(nvmlDevice_t device, unsigned int pid, nvmlAccountingStats_t *stats);
nvmlReturn_t nvmlDeviceGetAccountingStats(nvmlDevice_t device, unsigned int pid, nvmlAccountingStats_t *stats) {
  if (nvmlDeviceGetAccountingStatsFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetAccountingStatsFunc(device, pid, stats);
}


nvmlReturn_t (*nvmlDeviceGetAccountingPidsFunc)(nvmlDevice_t device, unsigned int *count, unsigned int *pids);
nvmlReturn_t nvmlDeviceGetAccountingPids(nvmlDevice_t device, unsigned int *count, unsigned int *pids) {
  if (nvmlDeviceGetAccountingPidsFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetAccountingPidsFunc(device, count, pids);
}

nvmlReturn_t (*nvmlDeviceGetAccountingBufferSizeFunc)(nvmlDevice_t device, unsigned int* bufferSize);
nvmlReturn_t nvmlDeviceGetAccountingBufferSize(nvmlDevice_t device, unsigned int* bufferSize) {
  if (nvmlDeviceGetAccountingBufferSizeFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetAccountingBufferSizeFunc(device, bufferSize);
}

nvmlReturn_t (*nvmlDeviceGetProcessUtilizationFunc)(nvmlDevice_t device, nvmlProcessUtilizationSample_t *utilization,
                                              unsigned int *processSamplesCount, unsigned long long lastSeenTimeStamp);
nvmlReturn_t nvmlDeviceGetProcessUtilization(nvmlDevice_t device, nvmlProcessUtilizationSample_t *utilization,
                                              unsigned int *processSamplesCount, unsigned long long lastSeenTimeStamp){
  if (nvmlDeviceGetProcessUtilizationFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetProcessUtilizationFunc(device, utilization, processSamplesCount, lastSeenTimeStamp);
}

nvmlReturn_t (*nvmlDeviceGetPciInfoFunc)(nvmlDevice_t device, nvmlPciInfo_t* pci);
nvmlReturn_t nvmlDeviceGetPciInfo(nvmlDevice_t device, nvmlPciInfo_t* pci) {
  if (nvmlDeviceGetPciInfoFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  return nvmlDeviceGetPciInfoFunc(device, pci);
}

nvmlReturn_t (*nvmlDeviceGetApplicationsClockFunc)(nvmlDevice_t device, nvmlClockType_t clockType, unsigned int* clockMHz);
nvmlReturn_t nvmlDeviceGetApplicationsClock(nvmlDevice_t device, nvmlClockType_t clockType, unsigned int* clockMHz){
	if (nvmlDeviceGetApplicationsClockFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	return nvmlDeviceGetApplicationsClockFunc(device, clockType, clockMHz);
}

nvmlReturn_t (*nvmlDeviceGetComputeRunningProcessesFunc)(nvmlDevice_t device, unsigned int* infoCount, nvmlProcessInfo_t* infos);
nvmlReturn_t nvmlDeviceGetComputeRunningProcesses(nvmlDevice_t device, unsigned int* infoCount, nvmlProcessInfo_t* infos){
	if (nvmlDeviceGetComputeRunningProcessesFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	return nvmlDeviceGetComputeRunningProcessesFunc(device, infoCount, infos);
}

nvmlReturn_t (*nvmlDeviceGetGraphicsRunningProcessesFunc)(nvmlDevice_t device, unsigned int* infoCount, nvmlProcessInfo_t* infos);
nvmlReturn_t nvmlDeviceGetGraphicsRunningProcesses(nvmlDevice_t device, unsigned int* infoCount, nvmlProcessInfo_t* infos){
	if (nvmlDeviceGetGraphicsRunningProcessesFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	return nvmlDeviceGetGraphicsRunningProcessesFunc(device, infoCount, infos);
}

nvmlReturn_t (*nvmlDeviceGetSamplesFunc)(nvmlDevice_t device, nvmlSamplingType_t type, unsigned long long lastSeenTimeStamp, nvmlValueType_t *sampleValType, unsigned int *sampleCount, nvmlSample_t *samples);

// Loads the "libnvidia-ml.so.1" shared library.
// Loads all symbols needed and initializes NVML.
// Call this before calling any other methods.
nvmlReturn_t nvmlInit_dl(void) {
  nvmlHandle = dlopen("libnvidia-ml.so.1", RTLD_LAZY);
  if (nvmlHandle == NULL) {
    return NVML_ERROR_LIBRARY_NOT_FOUND;
  }
  nvmlInitFunc = dlsym(nvmlHandle, "nvmlInit_v2");
  if (nvmlInitFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlShutdownFunc = dlsym(nvmlHandle, "nvmlShutdown");
  if (nvmlShutdownFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlErrorStringFunc = dlsym(nvmlHandle, "nvmlErrorString");
  if (nvmlErrorStringFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetHandleByIndexFunc = dlsym(nvmlHandle, "nvmlDeviceGetHandleByIndex_v2");
  if (nvmlDeviceGetHandleByIndexFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
	nvmlDeviceGetIndexFunc = dlsym(nvmlHandle, "nvmlDeviceGetIndex");
	if (nvmlDeviceGetIndexFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlSystemGetDriverVersionFunc = dlsym(nvmlHandle, "nvmlSystemGetDriverVersion");
  if (nvmlSystemGetDriverVersionFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
	nvmlSystemGetNVMLVersionFunc = dlsym(nvmlHandle, "nvmlSystemGetNVMLVersion");
	if(nvmlSystemGetNVMLVersionFunc == NULL){
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
  nvmlDeviceGetCountFunc = dlsym(nvmlHandle, "nvmlDeviceGetCount_v2");
  if (nvmlDeviceGetCountFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
	nvmlDeviceGetBrandFunc = dlsym(nvmlHandle, "nvmlDeviceGetBrand");
	if(nvmlDeviceGetBrandFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	nvmlDeviceGetBoardIdFunc = dlsym(nvmlHandle, "nvmlDeviceGetBoardId");
	if(nvmlDeviceGetBoardIdFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	nvmlDeviceGetComputeModeFunc = dlsym(nvmlHandle, "nvmlDeviceGetComputeMode");
	if(nvmlDeviceGetComputeModeFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	nvmlDeviceGetDisplayModeFunc = dlsym(nvmlHandle, "nvmlDeviceGetDisplayMode");
	if(nvmlDeviceGetDisplayModeFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	nvmlDeviceGetDisplayActiveFunc = dlsym(nvmlHandle, "nvmlDeviceGetDisplayActive");
	if(nvmlDeviceGetDisplayActiveFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	nvmlDeviceGetVbiosVersionFunc = dlsym(nvmlHandle, "nvmlDeviceGetVbiosVersion");
	if(nvmlDeviceGetDisplayActiveFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	nvmlDeviceGetCurrentClocksThrottleReasonsFunc = dlsym(nvmlHandle, "nvmlDeviceGetCurrentClocksThrottleReasons");
	if(nvmlDeviceGetCurrentClocksThrottleReasonsFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	nvmlDeviceGetSerialFunc = dlsym(nvmlHandle, "nvmlDeviceGetSerial");
	if(nvmlDeviceGetSerialFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	nvmlDeviceGetBAR1MemoryInfoFunc = dlsym(nvmlHandle, "nvmlDeviceGetBAR1MemoryInfo");
	if(nvmlDeviceGetBAR1MemoryInfoFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	nvmlDeviceGetPcieThroughputFunc = dlsym(nvmlHandle, "nvmlDeviceGetPcieThroughput");
	if(nvmlDeviceGetPcieThroughputFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	nvmlDeviceGetCurrPcieLinkGenerationFunc = dlsym(nvmlHandle, "nvmlDeviceGetCurrPcieLinkGeneration");
	if(nvmlDeviceGetCurrPcieLinkGenerationFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	nvmlDeviceGetMaxPcieLinkGenerationFunc = dlsym(nvmlHandle, "nvmlDeviceGetMaxPcieLinkGeneration");
	if(nvmlDeviceGetMaxPcieLinkGenerationFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	nvmlDeviceGetMaxPcieLinkWidthFunc = dlsym(nvmlHandle, "nvmlDeviceGetMaxPcieLinkWidth");
	if(nvmlDeviceGetMaxPcieLinkWidthFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
	nvmlDeviceGetCurrPcieLinkWidthFunc = dlsym(nvmlHandle, "nvmlDeviceGetCurrPcieLinkWidth");
	if(nvmlDeviceGetCurrPcieLinkWidthFunc == NULL) {
		return NVML_ERROR_FUNCTION_NOT_FOUND;
	}
  nvmlDeviceGetMinorNumberFunc = dlsym(nvmlHandle, "nvmlDeviceGetMinorNumber");
  if (nvmlDeviceGetMinorNumberFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetUUIDFunc = dlsym(nvmlHandle, "nvmlDeviceGetUUID");
  if (nvmlDeviceGetUUIDFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetNameFunc = dlsym(nvmlHandle, "nvmlDeviceGetName");
  if (nvmlDeviceGetNameFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetBrandFunc = dlsym(nvmlHandle, "nvmlDeviceGetBrand");
  if (nvmlDeviceGetBrandFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetPersistenceModeFunc = dlsym(nvmlHandle, "nvmlDeviceGetPersistenceMode");
  if (nvmlDeviceGetPersistenceModeFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceSetPersistenceModeFunc = dlsym(nvmlHandle, "nvmlDeviceSetPersistenceMode");
  if (nvmlDeviceSetPersistenceModeFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetComputeModeFunc = dlsym(nvmlHandle, "nvmlDeviceGetComputeMode");
  if (nvmlDeviceGetComputeModeFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceSetComputeModeFunc = dlsym(nvmlHandle, "nvmlDeviceSetComputeMode");
  if (nvmlDeviceSetComputeModeFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetPerformanceStateFunc = dlsym(nvmlHandle, "nvmlDeviceGetPerformanceState");
  if (nvmlDeviceGetPerformanceStateFunc== NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetClockInfoFunc = dlsym(nvmlHandle, "nvmlDeviceGetClockInfo");
  if (nvmlDeviceGetClockInfoFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetMaxClockInfoFunc = dlsym(nvmlHandle, "nvmlDeviceGetMaxClockInfo");
  if (nvmlDeviceGetMaxClockInfoFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetMemoryInfoFunc = dlsym(nvmlHandle, "nvmlDeviceGetMemoryInfo");
  if (nvmlDeviceGetMemoryInfoFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetUtilizationRatesFunc = dlsym(nvmlHandle, "nvmlDeviceGetUtilizationRates");
  if (nvmlDeviceGetUtilizationRatesFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetPowerUsageFunc = dlsym(nvmlHandle, "nvmlDeviceGetPowerUsage");
  if (nvmlDeviceGetPowerUsageFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetPowerManagementLimitConstraintsFunc = dlsym(nvmlHandle, "nvmlDeviceGetPowerManagementLimitConstraints");
  if (nvmlDeviceGetPowerManagementLimitConstraintsFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetPowerManagementDefaultLimitFunc = dlsym(nvmlHandle, "nvmlDeviceGetPowerManagementDefaultLimit");
  if (nvmlDeviceGetPowerManagementDefaultLimitFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetPcieThroughputFunc = dlsym(nvmlHandle, "nvmlDeviceGetPcieThroughput");
  if (nvmlDeviceGetPcieThroughputFunc== NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetCurrPcieLinkGenerationFunc = dlsym(nvmlHandle, "nvmlDeviceGetCurrPcieLinkGeneration");
  if (nvmlDeviceGetCurrPcieLinkGenerationFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetCurrPcieLinkWidthFunc = dlsym(nvmlHandle, "nvmlDeviceGetCurrPcieLinkWidth");
  if (nvmlDeviceGetCurrPcieLinkWidthFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetMaxPcieLinkGenerationFunc = dlsym(nvmlHandle, "nvmlDeviceGetMaxPcieLinkGeneration");
  if (nvmlDeviceGetMaxPcieLinkGenerationFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetMaxPcieLinkWidthFunc = dlsym(nvmlHandle, "nvmlDeviceGetMaxPcieLinkWidth");
  if (nvmlDeviceGetMaxPcieLinkWidthFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetPowerManagementLimitFunc = dlsym(nvmlHandle, "nvmlDeviceGetPowerManagementLimit");
	if (nvmlDeviceGetPowerManagementLimitFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetEnforcedPowerLimitFunc = dlsym(nvmlHandle, "nvmlDeviceGetEnforcedPowerLimit");
	if (nvmlDeviceGetEnforcedPowerLimitFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetTemperatureFunc = dlsym(nvmlHandle, "nvmlDeviceGetTemperature");
  if (nvmlDeviceGetTemperatureFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
	nvmlDeviceGetTemperatureThresholdFunc = dlsym(nvmlHandle, "nvmlDeviceGetTemperatureThreshold");
	if (nvmlDeviceGetTemperatureThresholdFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetFanSpeedFunc = dlsym(nvmlHandle, "nvmlDeviceGetFanSpeed");
  if (nvmlDeviceGetFanSpeedFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetSamplesFunc = dlsym(nvmlHandle, "nvmlDeviceGetSamples");
  if (nvmlDeviceGetSamplesFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetEncoderUtilizationFunc = dlsym(nvmlHandle, "nvmlDeviceGetEncoderUtilization");
  if (nvmlDeviceGetEncoderUtilizationFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetDecoderUtilizationFunc = dlsym(nvmlHandle, "nvmlDeviceGetDecoderUtilization");
  if (nvmlDeviceGetDecoderUtilizationFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlSystemGetProcessNameFunc = dlsym(nvmlHandle, "nvmlSystemGetProcessName");
  if (nvmlSystemGetProcessNameFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetAccountingModeFunc = dlsym(nvmlHandle, "nvmlDeviceGetAccountingMode");
  if (nvmlDeviceGetAccountingModeFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetAccountingStatsFunc = dlsym(nvmlHandle, "nvmlDeviceGetAccountingStats");
  if (nvmlDeviceGetAccountingStatsFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetAccountingPidsFunc = dlsym(nvmlHandle, "nvmlDeviceGetAccountingPids");
  if (nvmlDeviceGetAccountingPidsFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetAccountingBufferSizeFunc = dlsym(nvmlHandle, "nvmlDeviceGetAccountingBufferSize");
  if (nvmlDeviceGetAccountingBufferSizeFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetProcessUtilizationFunc = dlsym(nvmlHandle, "nvmlDeviceGetProcessUtilization");
  if (nvmlDeviceGetProcessUtilizationFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetPciInfoFunc = dlsym(nvmlHandle, "nvmlDeviceGetPciInfo");
  if (nvmlDeviceGetPciInfoFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetApplicationsClockFunc = dlsym(nvmlHandle, "nvmlDeviceGetApplicationsClock");
	if (nvmlDeviceGetApplicationsClockFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetClockInfoFunc = dlsym(nvmlHandle, "nvmlDeviceGetClockInfo");
  if (nvmlDeviceGetClockInfoFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetComputeRunningProcessesFunc = dlsym(nvmlHandle, "nvmlDeviceGetComputeRunningProcesses");
	if (nvmlDeviceGetComputeRunningProcessesFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlDeviceGetGraphicsRunningProcessesFunc = dlsym(nvmlHandle, "nvmlDeviceGetGraphicsRunningProcesses");
	if (nvmlDeviceGetGraphicsRunningProcessesFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlReturn_t result = nvmlInitFunc();
  if (result != NVML_SUCCESS) {
    dlclose(nvmlHandle);
    nvmlHandle = NULL;
    return result;
  }
  return NVML_SUCCESS;
}

// Shuts down NVML and decrements the reference count on the dynamically loaded
// "libnvidia-ml.so.1" library.
// Call this once NVML is no longer being used.
nvmlReturn_t nvmlShutdown_dl(void) {
  if (nvmlHandle == NULL) {
    return NVML_SUCCESS;
  }
  if (nvmlShutdownFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }
  nvmlReturn_t r = nvmlShutdownFunc();
  if (r != NVML_SUCCESS) {
    return r;
  }
  return (dlclose(nvmlHandle) ? NVML_ERROR_UNKNOWN : NVML_SUCCESS);
}

// This function is here because the API provided by NVML is not very user
// friendly. This function can be used to get average utilization.gpu or
// power.draw.
//
// `device`: The identifier of the target device.
// `type`: Type of sampling event. Only NVML_TOTAL_POWER_SAMPLES and NVML_GPU_UTILIZATION_SAMPLES are supported.
// `lastSeenTimeStamp`: Return average using samples with timestamp greather than this timestamp. Unix epoch in micro seconds.
// `averageUsage`: Reference in which average is returned.
//
// In my experiments, I found that NVML_GPU_UTILIZATION_SAMPLES buffer stores
// 100 samples that are uniformly spread with ~6 samples per second. So the
// buffer stores last ~16s of data.
// NVML_TOTAL_POWER_SAMPLES buffer stores 120 samples, but in different runs I
// noticed them to be non-uniformly separated. Sometimes 120 samples only
// consisted of 10s of data and sometimes they were spread over 60s.
//
nvmlReturn_t nvmlDeviceGetAverageUsage(nvmlDevice_t device, nvmlSamplingType_t type, unsigned long long lastSeenTimeStamp, unsigned int* averageUsage) {
  if (nvmlHandle == NULL) {
    return NVML_ERROR_LIBRARY_NOT_FOUND;
  }
  if (nvmlDeviceGetSamplesFunc == NULL) {
    return NVML_ERROR_FUNCTION_NOT_FOUND;
  }

  // We don't really use this because both the metrics we support
  // averagePowerUsage and averageGPUUtilization are unsigned int.
  nvmlValueType_t sampleValType;

  // This will be set to the number of samples that can be queried. We would
  // need to allocate an array of this size to store the samples.
  unsigned int sampleCount;

  // Invoking this method with `samples` set to NULL sets the sampleCount.
  nvmlReturn_t r = nvmlDeviceGetSamplesFunc(device, type, lastSeenTimeStamp, &sampleValType, &sampleCount, NULL);
  if (r != NVML_SUCCESS) {
    return r;
  }

  // Allocate memory to store sampleCount samples.
  // In my experiments, the sampleCount at this stage was always 120 for
  // NVML_TOTAL_POWER_SAMPLES and 100 for NVML_GPU_UTILIZATION_SAMPLES
  nvmlSample_t* samples = (nvmlSample_t*) malloc(sampleCount * sizeof(nvmlSample_t));

  r = nvmlDeviceGetSamplesFunc(device, type, lastSeenTimeStamp, &sampleValType, &sampleCount, samples);
  if (r != NVML_SUCCESS) {
    free(samples);
    return r;
  }

  int i = 0;
  unsigned int sum = 0;
  for (; i < sampleCount; i++) {
    sum += samples[i].sampleValue.uiVal;
  }
  *averageUsage = sum/sampleCount;

  free(samples);
  return r;
}
*/
import "C"

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	szDriver        = C.NVML_SYSTEM_DRIVER_VERSION_BUFFER_SIZE
	szName          = C.NVML_DEVICE_NAME_BUFFER_SIZE
	szUUID          = C.NVML_DEVICE_UUID_BUFFER_SIZE
	ctGraphics      = C.NVML_CLOCK_GRAPHICS
	ctSM            = C.NVML_CLOCK_SM
	ctMemory        = C.NVML_CLOCK_MEM
	ctVideo         = C.NVML_CLOCK_VIDEO
	pcieUtilTx      = C.NVML_PCIE_UTIL_TX_BYTES
	pcieUtilRx      = C.NVML_PCIE_UTIL_RX_BYTES
	szNVML          = C.NVML_SYSTEM_NVML_VERSION_BUFFER_SIZE
	szVBiosVersion  = C.NVML_DEVICE_VBIOS_VERSION_BUFFER_SIZE
	szDeviceSerial  = C.NVML_DEVICE_SERIAL_BUFFER_SIZE
)

var errLibraryNotLoaded = errors.New("could not load NVML library")

// Initialize initializes NVML.
// Call this before calling any other methods.
func Initialize() error {
	return errorString(C.nvmlInit_dl())
}

// Shutdown shuts down NVML.
// Call this once NVML is no longer being used.
func Shutdown() error {
	return errorString(C.nvmlShutdown_dl())
}

// errorString takes a nvmlReturn_t and converts it into a golang error.
// It uses a nvml method to convert to a user friendly error message.
func errorString(ret C.nvmlReturn_t) error {
	if ret == C.NVML_SUCCESS {
		return nil
	}
	// We need to special case this because if nvml library is not found
	// nvmlErrorString() method will not work.
	if ret == C.NVML_ERROR_LIBRARY_NOT_FOUND || C.nvmlHandle == nil {
		return errLibraryNotLoaded
	}
	err := C.GoString(C.nvmlErrorString(ret))
	return fmt.Errorf("NVML: %v", err)
}

// SystemDriverVersion returns the the driver version on the system.
func SystemDriverVersion() (string, error) {
	if C.nvmlHandle == nil {
		return "", errLibraryNotLoaded
	}
	var driver [szDriver]C.char
	r := C.nvmlSystemGetDriverVersion(&driver[0], szDriver)
	return C.GoString(&driver[0]), errorString(r)
}

// SystemNVMLVersion returns the NVML Library Version being used.
func SystemNVMLVersion() (string, error) {
	if C.nvmlHandle == nil {
		return "", errLibraryNotLoaded
	}
	var nvml [szNVML]C.char
	r := C.nvmlSystemGetNVMLVersion(&nvml[0], szNVML)
	return C.GoString(&nvml[0]), errorString(r)
}

// DeviceCount returns the number of nvidia devices on the system.
func DeviceCount() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var n C.uint
	r := C.nvmlDeviceGetCount(&n)
	return uint(n), errorString(r)
}

// Device is the handle for the device.
// This handle is obtained by calling DeviceHandleByIndex().
type Device struct {
	dev C.nvmlDevice_t
}

// Utilization is Structure to store utilization value and process Id
type Utilization struct {
	Pid       uint   //!< PID of process
	timeStamp uint64 //!< CPU Timestamp in microseconds
	SMUtil    uint   //!< SM (3D/Compute) Util Value
	MemUtil   uint   //!< Frame Buffer Memory Util Value
	EncUtil   uint   //!< Encoder Util Value
	DecUtil   uint   //!< Decoder Util Value
}

// AccountingStats is a Structire to Store accounting Stats for every process
type AccountingStats struct {
	GPUUtilization uint
	//!< Percent of time over the process's lifetime during which one or more kernels was executing on the GPU.
	//! Utilization stats just like returned by \ref nvmlDeviceGetUtilizationRates but for the life time of a
	//! process (not just the last sample period).
	//! Set to NVML_VALUE_NOT_AVAILABLE if nvmlDeviceGetUtilizationRates is not supported

	MemoryUtilization uint
	//!< Percent of time over the process's lifetime during which global (device) memory was being read or written.
	//! Set to NVML_VALUE_NOT_AVAILABLE if nvmlDeviceGetUtilizationRates is not supported

	MaxMemoryUsage uint64
	//!< Maximum total memory in bytes that was ever allocated by the process.
	//! Set to NVML_VALUE_NOT_AVAILABLE if nvmlProcessInfo_t->usedGpuMemory is not supported

	Time uint64
	//!< Amount of time in ms during which the compute context was active. The time is reported as 0 if
	//!< the process is not terminated

	StartTime uint64
	//!< CPU Timestamp in usec representing start time for the process

	IsRunning bool
	//!< Flag to represent if the process is running (1 for running, 0 for terminated)

	Reserved [5]uint
	// Reserved for
}

// Process is the exported handle for a process instance
type Process interface {
	PID() uint
	Memory() uint64
}

type process struct {
	pid           uint
	usedGpuMemory uint64
}

func (proc process) PID() uint {
	return proc.pid
}

func (proc process) Memory() uint64 {
	return proc.usedGpuMemory
}

// DeviceBrand is the equivalent for nvmlBrandType_t.
type DeviceBrand int

// Enumeration mapping for DeviceBrand to nvmlBrandType_t
const (
	DeviceBrandUnknown DeviceBrand = C.NVML_BRAND_UNKNOWN
	DeviceBrandTesla   DeviceBrand = C.NVML_BRAND_TESLA
	DeviceBrandNVS     DeviceBrand = C.NVML_BRAND_NVS
	DeviceBrandGRID    DeviceBrand = C.NVML_BRAND_GRID
	DeviceBrandGeForce DeviceBrand = C.NVML_BRAND_GEFORCE
)

func (b DeviceBrand) String() string {
	switch b {
	case DeviceBrandTesla:
		return "Tesla"
	case DeviceBrandNVS:
		return "NVS"
	case DeviceBrandGRID:
		return "GRID"
	case DeviceBrandGeForce:
		return "Geforce"
	default:
		return "unknown"
	}
}

// ComputeMode is the quivalent for nvmlComputeMode_t.
type ComputeMode int

// Enumeration mapping for ComputeMode to nvmlComputeMode_t
const (
	ComputeModeDefault          ComputeMode = C.NVML_COMPUTEMODE_DEFAULT
	ComputeModeExclusiveThread  ComputeMode = C.NVML_COMPUTEMODE_EXCLUSIVE_THREAD
	ComputeModeProhibited       ComputeMode = C.NVML_COMPUTEMODE_PROHIBITED
	ComputeModeExclusiveProcess ComputeMode = C.NVML_COMPUTEMODE_EXCLUSIVE_PROCESS
)

func (cm ComputeMode) String() string {
	switch cm {
	case ComputeModeProhibited:
		return "prohibited"
	case ComputeModeExclusiveThread:
		return "exclusive thread"
	case ComputeModeExclusiveProcess:
		return "exclusive process"
	case ComputeModeDefault:
		return "default"
	default:
		return "unkown"
	}
}

// EnableState is the quivalent for nvmlEnableState_t.
type EnableState int

// Enumeration mapping for ComputeMode to nvmlComputeMode_t
const (
	EnableStateFeatureEnabled  EnableState = C.NVML_FEATURE_ENABLED
	EnableStateFeatureDisabled EnableState = C.NVML_FEATURE_DISABLED
)

func (es EnableState) String() string {
	switch es {
	case EnableStateFeatureEnabled:
		return "enabled"
	case EnableStateFeatureDisabled:
		return "disabled"
	default:
		return "unknown"
	}
}

// PowerState is the equivalent to nvmlPstates_t
type PowerState int

// Enumeration mapping for PowerState to nvmlPstates_t
const (
	PowerState0       PowerState = C.NVML_PSTATE_0
	PowerState1       PowerState = C.NVML_PSTATE_1
	PowerState2       PowerState = C.NVML_PSTATE_2
	PowerState3       PowerState = C.NVML_PSTATE_3
	PowerState4       PowerState = C.NVML_PSTATE_4
	PowerState5       PowerState = C.NVML_PSTATE_5
	PowerState6       PowerState = C.NVML_PSTATE_6
	PowerState7       PowerState = C.NVML_PSTATE_7
	PowerState8       PowerState = C.NVML_PSTATE_8
	PowerState9       PowerState = C.NVML_PSTATE_9
	PowerState10      PowerState = C.NVML_PSTATE_10
	PowerState11      PowerState = C.NVML_PSTATE_11
	PowerState12      PowerState = C.NVML_PSTATE_12
	PowerState13      PowerState = C.NVML_PSTATE_13
	PowerState14      PowerState = C.NVML_PSTATE_14
	PowerState15      PowerState = C.NVML_PSTATE_15
	PowerStateUnknown PowerState = C.NVML_PSTATE_UNKNOWN
)

type ThrottlingReason int
// Enumeration of reasons for throttling
const (
    ThrottlingReasonNone = 0
    ThrottlingReasonIdle = 1
    ThrottlingReasonApplicationClock = 2
    ThrottlingReasonUserDefinedClocks = 3
    ThrottlingReasonSwPowerCap = 4
    ThrottlingReasonHwSlowdown = 5
    ThrottlingReasonSyncBoost = 6
    ThrottlingReasonSwThermalSlowdown = 7
    ThrottlingReasonHwThermalSlowdown = 8
    ThrottlingReasonHwPowerBrakeSlowdown = 9
    ThrottlingReasonDisplayClockSetting = 10
)



func (ps PowerState) String() string {
	switch ps {
	case PowerState0:
		return "P0 - Maximum Performance"
	case PowerState1:
		return "P1"
	case PowerState2:
		return "P2"
	case PowerState3:
		return "P3"
	case PowerState4:
		return "P4"
	case PowerState5:
		return "P5"
	case PowerState6:
		return "P6"
	case PowerState7:
		return "P7"
	case PowerState8:
		return "P8"
	case PowerState9:
		return "P9"
	case PowerState10:
		return "P10"
	case PowerState11:
		return "P11"
	case PowerState12:
		return "P12"
	case PowerState13:
		return "P13"
	case PowerState14:
		return "P14"
	case PowerState15:
		return "P15 - Minimum Performance"
	default:
		return "unknown"
	}
}

// ClockType is the equivalent nvmlClockType_t
type ClockType int

//Enumeration mapping for ClockType to nvmlClockType_t
const (
	ClockTypeGraphics ClockType = C.NVML_CLOCK_GRAPHICS
	ClockTypeSM       ClockType = C.NVML_CLOCK_SM
	ClockTypeMem      ClockType = C.NVML_CLOCK_MEM
	ClockTypeVideo    ClockType = C.NVML_CLOCK_VIDEO
)

func (ct ClockType) String() string {
	switch ct {
	case ClockTypeGraphics:
		return "graphics"
	case ClockTypeSM:
		return "sm"
	case ClockTypeMem:
		return "memory"
	case ClockTypeVideo:
		return "video"
	default:
		return "unknown"
	}
}

// DeviceHandleByIndex returns the device handle for a particular index.
// The indices range from 0 to DeviceCount()-1. The order in which NVML
// enumerates devices has no guarantees of consistency between reboots.
func DeviceHandleByIndex(idx uint) (Device, error) {
	if C.nvmlHandle == nil {
		return Device{}, errLibraryNotLoaded
	}
	var dev C.nvmlDevice_t
	r := C.nvmlDeviceGetHandleByIndex(C.uint(idx), &dev)
	return Device{dev}, errorString(r)
}

// Index return the index of the device
func (d Device) Index() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var index C.uint
	r := C.nvmlDeviceGetIndex(d.dev, &index)
	return uint(index), errorString(r)
}

// Brand returns the Product Brand of the device.
func (d Device) Brand() (DeviceBrand, error) {
	if C.nvmlHandle == nil {
		return DeviceBrandUnknown, errLibraryNotLoaded
	}
	var brand C.nvmlBrandType_t
	r := C.nvmlDeviceGetBrand(d.dev, &brand)
	return DeviceBrand(brand), errorString(r)
}

// BoardID returns the Devices Board ID.
// Devices with the same boardId indicate GPUs connected to the same PLX.
func (d Device) BoardID() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var boardid C.uint
	r := C.nvmlDeviceGetBoardId(d.dev, &boardid)
	return uint(boardid), errorString(r)
}

// ComputeMode returns the current Compute Mode of the device.
func (d Device) ComputeMode() (ComputeMode, error) {
	if C.nvmlHandle == nil {
		return ComputeModeDefault, errLibraryNotLoaded
	}
	var cm C.nvmlComputeMode_t
	r := C.nvmlDeviceGetComputeMode(d.dev, &cm)
	return ComputeMode(cm), errorString(r)
}

// DisplayMode returns the current Display Mode of the device.
func (d Device) DisplayMode() (EnableState, error) {
	if C.nvmlHandle == nil {
		return -1, errLibraryNotLoaded
	}
	var es C.nvmlEnableState_t
	r := C.nvmlDeviceGetDisplayMode(d.dev, &es)
	return EnableState(es), errorString(r)
}

// DisplayActive returns if there currently is an active display attached.
func (d Device) DisplayActive() (EnableState, error) {
	if C.nvmlHandle == nil {
		return -1, errLibraryNotLoaded
	}
	var es C.nvmlEnableState_t
	r := C.nvmlDeviceGetDisplayActive(d.dev, &es)
	return EnableState(es), errorString(r)
}

// VBiosVersion returns VBIOS version of the device.
func (d Device) VBiosVersion() (string, error) {
	if C.nvmlHandle == nil {
		return "", errLibraryNotLoaded
	}
	var version [szVBiosVersion]C.char
	r := C.nvmlDeviceGetVbiosVersion(d.dev, &version[0], szVBiosVersion)
	return C.GoString(&version[0]), errorString(r)
}

// CurrentClocksThrottleReasons returns reasons (bitmap) for the GPU being throttled
func (d Device) CurrentClocksThrottleReasons() (uint64, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
    var bitmap C.ulonglong
	r := C.nvmlDeviceGetCurrentClocksThrottleReasons(d.dev, &bitmap)
	return uint64(bitmap), errorString(r)
}


func (d Device) MostSeriousClocksThrottleReason() (int, error) {
    var bitmap, err = d.CurrentClocksThrottleReasons ()
    if (bitmap & C.nvmlClocksThrottleReasonDisplayClockSetting) != 0 {
        return ThrottlingReasonDisplayClockSetting, err
    }
    if (bitmap & C.nvmlClocksThrottleReasonHwPowerBrakeSlowdown) != 0 {
        return ThrottlingReasonHwPowerBrakeSlowdown, err
    }
    if (bitmap & C.nvmlClocksThrottleReasonHwThermalSlowdown) != 0 {
        return ThrottlingReasonHwThermalSlowdown, err
    }
    if (bitmap & C.nvmlClocksThrottleReasonSwThermalSlowdown) != 0 {
        return ThrottlingReasonSwThermalSlowdown, err
    }
    if (bitmap & C.nvmlClocksThrottleReasonSyncBoost) != 0 {
        return ThrottlingReasonSyncBoost, err
    }
    if (bitmap & C.nvmlClocksThrottleReasonHwSlowdown) != 0 {
        return ThrottlingReasonHwSlowdown, err
    }
    if (bitmap & C.nvmlClocksThrottleReasonSwPowerCap) != 0 {
        return ThrottlingReasonSwPowerCap, err
    }
    if (bitmap & C.nvmlClocksThrottleReasonUserDefinedClocks) != 0 {
        return ThrottlingReasonUserDefinedClocks, err
    }
    if (bitmap & C.nvmlClocksThrottleReasonApplicationsClocksSetting) != 0 {
        return ThrottlingReasonApplicationClock, err
    }
    if (bitmap & C.nvmlClocksThrottleReasonGpuIdle) != 0 {
        return ThrottlingReasonIdle, err
    }
    return ThrottlingReasonNone, err
}

// Serial returns the globally unique board serial number associated with this device's board.
func (d Device) Serial() (string, error) {
	if C.nvmlHandle == nil {
		return "", errLibraryNotLoaded
	}
	var serial [szDeviceSerial]C.char
	r := C.nvmlDeviceGetSerial(d.dev, &serial[0], szDeviceSerial)
	return C.GoString(&serial[0]), errorString(r)
}

// MinorNumber returns the minor number for the device.
// The minor number for the device is such that the Nvidia device node
// file for each GPU will have the form /dev/nvidia[minor number].
func (d Device) MinorNumber() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var n C.uint
	r := C.nvmlDeviceGetMinorNumber(d.dev, &n)
	return uint(n), errorString(r)
}

// BusID returns the BusID
func (d Device) BusID() (string, error) {
	if C.nvmlHandle == nil {
		return "", errLibraryNotLoaded
	}
	var pci C.nvmlPciInfo_t

	r := C.nvmlDeviceGetPciInfo(d.dev, &pci)
	return C.GoString(&pci.busId[0]), errorString(r)
}

// UUID returns the globally unique immutable UUID associated with this device.
func (d Device) UUID() (string, error) {
	if C.nvmlHandle == nil {
		return "", errLibraryNotLoaded
	}
	var uuid [szUUID]C.char
	r := C.nvmlDeviceGetUUID(d.dev, &uuid[0], szUUID)
	return C.GoString(&uuid[0]), errorString(r)
}

// Name returns the product name of the device.
func (d Device) Name() (string, error) {
	if C.nvmlHandle == nil {
		return "", errLibraryNotLoaded
	}
	var name [szName]C.char
	r := C.nvmlDeviceGetName(d.dev, &name[0], szName)
	return C.GoString(&name[0]), errorString(r)
}

// PersistenceMode returns the current driver persistence mode of the device.
func (d Device) PersistenceMode() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var pm C.nvmlEnableState_t
	r := C.nvmlDeviceGetPersistenceMode(d.dev, &pm)
	return uint(pm), errorString(r)
}

// SetPersistenceMode sets the current driver persistence mode of the device.
func (d Device) SetPersistenceMode(mode uint) error {
	if C.nvmlHandle == nil {
		return errLibraryNotLoaded
	}
	r := C.nvmlDeviceSetPersistenceMode(d.dev, C.nvmlEnableState_t(mode))
	return errorString(r)
}

// SetComputeMode sets the current compute mode of the device.
func (d Device) SetComputeMode(mode uint) error {
	if C.nvmlHandle == nil {
		return errLibraryNotLoaded
	}
	r := C.nvmlDeviceSetComputeMode(d.dev, C.nvmlComputeMode_t(mode))
	return errorString(r)
}

// PerformanceState returns the current performance state of the device.
func (d Device) PerformanceState() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var pstate C.nvmlPstates_t
	r := C.nvmlDeviceGetPerformanceState(d.dev, &pstate)
	return uint(pstate), errorString(r)
}

// GrClock returns the application graphics clock of the device.
func (d Device) GrClock() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var clockMHz C.uint
	r := C.nvmlDeviceGetClockInfo(d.dev, ctGraphics, &clockMHz)
	return uint(clockMHz), errorString(r)
}

// SMClock returns the application SM clock of the device.
func (d Device) SMClock() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var clockMHz C.uint
	r := C.nvmlDeviceGetClockInfo(d.dev, ctSM, &clockMHz)
	return uint(clockMHz), errorString(r)
}

// MemClock returns the memory clock of the device.
func (d Device) MemClock() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var clockMHz C.uint
	r := C.nvmlDeviceGetClockInfo(d.dev, ctMemory, &clockMHz)
	return uint(clockMHz), errorString(r)
}

// VideoClock returns the memory clock of the device.
func (d Device) VideoClock() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var clockMHz C.uint
	r := C.nvmlDeviceGetClockInfo(d.dev, ctVideo, &clockMHz)
	return uint(clockMHz), errorString(r)
}

// GrMaxClock returns the application graphics clock of the device.
func (d Device) GrMaxClock() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var clockMHz C.uint
	r := C.nvmlDeviceGetMaxClockInfo(d.dev, ctGraphics, &clockMHz)
	return uint(clockMHz), errorString(r)
}

// SMMaxClock returns the application SM clock of the device.
func (d Device) SMMaxClock() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var clockMHz C.uint
	r := C.nvmlDeviceGetMaxClockInfo(d.dev, ctSM, &clockMHz)
	return uint(clockMHz), errorString(r)
}

// MemMaxClock returns the application memory clock of the device.
func (d Device) MemMaxClock() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var clockMHz C.uint
	r := C.nvmlDeviceGetMaxClockInfo(d.dev, ctMemory, &clockMHz)
	return uint(clockMHz), errorString(r)
}

// VideoClock returns the memory clock of the device.
func (d Device) VideoMaxClock() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var clockMHz C.uint
	r := C.nvmlDeviceGetMaxClockInfo(d.dev, ctVideo, &clockMHz)
	return uint(clockMHz), errorString(r)
}

// MemoryInfo returns the total and used memory (in bytes) of the device.
func (d Device) MemoryInfo() (uint64, uint64, error) {
	if C.nvmlHandle == nil {
		return 0, 0, errLibraryNotLoaded
	}
	var memory C.nvmlMemory_t
	r := C.nvmlDeviceGetMemoryInfo(d.dev, &memory)
	return uint64(memory.total), uint64(memory.used), errorString(r)
}

// Bar1MemoryInfo returns the total and used memory (in bytes) of the devices BAR1 Memory.
// BAR1 is used to map the FB (device memory) so that it can be directly accessed by
//  the CPU or by 3rd party devices (peer-to-peer on the PCIE bus).
func (d Device) Bar1MemoryInfo() (uint64, uint64, error) {
	if C.nvmlHandle == nil {
		return 0, 0, errLibraryNotLoaded
	}
	var bar1 C.nvmlBAR1Memory_t
	r := C.nvmlDeviceGetBAR1MemoryInfo(d.dev, &bar1)
	return uint64(bar1.bar1Total), uint64(bar1.bar1Used), errorString(r)
}

// UtilizationRates returns the percent of time over the past sample period during which:
// utilization.gpu: one or more kernels were executing on the GPU.
// utilization.memory: global (device) memory was being read or written.
func (d Device) UtilizationRates() (uint, uint, error) {
	if C.nvmlHandle == nil {
		return 0, 0, errLibraryNotLoaded
	}
	var utilization C.nvmlUtilization_t
	r := C.nvmlDeviceGetUtilizationRates(d.dev, &utilization)
	return uint(utilization.gpu), uint(utilization.memory), errorString(r)
}

// PowerUsage returns the power usage for this GPU and its associated circuitry
// in milliwatts. The reading is accurate to within +/- 5% of current power draw.
func (d Device) PowerUsage() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var n C.uint
	r := C.nvmlDeviceGetPowerUsage(d.dev, &n)
	return uint(n), errorString(r)
}

// PowerLimitConstraints retrieves information about possible values of power management limits on this device.
func (d Device) PowerLimitConstraints() (uint, uint, error) {
	if C.nvmlHandle == nil {
		return 0, 0, errLibraryNotLoaded
	}
	var min C.uint
	var max C.uint
	r := C.nvmlDeviceGetPowerManagementLimitConstraints(d.dev, &min, &max)
	return uint(min), uint(max), errorString(r)
}

// PowerLimits returns the devices power Limits
// management (first uint): The power limit defines the upper boundary for the
//  card's power draw. If the card's total power draw reaches this limit the
//  power management algorithm kicks in.
// enforced (second uint): Get the effective power limit that the driver
//  enforces after taking into account all limiters.
//  Note: This can be different from the management limit if
//  other limits are set elsewhere This includes the out of band power limit
//  interface
func (d Device) PowerLimits() (uint, uint, error) {
	var errors []string

	if C.nvmlHandle == nil {
		return 0, 0, errLibraryNotLoaded
	}
	var management, enforced C.uint
	r := C.nvmlDeviceGetEnforcedPowerLimit(d.dev, &enforced)
	if errorString(r) != nil {
		errors = append(errors, errorString(r).Error())
	}
	r = C.nvmlDeviceGetPowerManagementLimit(d.dev, &management)
	if errorString(r) != nil {
		errors = append(errors, errorString(r).Error())
	}

	if len(errors) > 0 {
		return uint(management), uint(enforced), fmt.Errorf(strings.Join(errors, "\n"))
	}
	return uint(management), uint(enforced), nil
}

// AveragePowerUsage returns the power usage for this GPU and its associated circuitry
// in milliwatts averaged over the samples collected in the last `since` duration.
func (d Device) AveragePowerUsage(since time.Duration) (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	lastTs := C.ulonglong(time.Now().Add(-1*since).UnixNano() / 1000)
	var n C.uint
	r := C.nvmlDeviceGetAverageUsage(d.dev, C.NVML_TOTAL_POWER_SAMPLES, lastTs, &n)
	return uint(n), errorString(r)
}

// PowerLimit returns the power limit for this GPU and its associated circuitry
// in milliwatts
func (d Device) PowerLimit() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var n C.uint
	r := C.nvmlDeviceGetPowerManagementLimit(d.dev, &n)
	return uint(n), errorString(r)
}

// DefaultPowerLimit returns the power limit for this GPU and its associated circuitry
// in milliwatts
func (d Device) DefaultPowerLimit() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var n C.uint
	r := C.nvmlDeviceGetPowerManagementDefaultLimit(d.dev, &n)
	return uint(n), errorString(r)
}

// AverageGPUUtilization returns the utilization.gpu metric (percent of time
// one of more kernels were executing on the GPU) averaged over the samples
// collected in the last `since` duration.
func (d Device) AverageGPUUtilization(since time.Duration) (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	lastTs := C.ulonglong(time.Now().Add(-1*since).UnixNano() / 1000)
	var n C.uint
	r := C.nvmlDeviceGetAverageUsage(d.dev, C.NVML_GPU_UTILIZATION_SAMPLES, lastTs, &n)
	return uint(n), errorString(r)
}

// PcieTxThroughput returns the tx throughput in KB/s
func (d Device) PcieTxThroughput() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var n C.uint
	r := C.nvmlDeviceGetPcieThroughput(d.dev, pcieUtilTx, &n)
	return uint(n), errorString(r)
}

// PcieRxThroughput returns the rx throughput in KB/s
func (d Device) PcieRxThroughput() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var n C.uint
	r := C.nvmlDeviceGetPcieThroughput(d.dev, pcieUtilRx, &n)
	return uint(n), errorString(r)
}

// PcieGeneration returns the current PCIe link generation
func (d Device) PcieGeneration() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var n C.uint
	r := C.nvmlDeviceGetCurrPcieLinkGeneration(d.dev, &n)
	return uint(n), errorString(r)
}

// PcieWidth returns the current PCIe link width
func (d Device) PcieWidth() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var n C.uint
	r := C.nvmlDeviceGetCurrPcieLinkWidth(d.dev, &n)
	return uint(n), errorString(r)
}

// PcieMaxGeneration returns the current PCIe link generation
func (d Device) PcieMaxGeneration() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var n C.uint
	r := C.nvmlDeviceGetMaxPcieLinkGeneration(d.dev, &n)
	return uint(n), errorString(r)
}

// PcieMaxWidth returns the current PCIe link width
func (d Device) PcieMaxWidth() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var n C.uint
	r := C.nvmlDeviceGetMaxPcieLinkWidth(d.dev, &n)
	return uint(n), errorString(r)
}

// Temperature returns the temperature for this GPU in Celsius.
func (d Device) Temperature() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var n C.uint
	r := C.nvmlDeviceGetTemperature(d.dev, C.NVML_TEMPERATURE_GPU, &n)
	return uint(n), errorString(r)
}

// TemperatureThresholds returns the temperature thresholds for this device in Celcius
// first return argument is the shutdown threshold, second is the slowdown threshold
func (d Device) TemperatureThresholds() (uint, uint, error) {
	var errors []string

	if C.nvmlHandle == nil {
		return 0, 0, errLibraryNotLoaded
	}

	//shutdown type
	var shutdown C.uint
	r := C.nvmlDeviceGetTemperatureThreshold(d.dev, C.NVML_TEMPERATURE_THRESHOLD_SHUTDOWN, &shutdown)
	if errorString(r) != nil {
		errors = append(errors, errorString(r).Error())
	}

	//slowdown type
	var slowdown C.uint
	r = C.nvmlDeviceGetTemperatureThreshold(d.dev, C.NVML_TEMPERATURE_THRESHOLD_SLOWDOWN, &slowdown)
	if errorString(r) != nil {
		errors = append(errors, errorString(r).Error())
	}

	if len(errors) > 0 {
		return uint(shutdown), uint(slowdown), fmt.Errorf(strings.Join(errors, "\n"))
	}
	return uint(shutdown), uint(slowdown), nil
}

// FanSpeed returns the temperature for this GPU in the percentage of its full
// speed, with 100 being the maximum.
func (d Device) FanSpeed() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var n C.uint
	r := C.nvmlDeviceGetFanSpeed(d.dev, &n)
	return uint(n), errorString(r)
}

// EncoderUtilization returns the percent of time over the last sample period during which the GPU video encoder was being used.
// The sampling period is variable and is returned in the second return argument in microseconds.
func (d Device) EncoderUtilization() (uint, uint, error) {
	if C.nvmlHandle == nil {
		return 0, 0, errLibraryNotLoaded
	}
	var n, sp C.uint
	r := C.nvmlDeviceGetEncoderUtilization(d.dev, &n, &sp)
	return uint(n), uint(sp), errorString(r)
}

// DecoderUtilization returns the percent of time over the last sample period during which the GPU video decoder was being used.
// The sampling period is variable and is returned in the second return argument in microseconds.
func (d Device) DecoderUtilization() (uint, uint, error) {
	if C.nvmlHandle == nil {
		return 0, 0, errLibraryNotLoaded
	}
	var n, sp C.uint
	r := C.nvmlDeviceGetDecoderUtilization(d.dev, &n, &sp)
	return uint(n), uint(sp), errorString(r)
}

// DeviceGetAccountingMode Queries process's accounting stats
// @return mode                                 Reference in which to return the current accounting mode
func (d Device) AccountingMode() (C.nvmlEnableState_t, error) {
	var stats C.nvmlEnableState_t
	if C.nvmlHandle == nil {
		return stats, errLibraryNotLoaded
	}
	r := C.nvmlDeviceGetAccountingMode(d.dev, &stats)
	return stats, errorString(r)
}

// DeviceGetAccountingStats Queries process's accounting stats.
// @param pid                                  Process Id of the target process to query stats for
// @return stats                               Reference in which to return the process's accounting stats
func (d Device) AccountingStats(pid uint) (*AccountingStats, error) {
	if C.nvmlHandle == nil {
		return nil, errLibraryNotLoaded
	}
	var stats C.nvmlAccountingStats_t
	r := C.nvmlDeviceGetAccountingStats(d.dev, C.uint(pid), &stats)

	accountingStats := &AccountingStats{
		GPUUtilization:    uint(stats.gpuUtilization),
		MemoryUtilization: uint(stats.memoryUtilization),
		MaxMemoryUsage:    uint64(stats.maxMemoryUsage),
		Time:              uint64(stats.time),
		StartTime:         uint64(stats.startTime),
		IsRunning:         uint(stats.isRunning) == 1,
	}

	return accountingStats, errorString(r)
}

// DeviceGetAccountingPids Queries list of processes that can be queried for accounting stats. The list of processes returned
// @param count                                Maxnum pids
// @return pids                                Pids result
// @return count                               Queried pids num
func (d Device) AccountingPids(count uint) ([]C.uint, uint, error) {
	// init pids
	cCount := C.uint(count)
	if C.nvmlHandle == nil {
		return nil, 0, errLibraryNotLoaded
	}
	if count == 0 {
		r := C.nvmlDeviceGetAccountingPids(d.dev, &cCount, nil)
		return nil, uint(cCount), errorString(r)
	}

	pids := make([]C.uint, count)
	for index := range pids {
		pids[index] = 0
	}

	r := C.nvmlDeviceGetAccountingPids(d.dev, &cCount, &pids[0])
	return pids, uint(cCount), errorString(r)
}

// DeviceGetAccountingBufferSize Returns the number of processes that the circular buffer with accounting pids can hold.
// @return buffersize                        buffersize
func (d Device) AccountingBufferSize() (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var bufferSize C.uint
	r := C.nvmlDeviceGetAccountingBufferSize(d.dev, &bufferSize)
	return uint(bufferSize), errorString(r)
}

// DeviceGetProcessUtilization Retrieves the current utilization and process ID
// @param processCount                      Maxnum process buffersize
// @param since                             The last query time for process
// @return utilizations                     The utilizations for all process
// @return processCount                     The queried utilizations
func (d Device) ProcessUtilization(processCount uint, since time.Duration) ([]*Utilization, error) {
	if C.nvmlHandle == nil {
		return nil, errLibraryNotLoaded
	}
	if processCount <= 0 {
		return nil, errors.New("Process Count Less than zero")
	}

	cUtilizations := make([]C.nvmlProcessUtilizationSample_t, processCount)
	var runningProcess C.uint = C.uint(processCount * C.sizeof_nvmlProcessUtilizationSample_t)

	lastTS := C.ulonglong(time.Now().Add(-1*since).UnixNano() / 1000)
	r := C.nvmlDeviceGetProcessUtilization(d.dev, &cUtilizations[0], &runningProcess, lastTS)
	if errorString(r) != nil {
		return nil, errorString(r)
	}

	statisticsProcess := uint(runningProcess)
	if processCount < uint(runningProcess) {
		statisticsProcess = processCount
	}

	utilizations := make([]*Utilization, statisticsProcess)
	utilCount := 0
	for _, utilization := range cUtilizations[:statisticsProcess] {
		if utilization.pid <= 0 {
			continue
		}
		u := &Utilization{
			Pid:       uint(utilization.pid),
			timeStamp: uint64(utilization.timeStamp),
			SMUtil:    uint(utilization.smUtil),
			MemUtil:   uint(utilization.memUtil),
			EncUtil:   uint(utilization.encUtil),
			DecUtil:   uint(utilization.decUtil),
		}
		utilizations[utilCount] = u
		utilCount++
	}

	return utilizations[:utilCount], errorString(r)
}

// SystemGetProcessName GetProcessName by pid
// @param pid                      Process's id
// @param buffersize               The process name's buffersize
// @return name                    Process name
func SystemGetProcessName(pid, buffersize uint) (string, error) {
	if C.nvmlHandle == nil {
		return "", errLibraryNotLoaded
	}
	c := make([]C.char, buffersize)
	r := C.nvmlSystemGetProcessName(C.uint(pid), &c[0], C.uint(buffersize))
	return C.GoString(&c[0]), errorString(r)
}

// PCIeThroughput returns the current PCIe tx and rx bytes
// first uint is tx, second is rx in KB/s
func (d Device) PCIeThroughput() (uint, uint, error) {
	var errors []string

	if C.nvmlHandle == nil {
		return 0, 0, errLibraryNotLoaded
	}
	var rx, tx C.uint
	r := C.nvmlDeviceGetPcieThroughput(d.dev, C.NVML_PCIE_UTIL_TX_BYTES, &tx)
	if errorString(r) != nil {
		errors = append(errors, "Unable to query PCIe TX utilization: "+errorString(r).Error())
	}
	r = C.nvmlDeviceGetPcieThroughput(d.dev, C.NVML_PCIE_UTIL_RX_BYTES, &rx)
	if errorString(r) != nil {
		errors = append(errors, "Unable to query PCIe RX utilization: "+errorString(r).Error())
	}

	if len(errors) > 0 {
		return uint(tx), uint(rx), fmt.Errorf(strings.Join(errors, "\n"))
	}
	return uint(tx), uint(rx), nil
}

// PCIeLinkGen returns the current PCIe Link generation
// first uint ist the current generation, second is the maximum supported generation
func (d Device) PCIeLinkGen() (uint, uint, error) {
	var errors []string

	if C.nvmlHandle == nil {
		return 0, 0, errLibraryNotLoaded
	}
	var curr, max C.uint
	r := C.nvmlDeviceGetMaxPcieLinkGeneration(d.dev, &max)
	if errorString(r) != nil {
		errors = append(errors, "Unable to query PCIe max link generation: "+errorString(r).Error())
	}
	r = C.nvmlDeviceGetCurrPcieLinkGeneration(d.dev, &curr)
	if errorString(r) != nil {
		errors = append(errors, "Unable to query PCIe current link generation: "+errorString(r).Error())
	}

	if len(errors) > 0 {
		return uint(curr), uint(max), fmt.Errorf(strings.Join(errors, "\n"))
	}
	return uint(curr), uint(max), nil
}

// PCIeLinkWidth returns the current PCIe Link generation
// first uint ist the current width, second is the maximum supported width
func (d Device) PCIeLinkWidth() (uint, uint, error) {
	var errors []string

	if C.nvmlHandle == nil {
		return 0, 0, errLibraryNotLoaded
	}
	var curr, max C.uint
	r := C.nvmlDeviceGetMaxPcieLinkWidth(d.dev, &max)
	if errorString(r) != nil {
		errors = append(errors, "Unable to query PCIe max link wdith: "+errorString(r).Error())
	}
	r = C.nvmlDeviceGetCurrPcieLinkWidth(d.dev, &curr)
	if errorString(r) != nil {
		errors = append(errors, "Unable to query PCIe current link width: "+errorString(r).Error())
	}

	if len(errors) > 0 {
		return uint(curr), uint(max), fmt.Errorf(strings.Join(errors, "\n"))
	}
	return uint(curr), uint(max), nil
}

// ApplicationClock returns the current clock of a device application in MHz.
// the application should be specified in ct
func (d Device) ApplicationClock(ct ClockType) (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var clock C.uint
	r := C.nvmlDeviceGetApplicationsClock(d.dev, C.nvmlClockType_t(ct), &clock)
	return uint(clock), errorString(r)
}

// Clock returns the current clock of a device application in MHz.
// the application should be specified in ct
func (d Device) Clock(ct ClockType) (uint, error) {
	if C.nvmlHandle == nil {
		return 0, errLibraryNotLoaded
	}
	var clock C.uint
	r := C.nvmlDeviceGetClockInfo(d.dev, C.nvmlClockType_t(ct), &clock)
	return uint(clock), errorString(r)
}

// ComputeProcesses returns information about processes with a compute context on a device
func (d Device) ComputeProcesses() ([]Process, error) {
	var size = C.uint(2)
	var cprocs []C.nvmlProcessInfo_t
	var r C.nvmlReturn_t
	if C.nvmlHandle == nil {
		return nil, errLibraryNotLoaded
	}
	for r = C.nvmlReturn_t(C.NVML_ERROR_INSUFFICIENT_SIZE); r == C.NVML_ERROR_INSUFFICIENT_SIZE; {
		cprocs = make([]C.nvmlProcessInfo_t, uint(size))
		r = C.nvmlDeviceGetComputeRunningProcesses(d.dev, &size, &cprocs[0])
	}
	if(uint(size) > 0){
		procs := make([]process, len(cprocs))
		for i, cproc := range cprocs {
			procs[i].pid = uint(cproc.pid)
			procs[i].usedGpuMemory = uint64(cproc.usedGpuMemory)
		}

		Procs := make([]Process, len(procs))
		for i, tmp := range procs {
			Procs[i] = tmp
		}

		return Procs, errorString(r)
	}
	return nil, errorString(r)
}

// GraphicsProcesses returns information about processes with a graphics context on a device
func (d Device) GraphicsProcesses() ([]Process, error) {
	var size = C.uint(2)
	var cprocs []C.nvmlProcessInfo_t
	var r C.nvmlReturn_t
	if C.nvmlHandle == nil {
		return nil, errLibraryNotLoaded
	}
	for r = C.nvmlReturn_t(C.NVML_ERROR_INSUFFICIENT_SIZE); r == C.NVML_ERROR_INSUFFICIENT_SIZE; {
		cprocs = make([]C.nvmlProcessInfo_t, uint(size))
		r = C.nvmlDeviceGetGraphicsRunningProcesses(d.dev, &size, &cprocs[0])
	}

	if(uint(size) > 0){
		procs := make([]process, len(cprocs))
		for i, cproc := range cprocs {
			procs[i].pid = uint(cproc.pid)
			procs[i].usedGpuMemory = uint64(cproc.usedGpuMemory)
		}

		Procs := make([]Process, len(procs))
		for i, tmp := range procs {
			Procs[i] = tmp
		}

		return Procs, errorString(r)
	}
	return nil, errorString(r)
}
