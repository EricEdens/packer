// Code generated by protoc-gen-goext. DO NOT EDIT.

package clickhouse

import (
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

func (m *User) SetName(v string) {
	m.Name = v
}

func (m *User) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *User) SetPermissions(v []*Permission) {
	m.Permissions = v
}

func (m *User) SetSettings(v *UserSettings) {
	m.Settings = v
}

func (m *User) SetQuotas(v []*UserQuota) {
	m.Quotas = v
}

func (m *Permission) SetDatabaseName(v string) {
	m.DatabaseName = v
}

func (m *UserSpec) SetName(v string) {
	m.Name = v
}

func (m *UserSpec) SetPassword(v string) {
	m.Password = v
}

func (m *UserSpec) SetPermissions(v []*Permission) {
	m.Permissions = v
}

func (m *UserSpec) SetSettings(v *UserSettings) {
	m.Settings = v
}

func (m *UserSpec) SetQuotas(v []*UserQuota) {
	m.Quotas = v
}

func (m *UserSettings) SetReadonly(v *wrappers.Int64Value) {
	m.Readonly = v
}

func (m *UserSettings) SetAllowDdl(v *wrappers.BoolValue) {
	m.AllowDdl = v
}

func (m *UserSettings) SetInsertQuorum(v *wrappers.Int64Value) {
	m.InsertQuorum = v
}

func (m *UserSettings) SetConnectTimeout(v *wrappers.Int64Value) {
	m.ConnectTimeout = v
}

func (m *UserSettings) SetReceiveTimeout(v *wrappers.Int64Value) {
	m.ReceiveTimeout = v
}

func (m *UserSettings) SetSendTimeout(v *wrappers.Int64Value) {
	m.SendTimeout = v
}

func (m *UserSettings) SetInsertQuorumTimeout(v *wrappers.Int64Value) {
	m.InsertQuorumTimeout = v
}

func (m *UserSettings) SetSelectSequentialConsistency(v *wrappers.BoolValue) {
	m.SelectSequentialConsistency = v
}

func (m *UserSettings) SetMaxReplicaDelayForDistributedQueries(v *wrappers.Int64Value) {
	m.MaxReplicaDelayForDistributedQueries = v
}

func (m *UserSettings) SetFallbackToStaleReplicasForDistributedQueries(v *wrappers.BoolValue) {
	m.FallbackToStaleReplicasForDistributedQueries = v
}

func (m *UserSettings) SetReplicationAlterPartitionsSync(v *wrappers.Int64Value) {
	m.ReplicationAlterPartitionsSync = v
}

func (m *UserSettings) SetDistributedProductMode(v UserSettings_DistributedProductMode) {
	m.DistributedProductMode = v
}

func (m *UserSettings) SetDistributedAggregationMemoryEfficient(v *wrappers.BoolValue) {
	m.DistributedAggregationMemoryEfficient = v
}

func (m *UserSettings) SetDistributedDdlTaskTimeout(v *wrappers.Int64Value) {
	m.DistributedDdlTaskTimeout = v
}

func (m *UserSettings) SetSkipUnavailableShards(v *wrappers.BoolValue) {
	m.SkipUnavailableShards = v
}

func (m *UserSettings) SetCompile(v *wrappers.BoolValue) {
	m.Compile = v
}

func (m *UserSettings) SetMinCountToCompile(v *wrappers.Int64Value) {
	m.MinCountToCompile = v
}

func (m *UserSettings) SetCompileExpressions(v *wrappers.BoolValue) {
	m.CompileExpressions = v
}

func (m *UserSettings) SetMinCountToCompileExpression(v *wrappers.Int64Value) {
	m.MinCountToCompileExpression = v
}

func (m *UserSettings) SetMaxBlockSize(v *wrappers.Int64Value) {
	m.MaxBlockSize = v
}

func (m *UserSettings) SetMinInsertBlockSizeRows(v *wrappers.Int64Value) {
	m.MinInsertBlockSizeRows = v
}

func (m *UserSettings) SetMinInsertBlockSizeBytes(v *wrappers.Int64Value) {
	m.MinInsertBlockSizeBytes = v
}

func (m *UserSettings) SetMaxInsertBlockSize(v *wrappers.Int64Value) {
	m.MaxInsertBlockSize = v
}

func (m *UserSettings) SetMinBytesToUseDirectIo(v *wrappers.Int64Value) {
	m.MinBytesToUseDirectIo = v
}

func (m *UserSettings) SetUseUncompressedCache(v *wrappers.BoolValue) {
	m.UseUncompressedCache = v
}

