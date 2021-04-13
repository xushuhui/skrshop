// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"
	"skrshop/internal/data/ent/authmsmenu"
	"skrshop/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AuthMsMenuQuery is the builder for querying AuthMsMenu entities.
type AuthMsMenuQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.AuthMsMenu
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AuthMsMenuQuery builder.
func (ammq *AuthMsMenuQuery) Where(ps ...predicate.AuthMsMenu) *AuthMsMenuQuery {
	ammq.predicates = append(ammq.predicates, ps...)
	return ammq
}

// Limit adds a limit step to the query.
func (ammq *AuthMsMenuQuery) Limit(limit int) *AuthMsMenuQuery {
	ammq.limit = &limit
	return ammq
}

// Offset adds an offset step to the query.
func (ammq *AuthMsMenuQuery) Offset(offset int) *AuthMsMenuQuery {
	ammq.offset = &offset
	return ammq
}

// Order adds an order step to the query.
func (ammq *AuthMsMenuQuery) Order(o ...OrderFunc) *AuthMsMenuQuery {
	ammq.order = append(ammq.order, o...)
	return ammq
}

// First returns the first AuthMsMenu entity from the query.
// Returns a *NotFoundError when no AuthMsMenu was found.
func (ammq *AuthMsMenuQuery) First(ctx context.Context) (*AuthMsMenu, error) {
	nodes, err := ammq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{authmsmenu.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ammq *AuthMsMenuQuery) FirstX(ctx context.Context) *AuthMsMenu {
	node, err := ammq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AuthMsMenu ID from the query.
// Returns a *NotFoundError when no AuthMsMenu ID was found.
func (ammq *AuthMsMenuQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ammq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{authmsmenu.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ammq *AuthMsMenuQuery) FirstIDX(ctx context.Context) int {
	id, err := ammq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AuthMsMenu entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one AuthMsMenu entity is not found.
// Returns a *NotFoundError when no AuthMsMenu entities are found.
func (ammq *AuthMsMenuQuery) Only(ctx context.Context) (*AuthMsMenu, error) {
	nodes, err := ammq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{authmsmenu.Label}
	default:
		return nil, &NotSingularError{authmsmenu.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ammq *AuthMsMenuQuery) OnlyX(ctx context.Context) *AuthMsMenu {
	node, err := ammq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AuthMsMenu ID in the query.
// Returns a *NotSingularError when exactly one AuthMsMenu ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ammq *AuthMsMenuQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ammq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{authmsmenu.Label}
	default:
		err = &NotSingularError{authmsmenu.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ammq *AuthMsMenuQuery) OnlyIDX(ctx context.Context) int {
	id, err := ammq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AuthMsMenus.
func (ammq *AuthMsMenuQuery) All(ctx context.Context) ([]*AuthMsMenu, error) {
	if err := ammq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ammq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ammq *AuthMsMenuQuery) AllX(ctx context.Context) []*AuthMsMenu {
	nodes, err := ammq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AuthMsMenu IDs.
func (ammq *AuthMsMenuQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ammq.Select(authmsmenu.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ammq *AuthMsMenuQuery) IDsX(ctx context.Context) []int {
	ids, err := ammq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ammq *AuthMsMenuQuery) Count(ctx context.Context) (int, error) {
	if err := ammq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ammq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ammq *AuthMsMenuQuery) CountX(ctx context.Context) int {
	count, err := ammq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ammq *AuthMsMenuQuery) Exist(ctx context.Context) (bool, error) {
	if err := ammq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ammq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ammq *AuthMsMenuQuery) ExistX(ctx context.Context) bool {
	exist, err := ammq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AuthMsMenuQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ammq *AuthMsMenuQuery) Clone() *AuthMsMenuQuery {
	if ammq == nil {
		return nil
	}
	return &AuthMsMenuQuery{
		config:     ammq.config,
		limit:      ammq.limit,
		offset:     ammq.offset,
		order:      append([]OrderFunc{}, ammq.order...),
		predicates: append([]predicate.AuthMsMenu{}, ammq.predicates...),
		// clone intermediate query.
		sql:  ammq.sql.Clone(),
		path: ammq.path,
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
//	client.AuthMsMenu.Query().
//		GroupBy(authmsmenu.FieldStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ammq *AuthMsMenuQuery) GroupBy(field string, fields ...string) *AuthMsMenuGroupBy {
	group := &AuthMsMenuGroupBy{config: ammq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ammq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ammq.sqlQuery(ctx), nil
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
//	client.AuthMsMenu.Query().
//		Select(authmsmenu.FieldStatus).
//		Scan(ctx, &v)
//
func (ammq *AuthMsMenuQuery) Select(field string, fields ...string) *AuthMsMenuSelect {
	ammq.fields = append([]string{field}, fields...)
	return &AuthMsMenuSelect{AuthMsMenuQuery: ammq}
}

func (ammq *AuthMsMenuQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ammq.fields {
		if !authmsmenu.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ammq.path != nil {
		prev, err := ammq.path(ctx)
		if err != nil {
			return err
		}
		ammq.sql = prev
	}
	return nil
}

func (ammq *AuthMsMenuQuery) sqlAll(ctx context.Context) ([]*AuthMsMenu, error) {
	var (
		nodes = []*AuthMsMenu{}
		_spec = ammq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AuthMsMenu{config: ammq.config}
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
	if err := sqlgraph.QueryNodes(ctx, ammq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ammq *AuthMsMenuQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ammq.querySpec()
	return sqlgraph.CountNodes(ctx, ammq.driver, _spec)
}

func (ammq *AuthMsMenuQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ammq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (ammq *AuthMsMenuQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   authmsmenu.Table,
			Columns: authmsmenu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: authmsmenu.FieldID,
			},
		},
		From:   ammq.sql,
		Unique: true,
	}
	if fields := ammq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, authmsmenu.FieldID)
		for i := range fields {
			if fields[i] != authmsmenu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ammq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ammq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ammq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ammq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, authmsmenu.ValidColumn)
			}
		}
	}
	return _spec
}

func (ammq *AuthMsMenuQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ammq.driver.Dialect())
	t1 := builder.Table(authmsmenu.Table)
	selector := builder.Select(t1.Columns(authmsmenu.Columns...)...).From(t1)
	if ammq.sql != nil {
		selector = ammq.sql
		selector.Select(selector.Columns(authmsmenu.Columns...)...)
	}
	for _, p := range ammq.predicates {
		p(selector)
	}
	for _, p := range ammq.order {
		p(selector, authmsmenu.ValidColumn)
	}
	if offset := ammq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ammq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AuthMsMenuGroupBy is the group-by builder for AuthMsMenu entities.
type AuthMsMenuGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ammgb *AuthMsMenuGroupBy) Aggregate(fns ...AggregateFunc) *AuthMsMenuGroupBy {
	ammgb.fns = append(ammgb.fns, fns...)
	return ammgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ammgb *AuthMsMenuGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ammgb.path(ctx)
	if err != nil {
		return err
	}
	ammgb.sql = query
	return ammgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ammgb *AuthMsMenuGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ammgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ammgb *AuthMsMenuGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ammgb.fields) > 1 {
		return nil, errors.New("ent: AuthMsMenuGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ammgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ammgb *AuthMsMenuGroupBy) StringsX(ctx context.Context) []string {
	v, err := ammgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ammgb *AuthMsMenuGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ammgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authmsmenu.Label}
	default:
		err = fmt.Errorf("ent: AuthMsMenuGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ammgb *AuthMsMenuGroupBy) StringX(ctx context.Context) string {
	v, err := ammgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ammgb *AuthMsMenuGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ammgb.fields) > 1 {
		return nil, errors.New("ent: AuthMsMenuGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ammgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ammgb *AuthMsMenuGroupBy) IntsX(ctx context.Context) []int {
	v, err := ammgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ammgb *AuthMsMenuGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ammgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authmsmenu.Label}
	default:
		err = fmt.Errorf("ent: AuthMsMenuGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ammgb *AuthMsMenuGroupBy) IntX(ctx context.Context) int {
	v, err := ammgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ammgb *AuthMsMenuGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ammgb.fields) > 1 {
		return nil, errors.New("ent: AuthMsMenuGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ammgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ammgb *AuthMsMenuGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ammgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ammgb *AuthMsMenuGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ammgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authmsmenu.Label}
	default:
		err = fmt.Errorf("ent: AuthMsMenuGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ammgb *AuthMsMenuGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ammgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ammgb *AuthMsMenuGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ammgb.fields) > 1 {
		return nil, errors.New("ent: AuthMsMenuGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ammgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ammgb *AuthMsMenuGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ammgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ammgb *AuthMsMenuGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ammgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authmsmenu.Label}
	default:
		err = fmt.Errorf("ent: AuthMsMenuGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ammgb *AuthMsMenuGroupBy) BoolX(ctx context.Context) bool {
	v, err := ammgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ammgb *AuthMsMenuGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ammgb.fields {
		if !authmsmenu.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ammgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ammgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ammgb *AuthMsMenuGroupBy) sqlQuery() *sql.Selector {
	selector := ammgb.sql
	columns := make([]string, 0, len(ammgb.fields)+len(ammgb.fns))
	columns = append(columns, ammgb.fields...)
	for _, fn := range ammgb.fns {
		columns = append(columns, fn(selector, authmsmenu.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(ammgb.fields...)
}

// AuthMsMenuSelect is the builder for selecting fields of AuthMsMenu entities.
type AuthMsMenuSelect struct {
	*AuthMsMenuQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (amms *AuthMsMenuSelect) Scan(ctx context.Context, v interface{}) error {
	if err := amms.prepareQuery(ctx); err != nil {
		return err
	}
	amms.sql = amms.AuthMsMenuQuery.sqlQuery(ctx)
	return amms.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (amms *AuthMsMenuSelect) ScanX(ctx context.Context, v interface{}) {
	if err := amms.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (amms *AuthMsMenuSelect) Strings(ctx context.Context) ([]string, error) {
	if len(amms.fields) > 1 {
		return nil, errors.New("ent: AuthMsMenuSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := amms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (amms *AuthMsMenuSelect) StringsX(ctx context.Context) []string {
	v, err := amms.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (amms *AuthMsMenuSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = amms.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authmsmenu.Label}
	default:
		err = fmt.Errorf("ent: AuthMsMenuSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (amms *AuthMsMenuSelect) StringX(ctx context.Context) string {
	v, err := amms.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (amms *AuthMsMenuSelect) Ints(ctx context.Context) ([]int, error) {
	if len(amms.fields) > 1 {
		return nil, errors.New("ent: AuthMsMenuSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := amms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (amms *AuthMsMenuSelect) IntsX(ctx context.Context) []int {
	v, err := amms.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (amms *AuthMsMenuSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = amms.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authmsmenu.Label}
	default:
		err = fmt.Errorf("ent: AuthMsMenuSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (amms *AuthMsMenuSelect) IntX(ctx context.Context) int {
	v, err := amms.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (amms *AuthMsMenuSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(amms.fields) > 1 {
		return nil, errors.New("ent: AuthMsMenuSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := amms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (amms *AuthMsMenuSelect) Float64sX(ctx context.Context) []float64 {
	v, err := amms.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (amms *AuthMsMenuSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = amms.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authmsmenu.Label}
	default:
		err = fmt.Errorf("ent: AuthMsMenuSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (amms *AuthMsMenuSelect) Float64X(ctx context.Context) float64 {
	v, err := amms.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (amms *AuthMsMenuSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(amms.fields) > 1 {
		return nil, errors.New("ent: AuthMsMenuSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := amms.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (amms *AuthMsMenuSelect) BoolsX(ctx context.Context) []bool {
	v, err := amms.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (amms *AuthMsMenuSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = amms.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{authmsmenu.Label}
	default:
		err = fmt.Errorf("ent: AuthMsMenuSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (amms *AuthMsMenuSelect) BoolX(ctx context.Context) bool {
	v, err := amms.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (amms *AuthMsMenuSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := amms.sqlQuery().Query()
	if err := amms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (amms *AuthMsMenuSelect) sqlQuery() sql.Querier {
	selector := amms.sql
	selector.Select(selector.Columns(amms.fields...)...)
	return selector
}
