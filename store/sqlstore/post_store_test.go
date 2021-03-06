// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package sqlstore

import (
	"testing"

	"github.com/mattermost/mattermost-server/v5/store/searchtest"
	"github.com/mattermost/mattermost-server/v5/store/storetest"
)

func TestPostStore(t *testing.T) {
	StoreTestWithSqlSupplier(t, storetest.TestPostStore)
}

func TestSearchPostStore(t *testing.T) {
	StoreTestWithSearchTestEngine(t, searchtest.TestSearchPostStore)
}
