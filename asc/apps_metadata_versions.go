/**
Copyright (C) 2020 Aaron Sky.

This file is part of asc-go, a package for working with Apple's
App Store Connect API.

asc-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

asc-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with asc-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asc

import (
	"context"
	"fmt"
)

// AppStoreVersionState defines model for AppStoreVersionState.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionstate
type AppStoreVersionState string

const (
	// AppStoreVersionStateDeveloperRejected is an app store version state for DeveloperRejected.
	AppStoreVersionStateDeveloperRejected AppStoreVersionState = "DEVELOPER_REJECTED"
	// AppStoreVersionStateDeveloperRemovedFromSale is an app store version state for DeveloperRemovedFromSale.
	AppStoreVersionStateDeveloperRemovedFromSale AppStoreVersionState = "DEVELOPER_REMOVED_FROM_SALE"
	// AppStoreVersionStateInvalidBinary is an app store version state for InvalidBinary.
	AppStoreVersionStateInvalidBinary AppStoreVersionState = "INVALID_BINARY"
	// AppStoreVersionStateInReview is an app store version state for InReview.
	AppStoreVersionStateInReview AppStoreVersionState = "IN_REVIEW"
	// AppStoreVersionStateMetadataRejected is an app store version state for MetadataRejected.
	AppStoreVersionStateMetadataRejected AppStoreVersionState = "METADATA_REJECTED"
	// AppStoreVersionStatePendingAppleRelease is an app store version state for PendingAppleRelease.
	AppStoreVersionStatePendingAppleRelease AppStoreVersionState = "PENDING_APPLE_RELEASE"
	// AppStoreVersionStatePendingContract is an app store version state for PendingContract.
	AppStoreVersionStatePendingContract AppStoreVersionState = "PENDING_CONTRACT"
	// AppStoreVersionStatePendingDeveloperRelease is an app store version state for PendingDeveloperRelease.
	AppStoreVersionStatePendingDeveloperRelease AppStoreVersionState = "PENDING_DEVELOPER_RELEASE"
	// AppStoreVersionStatePreorderReadyForSale is an app store version state for PreorderReadyForSale.
	AppStoreVersionStatePreorderReadyForSale AppStoreVersionState = "PREORDER_READY_FOR_SALE"
	// AppStoreVersionStatePrepareForSubmission is an app store version state for PrepareForSubmission.
	AppStoreVersionStatePrepareForSubmission AppStoreVersionState = "PREPARE_FOR_SUBMISSION"
	// AppStoreVersionStateProcessingForAppStore is an app store version state for ProcessingForAppStore.
	AppStoreVersionStateProcessingForAppStore AppStoreVersionState = "PROCESSING_FOR_APP_STORE"
	// AppStoreVersionStateReadyForSale is an app store version state for ReadyForSale.
	AppStoreVersionStateReadyForSale AppStoreVersionState = "READY_FOR_SALE"
	// AppStoreVersionStateRejected is an app store version state for Rejected.
	AppStoreVersionStateRejected AppStoreVersionState = "REJECTED"
	// AppStoreVersionStateRemovedFromSale is an app store version state for RemovedFromSale.
	AppStoreVersionStateRemovedFromSale AppStoreVersionState = "REMOVED_FROM_SALE"
	// AppStoreVersionStateReplacedWithNewVersion is an app store version state for ReplacedWithNewVersion.
	AppStoreVersionStateReplacedWithNewVersion AppStoreVersionState = "REPLACED_WITH_NEW_VERSION"
	// AppStoreVersionStateWaitingForExportCompliance is an app store version state for WaitingForExportCompliance.
	AppStoreVersionStateWaitingForExportCompliance AppStoreVersionState = "WAITING_FOR_EXPORT_COMPLIANCE"
	// AppStoreVersionStateWaitingForReview is an app store version state for WaitingForReview.
	AppStoreVersionStateWaitingForReview AppStoreVersionState = "WAITING_FOR_REVIEW"
)

// AppStoreVersionUpdateRequest defines model for AppStoreVersionUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionupdaterequest/data
type appStoreVersionUpdateRequest struct {
	Attributes    *AppStoreVersionUpdateRequestAttributes    `json:"attributes,omitempty"`
	ID            string                                     `json:"id"`
	Relationships *appStoreVersionUpdateRequestRelationships `json:"relationships,omitempty"`
	Type          string                                     `json:"type"`
}

// AppStoreVersionUpdateRequestAttributes are attributes for AppStoreVersionUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionupdaterequest/data/attributes
type AppStoreVersionUpdateRequestAttributes struct {
	Copyright           *string   `json:"copyright,omitempty"`
	Downloadable        *bool     `json:"downloadable,omitempty"`
	EarliestReleaseDate *DateTime `json:"earliestReleaseDate,omitempty"`
	ReleaseType         *string   `json:"releaseType,omitempty"`
	UsesIDFA            *bool     `json:"usesIdfa,omitempty"`
	VersionString       *string   `json:"versionString,omitempty"`
}

// appStoreVersionUpdateRequestRelationships are relationships for AppStoreVersionUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionupdaterequest/data/relationships
type appStoreVersionUpdateRequestRelationships struct {
	Build *relationshipDeclaration `json:"build,omitempty"`
}

// AgeRatingDeclaration defines model for AgeRatingDeclaration.
//
// https://developer.apple.com/documentation/appstoreconnectapi/ageratingdeclaration
type AgeRatingDeclaration struct {
	Attributes *AgeRatingDeclarationAttributes `json:"attributes,omitempty"`
	ID         string                          `json:"id"`
	Links      ResourceLinks                   `json:"links"`
	Type       string                          `json:"type"`
}

// AgeRatingDeclarationAttributes defines model for AgeRatingDeclaration.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/ageratingdeclaration/attributes
type AgeRatingDeclarationAttributes struct {
	AlcoholTobaccoOrDrugUseOrReferences         *string      `json:"alcoholTobaccoOrDrugUseOrReferences,omitempty"`
	Contests                                    *string      `json:"contests,omitempty"`
	Gambling                                    *bool        `json:"gambling,omitempty"`
	GamblingSimulated                           *string      `json:"gamblingSimulated,omitempty"`
	HorrorOrFearThemes                          *string      `json:"horrorOrFearThemes,omitempty"`
	KidsAgeBand                                 *KidsAgeBand `json:"kidsAgeBand,omitempty"`
	MatureOrSuggestiveThemes                    *string      `json:"matureOrSuggestiveThemes,omitempty"`
	MedicalOrTreatmentInformation               *string      `json:"medicalOrTreatmentInformation,omitempty"`
	ProfanityOrCrudeHumor                       *string      `json:"profanityOrCrudeHumor,omitempty"`
	SexualContentGraphicAndNudity               *string      `json:"sexualContentGraphicAndNudity,omitempty"`
	SexualContentOrNudity                       *string      `json:"sexualContentOrNudity,omitempty"`
	SeventeenPlus                               *bool        `json:"seventeenPlus,omitempty"`
	UnrestrictedWebAccess                       *bool        `json:"unrestrictedWebAccess,omitempty"`
	ViolenceCartoonOrFantasy                    *string      `json:"violenceCartoonOrFantasy,omitempty"`
	ViolenceRealistic                           *string      `json:"violenceRealistic,omitempty"`
	ViolenceRealisticProlongedGraphicOrSadistic *string      `json:"violenceRealisticProlongedGraphicOrSadistic,omitempty"`
}

// AppStoreVersion defines model for AppStoreVersion.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversion
type AppStoreVersion struct {
	Attributes    *AppStoreVersionAttributes    `json:"attributes,omitempty"`
	ID            string                        `json:"id"`
	Links         ResourceLinks                 `json:"links"`
	Relationships *AppStoreVersionRelationships `json:"relationships,omitempty"`
	Type          string                        `json:"type"`
}

// AppStoreVersionAttributes defines model for AppStoreVersion.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversion/attributes
type AppStoreVersionAttributes struct {
	AppStoreState       *AppStoreVersionState `json:"appStoreState,omitempty"`
	Copyright           *string               `json:"copyright,omitempty"`
	CreatedDate         *DateTime             `json:"createdDate,omitempty"`
	Downloadable        *bool                 `json:"downloadable,omitempty"`
	EarliestReleaseDate *DateTime             `json:"earliestReleaseDate,omitempty"`
	Platform            *Platform             `json:"platform,omitempty"`
	ReleaseType         *string               `json:"releaseType,omitempty"`
	UsesIDFA            *bool                 `json:"usesIdfa,omitempty"`
	VersionString       *string               `json:"versionString,omitempty"`
}

// AppStoreVersionRelationships defines model for AppStoreVersion.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversion/relationships
type AppStoreVersionRelationships struct {
	App                          *Relationship      `json:"app,omitempty"`
	AppStoreReviewDetail         *Relationship      `json:"appStoreReviewDetail,omitempty"`
	AppStoreVersionLocalizations *PagedRelationship `json:"appStoreVersionLocalizations,omitempty"`
	AppStoreVersionPhasedRelease *Relationship      `json:"appStoreVersionPhasedRelease,omitempty"`
	AppStoreVersionSubmission    *Relationship      `json:"appStoreVersionSubmission,omitempty"`
	Build                        *Relationship      `json:"build,omitempty"`
	IDFADeclaration              *Relationship      `json:"idfaDeclaration,omitempty"`
	RoutingAppCoverage           *Relationship      `json:"routingAppCoverage,omitempty"`
}

// AppStoreVersionResponse defines model for AppStoreVersionResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionresponse
type AppStoreVersionResponse struct {
	Data     AppStoreVersion                   `json:"data"`
	Included []AppStoreVersionResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks                     `json:"links"`
}

// AppStoreVersionsResponse defines model for AppStoreVersionsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionsresponse
type AppStoreVersionsResponse struct {
	Data     []AppStoreVersion                 `json:"data"`
	Included []AppStoreVersionResponseIncluded `json:"included,omitempty"`
	Links    PagedDocumentLinks                `json:"links"`
	Meta     *PagingInformation                `json:"meta,omitempty"`
}

// AppStoreVersionResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a AppStoreVersionResponse or AppStoreVersionsResponse.
type AppStoreVersionResponseIncluded included

// AppStoreVersionCreateRequest defines model for AppStoreVersionCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversioncreaterequest/data
type appStoreVersionCreateRequest struct {
	Attributes    AppStoreVersionCreateRequestAttributes    `json:"attributes"`
	Relationships appStoreVersionCreateRequestRelationships `json:"relationships"`
	Type          string                                    `json:"type"`
}

// AppStoreVersionCreateRequestAttributes are attributes for AppStoreVersionCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversioncreaterequest/data/attributes
type AppStoreVersionCreateRequestAttributes struct {
	Copyright           *string   `json:"copyright,omitempty"`
	EarliestReleaseDate *DateTime `json:"earliestReleaseDate,omitempty"`
	Platform            Platform  `json:"platform"`
	ReleaseType         *string   `json:"releaseType,omitempty"`
	UsesIDFA            *bool     `json:"usesIdfa,omitempty"`
	VersionString       string    `json:"versionString"`
}

// AppStoreVersionCreateRequestRelationships are relationships for AppStoreVersionCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversioncreaterequest/data/relationships
type appStoreVersionCreateRequestRelationships struct {
	App   relationshipDeclaration  `json:"app"`
	Build *relationshipDeclaration `json:"build,omitempty"`
}

// AppStoreVersionBuildLinkageResponse defines model for AppStoreVersionBuildLinkageResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionbuildlinkageresponse
type AppStoreVersionBuildLinkageResponse struct {
	Data  RelationshipData `json:"data"`
	Links DocumentLinks    `json:"links"`
}

// ListAppStoreVersionsQuery are query options for ListAppStoreVersions
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_store_versions_for_an_app
type ListAppStoreVersionsQuery struct {
	FieldsApps                          []string `url:"fields[apps],omitempty"`
	FieldsAppStoreVersionSubmissions    []string `url:"fields[appStoreVersionSubmissions],omitempty"`
	FieldsBuilds                        []string `url:"fields[builds],omitempty"`
	FieldsAppStoreVersions              []string `url:"fields[appStoreVersions],omitempty"`
	FieldsAppStoreReviewDetails         []string `url:"fields[appStoreReviewDetails],omitempty"`
	FieldsAgeRatingDeclarations         []string `url:"fields[ageRatingDeclarations],omitempty"`
	FieldsAppStoreVersionPhasedReleases []string `url:"fields[appStoreVersionPhasedReleases],omitempty"`
	FieldsRoutingAppCoverages           []string `url:"fields[routingAppCoverages],omitempty"`
	FieldsIDFADeclarations              []string `url:"fields[idfaDeclarations],omitempty"`
	Limit                               int      `url:"limit,omitempty"`
	Include                             []string `url:"include,omitempty"`
	FilterID                            []string `url:"filter[id],omitempty"`
	FilterVersionString                 []string `url:"filter[versionString],omitempty"`
	FilterPlatform                      []string `url:"filter[platform],omitempty"`
	FilterAppStoreState                 []string `url:"filter[appStoreState],omitempty"`
	Cursor                              string   `url:"cursor,omitempty"`
}

// GetAppStoreVersionQuery are query options for GetAppStoreVersion
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_store_version_information
type GetAppStoreVersionQuery struct {
	FieldsAppStoreVersions              []string `url:"fields[appStoreVersions],omitempty"`
	FieldsAppStoreVersionSubmissions    []string `url:"fields[appStoreVersionSubmissions],omitempty"`
	FieldsBuilds                        []string `url:"fields[builds],omitempty"`
	FieldsAppStoreReviewDetails         []string `url:"fields[appStoreReviewDetails],omitempty"`
	FieldsAppStoreVersionPhasedReleases []string `url:"fields[appStoreVersionPhasedReleases],omitempty"`
	FieldsRoutingAppCoverages           []string `url:"fields[routingAppCoverages],omitempty"`
	FieldsIDFADeclarations              []string `url:"fields[idfaDeclarations],omitempty"`
	FieldsAppStoreVersionLocalizations  []string `url:"fields[appStoreVersionLocalizations],omitempty"`
	Include                             []string `url:"include,omitempty"`
	LimitAppStoreVersionLocalizations   int      `url:"limit[appStoreVersionLocalizations],omitempty"`
}

// ListAppStoreVersionsForApp gets a list of all App Store versions of an app across all platforms.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_all_app_store_versions_for_an_app
func (s *AppsService) ListAppStoreVersionsForApp(ctx context.Context, id string, params *ListAppStoreVersionsQuery) (*AppStoreVersionsResponse, *Response, error) {
	url := fmt.Sprintf("apps/%s/appStoreVersions", id)
	res := new(AppStoreVersionsResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// GetAppStoreVersion gets information for a specific app store version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_app_store_version_information
func (s *AppsService) GetAppStoreVersion(ctx context.Context, id string, params *GetAppStoreVersionQuery) (*AppStoreVersionResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s", id)
	res := new(AppStoreVersionResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// CreateAppStoreVersion adds a new App Store version or platform to an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_an_app_store_version
func (s *AppsService) CreateAppStoreVersion(ctx context.Context, attributes AppStoreVersionCreateRequestAttributes, appID string, buildID *string) (*AppStoreVersionResponse, *Response, error) {
	req := appStoreVersionCreateRequest{
		Attributes: attributes,
		Relationships: appStoreVersionCreateRequestRelationships{
			App: *newRelationshipDeclaration(&appID, "apps"),
		},
		Type: "appStoreVersions",
	}

	if buildID != nil {
		req.Relationships.Build = newRelationshipDeclaration(buildID, "builds")
	}

	res := new(AppStoreVersionResponse)
	resp, err := s.client.post(ctx, "appStoreVersions", newRequestBody(req), res)

	return res, resp, err
}

// UpdateAppStoreVersion updates the app store version for a specific app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_app_store_version
func (s *AppsService) UpdateAppStoreVersion(ctx context.Context, id string, attributes *AppStoreVersionUpdateRequestAttributes, buildID *string) (*AppStoreVersionResponse, *Response, error) {
	req := appStoreVersionUpdateRequest{
		Attributes: attributes,
		ID:         id,
		Type:       "appStoreVersions",
	}

	if buildID != nil {
		req.Relationships = &appStoreVersionUpdateRequestRelationships{
			Build: newRelationshipDeclaration(buildID, "builds"),
		}
	}

	url := fmt.Sprintf("appStoreVersions/%s", id)
	res := new(AppStoreVersionResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// DeleteAppStoreVersion deletes an app store version that is associated with an app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/delete_an_app_store_version
func (s *AppsService) DeleteAppStoreVersion(ctx context.Context, id string) (*Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s", id)

	return s.client.delete(ctx, url, nil)
}

// GetBuildIDForAppStoreVersion gets the ID of the build that is attached to a specific App Store version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/get_the_build_id_for_an_app_store_version
func (s *AppsService) GetBuildIDForAppStoreVersion(ctx context.Context, id string) (*AppStoreVersionBuildLinkageResponse, *Response, error) {
	url := fmt.Sprintf("appStoreVersions/%s/relationships/build", id)
	res := new(AppStoreVersionBuildLinkageResponse)
	resp, err := s.client.get(ctx, url, nil, res)

	return res, resp, err
}

// UpdateBuildForAppStoreVersion changes the build that is attached to a specific App Store version.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_the_build_for_an_app_store_version
func (s *AppsService) UpdateBuildForAppStoreVersion(ctx context.Context, id string, buildID *string) (*AppStoreVersionBuildLinkageResponse, *Response, error) {
	linkage := newRelationshipDeclaration(buildID, "builds")
	url := fmt.Sprintf("appStoreVersions/%s/relationships/build", id)
	res := new(AppStoreVersionBuildLinkageResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(linkage), res)

	return res, resp, err
}

// UnmarshalJSON is a custom unmarshaller for the heterogenous data stored in AppStoreVersionResponseIncluded.
func (i *AppStoreVersionResponseIncluded) UnmarshalJSON(b []byte) error {
	typeName, inner, err := unmarshalInclude(b)
	i.Type = typeName
	i.inner = inner

	return err
}

// AgeRatingDeclaration returns the AgeRatingDeclaration stored within, if one is present.
func (i *AppStoreVersionResponseIncluded) AgeRatingDeclaration() *AgeRatingDeclaration {
	return extractIncludedAgeRatingDeclaration(i.inner)
}

// AppStoreVersionLocalization returns the AppStoreVersionLocalization stored within, if one is present.
func (i *AppStoreVersionResponseIncluded) AppStoreVersionLocalization() *AppStoreVersionLocalization {
	return extractIncludedAppStoreVersionLocalization(i.inner)
}

// Build returns the Build stored within, if one is present.
func (i *AppStoreVersionResponseIncluded) Build() *Build {
	return extractIncludedBuild(i.inner)
}

// AppStoreVersionPhasedRelease returns the AppStoreVersionPhasedRelease stored within, if one is present.
func (i *AppStoreVersionResponseIncluded) AppStoreVersionPhasedRelease() *AppStoreVersionPhasedRelease {
	return extractIncludedAppStoreVersionPhasedRelease(i.inner)
}

// RoutingAppCoverage returns the RoutingAppCoverage stored within, if one is present.
func (i *AppStoreVersionResponseIncluded) RoutingAppCoverage() *RoutingAppCoverage {
	return extractIncludedRoutingAppCoverage(i.inner)
}

// AppStoreReviewDetail returns the AppStoreReviewDetail stored within, if one is present.
func (i *AppStoreVersionResponseIncluded) AppStoreReviewDetail() *AppStoreReviewDetail {
	return extractIncludedAppStoreReviewDetail(i.inner)
}

// AppStoreVersionSubmission returns the AppStoreVersionSubmission stored within, if one is present.
func (i *AppStoreVersionResponseIncluded) AppStoreVersionSubmission() *AppStoreVersionSubmission {
	return extractIncludedAppStoreVersionSubmission(i.inner)
}

// IDFADeclaration returns the IDFADeclaration stored within, if one is present.
func (i *AppStoreVersionResponseIncluded) IDFADeclaration() *IDFADeclaration {
	return extractIncludedIDFADeclaration(i.inner)
}
