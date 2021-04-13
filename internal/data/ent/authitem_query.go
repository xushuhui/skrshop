// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"skrshop/internal/data/ent/authitem"
	"skrshop/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AuthItemQuery is the builder for querying AuthItem entities.
type AuthItemQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.AuthItem
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthItemQuery builder.
func (aiq *AuthItemQuery) Where(ps ...predicate.AuthItem) *AuthItemQuery {
	aiq.predicates = append(aiq.predicates, ps...)
	return aiq
}

// Limit adds a limit step to the query.
func (aiq *AuthItemQuery) Limit(limit int) *AuthItemQuery {
	aiq.limit = &limit
	return aiq
}

// Offset adds an offset step to the query.
func (aiq *AuthItemQuery) Offset(offset int) *AuthItemQuery {
	aiq.offset = &offset
	return aiq
}

// Order adds an order step to the query.
func (aiq *AuthItemQuery) Order(o ...OrderFunc) *AuthItemQuery {
	aiq.order = append(aiq.order, o...)
	return aiq
}

// First returns the first AuthItem entity from the query.
// Returns a *NotFoundError when no AuthItem was found.
func (aiq *AuthItemQuery) First(ctx context.Context) (*AuthItem, error) {
	nodes, err := aiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{authitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aiq *AuthItemQuery) FirstX(ctx context.Context) *AuthItem {
	node, err := aiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AuthItem ID from the query.
// Returns a *NotFoundError when no AuthItem ID was found.
func (aiq *AuthItemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{authitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aiq *AuthItemQuery) FirstIDX(ctx context.Context) int {
	id, err := aiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AuthItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AuthItem entity is not found.
// Returns a *NotFoundError when no AuthItem entities are found.
func (aiq *AuthItemQuery) Only(ctx context.Context) (*AuthItem, error) {
	nodes, err := aiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{authitem.Label}
	default:
		return nil, &NotSingularError{authitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aiq *AuthItemQuery) OnlyX(ctx context.Context) *AuthItem {
	node, err := aiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AuthItem ID in the query.
// Returns a *NotSingularError when exactly one AuthItem ID is not found.
// Returns a *NotFoundError when no entities are found.
func (aiq *AuthItemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{authitem.Label}
	default:
		err = &NotSingularError{authitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aiq *AuthItemQuery) OnlyIDX(ctx context.Context) int {
	id, err := aiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AuthItems.
func (aiq *AuthItemQuery) All(ctx context.Context) ([]*AuthItem, error) {
	if err := aiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aiq *AuthItemQuery) AllX(ctx context.Context) []*AuthItem {
	nodes, err := aiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AuthItem IDs.
func (aiq *AuthItemQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := aiq.Select(authitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aiq *AuthItemQuery) IDsX(ctx context.Context) []int {
	ids, err := aiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aiq *AuthItemQuery) Count(ctx context.Context) (int, error) {
	if err := aiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aiq *AuthItemQuery) CountX(ctx context.Context) int {
	count, err := aiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aiq *AuthItemQuery) Exist(ctx context.Context) (bool, error) {
	if err := aiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aiq *AuthItemQuery) ExistX(ctx context.Context) bool {
	exist, err := aiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aiq *AuthItemQuery) Clone() *AuthItemQuery {
	if aiq == nil {
		return nil
	}
	return &AuthItemQuery{
		config:     aiq.config,
		limit:      aiq.limit,
		offset:     aiq.offset,
		order:      append([]OrderFunc{}, aiq.order...),
		predicates: append([]predicate.AuthItem{}, aiq.predicates...),
		// clone intermediate query.
		sql:  aiq.sql.Clone(),
		path: aiq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Status int8 `json:"status,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AuthItem.Query().
//		GroupBy(authitem.FieldStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aiq *AuthItemQuery) GroupBy(field string, fields ...string) *AuthItemGroupBy {
	group := &AuthItemGroupBy{config: aiq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aiq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Status int8 `json:"status,omitempty"`
//	}
//
//	client.AuthItem.Query().
//		Select(authitem.FieldStatus).
//		Scan(ctx, &v)
//
func (aiq *AuthItemQuery) Select(field string, fields ...string) *AuthItemSelect {
	aiq.fields = append([]string{field}, fields...)
	return &AuthItemSelect{AuthItemQuery: aiq}
}

func (aiq *AuthItemQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aiq.fields {
		if !authitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aiq.path != nil {
		prev, err := aiq.path(ctx)
		if err != nil {
			return err
		}
		aiq.sql = prev
	}
	return nil
}

func (aiq *AuthItemQuery) sqlAll(ctx context.Context) ([]*AuthItem, error) {
	var (
		nodes = []*AuthItem{}
		_spec = aiq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AuthItem{config: aiq.config}
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
	if err := sqlgraph.QueryNodes(ctx, aiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (aiq *AuthItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aiq.querySpec()
	return sqlgraph.CountNodes(ctx, aiq.driver, _spec)
}

func (aiq *AuthItemQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (aiq *AuthItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authitem.Table,
			Columns: authitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: authitem.FieldID,
			},
		},
		From:   aiq.sql,
		Unique: true,
	}
	if fields := aiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authitem.FieldID)
		for i := range fields {
			if fields[i] != authitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, authitem.ValidColumn)
			}
		}
	}
	return _spec
}

func (aiq *AuthItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aiq.driver.Dialect())
	t1 := builder.Table(authitem.Table)
	selector := builder.Select(t1.Columns(authitem.Columns...)...).From(t1)
	if aiq.sql != nil {
		selector = aiq.sql
		selector.Select(selector.Columns(authitem.Columns...)...)
	}
	for _, p := range aiq.predicates {
		p(selector)
	}
	for _, p := range aiq.order {
		p(selector, authitem.ValidColumn)
	}
	if offset := aiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AuthItemGroupBy is the group-by builder for AuthItem entities.
type AuthItemGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aigb *AuthItemGroupBy) Aggregate(fns ...AggregateFunc) *AuthItemGroupBy {
	aigb.fns = append(aigb.fns, fns...)
	return aigb
}

// Scan applies the group-by query and scans the result into the given value.
func (aigb *AuthItemGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := aigb.path(ctx)
	if err != nil {
		return err
	}
	aigb.sql = query
	return aigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (aigb *AuthItemGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := aigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (aigb *AuthItemGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(aigb.fields) > 1 {
		return nil, errors.New("ent: AuthItemGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := aigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (aigb *AuthItemGroupBy) StringsX(ctx context.Context) []string {
	v, err := aigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aigb *AuthItemGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = aigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authitem.Label}
	default:
		err = fmt.Errorf("ent: AuthItemGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (aigb *AuthItemGroupBy) StringX(ctx context.Context) string {
	v, err := aigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (aigb *AuthItemGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(aigb.fields) > 1 {
		return nil, errors.New("ent: AuthItemGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := aigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (aigb *AuthItemGroupBy) IntsX(ctx context.Context) []int {
	v, err := aigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aigb *AuthItemGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = aigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authitem.Label}
	default:
		err = fmt.Errorf("ent: AuthItemGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (aigb *AuthItemGroupBy) IntX(ctx context.Context) int {
	v, err := aigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (aigb *AuthItemGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(aigb.fields) > 1 {
		return nil, errors.New("ent: AuthItemGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := aigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (aigb *AuthItemGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := aigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aigb *AuthItemGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = aigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authitem.Label}
	default:
		err = fmt.Errorf("ent: AuthItemGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (aigb *AuthItemGroupBy) Float64X(ctx context.Context) float64 {
	v, err := aigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (aigb *AuthItemGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(aigb.fields) > 1 {
		return nil, errors.New("ent: AuthItemGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := aigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (aigb *AuthItemGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := aigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (aigb *AuthItemGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = aigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authitem.Label}
	default:
		err = fmt.Errorf("ent: AuthItemGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (aigb *AuthItemGroupBy) BoolX(ctx context.Context) bool {
	v, err := aigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (aigb *AuthItemGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range aigb.fields {
		if !authitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := aigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (aigb *AuthItemGroupBy) sqlQuery() *sql.Selector {
	selector := aigb.sql
	columns := make([]string, 0, len(aigb.fields)+len(aigb.fns))
	columns = append(columns, aigb.fields...)
	for _, fn := range aigb.fns {
		columns = append(columns, fn(selector, authitem.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(aigb.fields...)
}

// AuthItemSelect is the builder for selecting fields of AuthItem entities.
type AuthItemSelect struct {
	*AuthItemQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ais *AuthItemSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ais.prepareQuery(ctx); err != nil {
		return err
	}
	ais.sql = ais.AuthItemQuery.sqlQuery(ctx)
	return ais.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ais *AuthItemSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ais.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ais *AuthItemSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ais.fields) > 1 {
		return nil, errors.New("ent: AuthItemSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ais.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ais *AuthItemSelect) StringsX(ctx context.Context) []string {
	v, err := ais.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ais *AuthItemSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ais.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authitem.Label}
	default:
		err = fmt.Errorf("ent: AuthItemSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ais *AuthItemSelect) StringX(ctx context.Context) string {
	v, err := ais.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ais *AuthItemSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ais.fields) > 1 {
		return nil, errors.New("ent: AuthItemSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ais.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ais *AuthItemSelect) IntsX(ctx context.Context) []int {
	v, err := ais.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ais *AuthItemSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ais.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authitem.Label}
	default:
		err = fmt.Errorf("ent: AuthItemSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ais *AuthItemSelect) IntX(ctx context.Context) int {
	v, err := ais.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ais *AuthItemSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ais.fields) > 1 {
		return nil, errors.New("ent: AuthItemSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ais.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ais *AuthItemSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ais.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ais *AuthItemSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ais.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authitem.Label}
	default:
		err = fmt.Errorf("ent: AuthItemSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ais *AuthItemSelect) Float64X(ctx context.Context) float64 {
	v, err := ais.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ais *AuthItemSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ais.fields) > 1 {
		return nil, errors.New("ent: AuthItemSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ais.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ais *AuthItemSelect) BoolsX(ctx context.Context) []bool {
	v, err := ais.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ais *AuthItemSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ais.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authitem.Label}
	default:
		err = fmt.Errorf("ent: AuthItemSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ais *AuthItemSelect) BoolX(ctx context.Context) bool {
	v, err := ais.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ais *AuthItemSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ais.sqlQuery().Query()
	if err := ais.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ais *AuthItemSelect) sqlQuery() sql.Querier {
	selector := ais.sql
	selector.Select(selector.Columns(ais.fields...)...)
	return selector
}
