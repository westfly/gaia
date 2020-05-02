// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package initialize

import (
	"github.com/westfly/gaia/internal/app/api"
	"github.com/westfly/gaia/internal/app/api/mock"
	"github.com/westfly/gaia/internal/app/bll/impl/bll"
	"github.com/westfly/gaia/internal/app/initialize/data"
	"github.com/westfly/gaia/internal/app/module/adapter"
	"github.com/westfly/gaia/internal/app/router"
	"github.com/google/wire"

	// mongoModel "github.com/westfly/gaia/internal/app/model/impl/mongo/model"
	gormModel "github.com/westfly/gaia/internal/app/model/impl/gorm/model"
)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	// 默认使用gorm存储注入，这里可使用 InitMongoDB & mongoModel.ModelSet 替换为 gorm 存储
	wire.Build(
		InitGormDB,
		gormModel.ModelSet,
		// InitMongoDB,
		// mongoModel.ModelSet,
		InitAuth,
		InitCasbin,
		InitGinEngine,
		bll.BllSet,
		api.APISet,
		mock.MockSet,
		router.RouterSet,
		adapter.CasbinAdapterSet,
		data.MenuSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
