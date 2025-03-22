package main

import "github.com/google/uuid"

func v5UUID(data string) string {
	return uuid.NewSHA1(uuid.NameSpaceURL, []byte(data)).String()
}
