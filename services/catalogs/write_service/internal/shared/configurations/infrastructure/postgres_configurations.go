package infrastructure

import (
	"emperror.dev/errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zapadapter"
	postgres "github.com/mehdihadeli/store-golang-microservice-sample/pkg/postgres_pgx"
	"go.uber.org/zap"
)

func (ic *infrastructureConfigurator) configPostgres() (*postgres.Pgx, error, func()) {
	pgxConn, err := postgres.NewPgxPoolConn(ic.cfg.Postgresql, zapadapter.NewLogger(zap.L()), pgx.LogLevelInfo)
	if err != nil {
		return nil, errors.WrapIf(err, "postgresql.NewPgxConn"), nil
	}

	ic.log.Infof("postgres connected: %v", pgxConn.ConnPool.Stat().TotalConns())

	return pgxConn, nil, func() {
		pgxConn.Close()
	}
}
