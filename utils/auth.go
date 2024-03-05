package utils

import (
	"context"
	"encoding/base64"
	"os"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var adminUsername = os.Getenv("ADMIN_USERNAME")
var adminPassword = os.Getenv("ADMIN_PASSWORD")

func isAdmin(ctx context.Context) error {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "missing credentials")
	}
	authHeaders, ok := md["authorization"]
	if !ok || len(authHeaders) == 0 {
		return status.Error(codes.Unauthenticated, "missing authorization header")
	}
	authHeaderParts := strings.Fields(authHeaders[0])
	if len(authHeaderParts) != 2 || authHeaderParts[0] != "Basic" {
		return status.Error(codes.Unauthenticated, "invalid Authentication header")
	}
	decodedCredentials, err := base64.StdEncoding.DecodeString(authHeaderParts[1])
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "failed to decode credentials: %v", err)
	}
	credentials := strings.Split(string(decodedCredentials), ":")
	if len(credentials) != 2 {
		return status.Errorf(codes.Unauthenticated, "invalid credentials format")
	}
	username := credentials[0]
	password := credentials[1]

	if username != adminUsername && password != adminPassword {
		return status.Errorf(codes.Unauthenticated, "Invalid Admin Credentials")
	}

	return nil
}
