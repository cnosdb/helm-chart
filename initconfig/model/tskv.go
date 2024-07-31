package model

/* type QueryTskvConfig struct {
	ReportingDisabled bool   `toml:"reporting_disabled"`
	NodeID            int64  `toml:"node_id,omitempty"`
	FileBufferSize    string `toml:"file_buffer_size,omitempty"`

	Global struct {
		NodeID       int64  `toml:"node_id,omitempty"`
		Host         string `toml:"host,omitempty"`
		ClusterName  string `toml:"cluster_name,omitempty"`
		StoreMetrics bool   `toml:"store_metrics"`
	} `toml:"global,omitempty"`

	Deployment struct {
		Mode   string `toml:"mode,omitempty"`
		CPU    int    `toml:"cpu,omitempty"`
		Memory int    `toml:"memory,omitempty"`
	} `toml:"deployment,omitempty"`

	Meta struct {
		ServiceAddr            []string `toml:"service_addr,omitempty"`
		ReportTimeInterval     string   `toml:"report_time_interval,omitempty"`
		UsageSchemaCacheSize   string   `toml:"usage_schema_cache_size,omitempty"`
		ClusterSchemaCacheSize string   `toml:"cluster_schema_cache_size,omitempty"`
	} `toml:"meta,omitempty"`

	Query struct {
		MaxServerConnections int    `toml:"max_server_connections,omitempty"`
		QuerySQLLimit        any    `toml:"query_sql_limit,omitempty"`
		WriteSQLLimit        any    `toml:"write_sql_limit,omitempty"`
		AuthEnabled          bool   `toml:"auth_enabled"`
		ReadTimeoutMS        any    `toml:"read_timeout_ms,omitempty"`
		WriteTimeoutMS       any    `toml:"write_timeout_ms,omitempty"`
		StreamTriggerCPU     int    `toml:"stream_trigger_cpu,omitempty"`
		StreamExecutorCPU    int    `toml:"stream_executor_cpu,omitempty"`
		SQLRecordTimeout     string `toml:"sql_record_timeout,omitempty"`
	} `toml:"query,omitempty"`

	Storage struct {
		Path                       string `toml:"path,omitempty"`
		MaxSummarySize             string `toml:"max_summary_size,omitempty"`
		BaseFileSize               string `toml:"base_file_size,omitempty"`
		FlushReqChannelCap         int    `toml:"flush_req_channel_cap,omitempty"`
		MaxCachedReaders           int    `toml:"max_cached_readers,omitempty"`
		MaxLevel                   int    `toml:"max_level,omitempty"`
		CompactTriggerFileNum      int    `toml:"compact_trigger_file_num,omitempty"`
		CompactTriggerColdDuration string `toml:"compact_trigger_cold_duration,omitempty"`
		MaxCompactSize             string `toml:"max_compact_size,omitempty"`
		MaxConcurrentCompaction    int    `toml:"max_concurrent_compaction,omitempty"`
		CompactFileCacheSize       string `toml:"compact_file_cache_size,omitempty"`
		StrictWrite                bool   `toml:"strict_write"`
		ReserveSpace               string `toml:"reserve_space,omitempty"`
		CopyintoTriggerFlushSize   string `toml:"copyinto_trigger_flush_size,omitempty"`
	} `toml:"storage,omitempty"`

	WAL struct {
		Enabled                     bool   `toml:"enabled"`
		Path                        string `toml:"path,omitempty"`
		WALReqChannelCap            int    `toml:"wal_req_channel_cap,omitempty"`
		MaxFileSize                 string `toml:"max_file_size,omitempty"`
		FlushTriggerTotalFileSize   string `toml:"flush_trigger_total_file_size,omitempty"`
		FlushTriggerMinimumInterval string `toml:"flush_trigger_minimum_interval,omitempty"`
		Sync                        bool   `toml:"sync"`
		SyncInterval                string `toml:"sync_interval,omitempty"`
	} `toml:"wal,omitempty"`

	Cache struct {
		MaxBufferSize            any    `toml:"max_buffer_size,omitempty"`
		MaxImmutableNumber       int    `toml:"max_immutable_number,omitempty"`
		Partition                int    `toml:"partition,omitempty"`
		MaxUsageSchemaBufferSize string `toml:"max_usage_schema_buffer_size,omitempty"`
	} `toml:"cache,omitempty"`

	Log struct {
		Level        string `toml:"level,omitempty"`
		Path         string `toml:"path,omitempty"`
		MaxFileCount int    `toml:"max_file_count,omitempty"`
		FileRotation string `toml:"file_rotation,omitempty"`
		TokioTrace   struct {
			Addr string `toml:"addr,omitempty"`
		} `toml:"tokio_trace,omitempty"`
	} `toml:"log,omitempty"`

	Security struct {
		TLSConfig struct {
			Certificate string `toml:"certificate,omitempty"`
			PrivateKey  string `toml:"private_key,omitempty"`
		} `toml:"tls_config,omitempty"`
	} `toml:"security,omitempty"`

	Service struct {
		HTTPListenPort      int  `toml:"http_listen_port,omitempty"`
		GRPCListenPort      int  `toml:"grpc_listen_port,omitempty"`
		GRPCEnableGzip      bool `toml:"grpc_enable_gzip"`
		FlightRPCListenPort int  `toml:"flight_rpc_listen_port,omitempty"`
		TCPListenPort       int  `toml:"tcp_listen_port,omitempty"`
		VectorListenPort    int  `toml:"vector_listen_port,omitempty"`
		EnableReport        bool `toml:"enable_report"`
	} `toml:"service,omitempty"`

	Cluster struct {
		Name                 string   `toml:"name,omitempty"`
		MetaServiceAddr      []string `toml:"meta_service_addr,omitempty"`
		HTTPListenPort       int      `toml:"http_listen_port,omitempty"`
		GRPCListenPort       int      `toml:"grpc_listen_port,omitempty"`
		FlightRPCListenPort  int      `toml:"flight_rpc_listen_port,omitempty"`
		TCPListenPort        int      `toml:"tcp_listen_port,omitempty"`
		VectorListenPort     int      `toml:"vector_listen_port,omitempty"`
		RaftLogsToKeep       int      `toml:"raft_logs_to_keep,omitempty"`
		UsingRaftReplication bool     `toml:"using_raft_replication"`
	} `toml:"cluster,omitempty"`

	NodeBasic struct {
		NodeID       int64 `toml:"node_id,omitempty"`
		StoreMetrics bool  `toml:"store_metrics"`
	} `toml:"node_basic,omitempty"`

	Heartbeat struct {
		ReportTimeIntervalSecs int `toml:"report_time_interval_secs,omitempty"`
	} `toml:"heartbeat,omitempty"`

	HintedOff struct {
		Enable  bool   `toml:"enable"`
		Path    string `toml:"path,omitempty"`
		Threads int    `toml:"threads,omitempty"`
	} `toml:"hinted_off,omitempty"`

	Trace struct {
		AutoGenerateSpan          bool   `toml:"auto_generate_span"`
		OTLPEndpoint              string `toml:"otlp_endpoint,omitempty"`
		MaxSpansPerTrace          int    `toml:"max_spans_per_trace,omitempty"`
		BatchReportIntervalMillis int    `toml:"batch_report_interval_millis,omitempty"`
		BatchReportMaxSpans       int    `toml:"batch_report_max_spans,omitempty"`
		Log                       struct {
			Path string `toml:"path,omitempty"`
		} `toml:"log,omitempty"`
		Jaeger struct {
			JaegerAgentEndpoint  string `toml:"jaeger_agent_endpoint,omitempty"`
			MaxConcurrentExports int    `toml:"max_concurrent_exports,omitempty"`
			MaxQueueSize         int    `toml:"max_queue_size,omitempty"`
		} `toml:"jaeger,omitempty"`
	} `toml:"trace,omitempty"`
} */

