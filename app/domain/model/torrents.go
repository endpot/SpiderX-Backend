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

// Torrent is an object representing the database table.
type Torrent struct {
	ID            uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	InfoHash      string    `boil:"info_hash" json:"info_hash" toml:"info_hash" yaml:"info_hash"`
	CategoryID    uint      `boil:"category_id" json:"category_id" toml:"category_id" yaml:"category_id"`
	Title         string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	SimpleDesc    string    `boil:"simple_desc" json:"simple_desc" toml:"simple_desc" yaml:"simple_desc"`
	Description   string    `boil:"description" json:"description" toml:"description" yaml:"description"`
	Size          uint64    `boil:"size" json:"size" toml:"size" yaml:"size"`
	CommentCount  uint      `boil:"comment_count" json:"comment_count" toml:"comment_count" yaml:"comment_count"`
	ViewCount     uint      `boil:"view_count" json:"view_count" toml:"view_count" yaml:"view_count"`
	SeederCount   uint      `boil:"seeder_count" json:"seeder_count" toml:"seeder_count" yaml:"seeder_count"`
	LeecherCount  uint      `boil:"leecher_count" json:"leecher_count" toml:"leecher_count" yaml:"leecher_count"`
	SnatcherCount uint      `boil:"snatcher_count" json:"snatcher_count" toml:"snatcher_count" yaml:"snatcher_count"`
	RewardBonus   uint      `boil:"reward_bonus" json:"reward_bonus" toml:"reward_bonus" yaml:"reward_bonus"`
	IsAnonymous   uint8     `boil:"is_anonymous" json:"is_anonymous" toml:"is_anonymous" yaml:"is_anonymous"`
	PositionLevel uint8     `boil:"position_level" json:"position_level" toml:"position_level" yaml:"position_level"`
	UploaderID    uint      `boil:"uploader_id" json:"uploader_id" toml:"uploader_id" yaml:"uploader_id"`
	OwnerID       uint      `boil:"owner_id" json:"owner_id" toml:"owner_id" yaml:"owner_id"`
	State         uint8     `boil:"state" json:"state" toml:"state" yaml:"state"`
	CreatedAt     null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt     null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt     null.Time `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *torrentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L torrentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TorrentColumns = struct {
	ID            string
	InfoHash      string
	CategoryID    string
	Title         string
	SimpleDesc    string
	Description   string
	Size          string
	CommentCount  string
	ViewCount     string
	SeederCount   string
	LeecherCount  string
	SnatcherCount string
	RewardBonus   string
	IsAnonymous   string
	PositionLevel string
	UploaderID    string
	OwnerID       string
	State         string
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     string
}{
	ID:            "id",
	InfoHash:      "info_hash",
	CategoryID:    "category_id",
	Title:         "title",
	SimpleDesc:    "simple_desc",
	Description:   "description",
	Size:          "size",
	CommentCount:  "comment_count",
	ViewCount:     "view_count",
	SeederCount:   "seeder_count",
	LeecherCount:  "leecher_count",
	SnatcherCount: "snatcher_count",
	RewardBonus:   "reward_bonus",
	IsAnonymous:   "is_anonymous",
	PositionLevel: "position_level",
	UploaderID:    "uploader_id",
	OwnerID:       "owner_id",
	State:         "state",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// Generated where

type whereHelperuint struct{ field string }

func (w whereHelperuint) EQ(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint) NEQ(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint) LT(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint) LTE(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint) GT(x uint) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint) GTE(x uint) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint) IN(slice []uint) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint) NIN(slice []uint) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperuint64 struct{ field string }

func (w whereHelperuint64) EQ(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint64) NEQ(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint64) LT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint64) LTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint64) GT(x uint64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint64) GTE(x uint64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint64) IN(slice []uint64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint64) NIN(slice []uint64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperuint8 struct{ field string }

func (w whereHelperuint8) EQ(x uint8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint8) NEQ(x uint8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint8) LT(x uint8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint8) LTE(x uint8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint8) GT(x uint8) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint8) GTE(x uint8) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint8) IN(slice []uint8) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint8) NIN(slice []uint8) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var TorrentWhere = struct {
	ID            whereHelperuint
	InfoHash      whereHelperstring
	CategoryID    whereHelperuint
	Title         whereHelperstring
	SimpleDesc    whereHelperstring
	Description   whereHelperstring
	Size          whereHelperuint64
	CommentCount  whereHelperuint
	ViewCount     whereHelperuint
	SeederCount   whereHelperuint
	LeecherCount  whereHelperuint
	SnatcherCount whereHelperuint
	RewardBonus   whereHelperuint
	IsAnonymous   whereHelperuint8
	PositionLevel whereHelperuint8
	UploaderID    whereHelperuint
	OwnerID       whereHelperuint
	State         whereHelperuint8
	CreatedAt     whereHelpernull_Time
	UpdatedAt     whereHelpernull_Time
	DeletedAt     whereHelpernull_Time
}{
	ID:            whereHelperuint{field: "`torrents`.`id`"},
	InfoHash:      whereHelperstring{field: "`torrents`.`info_hash`"},
	CategoryID:    whereHelperuint{field: "`torrents`.`category_id`"},
	Title:         whereHelperstring{field: "`torrents`.`title`"},
	SimpleDesc:    whereHelperstring{field: "`torrents`.`simple_desc`"},
	Description:   whereHelperstring{field: "`torrents`.`description`"},
	Size:          whereHelperuint64{field: "`torrents`.`size`"},
	CommentCount:  whereHelperuint{field: "`torrents`.`comment_count`"},
	ViewCount:     whereHelperuint{field: "`torrents`.`view_count`"},
	SeederCount:   whereHelperuint{field: "`torrents`.`seeder_count`"},
	LeecherCount:  whereHelperuint{field: "`torrents`.`leecher_count`"},
	SnatcherCount: whereHelperuint{field: "`torrents`.`snatcher_count`"},
	RewardBonus:   whereHelperuint{field: "`torrents`.`reward_bonus`"},
	IsAnonymous:   whereHelperuint8{field: "`torrents`.`is_anonymous`"},
	PositionLevel: whereHelperuint8{field: "`torrents`.`position_level`"},
	UploaderID:    whereHelperuint{field: "`torrents`.`uploader_id`"},
	OwnerID:       whereHelperuint{field: "`torrents`.`owner_id`"},
	State:         whereHelperuint8{field: "`torrents`.`state`"},
	CreatedAt:     whereHelpernull_Time{field: "`torrents`.`created_at`"},
	UpdatedAt:     whereHelpernull_Time{field: "`torrents`.`updated_at`"},
	DeletedAt:     whereHelpernull_Time{field: "`torrents`.`deleted_at`"},
}

// TorrentRels is where relationship names are stored.
var TorrentRels = struct {
}{}

// torrentR is where relationships are stored.
type torrentR struct {
}

// NewStruct creates a new relationship struct
func (*torrentR) NewStruct() *torrentR {
	return &torrentR{}
}

// torrentL is where Load methods for each relationship are stored.
type torrentL struct{}

var (
	torrentAllColumns            = []string{"id", "info_hash", "category_id", "title", "simple_desc", "description", "size", "comment_count", "view_count", "seeder_count", "leecher_count", "snatcher_count", "reward_bonus", "is_anonymous", "position_level", "uploader_id", "owner_id", "state", "created_at", "updated_at", "deleted_at"}
	torrentColumnsWithoutDefault = []string{"info_hash", "title", "simple_desc", "description", "created_at", "updated_at", "deleted_at"}
	torrentColumnsWithDefault    = []string{"id", "category_id", "size", "comment_count", "view_count", "seeder_count", "leecher_count", "snatcher_count", "reward_bonus", "is_anonymous", "position_level", "uploader_id", "owner_id", "state"}
	torrentPrimaryKeyColumns     = []string{"id"}
)

type (
	// TorrentSlice is an alias for a slice of pointers to Torrent.
	// This should generally be used opposed to []Torrent.
	TorrentSlice []*Torrent
	// TorrentHook is the signature for custom Torrent hook methods
	TorrentHook func(context.Context, boil.ContextExecutor, *Torrent) error

	torrentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	torrentType                 = reflect.TypeOf(&Torrent{})
	torrentMapping              = queries.MakeStructMapping(torrentType)
	torrentPrimaryKeyMapping, _ = queries.BindMapping(torrentType, torrentMapping, torrentPrimaryKeyColumns)
	torrentInsertCacheMut       sync.RWMutex
	torrentInsertCache          = make(map[string]insertCache)
	torrentUpdateCacheMut       sync.RWMutex
	torrentUpdateCache          = make(map[string]updateCache)
	torrentUpsertCacheMut       sync.RWMutex
	torrentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var torrentBeforeInsertHooks []TorrentHook
var torrentBeforeUpdateHooks []TorrentHook
var torrentBeforeDeleteHooks []TorrentHook
var torrentBeforeUpsertHooks []TorrentHook

var torrentAfterInsertHooks []TorrentHook
var torrentAfterSelectHooks []TorrentHook
var torrentAfterUpdateHooks []TorrentHook
var torrentAfterDeleteHooks []TorrentHook
var torrentAfterUpsertHooks []TorrentHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Torrent) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range torrentBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Torrent) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range torrentBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Torrent) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range torrentBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Torrent) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range torrentBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Torrent) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range torrentAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Torrent) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range torrentAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Torrent) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range torrentAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Torrent) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range torrentAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Torrent) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range torrentAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTorrentHook registers your hook function for all future operations.
func AddTorrentHook(hookPoint boil.HookPoint, torrentHook TorrentHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		torrentBeforeInsertHooks = append(torrentBeforeInsertHooks, torrentHook)
	case boil.BeforeUpdateHook:
		torrentBeforeUpdateHooks = append(torrentBeforeUpdateHooks, torrentHook)
	case boil.BeforeDeleteHook:
		torrentBeforeDeleteHooks = append(torrentBeforeDeleteHooks, torrentHook)
	case boil.BeforeUpsertHook:
		torrentBeforeUpsertHooks = append(torrentBeforeUpsertHooks, torrentHook)
	case boil.AfterInsertHook:
		torrentAfterInsertHooks = append(torrentAfterInsertHooks, torrentHook)
	case boil.AfterSelectHook:
		torrentAfterSelectHooks = append(torrentAfterSelectHooks, torrentHook)
	case boil.AfterUpdateHook:
		torrentAfterUpdateHooks = append(torrentAfterUpdateHooks, torrentHook)
	case boil.AfterDeleteHook:
		torrentAfterDeleteHooks = append(torrentAfterDeleteHooks, torrentHook)
	case boil.AfterUpsertHook:
		torrentAfterUpsertHooks = append(torrentAfterUpsertHooks, torrentHook)
	}
}

// One returns a single torrent record from the query.
func (q torrentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Torrent, error) {
	o := &Torrent{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: failed to execute a one query for torrents")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Torrent records from the query.
func (q torrentQuery) All(ctx context.Context, exec boil.ContextExecutor) (TorrentSlice, error) {
	var o []*Torrent

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "model: failed to assign all query results to Torrent slice")
	}

	if len(torrentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Torrent records in the query.
func (q torrentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to count torrents rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q torrentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "model: failed to check if torrents exists")
	}

	return count > 0, nil
}

// Torrents retrieves all the records using an executor.
func Torrents(mods ...qm.QueryMod) torrentQuery {
	mods = append(mods, qm.From("`torrents`"), qmhelper.WhereIsNull("`torrents`.`deleted_at`"))
	return torrentQuery{NewQuery(mods...)}
}

// FindTorrent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTorrent(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*Torrent, error) {
	torrentObj := &Torrent{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `torrents` where `id`=? and `deleted_at` is null", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, torrentObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "model: unable to select from torrents")
	}

	return torrentObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Torrent) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("model: no torrents provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(torrentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	torrentInsertCacheMut.RLock()
	cache, cached := torrentInsertCache[key]
	torrentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			torrentAllColumns,
			torrentColumnsWithDefault,
			torrentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(torrentType, torrentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(torrentType, torrentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `torrents` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `torrents` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `torrents` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, torrentPrimaryKeyColumns))
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
		return errors.Wrap(err, "model: unable to insert into torrents")
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

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == torrentMapping["id"] {
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
		return errors.Wrap(err, "model: unable to populate default values for torrents")
	}

CacheNoHooks:
	if !cached {
		torrentInsertCacheMut.Lock()
		torrentInsertCache[key] = cache
		torrentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Torrent.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Torrent) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	torrentUpdateCacheMut.RLock()
	cache, cached := torrentUpdateCache[key]
	torrentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			torrentAllColumns,
			torrentPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("model: unable to update torrents, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `torrents` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, torrentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(torrentType, torrentMapping, append(wl, torrentPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "model: unable to update torrents row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by update for torrents")
	}

	if !cached {
		torrentUpdateCacheMut.Lock()
		torrentUpdateCache[key] = cache
		torrentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q torrentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all for torrents")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected for torrents")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TorrentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), torrentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `torrents` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, torrentPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to update all in torrent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to retrieve rows affected all in update all torrent")
	}
	return rowsAff, nil
}

var mySQLTorrentUniqueColumns = []string{
	"id",
	"info_hash",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Torrent) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("model: no torrents provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(torrentColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTorrentUniqueColumns, o)

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

	torrentUpsertCacheMut.RLock()
	cache, cached := torrentUpsertCache[key]
	torrentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			torrentAllColumns,
			torrentColumnsWithDefault,
			torrentColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			torrentAllColumns,
			torrentPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("model: unable to upsert torrents, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "torrents", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `torrents` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(torrentType, torrentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(torrentType, torrentMapping, ret)
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
		return errors.Wrap(err, "model: unable to upsert for torrents")
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

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == torrentMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(torrentType, torrentMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "model: unable to retrieve unique values for torrents")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "model: unable to populate default values for torrents")
	}

CacheNoHooks:
	if !cached {
		torrentUpsertCacheMut.Lock()
		torrentUpsertCache[key] = cache
		torrentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Torrent record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Torrent) Delete(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if o == nil {
		return 0, errors.New("model: no Torrent provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), torrentPrimaryKeyMapping)
		sql = "DELETE FROM `torrents` WHERE `id`=?"
	} else {
		currTime := time.Now().In(boil.GetLocation())
		o.DeletedAt = null.TimeFrom(currTime)
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE `torrents` SET %s WHERE `id`=?",
			strmangle.SetParamNames("`", "`", 0, wl),
		)
		valueMapping, err := queries.BindMapping(torrentType, torrentMapping, append(wl, torrentPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
		args = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), valueMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete from torrents")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by delete for torrents")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q torrentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("model: no torrentQuery provided for delete all")
	}

	if hardDelete {
		queries.SetDelete(q.Query)
	} else {
		currTime := time.Now().In(boil.GetLocation())
		queries.SetUpdate(q.Query, M{"deleted_at": currTime})
	}

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from torrents")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for torrents")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TorrentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor, hardDelete bool) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(torrentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var (
		sql  string
		args []interface{}
	)
	if hardDelete {
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), torrentPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
		}
		sql = "DELETE FROM `torrents` WHERE " +
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, torrentPrimaryKeyColumns, len(o))
	} else {
		currTime := time.Now().In(boil.GetLocation())
		for _, obj := range o {
			pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), torrentPrimaryKeyMapping)
			args = append(args, pkeyArgs...)
			obj.DeletedAt = null.TimeFrom(currTime)
		}
		wl := []string{"deleted_at"}
		sql = fmt.Sprintf("UPDATE `torrents` SET %s WHERE "+
			strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, torrentPrimaryKeyColumns, len(o)),
			strmangle.SetParamNames("`", "`", 0, wl),
		)
		args = append([]interface{}{currTime}, args...)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "model: unable to delete all from torrent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "model: failed to get rows affected by deleteall for torrents")
	}

	if len(torrentAfterDeleteHooks) != 0 {
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
func (o *Torrent) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTorrent(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TorrentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TorrentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), torrentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `torrents`.* FROM `torrents` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, torrentPrimaryKeyColumns, len(*o)) +
		"and `deleted_at` is null"

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "model: unable to reload all in TorrentSlice")
	}

	*o = slice

	return nil
}

// TorrentExists checks if the Torrent row exists.
func TorrentExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `torrents` where `id`=? and `deleted_at` is null limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "model: unable to check if torrents exists")
	}

	return exists, nil
}
