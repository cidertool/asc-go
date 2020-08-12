package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/aaronsky/asc-go/asc"
	"github.com/aaronsky/asc-go/examples/util"
)

var (
	givenName           = flag.String("givenname", "", "Given (first) name")
	familyName          = flag.String("familyname", "", "Family (last) name")
	email               = flag.String("email", "", "Email address")
	shouldCancel        = flag.Bool("cancel", false, "Instead of inviting the user, cancel the invitation instead")
	roles               = flag.String("roles", "", "Roles that the user should be set to, comma-separated")
	allAppsVisible      = flag.Bool("allappsvisible", false, "Whether the user should have access to all apps in the team.")
	provisioningAllowed = flag.Bool("provisioningallowed", false, "Whether the user should be allowed to update provisioning settings.")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	auth, err := util.TokenConfig()
	if err != nil {
		log.Fatalf("client config failed: %s", err)
	}

	// Create the App Store Connect client
	client := asc.NewClient(auth.Client())

	if *shouldCancel {
		err = cancelUserInvitation(ctx, client)
	} else {
		err = inviteUser(ctx, client)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func inviteUser(ctx context.Context, client *asc.Client) error {
	userRoles := make([]asc.UserRole, 0)
	for _, role := range strings.Split(*roles, ",") {
		role = strings.Trim(role, " ")
		userRoles = append(userRoles, asc.UserRole(role))
	}

	invitation, _, err := client.Users.CreateInvitation(ctx, asc.UserInvitationCreateRequestAttributes{
		FirstName:           *givenName,
		LastName:            *familyName,
		Email:               asc.Email(*email),
		Roles:               userRoles,
		AllAppsVisible:      allAppsVisible,
		ProvisioningAllowed: provisioningAllowed,
	}, nil)
	if err != nil {
		return err
	}
	fmt.Printf(
		"Sent an invitation to %s %s at %s. They should check their email and confirm the invitation before it expires at %s.",
		*invitation.Data.Attributes.FirstName,
		*invitation.Data.Attributes.LastName,
		*invitation.Data.Attributes.Email,
		invitation.Data.Attributes.ExpirationDate.String(),
	)
	return nil
}

func cancelUserInvitation(ctx context.Context, client *asc.Client) error {
	invitations, _, err := client.Users.ListInvitations(ctx, &asc.ListInvitationsQuery{
		FilterEmail: []string{*email},
	})
	var invitation asc.UserInvitation
	if len(invitations.Data) > 0 {
		invitation = invitations.Data[0]
	}
	_, err = client.Users.CancelInvitation(ctx, invitation.ID)
	if err != nil {
		return err
	}
	fmt.Printf(
		"Cancelled the invitation issued for %s %s at %s.",
		*invitation.Attributes.FirstName,
		*invitation.Attributes.LastName,
		*invitation.Attributes.Email,
	)
	return nil
}