type QueryTskvConfig struct {
	ReportingDisabled bool   `toml:"reporting_disabled" json:"reporting_disabled"`
	NodeID            int64  `toml:"node_id" json:"node_id"`
	FileBufferSize    string `toml:"file_buffer_size,omitempty" json:"file_buffer_size,omitempty"`
	LicenseFile       string `json:"license_file" toml:"license_file"`
	Host              string `toml:"host,omitempty" json:"host,omitempty"`

	Global struct {
		NodeID       int64  `toml:"node_id" json:"node_id"`
		Host         string `toml:"host,omitempty" json:"host,omitempty"`
		ClusterName  string `toml:"cluster_name,omitempty" json:"cluster_name,omitempty"`
		StoreMetrics bool   `toml:"store_metrics" json:"store_metrics"`
	} `toml:"global,omitempty" json:"global,omitempty"`

	Deployment struct {
		Mode   string `toml:"mode,omitempty" json:"mode,omitempty"`
		CPU    int    `toml:"cpu,omitempty" json:"cpu,omitempty"`
		Memory int    `toml:"memory,omitempty" json:"memory,omitempty"`
	} `toml:"deployment,omitempty" json:"deployment,omitempty"`

	Meta struct {
		ServiceAddr            []string `toml:"service_addr,omitempty" json:"service_addr,omitempty"`
		ReportTimeInterval     string   `toml:"report_time_interval,omitempty" json:"report_time_interval,omitempty"`
		UsageSchemaCacheSize   string   `toml:"usage_schema_cache_size,omitempty" json:"usage_schema_cache_size,omitempty"`
		ClusterSchemaCacheSize string   `toml:"cluster_schema_cache_size,omitempty" json:"cluster_schema_cache_size,omitempty"`
	} `toml:"meta,omitempty" json:"meta,omitempty"`

	Query struct {
		MaxServerConnections int    `toml:"max_server_connections,omitempty" json:"max_server_connections,omitempty"`
		QuerySQLLimit        any    `toml:"query_sql_limit,omitempty" json:"query_sql_limit,omitempty"`
		WriteSQLLimit        any    `toml:"write_sql_limit,omitempty" json:"write_sql_limit,omitempty"`
		AuthEnabled          bool   `toml:"auth_enabled" json:"auth_enabled"`
		ReadTimeoutMS        any    `toml:"read_timeout_ms,omitempty" json:"read_timeout_ms,omitempty"`
		WriteTimeoutMS       any    `toml:"write_timeout_ms,omitempty" json:"write_timeout_ms,omitempty"`
		StreamTriggerCPU     int    `toml:"stream_trigger_cpu,omitempty" json:"stream_trigger_cpu,omitempty"`
		StreamExecutorCPU    int    `toml:"stream_executor_cpu,omitempty" json:"stream_executor_cpu,omitempty"`
		SQLRecordTimeout     string `toml:"sql_record_timeout,omitempty" json:"sql_record_timeout,omitempty"`
	} `toml:"query,omitempty" json:"query,omitempty"`

	Storage struct {
		Path                       string `toml:"path,omitempty" json:"path,omitempty"`
		MaxSummarySize             string `toml:"max_summary_size,omitempty" json:"max_summary_size,omitempty"`
		BaseFileSize               string `toml:"base_file_size,omitempty" json:"base_file_size,omitempty"`
		FlushReqChannelCap         int    `toml:"flush_req_channel_cap,omitempty" json:"flush_req_channel_cap,omitempty"`
		MaxCachedReaders           int    `toml:"max_cached_readers,omitempty" json:"max_cached_readers,omitempty"`
		MaxLevel                   int    `toml:"max_level,omitempty" json:"max_level,omitempty"`
		CompactTriggerFileNum      int    `toml:"compact_trigger_file_num,omitempty" json:"compact_trigger_file_num,omitempty"`
		CompactTriggerColdDuration string `toml:"compact_trigger_cold_duration,omitempty" json:"compact_trigger_cold_duration,omitempty"`
		MaxCompactSize             string `toml:"max_compact_size,omitempty" json:"max_compact_size,omitempty"`
		MaxConcurrentCompaction    int    `toml:"max_concurrent_compaction,omitempty" json:"max_concurrent_compaction,omitempty"`
		CompactFileCacheSize       string `toml:"compact_file_cache_size,omitempty" json:"compact_file_cache_size,omitempty"`
		StrictWrite                bool   `toml:"strict_write" json:"strict_write"`
		ReserveSpace               string `toml:"reserve_space,omitempty" json:"reserve_space,omitempty"`
		CopyintoTriggerFlushSize   string `toml:"copyinto_trigger_flush_size,omitempty" json:"copyinto_trigger_flush_size,omitempty"`
	} `toml:"storage,omitempty" json:"storage,omitempty"`

	WAL struct {
		Enabled                     bool   `toml:"enabled" json:"enabled"`
		Path                        string `toml:"path,omitempty" json:"path,omitempty"`
		WALReqChannelCap            int    `toml:"wal_req_channel_cap,omitempty" json:"wal_req_channel_cap,omitempty"`
		MaxFileSize                 string `toml:"max_file_size,omitempty" json:"max_file_size,omitempty"`
		FlushTriggerTotalFileSize   string `toml:"flush_trigger_total_file_size,omitempty" json:"flush_trigger_total_file_size,omitempty"`
		FlushTriggerMinimumInterval string `toml:"flush_trigger_minimum_interval,omitempty" json:"flush_trigger_minimum_interval,omitempty"`
		Sync                        bool   `toml:"sync" json:"sync"`
		SyncInterval                string `toml:"sync_interval,omitempty" json:"sync_interval,omitempty"`
	} `toml:"wal,omitempty" json:"wal,omitempty"`

	Cache struct {
		MaxBufferSize            any    `toml:"max_buffer_size,omitempty" json:"max_buffer_size,omitempty"`
		MaxImmutableNumber       int    `toml:"max_immutable_number,omitempty" json:"max_immutable_number,omitempty"`
		Partition                int    `toml:"partition,omitempty" json:"partition,omitempty"`
		MaxUsageSchemaBufferSize string `toml:"max_usage_schema_buffer_size,omitempty" json:"max_usage_schema_buffer_size,omitempty"`
	} `toml:"cache,omitempty" json:"cache,omitempty"`

	Log struct {
		Level        string `toml:"level,omitempty" json:"level,omitempty"`
		Path         string `toml:"path,omitempty" json:"path,omitempty"`
		MaxFileCount int    `toml:"max_file_count,omitempty" json:"max_file_count,omitempty"`
		FileRotation string `toml:"file_rotation,omitempty" json:"file_rotation,omitempty"`
		TokioTrace   struct {
			Addr string `toml:"addr,omitempty" json:"addr,omitempty"`
		} `toml:"tokio_trace,omitempty" json:"tokio_trace,omitempty"`
	} `toml:"log,omitempty" json:"log,omitempty"`

	Security struct {
		TLSConfig struct {
			Certificate string `toml:"certificate,omitempty" json:"certificate,omitempty"`
			PrivateKey  string `toml:"private_key,omitempty" json:"private_key,omitempty"`
		} `toml:"tls_config,omitempty" json:"tls_config,omitempty"`
	} `toml:"security,omitempty" json:"security,omitempty"`

	Service struct {
		HTTPListenPort      int  `toml:"http_listen_port,omitempty" json:"http_listen_port,omitempty"`
		GRPCListenPort      int  `toml:"grpc_listen_port,omitempty" json:"grpc_listen_port,omitempty"`
		GRPCEnableGzip      bool `toml:"grpc_enable_gzip" json:"grpc_enable_gzip"`
		FlightRPCListenPort int  `toml:"flight_rpc_listen_port,omitempty" json:"flight_rpc_listen_port,omitempty"`
		TCPListenPort       int  `toml:"tcp_listen_port,omitempty" json:"tcp_listen_port,omitempty"`
		VectorListenPort    int  `toml:"vector_listen_port,omitempty" json:"vector_listen_port,omitempty"`
		EnableReport        bool `toml:"enable_report" json:"enable_report"`
	} `toml:"service,omitempty" json:"service,omitempty"`

	Cluster struct {
		Name                 string   `toml:"name,omitempty" json:"name,omitempty"`
		MetaServiceAddr      []string `toml:"meta_service_addr,omitempty" json:"meta_service_addr,omitempty"`
		HTTPListenPort       int      `toml:"http_listen_port,omitempty" json:"http_listen_port,omitempty"`
		GRPCListenPort       int      `toml:"grpc_listen_port,omitempty" json:"grpc_listen_port,omitempty"`
		FlightRPCListenPort  int      `toml:"flight_rpc_listen_port,omitempty" json:"flight_rpc_listen_port,omitempty"`
		TCPListenPort        int      `toml:"tcp_listen_port,omitempty" json:"tcp_listen_port,omitempty"`
		VectorListenPort     int      `toml:"vector_listen_port,omitempty" json:"vector_listen_port,omitempty"`
		RaftLogsToKeep       int      `toml:"raft_logs_to_keep,omitempty" json:"raft_logs_to_keep,omitempty"`
		UsingRaftReplication bool     `toml:"using_raft_replication" json:"using_raft_replication"`
	} `toml:"cluster,omitempty" json:"cluster,omitempty"`

	NodeBasic struct {
		NodeID       int64 `toml:"node_id" json:"node_id"`
		StoreMetrics bool  `toml:"store_metrics" json:"store_metrics"`
	} `toml:"node_basic,omitempty" json:"node_basic,omitempty"`

	Heartbeat struct {
		ReportTimeIntervalSecs int `toml:"report_time_interval_secs,omitempty" json:"report_time_interval_secs,omitempty"`
	} `toml:"heartbeat,omitempty" json:"heartbeat,omitempty"`

	HintedOff struct {
		Enable  bool   `toml:"enable" json:"enable"`
		Path    string `toml:"path,omitempty" json:"path,omitempty"`
		Threads int    `toml:"threads,omitempty" json:"threads,omitempty"`
	} `toml:"hinted_off,omitempty" json:"hinted_off,omitempty"`

	Trace struct {
		AutoGenerateSpan          bool   `toml:"auto_generate_span" json:"auto_generate_span"`
		OTLPEndpoint              string `toml:"otlp_endpoint,omitempty" json:"otlp_endpoint,omitempty"`
		MaxSpansPerTrace          int    `toml:"max_spans_per_trace,omitempty" json:"max_spans_per_trace,omitempty"`
		BatchReportIntervalMillis int    `toml:"batch_report_interval_millis,omitempty" json:"batch_report_interval_millis,omitempty"`
		BatchReportMaxSpans       int    `toml:"batch_report_max_spans,omitempty" json:"batch_report_max_spans,omitempty"`
		Log                       struct {
			Path string `toml:"path,omitempty" json:"path,omitempty"`
		} `toml:"log,omitempty" json:"log,omitempty"`
		Jaeger struct {
			JaegerAgentEndpoint  string `toml:"jaeger_agent_endpoint,omitempty" json:"jaeger_agent_endpoint,omitempty"`
			MaxConcurrentExports int    `toml:"max_concurrent_exports,omitempty" json:"max_concurrent_exports,omitempty"`
			MaxQueueSize         int    `toml:"max_queue_size,omitempty" json:"max_queue_size,omitempty"`
		} `toml:"jaeger,omitempty" json:"jaeger,omitempty"`
	} `toml:"trace,omitempty" json:"trace,omitempty"`
}
