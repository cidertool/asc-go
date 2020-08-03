package asc

import (
	"fmt"
	"net/http"
)

// AppStoreAgeRating defines model for AppStoreAgeRating.
type AppStoreAgeRating string

// List of AppStoreAgeRating
const (
	AppStoreAgeRatingFourPlus      AppStoreAgeRating = "FOUR_PLUS"
	AppStoreAgeRatingNinePlus      AppStoreAgeRating = "NINE_PLUS"
	AppStoreAgeRatingSeventeenPlus AppStoreAgeRating = "SEVENTEEN_PLUS"
	AppStoreAgeRatingTwelvePlus    AppStoreAgeRating = "TWELVE_PLUS"
)

// BrazilAgeRating defines model for BrazilAgeRating.
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
type KidsAgeBand string

// List of KidsAgeBand
const (
	KidsAgeBandFiveAndUnder KidsAgeBand = "FIVE_AND_UNDER"
	KidsAgeBandNineToEleven KidsAgeBand = "NINE_TO_ELEVEN"
	KidsAgeBandSixToEight   KidsAgeBand = "SIX_TO_EIGHT"
)

// AgeRatingDeclarationUpdateRequest defines model for AgeRatingDeclarationUpdateRequest.
type AgeRatingDeclarationUpdateRequest struct {
	Attributes *AgeRatingDeclarationUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                       `json:"id"`
	Type       string                                       `json:"type"`
}

// AgeRatingDeclarationUpdateRequestAttributes are attributes for AgeRatingDeclarationUpdateRequest
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
type AgeRatingDeclarationResponse struct {
	Data  AgeRatingDeclaration `json:"data"`
	Links DocumentLinks        `json:"links"`
}

// UpdateAgeRatingDeclaration provides age-related information so the App Store can determine the age rating for your app.
func (s *AppsService) UpdateAgeRatingDeclaration(id string, body *AgeRatingDeclarationUpdateRequest) (*AgeRatingDeclarationResponse, *http.Response, error) {
	url := fmt.Sprintf("ageRatingDeclarations/%s", id)
	res := new(AgeRatingDeclarationResponse)
	resp, err := s.client.patch(url, body, res)
	return res, resp, err
}
