// Copyright 2023 The Forgejo Authors. All rights reserved.
// SPDX-License-Identifier: MIT
package activitypub

import (
	"context"
	"errors"
	"strconv"
	"strings"

	repo_model "code.gitea.io/gitea/models/repo"
	"code.gitea.io/gitea/modules/setting"

	ap "github.com/go-ap/activitypub"
)

// Returns the owner and name corresponding to a Repository actor IRI
func RepositoryIRIToName(repoIRI ap.IRI) (string, string, error) {
	repoIRISplit := strings.Split(repoIRI.String(), "/")
	if len(repoIRISplit) < 5 {
		return "", "", errors.New("not a Repository actor IRI")
	}

	instance := repoIRISplit[2]
	username := repoIRISplit[len(repoIRISplit)-2]
	reponame := repoIRISplit[len(repoIRISplit)-1]
	if instance == setting.Domain {
		// Local repo
		return username, reponame, nil
	}
	// Remote repo
	return username + "@" + instance, reponame, nil
}

// Returns the repository corresponding to a Repository actor IRI
func RepositoryIRIToRepository(ctx context.Context, repoIRI ap.IRI) (*repo_model.Repository, error) {
	username, reponame, err := RepositoryIRIToName(repoIRI)
	if err != nil {
		return nil, err
	}

	return repo_model.GetRepositoryByOwnerAndName(ctx, username, reponame)
}

// Returns the owner, repo name, and idx of a Ticket object IRI
func TicketIRIToName(ticketIRI ap.IRI) (string, string, int64, error) {
	ticketIRISplit := strings.Split(ticketIRI.String(), "/")
	if len(ticketIRISplit) < 5 {
		return "", "", 0, errors.New("not a Ticket object IRI")
	}

	instance := ticketIRISplit[2]
	username := ticketIRISplit[len(ticketIRISplit)-3]
	reponame := ticketIRISplit[len(ticketIRISplit)-2]
	idx, err := strconv.ParseInt(ticketIRISplit[len(ticketIRISplit)-1], 10, 64)
	if err != nil {
		return "", "", 0, err
	}
	if instance == setting.Domain {
		// Local repo
		return username, reponame, idx, nil
	}
	// Remote repo
	return username + "@" + instance, reponame, idx, nil
}

// Returns the owner, repo name, and idx of a Branch object IRI
func BranchIRIToName(ticketIRI ap.IRI) (string, string, string, error) {
	ticketIRISplit := strings.Split(ticketIRI.String(), "/")
	if len(ticketIRISplit) < 5 {
		return "", "", "", errors.New("not a Branch object IRI")
	}

	instance := ticketIRISplit[2]
	username := ticketIRISplit[len(ticketIRISplit)-3]
	reponame := ticketIRISplit[len(ticketIRISplit)-2]
	branch := ticketIRISplit[len(ticketIRISplit)-1]
	if instance == setting.Domain {
		// Local repo
		return username, reponame, branch, nil
	}
	// Remote repo
	return username + "@" + instance, reponame, branch, nil
}
