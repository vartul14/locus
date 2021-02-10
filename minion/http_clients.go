package minion

import (
	"sync"
	"vartul14/locus/minion/google_client"
)

var (
	//Once var declaration
	googleClientSvcOnce sync.Once

	//Global client initialise
	googleClientSvcClient *google_client.GoogleClient
)

func GetGoogleClient() google_client.Client {
	googleClientSvcOnce.Do(func() {
		googleClientSvcClient = &google_client.GoogleClient{
			HostAndPort: "https://maps.googleapis.com",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
	})

	return googleClientSvcClient
}