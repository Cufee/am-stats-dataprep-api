package types

import "time"

type DocumentMeta struct {
	CreationTime time.Time `json:"creationTime" firestore:"creationTime"`
	LastUpdate   time.Time `json:"lastUpdate" firestore:"lastUpdate"`
	LastUsed     time.Time `json:"lastUsed" firestore:"lastUsed"`
}
