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
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "kube-config"), defaultConfig.KubeConfigPath, "Path to kubernetes client config file.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "master"), defaultConfig.MasterURL, "")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "workers"), defaultConfig.Workers, "Number of threads to process workflows")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "workflow-reeval-duration"), defaultConfig.WorkflowReEval.String(), "Frequency of re-evaluating workflows")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "downstream-eval-duration"), defaultConfig.DownstreamEval.String(), "Frequency of re-evaluating downstream tasks")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "limit-namespace"), defaultConfig.LimitNamespace, "Namespaces to watch for this propeller")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "namespace-filter"), defaultConfig.NamespaceFilter.String(), "Regular expression to apply fine grain filter when watching all namespaces")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "prof-port"), defaultConfig.ProfilerPort.String(), "Profiler port")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "metadata-prefix"), defaultConfig.MetadataPrefix, "MetadataPrefix should be used if all the metadata for Flyte executions should be stored under a specific prefix in CloudStorage. If not specified,  the data will be stored in the base container directly.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "queue.type"), defaultConfig.Queue.Type, "Type of composite queue to use for the WorkQueue")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "queue.queue.type"), defaultConfig.Queue.Queue.Type, "Type of RateLimiter to use for the WorkQueue")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "queue.queue.base-delay"), defaultConfig.Queue.Queue.BaseDelay.String(), "base backoff delay for failure")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "queue.queue.max-delay"), defaultConfig.Queue.Queue.MaxDelay.String(), "Max backoff delay for failure")
	cmdFlags.Int64(fmt.Sprintf("%v%v", prefix, "queue.queue.rate"), defaultConfig.Queue.Queue.Rate, "Bucket Refill rate per second")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "queue.queue.capacity"), defaultConfig.Queue.Queue.Capacity, "Bucket capacity as number of items")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "queue.sub-queue.type"), defaultConfig.Queue.Sub.Type, "Type of RateLimiter to use for the WorkQueue")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "queue.sub-queue.base-delay"), defaultConfig.Queue.Sub.BaseDelay.String(), "base backoff delay for failure")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "queue.sub-queue.max-delay"), defaultConfig.Queue.Sub.MaxDelay.String(), "Max backoff delay for failure")
	cmdFlags.Int64(fmt.Sprintf("%v%v", prefix, "queue.sub-queue.rate"), defaultConfig.Queue.Sub.Rate, "Bucket Refill rate per second")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "queue.sub-queue.capacity"), defaultConfig.Queue.Sub.Capacity, "Bucket capacity as number of items")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "queue.batching-interval"), defaultConfig.Queue.BatchingInterval.String(), "Duration for which downstream updates are buffered")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "queue.batch-size"), defaultConfig.Queue.BatchSize, "Number of downstream triggered top-level objects to re-enqueue every duration. -1 indicates all available.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "metrics-prefix"), defaultConfig.MetricsPrefix, "An optional prefix for all published metrics.")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "enable-admin-launcher"), defaultConfig.EnableAdminLauncher, " Enable remote Workflow launcher to Admin")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "max-workflow-retries"), defaultConfig.MaxWorkflowRetries, "Maximum number of retries per workflow")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "max-ttl-hours"), defaultConfig.MaxTTLInHours, "Maximum number of hours a completed workflow should be retained. Number between 1-23 hours")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "gc-interval"), defaultConfig.GCInterval.String(), "Run periodic GC every 30 minutes")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "leader-election.enabled"), defaultConfig.LeaderElection.Enabled, "Enables/Disables leader election.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "leader-election.lock-config-map.Namespace"), defaultConfig.LeaderElection.LockConfigMap.Namespace, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "leader-election.lock-config-map.Name"), defaultConfig.LeaderElection.LockConfigMap.Name, "")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "leader-election.lease-duration"), defaultConfig.LeaderElection.LeaseDuration.String(), "Duration that non-leader candidates will wait to force acquire leadership. This is measured against time of last observed ack.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "leader-election.renew-deadline"), defaultConfig.LeaderElection.RenewDeadline.String(), "Duration that the acting master will retry refreshing leadership before giving up.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "leader-election.retry-period"), defaultConfig.LeaderElection.RetryPeriod.String(), "Duration the LeaderElector clients should wait between tries of actions.")
	cmdFlags.Bool(fmt.Sprintf("%v%v", prefix, "publish-k8s-events"), defaultConfig.PublishK8sEvents, "Enable events publishing to K8s events API.")
	cmdFlags.Int64(fmt.Sprintf("%v%v", prefix, "max-output-size-bytes"), defaultConfig.MaxDatasetSizeBytes, "Maximum size of outputs per task")
	cmdFlags.Int(fmt.Sprintf("%v%v", prefix, "kube-client-config.burst"), defaultConfig.KubeConfig.Burst, "Max burst rate for throttle. 0 defaults to 10")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "kube-client-config.timeout"), defaultConfig.KubeConfig.Timeout.String(), "Max duration allowed for every request to KubeAPI before giving up. 0 implies no timeout.")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "node-config.default-deadlines.node-execution-deadline"), defaultConfig.NodeConfig.DefaultDeadlines.DefaultNodeExecutionDeadline.String(), "Default value of node execution timeout")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "node-config.default-deadlines.node-active-deadline"), defaultConfig.NodeConfig.DefaultDeadlines.DefaultNodeActiveDeadline.String(), "Default value of node timeout")
	cmdFlags.String(fmt.Sprintf("%v%v", prefix, "node-config.default-deadlines.workflow-active-deadline"), defaultConfig.NodeConfig.DefaultDeadlines.DefaultWorkflowActiveDeadline.String(), "Default value of workflow timeout")
	cmdFlags.Int64(fmt.Sprintf("%v%v", prefix, "node-config.max-node-retries-system-failures"), defaultConfig.NodeConfig.MaxNodeRetriesForSystemFailures, "Maximum number of retries per node for node failure due to infra issues")
	return cmdFlags
}
