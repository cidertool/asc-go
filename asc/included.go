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
	"encoding/json"
	"fmt"
)

// ErrInvalidIncluded happens when an invalid "included" type is returned by the App Store Connect API.
// If this is encountered, it should be reported as a bug to the cidertool/asc-go repository issue
// tracker.
type ErrInvalidIncluded struct {
	Type string
}

func (e ErrInvalidIncluded) Error() string {
	return fmt.Sprintf("type %s not recognized as includable model", e.Type)
}

type included struct {
	Type  string
	inner interface{} // nolint: structcheck
}

func extractIncludedAgeRatingDeclaration(i interface{}) *AgeRatingDeclaration {
	if v, ok := i.(AgeRatingDeclaration); ok {
		return &v
	}

	return nil
}

func extractIncludedApp(i interface{}) *App {
	if v, ok := i.(App); ok {
		return &v
	}

	return nil
}

func extractIncludedAppCategory(i interface{}) *AppCategory {
	if v, ok := i.(AppCategory); ok {
		return &v
	}

	return nil
}

func extractIncludedAppEncryptionDeclaration(i interface{}) *AppEncryptionDeclaration {
	if v, ok := i.(AppEncryptionDeclaration); ok {
		return &v
	}

	return nil
}

func extractIncludedAppInfo(i interface{}) *AppInfo {
	if v, ok := i.(AppInfo); ok {
		return &v
	}

	return nil
}

func extractIncludedAppInfoLocalization(i interface{}) *AppInfoLocalization {
	if v, ok := i.(AppInfoLocalization); ok {
		return &v
	}

	return nil
}

func extractIncludedAppPreOrder(i interface{}) *AppPreOrder {
	if v, ok := i.(AppPreOrder); ok {
		return &v
	}

	return nil
}

func extractIncludedAppPreviewSet(i interface{}) *AppPreviewSet {
	if v, ok := i.(AppPreviewSet); ok {
		return &v
	}

	return nil
}

func extractIncludedAppPrice(i interface{}) *AppPrice {
	if v, ok := i.(AppPrice); ok {
		return &v
	}

	return nil
}

func extractIncludedAppScreenshotSet(i interface{}) *AppScreenshotSet {
	if v, ok := i.(AppScreenshotSet); ok {
		return &v
	}

	return nil
}

func extractIncludedAppStoreReviewDetail(i interface{}) *AppStoreReviewDetail {
	if v, ok := i.(AppStoreReviewDetail); ok {
		return &v
	}

	return nil
}

func extractIncludedAppStoreVersion(i interface{}) *AppStoreVersion {
	if v, ok := i.(AppStoreVersion); ok {
		return &v
	}

	return nil
}

func extractIncludedAppStoreVersionLocalization(i interface{}) *AppStoreVersionLocalization {
	if v, ok := i.(AppStoreVersionLocalization); ok {
		return &v
	}

	return nil
}

func extractIncludedAppStoreVersionPhasedRelease(i interface{}) *AppStoreVersionPhasedRelease {
	if v, ok := i.(AppStoreVersionPhasedRelease); ok {
		return &v
	}

	return nil
}

func extractIncludedAppStoreVersionSubmission(i interface{}) *AppStoreVersionSubmission {
	if v, ok := i.(AppStoreVersionSubmission); ok {
		return &v
	}

	return nil
}

func extractIncludedBetaAppLocalization(i interface{}) *BetaAppLocalization {
	if v, ok := i.(BetaAppLocalization); ok {
		return &v
	}

	return nil
}

func extractIncludedBetaAppReviewDetail(i interface{}) *BetaAppReviewDetail {
	if v, ok := i.(BetaAppReviewDetail); ok {
		return &v
	}

	return nil
}

func extractIncludedBetaAppReviewSubmission(i interface{}) *BetaAppReviewSubmission {
	if v, ok := i.(BetaAppReviewSubmission); ok {
		return &v
	}

	return nil
}

func extractIncludedBetaBuildLocalization(i interface{}) *BetaBuildLocalization {
	if v, ok := i.(BetaBuildLocalization); ok {
		return &v
	}

	return nil
}

