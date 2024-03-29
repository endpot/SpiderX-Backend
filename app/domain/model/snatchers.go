// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Snatcher is an object representing the database table.
type Snatcher struct {
	ID              int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	TorrentID       uint      `boil:"torrent_id" json:"torrent_id" toml:"torrent_id" yaml:"torrent_id"`
	UserID          uint      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Ipv4            string    `boil:"ipv4" json:"ipv4" toml:"ipv4" yaml:"ipv4"`
	Ipv6            string    `boil:"ipv6" json:"ipv6" toml:"ipv6" yaml:"ipv6"`
	Port            uint16    `boil:"port" json:"port" toml:"port" yaml:"port"`
	UploadedBytes   uint64    `boil:"uploaded_bytes" json:"uploaded_bytes" toml:"uploaded_bytes" yaml:"uploaded_bytes"`
	DownloadedBytes uint64    `boil:"downloaded_bytes" json:"downloaded_bytes" toml:"downloaded_bytes" yaml:"downloaded_bytes"`
	LeftBytes       uint64    `boil:"left_bytes" json:"left_bytes" toml:"left_bytes" yaml:"left_bytes"`
	SeedTime        uint      `boil:"seed_time" json:"seed_time" toml:"seed_time" yaml:"seed_time"`
	LeechTime       uint      `boil:"leech_time" json:"leech_time" toml:"leech_time" yaml:"leech_time"`
	IsFinished      uint8     `boil:"is_finished" json:"is_finished" toml:"is_finished" yaml:"is_finished"`
	CreatedAt       null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt       null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	FinishedAt      null.Time `boil:"finished_at" json:"finished_at,omitempty" toml:"finished_at" yaml:"finished_at,omitempty"`

	R *snatcherR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L snatcherL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SnatcherColumns = struct {
	ID              string
	TorrentID       string
	UserID          string
	Ipv4            string
	Ipv6            string
	Port            string
	UploadedBytes   string
	DownloadedBytes string
	LeftBytes       string
	SeedTime        string
	LeechTime       string
	IsFinished      string
	CreatedAt       string
	UpdatedAt       string
	FinishedAt      string
}{
	ID:              "id",
	TorrentID:       "torrent_id",
	UserID:          "user_id",
	Ipv4:            "ipv4",
	Ipv6:            "ipv6",
	Port:            "port",
	UploadedBytes:   "uploaded_bytes",
	DownloadedBytes: "downloaded_bytes",
	LeftBytes:       "left_bytes",
	SeedTime:        "seed_time",
	LeechTime:       "leech_time",
	IsFinished:      "is_finished",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
	FinishedAt:      "finished_at",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var SnatcherWhere = struct {
	ID              whereHelperint
	TorrentID       whereHelperuint
	UserID          whereHelperuint
	Ipv4            whereHelperstring
	Ipv6            whereHelperstring
	Port            whereHelperuint16
	UploadedBytes   whereHelperuint64
	DownloadedBytes whereHelperuint64
	LeftBytes       whereHelperuint64
	SeedTime        whereHelperuint
	LeechTime       whereHelperuint
	IsFinished      whereHelperuint8
	CreatedAt       whereHelpernull_Time
	UpdatedAt       whereHelpernull_Time
	FinishedAt      whereHelpernull_Time
}{
	ID:              whereHelperint{field: "`snatchers`.`id`"},
	TorrentID:       whereHelperuint{field: "`snatchers`.`torrent_id`"},
	UserID:          whereHelperuint{field: "`snatchers`.`user_id`"},
	Ipv4:            whereHelperstring{field: "`snatchers`.`ipv4`"},
	Ipv6:            whereHelperstring{field: "`snatchers`.`ipv6`"},
	Port:            whereHelperuint16{field: "`snatchers`.`port`"},
	UploadedBytes:   whereHelperuint64{field: "`snatchers`.`uploaded_bytes`"},
	DownloadedBytes: whereHelperuint64{field: "`snatchers`.`downloaded_bytes`"},
	LeftBytes:       whereHelperuint64{field: "`snatchers`.`left_bytes`"},
	SeedTime:        whereHelperuint{field: "`snatchers`.`seed_time`"},
	LeechTime:       whereHelperuint{field: "`snatchers`.`leech_time`"},
	IsFinished:      whereHelperuint8{field: "`snatchers`.`is_finished`"},
	CreatedAt:       whereHelpernull_Time{field: "`snatchers`.`created_at`"},
	UpdatedAt:       whereHelpernull_Time{field: "`snatchers`.`updated_at`"},
	FinishedAt:      whereHelpernull_Time{field: "`snatchers`.`finished_at`"},
}

// SnatcherRels is where relationship names are stored.
var SnatcherRels = struct {
}{}

// snatcherR is where relationships are stored.
type snatcherR struct {
}

// NewStruct creates a new relationship struct
func (*snatcherR) NewStruct() *snatcherR {
	return &snatcherR{}
}

// snatcherL is where Load methods for each relationship are stored.
type snatcherL struct{}

var (
	snatcherAllColumns            = []string{"id", "torrent_id", "user_id", "ipv4", "ipv6", "port", "uploaded_bytes", "downloaded_bytes", "left_bytes", "seed_time", "leech_time", "is_finished", "created_at", "updated_at", "finished_at"}
	snatcherColumnsWithoutDefault = []string{"ipv4", "ipv6", "finished_at"}
	snatcherColumnsWithDefault    = []string{"id", "torrent_id", "user_id", "port", "uploaded_bytes", "downloaded_bytes", "left_bytes", "seed_time", "leech_time", "is_finished", "created_at", "updated_at"}
	snatcherPrimaryKeyColumns     = []string{"id"}
)

type (
	// SnatcherSlice is an alias for a slice of pointers to Snatcher.
	// This should generally be used opposed to []Snatcher.
	SnatcherSlice []*Snatcher
	// SnatcherHook is the signature for custom Snatcher hook methods
	SnatcherHook func(context.Context, boil.ContextExecutor, *Snatcher) error

	snatcherQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	snatcherType                 = reflect.TypeOf(&Snatcher{})
	snatcherMapping              = queries.MakeStructMapping(snatcherType)
	snatcherPrimaryKeyMapping, _ = queries.BindMapping(snatcherType, snatcherMapping, snatcherPrimaryKeyColumns)
	snatcherInsertCacheMut       sync.RWMutex
	snatcherInsertCache          = make(map[string]insertCache)
	snatcherUpdateCacheMut       sync.RWMutex
	snatcherUpdateCache          = make(map[string]updateCache)
	snatcherUpsertCacheMut       sync.RWMutex
	snatcherUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var snatcherBeforeInsertHooks []SnatcherHook
var snatcherBeforeUpdateHooks []SnatcherHook
var snatcherBeforeDeleteHooks []SnatcherHook
var snatcherBeforeUpsertHooks []SnatcherHook

var snatcherAfterInsertHooks []SnatcherHook
var snatcherAfterSelectHooks []SnatcherHook
var snatcherAfterUpdateHooks []SnatcherHook
var snatcherAfterDeleteHooks []SnatcherHook
var snatcherAfterUpsertHooks []SnatcherHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Snatcher) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range snatcherBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Snatcher) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range snatcherBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Snatcher) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range snatcherBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Snatcher) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range snatcherBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Snatcher) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range snatcherAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Snatcher) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range snatcherAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Snatcher) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range snatcherAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Snatcher) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range snatcherAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Snatcher) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range snatcherAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSnatcherHook registers your hook function for all future operations.
func AddSnatcherHook(hookPoint boil.HookPoint, snatcherHook SnatcherHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		snatcherBeforeInsertHooks = append(snatcherBeforeInsertHooks, snatcherHook)
	case boil.BeforeUpdateHook:
		snatcherBeforeUpdateHooks = append(snatcherBeforeUpdateHooks, snatcherHook)
	case boil.BeforeDeleteHook:
		snatcherBeforeDeleteHooks = append(snatcherBeforeDeleteHooks, snatcherHook)
	case boil.BeforeUpsertHook:
		snatcherBeforeUpsertHooks = append(snatcherBeforeUpsertHooks, snatcherHook)
	case boil.AfterInsertHook:
		snatcherAfterInsertHooks = append(snatcherAfterInsertHooks, snatcherHook)
	case boil.AfterSelectHook:
		snatcherAfterSelectHooks = append(snatcherAfterSelectHooks, snatcherHook)
	case boil.AfterUpdateHook:
		snatcherAfterUpdateHooks = append(snatcherAfterUpdateHooks, snatcherHook)
	case boil.AfterDeleteHook:
		snatcherAfterDeleteHooks = append(snatcherAfterDeleteHooks, snatcherHook)
	case boil.AfterUpsertHook:
		snatcherAfterUpsertHooks = append(snatcherAfterUpsertHooks, snatcherHook)
	}
}

