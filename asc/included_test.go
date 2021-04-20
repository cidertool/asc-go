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
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockPayloadIncluded struct {
	Included []mockIncluded `json:"included,omitempty"`
}

type mockIncluded included

func (i *mockIncluded) UnmarshalJSON(b []byte) error {
	typeName, inner, err := unmarshalInclude(b)
	i.Type = typeName
	i.inner = inner

	return err
}

func TestIncluded(t *testing.T) {
	t.Parallel()

	var payload *mockPayloadIncluded

	payload = nil
	jsonNone := `{}`
	err := json.Unmarshal([]byte(jsonNone), &payload)
	assert.NoError(t, err)
	assert.Empty(t, payload.Included)

	payload = nil
	jsonGood := `{"included":[{"type":"apps","id":"10"}]}`
	err = json.Unmarshal([]byte(jsonGood), &payload)
	assert.NoError(t, err)
	assert.NotEmpty(t, payload.Included)

	payload = nil
	jsonInvalidType := `{"included":[{"type":"dog"}]}`
	err = json.Unmarshal([]byte(jsonInvalidType), &payload)
	assert.Error(t, err)

	payload = nil
	jsonInvalidStructure := `{"included":[{"type":-1}]}`
	err = json.Unmarshal([]byte(jsonInvalidStructure), &payload)
	assert.Error(t, err)
}

func TestKnownIncludeTypes(t *testing.T) {
	t.Parallel()

	knownTypes := []string{"ageRatingDeclarations", "apps", "appCategories", "appEncryptionDeclarations",
		"appInfos", "appInfoLocalizations", "appPreOrders", "appPreviewSets", "appPrices", "appScreenshotSets",
		"appStoreReviewDetails", "appStoreVersions", "appStoreVersionLocalizations", "appStoreVersionPhasedReleases",
		"appStoreVersionSubmissions", "betaAppLocalizations", "betaAppReviewDetails", "betaAppReviewSubmissions",
		"betaBuildLocalizations", "betaGroups", "betaLicenseAgreements", "betaTesters", "builds", "buildBetaDetails",
		"buildIcons", "bundleIds", "bundleIdCapabilities", "certificates", "devices", "diagnosticSignatures",
		"endUserLicenseAgreements", "gameCenterEnabledVersions", "idfaDeclarations", "inAppPurchases", "perfPowerMetrics",
		"preReleaseVersions", "profiles", "routingAppCoverages", "territories"}

	var payload *mockPayloadIncluded

	var err error

	for _, typeName := range knownTypes {
		payload = nil
		jsonString := fmt.Sprintf(`{"included":[{"type":"%s","id":"10"}]}`, typeName)
		err = json.Unmarshal([]byte(jsonString), &payload)
		assert.NoError(t, err)
		assert.NotEmpty(t, payload.Included)
	}
}