func (m *UserSettings) SetMergeTreeMaxRowsToUseCache(v *wrappers.Int64Value) {
	m.MergeTreeMaxRowsToUseCache = v
}

func (m *UserSettings) SetMergeTreeMaxBytesToUseCache(v *wrappers.Int64Value) {
	m.MergeTreeMaxBytesToUseCache = v
}

func (m *UserSettings) SetMergeTreeMinRowsForConcurrentRead(v *wrappers.Int64Value) {
	m.MergeTreeMinRowsForConcurrentRead = v
}

func (m *UserSettings) SetMergeTreeMinBytesForConcurrentRead(v *wrappers.Int64Value) {
	m.MergeTreeMinBytesForConcurrentRead = v
}

func (m *UserSettings) SetMaxBytesBeforeExternalGroupBy(v *wrappers.Int64Value) {
	m.MaxBytesBeforeExternalGroupBy = v
}

func (m *UserSettings) SetMaxBytesBeforeExternalSort(v *wrappers.Int64Value) {
	m.MaxBytesBeforeExternalSort = v
}

func (m *UserSettings) SetGroupByTwoLevelThreshold(v *wrappers.Int64Value) {
	m.GroupByTwoLevelThreshold = v
}

func (m *UserSettings) SetGroupByTwoLevelThresholdBytes(v *wrappers.Int64Value) {
	m.GroupByTwoLevelThresholdBytes = v
}

func (m *UserSettings) SetPriority(v *wrappers.Int64Value) {
	m.Priority = v
}

func (m *UserSettings) SetMaxThreads(v *wrappers.Int64Value) {
	m.MaxThreads = v
}

func (m *UserSettings) SetMaxMemoryUsage(v *wrappers.Int64Value) {
	m.MaxMemoryUsage = v
}

func (m *UserSettings) SetMaxMemoryUsageForUser(v *wrappers.Int64Value) {
	m.MaxMemoryUsageForUser = v
}

func (m *UserSettings) SetMaxNetworkBandwidth(v *wrappers.Int64Value) {
	m.MaxNetworkBandwidth = v
}

func (m *UserSettings) SetMaxNetworkBandwidthForUser(v *wrappers.Int64Value) {
	m.MaxNetworkBandwidthForUser = v
}

func (m *UserSettings) SetForceIndexByDate(v *wrappers.BoolValue) {
	m.ForceIndexByDate = v
}

func (m *UserSettings) SetForcePrimaryKey(v *wrappers.BoolValue) {
	m.ForcePrimaryKey = v
}

func (m *UserSettings) SetMaxRowsToRead(v *wrappers.Int64Value) {
	m.MaxRowsToRead = v
}

func (m *UserSettings) SetMaxBytesToRead(v *wrappers.Int64Value) {
	m.MaxBytesToRead = v
}

func (m *UserSettings) SetReadOverflowMode(v UserSettings_OverflowMode) {
	m.ReadOverflowMode = v
}

func (m *UserSettings) SetMaxRowsToGroupBy(v *wrappers.Int64Value) {
	m.MaxRowsToGroupBy = v
}

func (m *UserSettings) SetGroupByOverflowMode(v UserSettings_GroupByOverflowMode) {
	m.GroupByOverflowMode = v
}

func (m *UserSettings) SetMaxRowsToSort(v *wrappers.Int64Value) {
	m.MaxRowsToSort = v
}

func (m *UserSettings) SetMaxBytesToSort(v *wrappers.Int64Value) {
	m.MaxBytesToSort = v
}

func (m *UserSettings) SetSortOverflowMode(v UserSettings_OverflowMode) {
	m.SortOverflowMode = v
}

func (m *UserSettings) SetMaxResultRows(v *wrappers.Int64Value) {
	m.MaxResultRows = v
}

func (m *UserSettings) SetMaxResultBytes(v *wrappers.Int64Value) {
	m.MaxResultBytes = v
}

func (m *UserSettings) SetResultOverflowMode(v UserSettings_OverflowMode) {
	m.ResultOverflowMode = v
}

func (m *UserSettings) SetMaxRowsInDistinct(v *wrappers.Int64Value) {
	m.MaxRowsInDistinct = v
}

func (m *UserSettings) SetMaxBytesInDistinct(v *wrappers.Int64Value) {
	m.MaxBytesInDistinct = v
}

func (m *UserSettings) SetDistinctOverflowMode(v UserSettings_OverflowMode) {
	m.DistinctOverflowMode = v
}

