package model

/* type MetaConfig struct {
	ID                       int64  `toml:"id"`
	Host                     string `toml:"host,omitempty"`
	Port                     int    `toml:"port,omitempty"`
	SnapshotPath             string `toml:"snapshot_path,omitempty"`
	JournalPath              string `toml:"journal_path,omitempty"`
	SnapshotPerEvents        int    `toml:"snapshot_per_events,omitempty"`
	DataPath                 string `toml:"data_path,omitempty"`
	ClusterName              string `toml:"cluster_name,omitempty"`
	GrpcEnableGzip           bool   `toml:"grpc_enable_gzip"`
	LmdbMaxMapSize           int    `toml:"lmdb_max_map_size,omitempty"`
	HeartbeatInterval        int    `toml:"heartbeat_interval,omitempty"`
	RaftLogsToKeep           int    `toml:"raft_logs_to_keep,omitempty"`
	InstallSnapshotTimeout   int    `toml:"install_snapshot_timeout,omitempty"`
	SendAppendEntriesTimeout int    `toml:"send_append_entries_timeout,omitempty"`
	UsageSchemaCacheSize     int    `toml:"usage_schema_cache_size,omitempty"`
	ClusterSchemaCacheSize   int    `toml:"cluster_schema_cache_size,omitempty"`
	Log                      struct {
		Level string `toml:"level,omitempty"`
		Path  string `toml:"path,omitempty"`
	} `toml:"log,omitempty"`
	MetaInit struct {
		ClusterName     string   `toml:"cluster_name,omitempty"`
		AdminUser       string   `toml:"admin_user,omitempty"`
		SystemTenant    string   `toml:"system_tenant,omitempty"`
		DefaultDatabase []string `toml:"default_database,omitempty"`
	} `toml:"meta_init,omitempty"`
	Heartbeat struct {
		HeartbeatRecheckInterval int `toml:"heartbeat_recheck_interval,omitempty"`
		HeartbeatExpiredInterval int `toml:"heartbeat_expired_interval,omitempty"`
	} `toml:"heartbeat,omitempty"`
} */

type MetaConfig struct {
	ID                        int64  `json:"id" toml:"id"`
	LicenseFile               string `json:"license_file" toml:"license_file"`
	Host                      string `json:"host,omitempty" toml:"host,omitempty"`
	Port                      int    `json:"port,omitempty" toml:"port,omitempty"`
	SnapshotPath              string `json:"snapshot_path,omitempty" toml:"snapshot_path,omitempty"`
	JournalPath               string `json:"journal_path,omitempty" toml:"journal_path,omitempty"`
	SnapshotPerEvents         int    `json:"snapshot_per_events,omitempty" toml:"snapshot_per_events,omitempty"`
	DataPath                  string `json:"data_path,omitempty" toml:"data_path,omitempty"`
	ClusterName               string `json:"cluster_name,omitempty" toml:"cluster_name,omitempty"`
	AutoMigrateVnodesDuration int    `json:"auto_migrate_vnodes_duration" toml:"auto_migrate_vnodes_duration"`
	GrpcEnableGzip            bool   `json:"grpc_enable_gzip" toml:"grpc_enable_gzip"`
	LmdbMaxMapSize            int    `json:"lmdb_max_map_size,omitempty" toml:"lmdb_max_map_size,omitempty"`
	HeartbeatInterval         int    `json:"heartbeat_interval,omitempty" toml:"heartbeat_interval,omitempty"`
	RaftLogsToKeep            int    `json:"raft_logs_to_keep,omitempty" toml:"raft_logs_to_keep,omitempty"`
	InstallSnapshotTimeout    int    `json:"install_snapshot_timeout,omitempty" toml:"install_snapshot_timeout,omitempty"`
	SendAppendEntriesTimeout  int    `json:"send_append_entries_timeout,omitempty" toml:"send_append_entries_timeout,omitempty"`
	UsageSchemaCacheSize      int    `json:"usage_schema_cache_size,omitempty" toml:"usage_schema_cache_size,omitempty"`
	ClusterSchemaCacheSize    int    `json:"cluster_schema_cache_size,omitempty" toml:"cluster_schema_cache_size,omitempty"`
	Log                       struct {
		Level string `json:"level,omitempty" toml:"level,omitempty"`
		Path  string `json:"path,omitempty" toml:"path,omitempty"`
	} `json:"log,omitempty" toml:"log,omitempty"`
	MetaInit struct {
		ClusterName     string   `json:"cluster_name,omitempty" toml:"cluster_name,omitempty"`
		AdminUser       string   `json:"admin_user,omitempty" toml:"admin_user,omitempty"`
		SystemTenant    string   `json:"system_tenant,omitempty" toml:"system_tenant,omitempty"`
		DefaultDatabase []string `json:"default_database,omitempty" toml:"default_database,omitempty"`
	} `json:"meta_init,omitempty" toml:"meta_init,omitempty"`
	Heartbeat struct {
		HeartbeatRecheckInterval int `json:"heartbeat_recheck_interval,omitempty" toml:"heartbeat_recheck_interval,omitempty"`
		HeartbeatExpiredInterval int `json:"heartbeat_expired_interval,omitempty" toml:"heartbeat_expired_interval,omitempty"`
	} `json:"heartbeat,omitempty" toml:"heartbeat,omitempty"`
}
