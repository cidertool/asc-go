package asc

import (
	"context"
	"fmt"
)

// AppStoreAgeRating defines model for AppStoreAgeRating.
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreagerating
type AppStoreAgeRating string

// List of AppStoreAgeRating
const (
	AppStoreAgeRatingFourPlus      AppStoreAgeRating = "FOUR_PLUS"
	AppStoreAgeRatingNinePlus      AppStoreAgeRating = "NINE_PLUS"
	AppStoreAgeRatingSeventeenPlus AppStoreAgeRating = "SEVENTEEN_PLUS"
	AppStoreAgeRatingTwelvePlus    AppStoreAgeRating = "TWELVE_PLUS"
)

// BrazilAgeRating defines model for BrazilAgeRating.
//
// https://developer.apple.com/documentation/appstoreconnectapi/brazilagerating
type BrazilAgeRating string

// List of BrazilAgeRating
const (
	BrazilAgeRatingEighteen BrazilAgeRating = "EIGHTEEN"
	BrazilAgeRatingFourteen BrazilAgeRating = "FOURTEEN"
	BrazilAgeRatingL        BrazilAgeRating = "L"
	BrazilAgeRatingSixteen  BrazilAgeRating = "SIXTEEN"
	BrazilAgeRatingTen      BrazilAgeRating = "TEN"
	BrazilAgeRatingTwelve   BrazilAgeRating = "TWELVE"
)

// KidsAgeBand defines model for KidsAgeBand.
//
// https://developer.apple.com/documentation/appstoreconnectapi/kidsageband
type KidsAgeBand string

// List of KidsAgeBand
const (
	KidsAgeBandFiveAndUnder KidsAgeBand = "FIVE_AND_UNDER"
	KidsAgeBandNineToEleven KidsAgeBand = "NINE_TO_ELEVEN"
	KidsAgeBandSixToEight   KidsAgeBand = "SIX_TO_EIGHT"
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
	GamblingAndContests                         *bool        `json:"gamblingAndContests,omitempty"`
	GamblingSimulated                           *string      `json:"gamblingSimulated,omitempty"`
	HorrorOrFearThemes                          *string      `json:"horrorOrFearThemes,omitempty"`
	KidsAgeBand                                 *KidsAgeBand `json:"kidsAgeBand,omitempty"`
	MatureOrSuggestiveThemes                    *string      `json:"matureOrSuggestiveThemes,omitempty"`
	MedicalOrTreatmentInformation               *string      `json:"medicalOrTreatmentInformation,omitempty"`
	ProfanityOrCrudeHumor                       *string      `json:"profanityOrCrudeHumor,omitempty"`
	SexualContentGraphicAndNudity               *string      `json:"sexualContentGraphicAndNudity,omitempty"`
	SexualContentOrNudity                       *string      `json:"sexualContentOrNudity,omitempty"`
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
