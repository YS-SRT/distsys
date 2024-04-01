package utils

import (
	"distsys/base"

	"go.uber.org/zap"
)

func VerifyDBResult(err error, rows int64, actFailedMsg string) *base.CustomError {
	switch {
	case err != nil:
		zap.L().Error("DB operation error", zap.Error(err))
		return base.InternalCustomError(err)
	case rows <= 0:
		//zap.L().Info("DB operation affected zero rows")
		return base.BuildCustomInternalError(actFailedMsg)
	default:
		return nil
	}

}
