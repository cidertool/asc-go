package asc

import (
	"encoding/json"
	"fmt"
)

type included struct {
	Type  string
	inner interface{}
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

func unmarshalInclude(b []byte) (typeName string, i interface{}, err error) {
	var typeRef struct {
		Type string `json:"type"`
	}
	err = json.Unmarshal(b, &typeRef)
	if err != nil {
		return "", nil, err
	}
	switch typeRef.Type {
	case "ageRatingDeclarations":
		var v AgeRatingDeclaration
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "apps":
		var v App
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "appCategories":
		var v AppCategory
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "appEncryptionDeclarations":
		var v AppEncryptionDeclaration
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "appInfos":
		var v AppInfo
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "appInfoLocalizations":
		var v AppInfoLocalization
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "appPreOrders":
		var v AppPreOrder
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "appPreviewSets":
		var v AppPreviewSet
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "appPrices":
		var v AppPrice
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "appScreenshotSets":
		var v AppScreenshotSet
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "appStoreReviewDetails":
		var v AppStoreReviewDetail
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "appStoreVersions":
		var v AppStoreVersion
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "appStoreVersionLocalizations":
		var v AppStoreVersionLocalization
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "appStoreVersionPhasedReleases":
		var v AppStoreVersionPhasedRelease
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "appStoreVersionSubmissions":
		var v AppStoreVersionSubmission
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "betaAppLocalizations":
		var v BetaAppLocalization
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "betaAppReviewDetails":
		var v BetaAppReviewDetail
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "betaAppReviewSubmissions":
		var v BetaAppReviewSubmission
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "betaBuildLocalizations":
		var v BetaBuildLocalization
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "betaGroups":
		var v BetaGroup
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "betaLicenseAgreements":
		var v BetaLicenseAgreement
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "betaTesters":
		var v BetaTester
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "builds":
		var v Build
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "buildBetaDetails":
		var v BuildBetaDetail
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "buildIcons":
		var v BuildIcon
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "bundleIds":
		var v BundleID
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "bundleIdCapabilities":
		var v BundleIDCapability
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "certificates":
		var v Certificate
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "devices":
		var v Device
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "diagnosticSignatures":
		var v DiagnosticSignature
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "endUserLicenseAgreements":
		var v EndUserLicenseAgreement
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "gameCenterEnabledVersions":
		var v GameCenterEnabledVersion
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "idfaDeclarations":
		var v IDFADeclaration
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "inAppPurchases":
		var v InAppPurchase
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "perfPowerMetrics":
		var v PerfPowerMetric
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "preReleaseVersions":
		var v PrereleaseVersion
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "profiles":
		var v Profile
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "routingAppCoverages":
		var v RoutingAppCoverage
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	case "territories":
		var v Territory
		err := json.Unmarshal(b, &v)
		return typeRef.Type, v, err
	default:
		return typeRef.Type, nil, fmt.Errorf("type %s not recognized as includable model", typeRef.Type)
	}

}
