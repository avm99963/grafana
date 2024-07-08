package sql

import (
	"context"
	"strconv"
	"testing"

	"github.com/grafana/grafana/pkg/infra/db"
	"github.com/grafana/grafana/pkg/services/featuremgmt"
	"github.com/grafana/grafana/pkg/setting"
	"github.com/grafana/grafana/pkg/storage/unified/resource"
	"github.com/grafana/grafana/pkg/storage/unified/sql/db/dbimpl"
	"github.com/grafana/grafana/pkg/tests/testsuite"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	testsuite.Run(m)
}

func TestBackendHappyPath(t *testing.T) {
	ctx := context.Background()
	dbstore := db.InitTestDB(t)

	rdb, err := dbimpl.ProvideResourceDB(dbstore, setting.NewCfg(), featuremgmt.WithFeatures(featuremgmt.FlagUnifiedStorage), nil)
	assert.NoError(t, err)
	store, err := NewBackendStore(backendOptions{
		DB: rdb,
	})

	assert.NoError(t, err)
	assert.NotNil(t, store)

	stream, err := store.WatchWriteEvents(ctx)
	assert.NoError(t, err)

	t.Run("Add 3 resources", func(t *testing.T) {
		for i := 1; i <= 3; i++ {
			rv, err := store.WriteEvent(ctx, resource.WriteEvent{
				Type:  resource.WatchEvent_ADDED,
				Value: []byte("initial value " + strconv.Itoa(i)),
				Key: &resource.ResourceKey{
					Namespace: "namespace",
					Group:     "group",
					Resource:  "resource",
					Name:      "item" + strconv.Itoa(i),
				},
			})
			assert.NoError(t, err)
			assert.Equal(t, int64(i), rv)
		}
	})

	t.Run("Update item2", func(t *testing.T) {
		rv, err := store.WriteEvent(ctx, resource.WriteEvent{
			Type:  resource.WatchEvent_MODIFIED,
			Value: []byte("updated value"),
			Key: &resource.ResourceKey{
				Namespace: "namespace",
				Group:     "group",
				Resource:  "resource",
				Name:      "item2",
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, int64(4), rv)
	})

	t.Run("Delete item1", func(t *testing.T) {
		rv, err := store.WriteEvent(ctx, resource.WriteEvent{
			Type: resource.WatchEvent_DELETED,
			Key: &resource.ResourceKey{
				Namespace: "namespace",
				Group:     "group",
				Resource:  "resource",
				Name:      "item1",
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, int64(5), rv)
	})

	t.Run("Read latest item 2", func(t *testing.T) {
		resp, err := store.Read(ctx, &resource.ReadRequest{
			Key: &resource.ResourceKey{
				Namespace: "namespace",
				Group:     "group",
				Resource:  "resource",
				Name:      "item2",
			},
		})
		assert.NoError(t, err)
		assert.Equal(t, int64(4), resp.ResourceVersion)
		assert.Equal(t, "updated value", string(resp.Value))
	})

	t.Run("Read early verion of item2", func(t *testing.T) {
		resp, err := store.Read(ctx, &resource.ReadRequest{
			Key: &resource.ResourceKey{
				Namespace: "namespace",
				Group:     "group",
				Resource:  "resource",
				Name:      "item2",
			},
			ResourceVersion: 3, // item2 was created at rv=2 and updated at rv=4
		})
		assert.NoError(t, err)
		assert.Equal(t, int64(2), resp.ResourceVersion)
		assert.Equal(t, "initial value 2", string(resp.Value))
	})

	t.Run("PrepareList latest", func(t *testing.T) {
		resp, err := store.PrepareList(ctx, &resource.ListRequest{})
		assert.NoError(t, err)
		assert.Len(t, resp.Items, 2)
		assert.Equal(t, "updated value", string(resp.Items[0].Value))
		assert.Equal(t, "initial value 3", string(resp.Items[1].Value))
	})

	t.Run("Watch events", func(t *testing.T) {
		event := <-stream
		assert.Equal(t, "item1", event.Key.Name)
		assert.Equal(t, int64(1), event.ResourceVersion)
		assert.Equal(t, resource.WatchEvent_ADDED, event.Type)
		event = <-stream
		assert.Equal(t, "item2", event.Key.Name)
		assert.Equal(t, int64(2), event.ResourceVersion)
		assert.Equal(t, resource.WatchEvent_ADDED, event.Type)

		event = <-stream
		assert.Equal(t, "item3", event.Key.Name)
		assert.Equal(t, int64(3), event.ResourceVersion)
		assert.Equal(t, resource.WatchEvent_ADDED, event.Type)

		event = <-stream
		assert.Equal(t, "item2", event.Key.Name)
		assert.Equal(t, int64(4), event.ResourceVersion)
		assert.Equal(t, resource.WatchEvent_MODIFIED, event.Type)

		event = <-stream
		assert.Equal(t, "item1", event.Key.Name)
		assert.Equal(t, int64(5), event.ResourceVersion)
		assert.Equal(t, resource.WatchEvent_DELETED, event.Type)
	})
}

func TestBackendWatchWriteEventsFromHead(t *testing.T) {
	ctx := context.Background()
	dbstore := db.InitTestDB(t)

	rdb, err := dbimpl.ProvideResourceDB(dbstore, setting.NewCfg(), featuremgmt.WithFeatures(featuremgmt.FlagUnifiedStorage), nil)
	assert.NoError(t, err)
	store, err := NewBackendStore(backendOptions{
		DB: rdb,
	})

	assert.NoError(t, err)
	assert.NotNil(t, store)

	// Create a few resources before initing the watch
	_, err = store.WriteEvent(ctx, resource.WriteEvent{
		Type:  resource.WatchEvent_ADDED,
		Value: []byte("initial value 0"),
		Key: &resource.ResourceKey{
			Namespace: "namespace",
			Group:     "group",
			Resource:  "resource",
			Name:      "item 0",
		},
	})
	assert.NoError(t, err)

	// Start the watch
	stream, err := store.WatchWriteEvents(ctx)
	assert.NoError(t, err)

	// Create one more event
	_, err = store.WriteEvent(ctx, resource.WriteEvent{
		Type:  resource.WatchEvent_ADDED,
		Value: []byte("initial value 2"),
		Key: &resource.ResourceKey{
			Namespace: "namespace",
			Group:     "group",
			Resource:  "resource",
			Name:      "item2",
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, "item2", (<-stream).Key.Name)
}