func extractIncludedBetaGroup(i interface{}) *BetaGroup {
	if v, ok := i.(BetaGroup); ok {
		return &v
	}

	return nil
}

func extractIncludedBetaLicenseAgreement(i interface{}) *BetaLicenseAgreement {
	if v, ok := i.(BetaLicenseAgreement); ok {
		return &v
	}

	return nil
}

func extractIncludedBetaTester(i interface{}) *BetaTester {
	if v, ok := i.(BetaTester); ok {
		return &v
	}

	return nil
}

func extractIncludedBuild(i interface{}) *Build {
	if v, ok := i.(Build); ok {
		return &v
	}

	return nil
}

func extractIncludedBuildBetaDetail(i interface{}) *BuildBetaDetail {
	if v, ok := i.(BuildBetaDetail); ok {
		return &v
	}

	return nil
}

func extractIncludedBuildIcon(i interface{}) *BuildIcon {
	if v, ok := i.(BuildIcon); ok {
		return &v
	}

	return nil
}

func extractIncludedBundleID(i interface{}) *BundleID {
	if v, ok := i.(BundleID); ok {
		return &v
	}

	return nil
}

func extractIncludedBundleIDCapability(i interface{}) *BundleIDCapability {
	if v, ok := i.(BundleIDCapability); ok {
		return &v
	}

	return nil
}

func extractIncludedCertificate(i interface{}) *Certificate {
	if v, ok := i.(Certificate); ok {
		return &v
	}

	return nil
}

func extractIncludedDevice(i interface{}) *Device {
	if v, ok := i.(Device); ok {
		return &v
	}

	return nil
}

func extractIncludedDiagnosticSignature(i interface{}) *DiagnosticSignature {
	if v, ok := i.(DiagnosticSignature); ok {
		return &v
	}

	return nil
}

func extractIncludedEndUserLicenseAgreement(i interface{}) *EndUserLicenseAgreement {
	if v, ok := i.(EndUserLicenseAgreement); ok {
		return &v
	}

	return nil
}

func extractIncludedGameCenterEnabledVersion(i interface{}) *GameCenterEnabledVersion {
	if v, ok := i.(GameCenterEnabledVersion); ok {
		return &v
	}

	return nil
}

func extractIncludedIDFADeclaration(i interface{}) *IDFADeclaration {
	if v, ok := i.(IDFADeclaration); ok {
		return &v
	}

	return nil
}

func extractIncludedInAppPurchase(i interface{}) *InAppPurchase {
	if v, ok := i.(InAppPurchase); ok {
		return &v
	}

	return nil
}

func extractIncludedPerfPowerMetric(i interface{}) *PerfPowerMetric {
	if v, ok := i.(PerfPowerMetric); ok {
		return &v
	}

	return nil
}

func extractIncludedPrereleaseVersion(i interface{}) *PrereleaseVersion {
	if v, ok := i.(PrereleaseVersion); ok {
		return &v
	}

	return nil
}

func extractIncludedProfile(i interface{}) *Profile {
	if v, ok := i.(Profile); ok {
		return &v
	}

	return nil
}

func extractIncludedRoutingAppCoverage(i interface{}) *RoutingAppCoverage {
	if v, ok := i.(RoutingAppCoverage); ok {
		return &v
	}

	return nil
}

func extractIncludedTerritory(i interface{}) *Territory {
	if v, ok := i.(Territory); ok {
		return &v
	}

	return nil
}

func unmarshalInclude(b []byte) (string, interface{}, error) {
	var typeRef struct {
		Type string `json:"type"`
	}

	err := json.Unmarshal(b, &typeRef)
	if err != nil {
		return "", nil, err
	}

	typeName := typeRef.Type

	return supportedIncludeTypes()(typeName, b)
}

type includeTypeUnmarshallers map[string]func([]byte) (string, interface{}, error)