// One returns a single snatcher record from the query.
func (q snatcherQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Snatcher, error) {
	o := &Snatcher{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for snatchers")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Snatcher records from the query.
func (q snatcherQuery) All(ctx context.Context, exec boil.ContextExecutor) (SnatcherSlice, error) {
	var o []*Snatcher

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to Snatcher slice")
	}

	if len(snatcherAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Snatcher records in the query.
func (q snatcherQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count snatchers rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q snatcherQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if snatchers exists")
	}

	return count > 0, nil
}

// Snatchers retrieves all the records using an executor.
func Snatchers(mods ...qm.QueryMod) snatcherQuery {
	mods = append(mods, qm.From("`snatchers`"))
	return snatcherQuery{NewQuery(mods...)}
}

// FindSnatcher retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSnatcher(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Snatcher, error) {
	snatcherObj := &Snatcher{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `snatchers` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, snatcherObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from snatchers")
	}

	return snatcherObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Snatcher) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no snatchers provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(snatcherColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	snatcherInsertCacheMut.RLock()
	cache, cached := snatcherInsertCache[key]
	snatcherInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			snatcherAllColumns,
			snatcherColumnsWithDefault,
			snatcherColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(snatcherType, snatcherMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(snatcherType, snatcherMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `snatchers` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `snatchers` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `snatchers` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, snatcherPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to insert into snatchers")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == snatcherMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for snatchers")
	}

CacheNoHooks:
	if !cached {
		snatcherInsertCacheMut.Lock()
		snatcherInsertCache[key] = cache
		snatcherInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Snatcher.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Snatcher) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	snatcherUpdateCacheMut.RLock()
	cache, cached := snatcherUpdateCache[key]
	snatcherUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			snatcherAllColumns,
			snatcherPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update snatchers, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `snatchers` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, snatcherPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(snatcherType, snatcherMapping, append(wl, snatcherPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update snatchers row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for snatchers")
	}

	if !cached {
		snatcherUpdateCacheMut.Lock()
		snatcherUpdateCache[key] = cache
		snatcherUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q snatcherQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for snatchers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for snatchers")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SnatcherSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("model: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), snatcherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `snatchers` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, snatcherPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in snatcher slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all snatcher")
	}
	return rowsAff, nil
}

var mySQLSnatcherUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Snatcher) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no snatchers provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(snatcherColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLSnatcherUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	snatcherUpsertCacheMut.RLock()
	cache, cached := snatcherUpsertCache[key]
	snatcherUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			snatcherAllColumns,
			snatcherColumnsWithDefault,
			snatcherColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			snatcherAllColumns,
			snatcherPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert snatchers, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "snatchers", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `snatchers` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(snatcherType, snatcherMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(snatcherType, snatcherMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "model: unable to upsert for snatchers")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == snatcherMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(snatcherType, snatcherMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for snatchers")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for snatchers")
	}

CacheNoHooks:
	if !cached {
		snatcherUpsertCacheMut.Lock()
		snatcherUpsertCache[key] = cache
		snatcherUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Snatcher record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Snatcher) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no Snatcher provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), snatcherPrimaryKeyMapping)
	sql := "DELETE FROM `snatchers` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from snatchers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for snatchers")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q snatcherQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no snatcherQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from snatchers")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for snatchers")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SnatcherSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(snatcherBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), snatcherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `snatchers` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, snatcherPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from snatcher slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for snatchers")
	}

	if len(snatcherAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Snatcher) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSnatcher(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SnatcherSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SnatcherSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), snatcherPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `snatchers`.* FROM `snatchers` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, snatcherPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in SnatcherSlice")
	}

	*o = slice

	return nil
}

// SnatcherExists checks if the Snatcher row exists.
func SnatcherExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `snatchers` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if snatchers exists")
	}

	return exists, nil
}
