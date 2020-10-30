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

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/cidertool/asc-go/asc"
	"github.com/cidertool/asc-go/examples/util"
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
	if err != nil {
		fmt.Println(err)
	}
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
