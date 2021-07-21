// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots.

package config

import (
	"encoding/json"
	"reflect"

	"fmt"

	"github.com/spf13/pflag"
)

// If v is a pointer, it will get its element value or the zero value of the element type.
// If v is not a pointer, it will return it as is.
func (Config) elemValueOrNil(v interface{}) interface{} {
	if t := reflect.TypeOf(v); t.Kind() == reflect.Ptr {
		if reflect.ValueOf(v).IsNil() {
			return reflect.Zero(t.Elem()).Interface()
		} else {
			return reflect.ValueOf(v).Interface()
		}
	} else if v == nil {
		return reflect.Zero(t).Interface()
	}

	return v
}

func (Config) mustMarshalJSON(v json.Marshaler) string {
	raw, err := v.MarshalJSON()
	if err != nil {
		panic(err)
	}

	return string(raw)
}

// GetPFlagSet will return strongly types pflags for all fields in Config and its nested types. The format of the
// flags is json-name.json-sub-name... etc.
func (cfg Config) GetPFlagSet(prefix string) *pflag.FlagSet {
	cmdFlags := pflag.NewFlagSet("Config", pflag.ExitOnError)
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "resyncPeriod"), "30s", "Determines the resync period for all watchers.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "metricsPrefix"), "flyteK8sSchedulerExtension", "Prefix for metrics propagated to prometheus")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "profilerPort"), "10254", "Profiler port")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "workers"), 4, "Number of routines to process custom resource")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "unrestrictedCpuAutoscalingLimit"), 30, "Max number of cpus for unrestricted auto-scaling.")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "unrestrictedMemoryAutoscalingLimit"), 30, "Max amount of memory for unrestricted auto-scaling.")
	return cmdFlags
}
