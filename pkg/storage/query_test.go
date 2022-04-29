package storage

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"
)

var _ Document = TestDocument{}
var _ json.Marshaler = TestDocument{}

type TestDocument struct {
	SomeName string
}

func (t TestDocument) MarshalJSON() ([]byte, error) {
	raw := map[string]interface{}{
		"_id":  t.SomeName + "id",
		"name": t.SomeName,
	}
	return json.Marshal(raw)
}

func (t TestDocument) DefaultDocumentFilter() map[string]interface{} {
	return map[string]interface{}{"name": t.SomeName}
}

func TestInsertOptions_ToPluginOptions(t *testing.T) {
	// Test that when we insert documents that we first go through
	// its json representation, so that the fields have the right
	// names based on the json tag, and that any custom json marshaling
	// is used.

	docA := TestDocument{SomeName: "a"}
	docB := TestDocument{SomeName: "b"}
	opts := InsertOptions{Documents: []interface{}{docA, docB}}
	gotOpts, err := opts.ToPluginOptions("mydocs")
	require.NoError(t, err)

	wantRawDocs := []bson.M{
		{"_id": "aid", "name": "a"},
		{"_id": "bid", "name": "b"},
	}

	require.Equal(t, wantRawDocs, gotOpts.Documents)
}

func TestUpdateOptions_ToPluginOptions(t *testing.T) {
	// Test that when we update documents that we first go through
	// its json representation, so that the fields have the right
	// names based on the json tag, and that any custom json marshaling
	// is used.

	doc := TestDocument{SomeName: "a"}
	opts := UpdateOptions{Document: doc}
	gotOpts, err := opts.ToPluginOptions("mydocs")
	require.NoError(t, err)

	wantRawDoc := bson.M{"_id": "aid", "name": "a"}
	require.Equal(t, wantRawDoc, gotOpts.Document)
}

func TestFindOptions_ToPluginOptions(t *testing.T) {
	so := FindOptions{
		Sort: []string{"-_id", "name"},
	}
	po := so.ToPluginOptions("mythings")
	wantSortDoc := bson.D{
		{"_id", -1},
		{"name", 1}}
	require.Equal(t, wantSortDoc, po.Sort)
}
