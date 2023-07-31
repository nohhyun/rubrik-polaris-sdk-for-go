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

// addK8sResourceset GraphQL query
var addK8sResourcesetQuery = `mutation SdkGolangAddK8sResourceset($config: K8sResourceSetAddInput!) {
  addK8sResourceSet(input: {config: $config}) {
    id
    name
    k8SClusterUuid
    k8SNamespace
    rsType
    definition
    kubernetesNamespace
    kubernetesClusterUuid
    hookConfigs
  }
}`

// createK8sResourceSnapshot GraphQL query
var createK8sResourceSnapshotQuery = `mutation SdkGolangCreateK8sResourceSnapshot($resourceSetId: String!, $jobConfig: BaseOnDemandSnapshotConfigInput!) {
    createK8sResourceSetSnapshot(input: {id: $resourceSetId, config: $jobConfig}) {
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

// deleteK8sResourceset GraphQL query
var deleteK8sResourcesetQuery = `mutation SdkGolangDeleteK8sResourceset($id:String!, $preserveSnapshots: Boolean) {
  deleteK8sResourceSet(input: {id: $id, preserveSnapshots: $preserveSnapshots}) {
    success
  }
}`

// exportK8sResourcesetSnapshot GraphQL query
var exportK8sResourcesetSnapshotQuery = `mutation SdkGolangExportK8sResourcesetSnapshot($id: String!, $jobConfig: K8sExportParametersInput!) {
  exportK8sResourceSetSnapshot(input: {id: $id, jobConfig: $jobConfig}) {
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

// jobInstance GraphQL query
var jobInstanceQuery = `query SdkGolangJobInstance($id:String!, $clusterUuid: String!) {
  jobInstance(input: {id:$id, clusterUuid: $clusterUuid}) {
    archived
    endTime
    errorInfo
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

// k8sResourceset GraphQL query
var k8sResourcesetQuery = `query SdkGolangK8sResourceset($fid: UUID!) {
  kubernetesResourceSet(fid: $fid) {
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
