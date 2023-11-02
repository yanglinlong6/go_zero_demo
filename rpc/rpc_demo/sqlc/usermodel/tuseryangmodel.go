package usermodel

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TUserYangModel = (*customTUserYangModel)(nil)

type (
	// TUserYangModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTUserYangModel.
	TUserYangModel interface {
		tUserYangModel
	}

	customTUserYangModel struct {
		*defaultTUserYangModel
	}
)

// NewTUserYangModel returns a model for the database table.
func NewTUserYangModel(conn sqlx.SqlConn) TUserYangModel {
	return &customTUserYangModel{
		defaultTUserYangModel: newTUserYangModel(conn),
	}
}
