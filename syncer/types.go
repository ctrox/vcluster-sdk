package syncer

import (
	"github.com/loft-sh/vcluster-sdk/syncer/context"
	"github.com/loft-sh/vcluster-sdk/syncer/translator"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Object interface {
	Name() string
	Resource() client.Object
}

type Syncer interface {
	Object
	translator.NameTranslator

	SyncDown(ctx *context.SyncContext, vObj client.Object) (ctrl.Result, error)
	Sync(ctx *context.SyncContext, pObj client.Object, vObj client.Object) (ctrl.Result, error)
}

type UpSyncer interface {
	SyncUp(ctx *context.SyncContext, pObj client.Object) (ctrl.Result, error)
}

type FakeSyncer interface {
	Object

	FakeSyncUp(ctx *context.SyncContext, req types.NamespacedName) (ctrl.Result, error)
	FakeSync(ctx *context.SyncContext, vObj client.Object) (ctrl.Result, error)
}

type Starter interface {
	ReconcileStart(ctx *context.SyncContext, req ctrl.Request) (bool, error)
	ReconcileEnd()
}

// IndicesRegisterer registers additional indices for the controller
type IndicesRegisterer interface {
	RegisterIndices(ctx *context.RegisterContext) error
}

// ControllerModifier is used to modify the created controller for the syncer
type ControllerModifier interface {
	ModifyController(ctx *context.RegisterContext, builder *builder.Builder) (*builder.Builder, error)
}