func supportedIncludeTypes() func(string, []byte) (string, interface{}, error) {
	allCases := includeTypeUnmarshallers{
		"ageRatingDeclarations": func(b []byte) (string, interface{}, error) {
			var v AgeRatingDeclaration
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"apps": func(b []byte) (string, interface{}, error) {
			var v App
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"appCategories": func(b []byte) (string, interface{}, error) {
			var v AppCategory
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"appEncryptionDeclarations": func(b []byte) (string, interface{}, error) {
			var v AppEncryptionDeclaration
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"appInfos": func(b []byte) (string, interface{}, error) {
			var v AppInfo
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"appInfoLocalizations": func(b []byte) (string, interface{}, error) {
			var v AppInfoLocalization
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"appPreOrders": func(b []byte) (string, interface{}, error) {
			var v AppPreOrder
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"appPreviewSets": func(b []byte) (string, interface{}, error) {
			var v AppPreviewSet
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"appPrices": func(b []byte) (string, interface{}, error) {
			var v AppPrice
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"appScreenshotSets": func(b []byte) (string, interface{}, error) {
			var v AppScreenshotSet
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"appStoreReviewDetails": func(b []byte) (string, interface{}, error) {
			var v AppStoreReviewDetail
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"appStoreVersions": func(b []byte) (string, interface{}, error) {
			var v AppStoreVersion
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"appStoreVersionLocalizations": func(b []byte) (string, interface{}, error) {
			var v AppStoreVersionLocalization
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"appStoreVersionPhasedReleases": func(b []byte) (string, interface{}, error) {
			var v AppStoreVersionPhasedRelease
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"appStoreVersionSubmissions": func(b []byte) (string, interface{}, error) {
			var v AppStoreVersionSubmission
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"betaAppLocalizations": func(b []byte) (string, interface{}, error) {
			var v BetaAppLocalization
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"betaAppReviewDetails": func(b []byte) (string, interface{}, error) {
			var v BetaAppReviewDetail
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"betaAppReviewSubmissions": func(b []byte) (string, interface{}, error) {
			var v BetaAppReviewSubmission
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"betaBuildLocalizations": func(b []byte) (string, interface{}, error) {
			var v BetaBuildLocalization
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"betaGroups": func(b []byte) (string, interface{}, error) {
			var v BetaGroup
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"betaLicenseAgreements": func(b []byte) (string, interface{}, error) {
			var v BetaLicenseAgreement
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"betaTesters": func(b []byte) (string, interface{}, error) {
			var v BetaTester
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"builds": func(b []byte) (string, interface{}, error) {
			var v Build
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"buildBetaDetails": func(b []byte) (string, interface{}, error) {
			var v BuildBetaDetail
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"buildIcons": func(b []byte) (string, interface{}, error) {
			var v BuildIcon
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"bundleIds": func(b []byte) (string, interface{}, error) {
			var v BundleID
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"bundleIdCapabilities": func(b []byte) (string, interface{}, error) {
			var v BundleIDCapability
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"certificates": func(b []byte) (string, interface{}, error) {
			var v Certificate
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"devices": func(b []byte) (string, interface{}, error) {
			var v Device
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"diagnosticSignatures": func(b []byte) (string, interface{}, error) {
			var v DiagnosticSignature
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"endUserLicenseAgreements": func(b []byte) (string, interface{}, error) {
			var v EndUserLicenseAgreement
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"gameCenterEnabledVersions": func(b []byte) (string, interface{}, error) {
			var v GameCenterEnabledVersion
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"idfaDeclarations": func(b []byte) (string, interface{}, error) {
			var v IDFADeclaration
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"inAppPurchases": func(b []byte) (string, interface{}, error) {
			var v InAppPurchase
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"perfPowerMetrics": func(b []byte) (string, interface{}, error) {
			var v PerfPowerMetric
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"preReleaseVersions": func(b []byte) (string, interface{}, error) {
			var v PrereleaseVersion
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"profiles": func(b []byte) (string, interface{}, error) {
			var v Profile
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"routingAppCoverages": func(b []byte) (string, interface{}, error) {
			var v RoutingAppCoverage
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
		"territories": func(b []byte) (string, interface{}, error) {
			var v Territory
			err := json.Unmarshal(b, &v)

			return v.Type, v, err
		},
	}

	return func(typeName string, b []byte) (string, interface{}, error) {
		if deser, ok := allCases[typeName]; ok {
			return deser(b)
		}

		return typeName, nil, ErrInvalidIncluded{Type: typeName}
	}
}
