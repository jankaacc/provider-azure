/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	activedirectoryadministrator "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/activedirectoryadministrator"
	configuration "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/configuration"
	database "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/database"
	firewallrule "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/firewallrule"
	flexibleserver "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/flexibleserver"
	flexibleserverconfiguration "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/flexibleserverconfiguration"
	flexibleserverdatabase "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/flexibleserverdatabase"
	flexibleserverfirewallrule "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/flexibleserverfirewallrule"
	server "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/server"
	serverkey "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/serverkey"
	virtualnetworkrule "github.com/upbound/provider-azure/internal/controller/dbforpostgresql/virtualnetworkrule"
)

// Setup_dbforpostgresql creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup_dbforpostgresql(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		activedirectoryadministrator.Setup,
		configuration.Setup,
		database.Setup,
		firewallrule.Setup,
		flexibleserver.Setup,
		flexibleserverconfiguration.Setup,
		flexibleserverdatabase.Setup,
		flexibleserverfirewallrule.Setup,
		server.Setup,
		serverkey.Setup,
		virtualnetworkrule.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
