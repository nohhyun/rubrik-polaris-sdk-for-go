// Code generated by queries_gen.go DO NOT EDIT

// MIT License
//
// Copyright (c) 2021 Rubrik
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package infinityk8s

// activitySeries GraphQL query
var activitySeriesQuery = `query SdkGolangActivitySeries(
    $input: ActivitySeriesInput!
) {
    activitySeries(input: $input) {
        activityConnection {
            nodes {
                activityInfo
                message
                status
                time
                severity
            }
            pageInfo {
                endCursor
                hasNextPage
            }
            count
        }
    }
}`

// addK8sProtectionSet GraphQL query
var addK8sProtectionSetQuery = `mutation SdkGolangAddK8sProtectionSet($config: K8sProtectionSetAddInput!) {
  addK8sProtectionSet(input: {config: $config}) {
    id
    name
    rsType
    definition
    kubernetesNamespace
    kubernetesClusterUuid
    hookConfigs
  }
}`

// createK8sProtectionSetSnapshot GraphQL query
var createK8sProtectionSetSnapshotQuery = `mutation SdkGolangCreateK8sProtectionSetSnapshot($protectionSetId: String!, $jobConfig: BaseOnDemandSnapshotConfigInput!) {
    createK8sProtectionSetSnapshot(input: {id: $protectionSetId, config: $jobConfig}) {
        endTime
        error {
            message
        }
        id
        nodeId
        progress
        startTime
        status
        links {
            href
            rel
        }
    }
}`

// deleteK8sProtectionSet GraphQL query
var deleteK8sProtectionSetQuery = `mutation SdkGolangDeleteK8sProtectionSet($id:String!, $preserveSnapshots: Boolean) {
  deleteK8sProtectionSet(input: {id: $id, preserveSnapshots: $preserveSnapshots}) {
    success
  }
}`

// exportK8sProtectionSetSnapshot GraphQL query
var exportK8sProtectionSetSnapshotQuery = `mutation SdkGolangExportK8sProtectionSetSnapshot($id: String!, $jobConfig: K8sExportParametersInput!) {
  exportK8sProtectionSetSnapshot(input: {id: $id, jobConfig: $jobConfig}) {
    endTime
    error {
      message
    }
    id
    links {
      rel
      href
    }
    nodeId
    progress
    startTime
    status
  }
}`

// getProtectionSetSnapshot GraphQL query
var getProtectionSetSnapshotQuery = `query SdkGolangK8sProtectionSetSnapshots($fid: String!) {
    k8sProtectionSetSnapshots(input: {id: $fid}) {
        data {
            baseSnapshotSummary {
                id
                slaId
            }
        }
    }
}`

// jobInstance GraphQL query
var jobInstanceQuery = `query SdkGolangJobInstance($id:String!, $clusterUuid: String!) {
  jobInstance(input: {id:$id, clusterUuid: $clusterUuid}) {
    archived
    endTime
    errorInfo
    eventSeriesId
    id
    isDisabled
    jobProgress
    jobType
    nodeId
    result
    startTime
    status
    childJobDebugInfo
    opentracingContext
  }
}`

// k8sObjectFid GraphQL query
var k8sObjectFidQuery = `query SdkGolangK8sObjectFid($k8SObjectInternalIdArg: UUID!, $clusterUuid: UUID!) {
  k8sObjectFid(K8sObjectInternalIDArg: $k8SObjectInternalIdArg, clusterUuid: $clusterUuid)
}`

// k8sObjectInternalId GraphQL query
var k8sObjectInternalIdQuery = `query SdkGolangK8sObjectInternalId($fid: UUID!) {
  k8sObjectInternalId(fid: $fid)
}`

// k8sProtectionSet GraphQL query
var k8sProtectionSetQuery = `query SdkGolangK8sProtectionSet($fid: UUID!) {
  kubernetesProtectionSet(fid: $fid) {
    cdmId
    clusterUuid
    configuredSlaDomain {
      id
      name
      version
    }
    effectiveRetentionSlaDomain {
      id
      name
      version
    }
    effectiveSlaDomain {
      id
      name
      version
    }
    effectiveSlaSourceObject {
      fid
      name
      objectType
    }
    id
    isRelic
    k8sClusterUuid
    name
    namespace
    objectType
    pendingSla {
      id
      name
      version
    }
    primaryClusterLocation {
      clusterUuid
      createDate
      id
      isActive
      isArchived
      name
      type
    }
    primaryClusterUuid
    replicatedObjectCount
    rsName
    rsType
    slaAssignment
    slaPauseStatus
  }
}`

// restoreK8sProtectionSetSnapshot GraphQL query
var restoreK8sProtectionSetSnapshotQuery = `mutation SdkGolangRestoreK8sProtectionSetSnapshot($id: String!, $jobConfig: K8sRestoreParametersInput!) {
  restoreK8sProtectionSetSnapshot(input: {id: $id, jobConfig: $jobConfig}) {
    endTime
    error {
      message
    }
    id
    links {
      rel
      href
    }
    nodeId
    progress
    startTime
    status
  }
}`

// updateK8sProtectionSet GraphQL query
var updateK8sProtectionSetQuery = `mutation SdkGolangUpdateK8sProtectionSet(
    $id: String!
    $updateConfig: K8sProtectionSetUpdateConfigInput!
) {
    updateK8sProtectionSet(input: { id: $id, config: $updateConfig }) {
        success
    }
}`
