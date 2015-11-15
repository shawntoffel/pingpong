package httpUtils

import (
	"github.com/pingpong/models"
	"net"
)

func GetRemoteIp(remoteAddress string) (string, *models.Error) {

	// Extract the ip from the request Remote Address
	extractedIp, _, splitError := net.SplitHostPort(remoteAddress)

	// Build an error if we cant extract the ip
	if splitError != nil {

		extractError := models.Error{
			Error:  "Could not determine client IP address",
			Detail: splitError.Error(),
		}

		return extractedIp, &extractError
	}

	return extractedIp, nil
}
