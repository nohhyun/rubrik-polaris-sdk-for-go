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

package gcp

// allFeaturePermissionsForGcpCloudAccount GraphQL query
var allFeaturePermissionsForGcpCloudAccountQuery = `query SdkGolangAllFeaturePermissionsForGcpCloudAccount($feature: CloudAccountFeature!) {
    result: allFeaturePermissionsForGcpCloudAccount(feature: $feature){
        permission
    }
}`

// allGcpCloudAccountProjectsByFeature GraphQL query
var allGcpCloudAccountProjectsByFeatureQuery = `query SdkGolangAllGcpCloudAccountProjectsByFeature($feature: CloudAccountFeature!, $projectSearchText: String!) {
    result: allGcpCloudAccountProjectsByFeature(feature: $feature, projectStatusFilters: [], projectSearchText: $projectSearchText) {
        project {
            id
            name
            projectId
            projectNumber
            roleId
            usesGlobalConfig
        }
        featureDetail {
            feature
            status
        }
    }
}`

// gcpCloudAccountAddManualAuthProject GraphQL query
var gcpCloudAccountAddManualAuthProjectQuery = `mutation SdkGolangGcpCloudAccountAddManualAuthProject($gcpNativeProjectId: String!, $gcpProjectName: String!, $gcpProjectNumber: Long!, $organizationName: String, $serviceAccountJwtConfig: String, $feature: CloudAccountFeature!) {
    gcpCloudAccountAddManualAuthProject(input: {
        gcpNativeProjectId:      $gcpNativeProjectId,
        gcpProjectName:          $gcpProjectName,
        gcpProjectNumber:        $gcpProjectNumber,
        organizationName:        $organizationName,
        serviceAccountJwtConfig: $serviceAccountJwtConfig,
        features:                [$feature],
    })
}`

// gcpCloudAccountDeleteProjects GraphQL query
var gcpCloudAccountDeleteProjectsQuery = `mutation SdkGolangGcpCloudAccountDeleteProjects($nativeProtectionProjectId: UUID!) {
    result: gcpCloudAccountDeleteProjects(input: {
        nativeProtectionProjectIds: [$nativeProtectionProjectId],
        sharedVpcHostProjectIds:    [],
        cloudAccountsProjectIds:    [],
        skipResourceDeletion:       true,
    }) {
        gcpProjectDeleteStatuses {
            projectUuid
            success
            error
        }
    }
}`

// gcpGetDefaultCredentialsServiceAccount GraphQL query
var gcpGetDefaultCredentialsServiceAccountQuery = `query SdkGolangGcpGetDefaultCredentialsServiceAccount {
    gcpGetDefaultCredentialsServiceAccount
}`

// gcpNativeDisableProject GraphQL query
var gcpNativeDisableProjectQuery = `mutation SdkGolangGcpNativeDisableProject($projectId: UUID!, $shouldDeleteNativeSnapshots: Boolean!) {
  gcpNativeDisableProject(input: {
    projectId:                   $projectId,
    shouldDeleteNativeSnapshots: $shouldDeleteNativeSnapshots
  }) {
    jobId
    error
  }
}`

// gcpNativeProject GraphQL query
var gcpNativeProjectQuery = `query SdkGolangGcpNativeProject($fid: UUID!) {
    gcpNativeProject(fid: $fid) {
        id
        name
        nativeId
        nativeName
        projectNumber
        organizationName
        slaAssignment
        configuredSlaDomain{
            id
            name
        }
        effectiveSlaDomain{
            id
            name
        }
    }
}`

// gcpNativeProjects GraphQL query
var gcpNativeProjectsQuery = `query SdkGolangGcpNativeProjects($after: String, $filter: String!) {
    result: gcpNativeProjects(after: $after, projectFilters: {nameOrNumberSubstringFilter: {nameOrNumberSubstring: $filter}}){
        count
        edges {
            node {
                id
                name
                nativeId
                nativeName
                projectNumber
                organizationName
                slaAssignment
                configuredSlaDomain{
                    id
                    name
                }
                effectiveSlaDomain{
                    id
                    name
                }
            }
        }
        pageInfo {
            endCursor
            hasNextPage
        }
    }
}`

// gcpSetDefaultServiceAccountJwtConfig GraphQL query
var gcpSetDefaultServiceAccountJwtConfigQuery = `mutation SdkGolangGcpSetDefaultServiceAccountJwtConfig($serviceAccountName: String!, $serviceAccountJwtConfig: String!) {
    gcpSetDefaultServiceAccountJwtConfig(input: {
        serviceAccountJwtConfig: $serviceAccountJwtConfig,
        serviceAccountName:      $serviceAccountName,
    })
}`

// upgradeGcpCloudAccountPermissionsWithoutOauth GraphQL query
var upgradeGcpCloudAccountPermissionsWithoutOauthQuery = `mutation SdkGolangUpgradeGcpCloudAccountPermissionsWithoutOauth($cloudAccountId: UUID!, $feature: CloudAccountFeature!) {
    result: upgradeGcpCloudAccountPermissionsWithoutOauth(input: {
        cloudAccountId: $cloudAccountId,
        feature:        $feature
    }) {
        status {
            projectUuid
            success
            error
        }
    }
}`
