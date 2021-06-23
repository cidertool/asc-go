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

// AppStoreAgeRating defines model for AppStoreAgeRating.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreagerating
type AppStoreAgeRating string

const (
	// AppStoreAgeRatingFourPlus is for an age rating of 4+.
	AppStoreAgeRatingFourPlus AppStoreAgeRating = "FOUR_PLUS"
	// AppStoreAgeRatingNinePlus is for an age rating of 9+.
	AppStoreAgeRatingNinePlus AppStoreAgeRating = "NINE_PLUS"
	// AppStoreAgeRatingSeventeenPlus is for an age rating of 17+.
	AppStoreAgeRatingSeventeenPlus AppStoreAgeRating = "SEVENTEEN_PLUS"
	// AppStoreAgeRatingTwelvePlus is for an age rating of 12+.
	AppStoreAgeRatingTwelvePlus AppStoreAgeRating = "TWELVE_PLUS"
)

// BrazilAgeRating defines model for BrazilAgeRating.
//
// https://developer.apple.com/documentation/appstoreconnectapi/brazilagerating
type BrazilAgeRating string

const (
	// BrazilAgeRatingEighteen is for an age rating of 18.
	BrazilAgeRatingEighteen BrazilAgeRating = "EIGHTEEN"
	// BrazilAgeRatingFourteen is for an age rating of 14.
	BrazilAgeRatingFourteen BrazilAgeRating = "FOURTEEN"
	// BrazilAgeRatingL is for an age rating of L, for general audiences.
	BrazilAgeRatingL BrazilAgeRating = "L"
	// BrazilAgeRatingSixteen is for an age rating of 16.
	BrazilAgeRatingSixteen BrazilAgeRating = "SIXTEEN"
	// BrazilAgeRatingTen is for an age rating of 10.
	BrazilAgeRatingTen BrazilAgeRating = "TEN"
	// BrazilAgeRatingTwelve is for an age rating of 12.
	BrazilAgeRatingTwelve BrazilAgeRating = "TWELVE"
)

// KidsAgeBand defines model for KidsAgeBand.
//
// https://developer.apple.com/documentation/appstoreconnectapi/kidsageband
type KidsAgeBand string

const (
	// KidsAgeBandFiveAndUnder is for an age rating of 5 and under.
	KidsAgeBandFiveAndUnder KidsAgeBand = "FIVE_AND_UNDER"
	// KidsAgeBandNineToEleven is for an age rating of 9 to 11.
	KidsAgeBandNineToEleven KidsAgeBand = "NINE_TO_ELEVEN"
	// KidsAgeBandSixToEight is for an age rating of 6 to 8.
	KidsAgeBandSixToEight KidsAgeBand = "SIX_TO_EIGHT"
)

// ageRatingDeclarationUpdateRequest defines model for AgeRatingDeclarationUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/ageratingdeclarationupdaterequest/data
type ageRatingDeclarationUpdateRequest struct {
	Attributes *AgeRatingDeclarationUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                       `json:"id"`
	Type       string                                       `json:"type"`
}

// AgeRatingDeclarationUpdateRequestAttributes are attributes for AgeRatingDeclarationUpdateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/ageratingdeclarationupdaterequest/data/attributes
type AgeRatingDeclarationUpdateRequestAttributes struct {
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

// AgeRatingDeclarationResponse defines model for AgeRatingDeclarationResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/ageratingdeclarationresponse
type AgeRatingDeclarationResponse struct {
	Data  AgeRatingDeclaration `json:"data"`
	Links DocumentLinks        `json:"links"`
}

// UpdateAgeRatingDeclaration provides age-related information so the App Store can determine the age rating for your app.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_an_age_rating_declaration
func (s *AppsService) UpdateAgeRatingDeclaration(ctx context.Context, id string, attributes *AgeRatingDeclarationUpdateRequestAttributes) (*AgeRatingDeclarationResponse, *Response, error) {
	req := ageRatingDeclarationUpdateRequest{
		Attributes: attributes,
		ID:         id,
		Type:       "ageRatingDeclarations",
	}
	url := fmt.Sprintf("ageRatingDeclarations/%s", id)
	res := new(AgeRatingDeclarationResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}
