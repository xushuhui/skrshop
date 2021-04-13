// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"skrshop/internal/data/ent/accountplatform"
	"skrshop/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountPlatformQuery is the builder for querying AccountPlatform entities.
type AccountPlatformQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.AccountPlatform
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AccountPlatformQuery builder.
func (apq *AccountPlatformQuery) Where(ps ...predicate.AccountPlatform) *AccountPlatformQuery {
	apq.predicates = append(apq.predicates, ps...)
	return apq
}

// Limit adds a limit step to the query.
func (apq *AccountPlatformQuery) Limit(limit int) *AccountPlatformQuery {
	apq.limit = &limit
	return apq
}

// Offset adds an offset step to the query.
func (apq *AccountPlatformQuery) Offset(offset int) *AccountPlatformQuery {
	apq.offset = &offset
	return apq
}

// Order adds an order step to the query.
func (apq *AccountPlatformQuery) Order(o ...OrderFunc) *AccountPlatformQuery {
	apq.order = append(apq.order, o...)
	return apq
}

// First returns the first AccountPlatform entity from the query.
// Returns a *NotFoundError when no AccountPlatform was found.
func (apq *AccountPlatformQuery) First(ctx context.Context) (*AccountPlatform, error) {
	nodes, err := apq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{accountplatform.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (apq *AccountPlatformQuery) FirstX(ctx context.Context) *AccountPlatform {
	node, err := apq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AccountPlatform ID from the query.
// Returns a *NotFoundError when no AccountPlatform ID was found.
func (apq *AccountPlatformQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = apq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{accountplatform.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (apq *AccountPlatformQuery) FirstIDX(ctx context.Context) int {
	id, err := apq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AccountPlatform entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AccountPlatform entity is not found.
// Returns a *NotFoundError when no AccountPlatform entities are found.
func (apq *AccountPlatformQuery) Only(ctx context.Context) (*AccountPlatform, error) {
	nodes, err := apq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{accountplatform.Label}
	default:
		return nil, &NotSingularError{accountplatform.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (apq *AccountPlatformQuery) OnlyX(ctx context.Context) *AccountPlatform {
	node, err := apq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AccountPlatform ID in the query.
// Returns a *NotSingularError when exactly one AccountPlatform ID is not found.
// Returns a *NotFoundError when no entities are found.
func (apq *AccountPlatformQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = apq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{accountplatform.Label}
	default:
		err = &NotSingularError{accountplatform.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (apq *AccountPlatformQuery) OnlyIDX(ctx context.Context) int {
	id, err := apq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AccountPlatforms.
func (apq *AccountPlatformQuery) All(ctx context.Context) ([]*AccountPlatform, error) {
	if err := apq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return apq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (apq *AccountPlatformQuery) AllX(ctx context.Context) []*AccountPlatform {
	nodes, err := apq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AccountPlatform IDs.
func (apq *AccountPlatformQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := apq.Select(accountplatform.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (apq *AccountPlatformQuery) IDsX(ctx context.Context) []int {
	ids, err := apq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (apq *AccountPlatformQuery) Count(ctx context.Context) (int, error) {
	if err := apq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return apq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (apq *AccountPlatformQuery) CountX(ctx context.Context) int {
	count, err := apq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (apq *AccountPlatformQuery) Exist(ctx context.Context) (bool, error) {
	if err := apq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return apq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (apq *AccountPlatformQuery) ExistX(ctx context.Context) bool {
	exist, err := apq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AccountPlatformQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (apq *AccountPlatformQuery) Clone() *AccountPlatformQuery {
	if apq == nil {
		return nil
	}
	return &AccountPlatformQuery{
		config:     apq.config,
		limit:      apq.limit,
		offset:     apq.offset,
		order:      append([]OrderFunc{}, apq.order...),
		predicates: append([]predicate.AccountPlatform{}, apq.predicates...),
		// clone intermediate query.
		sql:  apq.sql.Clone(),
		path: apq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UID int64 `json:"uid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AccountPlatform.Query().
//		GroupBy(accountplatform.FieldUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (apq *AccountPlatformQuery) GroupBy(field string, fields ...string) *AccountPlatformGroupBy {
	group := &AccountPlatformGroupBy{config: apq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := apq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return apq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UID int64 `json:"uid,omitempty"`
//	}
//
//	client.AccountPlatform.Query().
//		Select(accountplatform.FieldUID).
//		Scan(ctx, &v)
//
func (apq *AccountPlatformQuery) Select(field string, fields ...string) *AccountPlatformSelect {
	apq.fields = append([]string{field}, fields...)
	return &AccountPlatformSelect{AccountPlatformQuery: apq}
}

func (apq *AccountPlatformQuery) prepareQuery(ctx context.Context) error {
	for _, f := range apq.fields {
		if !accountplatform.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if apq.path != nil {
		prev, err := apq.path(ctx)
		if err != nil {
			return err
		}
		apq.sql = prev
	}
	return nil
}

func (apq *AccountPlatformQuery) sqlAll(ctx context.Context) ([]*AccountPlatform, error) {
	var (
		nodes = []*AccountPlatform{}
		_spec = apq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AccountPlatform{config: apq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, apq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (apq *AccountPlatformQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := apq.querySpec()
	return sqlgraph.CountNodes(ctx, apq.driver, _spec)
}

func (apq *AccountPlatformQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := apq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (apq *AccountPlatformQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   accountplatform.Table,
			Columns: accountplatform.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: accountplatform.FieldID,
			},
		},
		From:   apq.sql,
		Unique: true,
	}
	if fields := apq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accountplatform.FieldID)
		for i := range fields {
			if fields[i] != accountplatform.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := apq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := apq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := apq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := apq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, accountplatform.ValidColumn)
			}
		}
	}
	return _spec
}

func (apq *AccountPlatformQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(apq.driver.Dialect())
	t1 := builder.Table(accountplatform.Table)
	selector := builder.Select(t1.Columns(accountplatform.Columns...)...).From(t1)
	if apq.sql != nil {
		selector = apq.sql
		selector.Select(selector.Columns(accountplatform.Columns...)...)
	}
	for _, p := range apq.predicates {
		p(selector)
	}
	for _, p := range apq.order {
		p(selector, accountplatform.ValidColumn)
	}
	if offset := apq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := apq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AccountPlatformGroupBy is the group-by builder for AccountPlatform entities.
type AccountPlatformGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (apgb *AccountPlatformGroupBy) Aggregate(fns ...AggregateFunc) *AccountPlatformGroupBy {
	apgb.fns = append(apgb.fns, fns...)
	return apgb
}

// Scan applies the group-by query and scans the result into the given value.
func (apgb *AccountPlatformGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := apgb.path(ctx)
	if err != nil {
		return err
	}
	apgb.sql = query
	return apgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (apgb *AccountPlatformGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := apgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (apgb *AccountPlatformGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(apgb.fields) > 1 {
		return nil, errors.New("ent: AccountPlatformGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := apgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (apgb *AccountPlatformGroupBy) StringsX(ctx context.Context) []string {
	v, err := apgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (apgb *AccountPlatformGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = apgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{accountplatform.Label}
	default:
		err = fmt.Errorf("ent: AccountPlatformGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (apgb *AccountPlatformGroupBy) StringX(ctx context.Context) string {
	v, err := apgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (apgb *AccountPlatformGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(apgb.fields) > 1 {
		return nil, errors.New("ent: AccountPlatformGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := apgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (apgb *AccountPlatformGroupBy) IntsX(ctx context.Context) []int {
	v, err := apgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (apgb *AccountPlatformGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = apgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{accountplatform.Label}
	default:
		err = fmt.Errorf("ent: AccountPlatformGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (apgb *AccountPlatformGroupBy) IntX(ctx context.Context) int {
	v, err := apgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (apgb *AccountPlatformGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(apgb.fields) > 1 {
		return nil, errors.New("ent: AccountPlatformGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := apgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (apgb *AccountPlatformGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := apgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (apgb *AccountPlatformGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = apgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{accountplatform.Label}
	default:
		err = fmt.Errorf("ent: AccountPlatformGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (apgb *AccountPlatformGroupBy) Float64X(ctx context.Context) float64 {
	v, err := apgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (apgb *AccountPlatformGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(apgb.fields) > 1 {
		return nil, errors.New("ent: AccountPlatformGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := apgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (apgb *AccountPlatformGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := apgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (apgb *AccountPlatformGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = apgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{accountplatform.Label}
	default:
		err = fmt.Errorf("ent: AccountPlatformGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (apgb *AccountPlatformGroupBy) BoolX(ctx context.Context) bool {
	v, err := apgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (apgb *AccountPlatformGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range apgb.fields {
		if !accountplatform.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := apgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := apgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (apgb *AccountPlatformGroupBy) sqlQuery() *sql.Selector {
	selector := apgb.sql
	columns := make([]string, 0, len(apgb.fields)+len(apgb.fns))
	columns = append(columns, apgb.fields...)
	for _, fn := range apgb.fns {
		columns = append(columns, fn(selector, accountplatform.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(apgb.fields...)
}

// AccountPlatformSelect is the builder for selecting fields of AccountPlatform entities.
type AccountPlatformSelect struct {
	*AccountPlatformQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (aps *AccountPlatformSelect) Scan(ctx context.Context, v interface{}) error {
	if err := aps.prepareQuery(ctx); err != nil {
		return err
	}
	aps.sql = aps.AccountPlatformQuery.sqlQuery(ctx)
	return aps.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aps *AccountPlatformSelect) ScanX(ctx context.Context, v interface{}) {
	if err := aps.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (aps *AccountPlatformSelect) Strings(ctx context.Context) ([]string, error) {
	if len(aps.fields) > 1 {
		return nil, errors.New("ent: AccountPlatformSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := aps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aps *AccountPlatformSelect) StringsX(ctx context.Context) []string {
	v, err := aps.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (aps *AccountPlatformSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aps.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{accountplatform.Label}
	default:
		err = fmt.Errorf("ent: AccountPlatformSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aps *AccountPlatformSelect) StringX(ctx context.Context) string {
	v, err := aps.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (aps *AccountPlatformSelect) Ints(ctx context.Context) ([]int, error) {
	if len(aps.fields) > 1 {
		return nil, errors.New("ent: AccountPlatformSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := aps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aps *AccountPlatformSelect) IntsX(ctx context.Context) []int {
	v, err := aps.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (aps *AccountPlatformSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aps.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{accountplatform.Label}
	default:
		err = fmt.Errorf("ent: AccountPlatformSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aps *AccountPlatformSelect) IntX(ctx context.Context) int {
	v, err := aps.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (aps *AccountPlatformSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(aps.fields) > 1 {
		return nil, errors.New("ent: AccountPlatformSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := aps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aps *AccountPlatformSelect) Float64sX(ctx context.Context) []float64 {
	v, err := aps.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (aps *AccountPlatformSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aps.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{accountplatform.Label}
	default:
		err = fmt.Errorf("ent: AccountPlatformSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aps *AccountPlatformSelect) Float64X(ctx context.Context) float64 {
	v, err := aps.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (aps *AccountPlatformSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(aps.fields) > 1 {
		return nil, errors.New("ent: AccountPlatformSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := aps.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aps *AccountPlatformSelect) BoolsX(ctx context.Context) []bool {
	v, err := aps.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (aps *AccountPlatformSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aps.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{accountplatform.Label}
	default:
		err = fmt.Errorf("ent: AccountPlatformSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aps *AccountPlatformSelect) BoolX(ctx context.Context) bool {
	v, err := aps.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aps *AccountPlatformSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := aps.sqlQuery().Query()
	if err := aps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (aps *AccountPlatformSelect) sqlQuery() sql.Querier {
	selector := aps.sql
	selector.Select(selector.Columns(aps.fields...)...)
	return selector
}
