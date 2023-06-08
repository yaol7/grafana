package admission

import (
	"context"
	"encoding/json"
	"io"

	"github.com/grafana/thema"
	"github.com/grafana/thema/vmux"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apiserver/pkg/admission"

	"github.com/grafana/grafana/pkg/cuectx"
	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/kinds/dashboard"
)

const PluginNameSchemaValidate = "SchemaValidate"

// Register registers a plugin
func RegisterSchemaValidate(plugins *admission.Plugins) {
	plugins.Register(PluginNameSchemaValidate, func(config io.Reader) (admission.Interface, error) {
		return NewSchemaValidate(), nil
	})
}

type schemaValidate struct {
	log log.Logger
}

var _ admission.ValidationInterface = schemaValidate{}

// Validate makes an admission decision based on the request attributes.  It is NOT allowed to mutate.
func (sv schemaValidate) Validate(ctx context.Context, a admission.Attributes, o admission.ObjectInterfaces) (err error) {
	obj := a.GetObject()
	uobj, ok := obj.(*unstructured.Unstructured)
	if !ok {
		// not sure if this is ever expected to happen
		sv.log.Info("failed to cast object to unstructured")
		return nil
	}
	spec, err := json.Marshal(uobj.Object["spec"])
	if err != nil {
		sv.log.Info("failed to marshal spec", "err", err)
		return nil
	}

	// this needs to be generic, but for now, just do it for dashboards
	if obj.GetObjectKind().GroupVersionKind().Kind == "Dashboard" {
		dk, err := dashboard.NewKind(cuectx.GrafanaThemaRuntime())
		if err != nil {
			sv.log.Info("failed to create dashboard kind", "err", err)
			return nil
		}

		// TODO thema (or codegen, or both) request: rename JSONValueMux to
		// something that's a bit clearer; it's unmarshalling the JSON bytes to
		// a dashboard and validating that against any schema
		_, _, err = dk.JSONValueMux(spec)
		if err != nil {
			sv.log.Error("failed to validate dashboard", "err", err)
			return nil
		}

		// JSONValueMux validates that the dashboard matches *any* schema, so we
		// may need to translate to latest.
		//
		// TODO: None of this feels correct; there should be a dashboard.Kind
		// "validate latest" function, or something like that.
		sch, err := dk.Lineage().Schema(thema.LatestVersion(dk.Lineage()))
		if err != nil {
			sv.log.Error("failed to get latest schema", "err", err)
			return nil
		}
		cueVal, _ := vmux.NewJSONCodec("dashboard.json").Decode(cuectx.GrafanaCUEContext(), spec)
		_, err = sch.Validate(cueVal)
		if err != nil {
			sv.log.Info("failed to validate dashboard", "err", err)
		}
	}
	return nil
}

// Handles returns true if this admission controller can handle the given operation
// where operation can be one of CREATE, UPDATE, DELETE, or CONNECT
func (schemaValidate) Handles(operation admission.Operation) bool {
	switch operation {
	case admission.Connect:
		return false
	}
	return true
}

// NewSchemaValidate creates a NewSchemaValidate admission handler
func NewSchemaValidate() admission.Interface {
	return schemaValidate{
		log: log.New("admission.schema-validate"),
	}
}