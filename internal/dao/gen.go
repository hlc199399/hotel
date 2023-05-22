// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q             = new(Query)
	Order         *order
	Room          *room
	RoomCondition *roomCondition
	User          *user
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Order = &Q.Order
	Room = &Q.Room
	RoomCondition = &Q.RoomCondition
	User = &Q.User
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:            db,
		Order:         newOrder(db, opts...),
		Room:          newRoom(db, opts...),
		RoomCondition: newRoomCondition(db, opts...),
		User:          newUser(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Order         order
	Room          room
	RoomCondition roomCondition
	User          user
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:            db,
		Order:         q.Order.clone(db),
		Room:          q.Room.clone(db),
		RoomCondition: q.RoomCondition.clone(db),
		User:          q.User.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:            db,
		Order:         q.Order.replaceDB(db),
		Room:          q.Room.replaceDB(db),
		RoomCondition: q.RoomCondition.replaceDB(db),
		User:          q.User.replaceDB(db),
	}
}

type queryCtx struct {
	Order         IOrderDo
	Room          IRoomDo
	RoomCondition IRoomConditionDo
	User          IUserDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Order:         q.Order.WithContext(ctx),
		Room:          q.Room.WithContext(ctx),
		RoomCondition: q.RoomCondition.WithContext(ctx),
		User:          q.User.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
