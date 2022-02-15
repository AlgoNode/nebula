// Code generated by SQLBoiler 4.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testIPAddresses(t *testing.T) {
	t.Parallel()

	query := IPAddresses()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testIPAddressesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IPAddress{}
	if err = randomize.Struct(seed, o, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := IPAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIPAddressesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IPAddress{}
	if err = randomize.Struct(seed, o, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := IPAddresses().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := IPAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIPAddressesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IPAddress{}
	if err = randomize.Struct(seed, o, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := IPAddressSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := IPAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIPAddressesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IPAddress{}
	if err = randomize.Struct(seed, o, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := IPAddressExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if IPAddress exists: %s", err)
	}
	if !e {
		t.Errorf("Expected IPAddressExists to return true, but got false.")
	}
}

func testIPAddressesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IPAddress{}
	if err = randomize.Struct(seed, o, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	ipAddressFound, err := FindIPAddress(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if ipAddressFound == nil {
		t.Error("want a record, got nil")
	}
}

func testIPAddressesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IPAddress{}
	if err = randomize.Struct(seed, o, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = IPAddresses().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testIPAddressesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IPAddress{}
	if err = randomize.Struct(seed, o, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := IPAddresses().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testIPAddressesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	ipAddressOne := &IPAddress{}
	ipAddressTwo := &IPAddress{}
	if err = randomize.Struct(seed, ipAddressOne, ipAddressDBTypes, false, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}
	if err = randomize.Struct(seed, ipAddressTwo, ipAddressDBTypes, false, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = ipAddressOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = ipAddressTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := IPAddresses().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testIPAddressesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	ipAddressOne := &IPAddress{}
	ipAddressTwo := &IPAddress{}
	if err = randomize.Struct(seed, ipAddressOne, ipAddressDBTypes, false, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}
	if err = randomize.Struct(seed, ipAddressTwo, ipAddressDBTypes, false, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = ipAddressOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = ipAddressTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := IPAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func ipAddressBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *IPAddress) error {
	*o = IPAddress{}
	return nil
}

func ipAddressAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *IPAddress) error {
	*o = IPAddress{}
	return nil
}

func ipAddressAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *IPAddress) error {
	*o = IPAddress{}
	return nil
}

func ipAddressBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *IPAddress) error {
	*o = IPAddress{}
	return nil
}

func ipAddressAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *IPAddress) error {
	*o = IPAddress{}
	return nil
}

func ipAddressBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *IPAddress) error {
	*o = IPAddress{}
	return nil
}

func ipAddressAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *IPAddress) error {
	*o = IPAddress{}
	return nil
}

func ipAddressBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *IPAddress) error {
	*o = IPAddress{}
	return nil
}

func ipAddressAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *IPAddress) error {
	*o = IPAddress{}
	return nil
}

func testIPAddressesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &IPAddress{}
	o := &IPAddress{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, ipAddressDBTypes, false); err != nil {
		t.Errorf("Unable to randomize IPAddress object: %s", err)
	}

	AddIPAddressHook(boil.BeforeInsertHook, ipAddressBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	ipAddressBeforeInsertHooks = []IPAddressHook{}

	AddIPAddressHook(boil.AfterInsertHook, ipAddressAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	ipAddressAfterInsertHooks = []IPAddressHook{}

	AddIPAddressHook(boil.AfterSelectHook, ipAddressAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	ipAddressAfterSelectHooks = []IPAddressHook{}

	AddIPAddressHook(boil.BeforeUpdateHook, ipAddressBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	ipAddressBeforeUpdateHooks = []IPAddressHook{}

	AddIPAddressHook(boil.AfterUpdateHook, ipAddressAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	ipAddressAfterUpdateHooks = []IPAddressHook{}

	AddIPAddressHook(boil.BeforeDeleteHook, ipAddressBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	ipAddressBeforeDeleteHooks = []IPAddressHook{}

	AddIPAddressHook(boil.AfterDeleteHook, ipAddressAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	ipAddressAfterDeleteHooks = []IPAddressHook{}

	AddIPAddressHook(boil.BeforeUpsertHook, ipAddressBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	ipAddressBeforeUpsertHooks = []IPAddressHook{}

	AddIPAddressHook(boil.AfterUpsertHook, ipAddressAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	ipAddressAfterUpsertHooks = []IPAddressHook{}
}

func testIPAddressesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IPAddress{}
	if err = randomize.Struct(seed, o, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := IPAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testIPAddressesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IPAddress{}
	if err = randomize.Struct(seed, o, ipAddressDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(ipAddressColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := IPAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testIPAddressToManyMultiAddressesXIPAddresses(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a IPAddress
	var b, c MultiAddressesXIPAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, multiAddressesXIPAddressDBTypes, false, multiAddressesXIPAddressColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, multiAddressesXIPAddressDBTypes, false, multiAddressesXIPAddressColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.IPAddressID = a.ID
	c.IPAddressID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.MultiAddressesXIPAddresses().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.IPAddressID == b.IPAddressID {
			bFound = true
		}
		if v.IPAddressID == c.IPAddressID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := IPAddressSlice{&a}
	if err = a.L.LoadMultiAddressesXIPAddresses(ctx, tx, false, (*[]*IPAddress)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MultiAddressesXIPAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.MultiAddressesXIPAddresses = nil
	if err = a.L.LoadMultiAddressesXIPAddresses(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.MultiAddressesXIPAddresses); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testIPAddressToManyAddOpMultiAddressesXIPAddresses(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a IPAddress
	var b, c, d, e MultiAddressesXIPAddress

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, ipAddressDBTypes, false, strmangle.SetComplement(ipAddressPrimaryKeyColumns, ipAddressColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*MultiAddressesXIPAddress{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, multiAddressesXIPAddressDBTypes, false, strmangle.SetComplement(multiAddressesXIPAddressPrimaryKeyColumns, multiAddressesXIPAddressColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*MultiAddressesXIPAddress{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddMultiAddressesXIPAddresses(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.IPAddressID {
			t.Error("foreign key was wrong value", a.ID, first.IPAddressID)
		}
		if a.ID != second.IPAddressID {
			t.Error("foreign key was wrong value", a.ID, second.IPAddressID)
		}

		if first.R.IPAddress != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.IPAddress != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.MultiAddressesXIPAddresses[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.MultiAddressesXIPAddresses[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.MultiAddressesXIPAddresses().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testIPAddressesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IPAddress{}
	if err = randomize.Struct(seed, o, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testIPAddressesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IPAddress{}
	if err = randomize.Struct(seed, o, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := IPAddressSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testIPAddressesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &IPAddress{}
	if err = randomize.Struct(seed, o, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := IPAddresses().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	ipAddressDBTypes = map[string]string{`ID`: `integer`, `UpdatedAt`: `timestamp with time zone`, `CreatedAt`: `timestamp with time zone`, `Address`: `inet`, `Country`: `character varying`, `Asn`: `integer`}
	_                = bytes.MinRead
)

func testIPAddressesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(ipAddressPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(ipAddressAllColumns) == len(ipAddressPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &IPAddress{}
	if err = randomize.Struct(seed, o, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := IPAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, ipAddressDBTypes, true, ipAddressPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testIPAddressesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(ipAddressAllColumns) == len(ipAddressPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &IPAddress{}
	if err = randomize.Struct(seed, o, ipAddressDBTypes, true, ipAddressColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := IPAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, ipAddressDBTypes, true, ipAddressPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(ipAddressAllColumns, ipAddressPrimaryKeyColumns) {
		fields = ipAddressAllColumns
	} else {
		fields = strmangle.SetComplement(
			ipAddressAllColumns,
			ipAddressPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := IPAddressSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testIPAddressesUpsert(t *testing.T) {
	t.Parallel()

	if len(ipAddressAllColumns) == len(ipAddressPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := IPAddress{}
	if err = randomize.Struct(seed, &o, ipAddressDBTypes, true); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert IPAddress: %s", err)
	}

	count, err := IPAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, ipAddressDBTypes, false, ipAddressPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize IPAddress struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert IPAddress: %s", err)
	}

	count, err = IPAddresses().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