func (m *UserSettings) SetMaxRowsToTransfer(v *wrappers.Int64Value) {
	m.MaxRowsToTransfer = v
}

func (m *UserSettings) SetMaxBytesToTransfer(v *wrappers.Int64Value) {
	m.MaxBytesToTransfer = v
}

func (m *UserSettings) SetTransferOverflowMode(v UserSettings_OverflowMode) {
	m.TransferOverflowMode = v
}

func (m *UserSettings) SetMaxExecutionTime(v *wrappers.Int64Value) {
	m.MaxExecutionTime = v
}

func (m *UserSettings) SetTimeoutOverflowMode(v UserSettings_OverflowMode) {
	m.TimeoutOverflowMode = v
}

func (m *UserSettings) SetMaxColumnsToRead(v *wrappers.Int64Value) {
	m.MaxColumnsToRead = v
}

func (m *UserSettings) SetMaxTemporaryColumns(v *wrappers.Int64Value) {
	m.MaxTemporaryColumns = v
}

func (m *UserSettings) SetMaxTemporaryNonConstColumns(v *wrappers.Int64Value) {
	m.MaxTemporaryNonConstColumns = v
}

func (m *UserSettings) SetMaxQuerySize(v *wrappers.Int64Value) {
	m.MaxQuerySize = v
}

func (m *UserSettings) SetMaxAstDepth(v *wrappers.Int64Value) {
	m.MaxAstDepth = v
}

func (m *UserSettings) SetMaxAstElements(v *wrappers.Int64Value) {
	m.MaxAstElements = v
}

func (m *UserSettings) SetMaxExpandedAstElements(v *wrappers.Int64Value) {
	m.MaxExpandedAstElements = v
}

func (m *UserSettings) SetInputFormatValuesInterpretExpressions(v *wrappers.BoolValue) {
	m.InputFormatValuesInterpretExpressions = v
}

func (m *UserSettings) SetInputFormatDefaultsForOmittedFields(v *wrappers.BoolValue) {
	m.InputFormatDefaultsForOmittedFields = v
}

func (m *UserSettings) SetOutputFormatJsonQuote_64BitIntegers(v *wrappers.BoolValue) {
	m.OutputFormatJsonQuote_64BitIntegers = v
}

func (m *UserSettings) SetOutputFormatJsonQuoteDenormals(v *wrappers.BoolValue) {
	m.OutputFormatJsonQuoteDenormals = v
}

func (m *UserSettings) SetLowCardinalityAllowInNativeFormat(v *wrappers.BoolValue) {
	m.LowCardinalityAllowInNativeFormat = v
}

func (m *UserSettings) SetEmptyResultForAggregationByEmptySet(v *wrappers.BoolValue) {
	m.EmptyResultForAggregationByEmptySet = v
}

func (m *UserSettings) SetHttpConnectionTimeout(v *wrappers.Int64Value) {
	m.HttpConnectionTimeout = v
}

func (m *UserSettings) SetHttpReceiveTimeout(v *wrappers.Int64Value) {
	m.HttpReceiveTimeout = v
}

func (m *UserSettings) SetHttpSendTimeout(v *wrappers.Int64Value) {
	m.HttpSendTimeout = v
}

func (m *UserSettings) SetEnableHttpCompression(v *wrappers.BoolValue) {
	m.EnableHttpCompression = v
}

func (m *UserSettings) SetSendProgressInHttpHeaders(v *wrappers.BoolValue) {
	m.SendProgressInHttpHeaders = v
}

func (m *UserSettings) SetHttpHeadersProgressInterval(v *wrappers.Int64Value) {
	m.HttpHeadersProgressInterval = v
}

func (m *UserSettings) SetAddHttpCorsHeader(v *wrappers.BoolValue) {
	m.AddHttpCorsHeader = v
}

func (m *UserQuota) SetIntervalDuration(v *wrappers.Int64Value) {
	m.IntervalDuration = v
}

func (m *UserQuota) SetQueries(v *wrappers.Int64Value) {
	m.Queries = v
}

func (m *UserQuota) SetErrors(v *wrappers.Int64Value) {
	m.Errors = v
}

func (m *UserQuota) SetResultRows(v *wrappers.Int64Value) {
	m.ResultRows = v
}

func (m *UserQuota) SetReadRows(v *wrappers.Int64Value) {
	m.ReadRows = v
}

func (m *UserQuota) SetExecutionTime(v *wrappers.Int64Value) {
	m.ExecutionTime = v
}
