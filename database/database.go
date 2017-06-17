package database

import (
	"sync"

	"github.com/ocmdev/rita-blacklist2/list"
)

type (
	//Handle provides an interface for using a databae to hold
	//blacklist information
	Handle interface {
		//Init opens the connection to the backing database
		Init(connectionString string) error

		//GetRegisteredLists retrieves all of the lists registered with the database
		GetRegisteredLists() ([]list.Metadata, error)

		//RegisterList registers a new blacklist source with the database
		RegisterList(list.Metadata) error

		//RemoveList removes an existing blaclist source from the database
		RemoveList(list.Metadata) error

		//InsertEntries inserts entries from a list into the database
		InsertEntries(
			entryType list.BlacklistedEntryType,
			entries <-chan list.BlacklistedEntry,
			wg *sync.WaitGroup, errorsOut chan<- error,
		)

		//FindEntries finds entries of a given type and index
		FindEntries(dataType list.BlacklistedEntryType, index string) ([]DBEntry, error)
	}

	//DBEntry is the database safe version of BlacklistedEntry
	DBEntry struct {
		//Index is the main data held by this entry
		Index string
		//List is the source list
		List string
		//ExtraData contains extra information this blacklist source provides
		ExtraData map[string]interface{}
	}
)
