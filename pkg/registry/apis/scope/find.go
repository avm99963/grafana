package scope

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/grafana/grafana/pkg/infra/log"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"

	scope "github.com/grafana/grafana/pkg/apis/scope/v0alpha1"
)

type findREST struct {
	scopeNodeStorage *storage
	log              log.Logger
}

var (
	_ rest.Storage              = (*findREST)(nil)
	_ rest.SingularNameProvider = (*findREST)(nil)
	_ rest.Connecter            = (*findREST)(nil)
	_ rest.Scoper               = (*findREST)(nil)
	_ rest.StorageMetadata      = (*findREST)(nil)
)

func (r *findREST) New() runtime.Object {
	// This is added as the "ResponseType" regarless what ProducesObject() says :)
	return &scope.FindScopeNodeChildrenResults{}
}

func (r *findREST) Destroy() {}

func (r *findREST) NamespaceScoped() bool {
	return true
}

func (r *findREST) GetSingularName() string {
	return "FindScopeNodeChildrenResults" // Used for the
}

func (r *findREST) ProducesMIMETypes(verb string) []string {
	return []string{"application/json"} // and parquet!
}

func (r *findREST) ProducesObject(verb string) interface{} {
	return &scope.FindScopeNodeChildrenResults{}
}

func (r *findREST) ConnectMethods() []string {
	return []string{"GET"}
}

func (r *findREST) NewConnectOptions() (runtime.Object, bool, string) {
	return nil, false, "" // true means you can use the trailing path as a variable
}

func (r *findREST) Connect(ctx context.Context, name string, opts runtime.Object, responder rest.Responder) (http.Handler, error) {
	// See: /pkg/apiserver/builder/helper.go#L34
	// The name is set with a rewriter hack
	if name != "name" {
		return nil, errors.NewNotFound(schema.GroupResource{}, name)
	}

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		parent := req.URL.Query().Get("parent")
		query := req.URL.Query().Get("query")
		results := &scope.FindScopeNodeChildrenResults{}

		r.log.Info("finding scope node children", "namespace", request.NamespaceValue(ctx), "parent", parent, "query", query)

		raw, err := r.scopeNodeStorage.List(ctx, &internalversion.ListOptions{
			Limit: 1000,
		})

		if err != nil {
			responder.Error(err)
			return
		}

		all, ok := raw.(*scope.ScopeNodeList)
		if !ok {
			responder.Error(fmt.Errorf("expected ScopeNodeList"))
			return
		}

		r.log.Info("scope node query result", "length", len(all.Items))

		for _, item := range all.Items {
			filterAndAppendItem(item, parent, query, results)
		}

		responder.Object(200, results)
	}), nil
}

func filterAndAppendItem(item scope.ScopeNode, parent string, query string, results *scope.FindScopeNodeChildrenResults) {
	if parent != item.Spec.ParentName {
		return // Someday this will have an index in raw storage on parentName
	}

	// skip if query is passed and title doesn't match.
	// HasPrefix is not the end goal but something that that gets us started.
	if query != "" && !strings.HasPrefix(item.Spec.Title, query) {
		return
	}

	results.Items = append(results.Items, item)
}
