// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/samirgadkari/search/ent/doc"
	"github.com/samirgadkari/search/ent/predicate"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeDoc = "Doc"
)

// DocMutation represents an operation that mutates the Doc nodes in the graph.
type DocMutation struct {
	config
	op            Op
	typ           string
	id            *uint32
	wordInts      *[]byte
	inputDocId    *string
	userId        *string
	businessId    *string
	stars         *float32
	addstars      *float32
	useful        *int16
	adduseful     *int16
	funny         *int16
	addfunny      *int16
	cool          *int16
	addcool       *int16
	text          *string
	date          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Doc, error)
	predicates    []predicate.Doc
}

var _ ent.Mutation = (*DocMutation)(nil)

// docOption allows management of the mutation configuration using functional options.
type docOption func(*DocMutation)

// newDocMutation creates new mutation for the Doc entity.
func newDocMutation(c config, op Op, opts ...docOption) *DocMutation {
	m := &DocMutation{
		config:        c,
		op:            op,
		typ:           TypeDoc,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withDocID sets the ID field of the mutation.
func withDocID(id uint32) docOption {
	return func(m *DocMutation) {
		var (
			err   error
			once  sync.Once
			value *Doc
		)
		m.oldValue = func(ctx context.Context) (*Doc, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Doc.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withDoc sets the old Doc of the mutation.
func withDoc(node *Doc) docOption {
	return func(m *DocMutation) {
		m.oldValue = func(context.Context) (*Doc, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m DocMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m DocMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Doc entities.
func (m *DocMutation) SetID(id uint32) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *DocMutation) ID() (id uint32, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *DocMutation) IDs(ctx context.Context) ([]uint32, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uint32{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Doc.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetWordInts sets the "wordInts" field.
func (m *DocMutation) SetWordInts(b []byte) {
	m.wordInts = &b
}

// WordInts returns the value of the "wordInts" field in the mutation.
func (m *DocMutation) WordInts() (r []byte, exists bool) {
	v := m.wordInts
	if v == nil {
		return
	}
	return *v, true
}

// OldWordInts returns the old "wordInts" field's value of the Doc entity.
// If the Doc object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DocMutation) OldWordInts(ctx context.Context) (v []byte, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldWordInts is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldWordInts requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldWordInts: %w", err)
	}
	return oldValue.WordInts, nil
}

// ResetWordInts resets all changes to the "wordInts" field.
func (m *DocMutation) ResetWordInts() {
	m.wordInts = nil
}

// SetInputDocId sets the "inputDocId" field.
func (m *DocMutation) SetInputDocId(s string) {
	m.inputDocId = &s
}

// InputDocId returns the value of the "inputDocId" field in the mutation.
func (m *DocMutation) InputDocId() (r string, exists bool) {
	v := m.inputDocId
	if v == nil {
		return
	}
	return *v, true
}

// OldInputDocId returns the old "inputDocId" field's value of the Doc entity.
// If the Doc object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DocMutation) OldInputDocId(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldInputDocId is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldInputDocId requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldInputDocId: %w", err)
	}
	return oldValue.InputDocId, nil
}

// ResetInputDocId resets all changes to the "inputDocId" field.
func (m *DocMutation) ResetInputDocId() {
	m.inputDocId = nil
}

// SetUserId sets the "userId" field.
func (m *DocMutation) SetUserId(s string) {
	m.userId = &s
}

// UserId returns the value of the "userId" field in the mutation.
func (m *DocMutation) UserId() (r string, exists bool) {
	v := m.userId
	if v == nil {
		return
	}
	return *v, true
}

// OldUserId returns the old "userId" field's value of the Doc entity.
// If the Doc object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DocMutation) OldUserId(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserId is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserId requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserId: %w", err)
	}
	return oldValue.UserId, nil
}

// ResetUserId resets all changes to the "userId" field.
func (m *DocMutation) ResetUserId() {
	m.userId = nil
}

// SetBusinessId sets the "businessId" field.
func (m *DocMutation) SetBusinessId(s string) {
	m.businessId = &s
}

// BusinessId returns the value of the "businessId" field in the mutation.
func (m *DocMutation) BusinessId() (r string, exists bool) {
	v := m.businessId
	if v == nil {
		return
	}
	return *v, true
}

// OldBusinessId returns the old "businessId" field's value of the Doc entity.
// If the Doc object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DocMutation) OldBusinessId(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBusinessId is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBusinessId requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBusinessId: %w", err)
	}
	return oldValue.BusinessId, nil
}

// ResetBusinessId resets all changes to the "businessId" field.
func (m *DocMutation) ResetBusinessId() {
	m.businessId = nil
}

// SetStars sets the "stars" field.
func (m *DocMutation) SetStars(f float32) {
	m.stars = &f
	m.addstars = nil
}

// Stars returns the value of the "stars" field in the mutation.
func (m *DocMutation) Stars() (r float32, exists bool) {
	v := m.stars
	if v == nil {
		return
	}
	return *v, true
}

// OldStars returns the old "stars" field's value of the Doc entity.
// If the Doc object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DocMutation) OldStars(ctx context.Context) (v float32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStars is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStars requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStars: %w", err)
	}
	return oldValue.Stars, nil
}

// AddStars adds f to the "stars" field.
func (m *DocMutation) AddStars(f float32) {
	if m.addstars != nil {
		*m.addstars += f
	} else {
		m.addstars = &f
	}
}

// AddedStars returns the value that was added to the "stars" field in this mutation.
func (m *DocMutation) AddedStars() (r float32, exists bool) {
	v := m.addstars
	if v == nil {
		return
	}
	return *v, true
}

// ResetStars resets all changes to the "stars" field.
func (m *DocMutation) ResetStars() {
	m.stars = nil
	m.addstars = nil
}

// SetUseful sets the "useful" field.
func (m *DocMutation) SetUseful(i int16) {
	m.useful = &i
	m.adduseful = nil
}

// Useful returns the value of the "useful" field in the mutation.
func (m *DocMutation) Useful() (r int16, exists bool) {
	v := m.useful
	if v == nil {
		return
	}
	return *v, true
}

// OldUseful returns the old "useful" field's value of the Doc entity.
// If the Doc object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DocMutation) OldUseful(ctx context.Context) (v int16, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUseful is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUseful requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUseful: %w", err)
	}
	return oldValue.Useful, nil
}

// AddUseful adds i to the "useful" field.
func (m *DocMutation) AddUseful(i int16) {
	if m.adduseful != nil {
		*m.adduseful += i
	} else {
		m.adduseful = &i
	}
}

// AddedUseful returns the value that was added to the "useful" field in this mutation.
func (m *DocMutation) AddedUseful() (r int16, exists bool) {
	v := m.adduseful
	if v == nil {
		return
	}
	return *v, true
}

// ResetUseful resets all changes to the "useful" field.
func (m *DocMutation) ResetUseful() {
	m.useful = nil
	m.adduseful = nil
}

// SetFunny sets the "funny" field.
func (m *DocMutation) SetFunny(i int16) {
	m.funny = &i
	m.addfunny = nil
}

// Funny returns the value of the "funny" field in the mutation.
func (m *DocMutation) Funny() (r int16, exists bool) {
	v := m.funny
	if v == nil {
		return
	}
	return *v, true
}

// OldFunny returns the old "funny" field's value of the Doc entity.
// If the Doc object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DocMutation) OldFunny(ctx context.Context) (v int16, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFunny is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFunny requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFunny: %w", err)
	}
	return oldValue.Funny, nil
}

// AddFunny adds i to the "funny" field.
func (m *DocMutation) AddFunny(i int16) {
	if m.addfunny != nil {
		*m.addfunny += i
	} else {
		m.addfunny = &i
	}
}

// AddedFunny returns the value that was added to the "funny" field in this mutation.
func (m *DocMutation) AddedFunny() (r int16, exists bool) {
	v := m.addfunny
	if v == nil {
		return
	}
	return *v, true
}

// ResetFunny resets all changes to the "funny" field.
func (m *DocMutation) ResetFunny() {
	m.funny = nil
	m.addfunny = nil
}

// SetCool sets the "cool" field.
func (m *DocMutation) SetCool(i int16) {
	m.cool = &i
	m.addcool = nil
}

// Cool returns the value of the "cool" field in the mutation.
func (m *DocMutation) Cool() (r int16, exists bool) {
	v := m.cool
	if v == nil {
		return
	}
	return *v, true
}

// OldCool returns the old "cool" field's value of the Doc entity.
// If the Doc object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DocMutation) OldCool(ctx context.Context) (v int16, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCool is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCool requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCool: %w", err)
	}
	return oldValue.Cool, nil
}

// AddCool adds i to the "cool" field.
func (m *DocMutation) AddCool(i int16) {
	if m.addcool != nil {
		*m.addcool += i
	} else {
		m.addcool = &i
	}
}

// AddedCool returns the value that was added to the "cool" field in this mutation.
func (m *DocMutation) AddedCool() (r int16, exists bool) {
	v := m.addcool
	if v == nil {
		return
	}
	return *v, true
}

// ResetCool resets all changes to the "cool" field.
func (m *DocMutation) ResetCool() {
	m.cool = nil
	m.addcool = nil
}

// SetText sets the "text" field.
func (m *DocMutation) SetText(s string) {
	m.text = &s
}

// Text returns the value of the "text" field in the mutation.
func (m *DocMutation) Text() (r string, exists bool) {
	v := m.text
	if v == nil {
		return
	}
	return *v, true
}

// OldText returns the old "text" field's value of the Doc entity.
// If the Doc object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DocMutation) OldText(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldText is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldText requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldText: %w", err)
	}
	return oldValue.Text, nil
}

// ResetText resets all changes to the "text" field.
func (m *DocMutation) ResetText() {
	m.text = nil
}

// SetDate sets the "date" field.
func (m *DocMutation) SetDate(s string) {
	m.date = &s
}

// Date returns the value of the "date" field in the mutation.
func (m *DocMutation) Date() (r string, exists bool) {
	v := m.date
	if v == nil {
		return
	}
	return *v, true
}

// OldDate returns the old "date" field's value of the Doc entity.
// If the Doc object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *DocMutation) OldDate(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDate: %w", err)
	}
	return oldValue.Date, nil
}

// ResetDate resets all changes to the "date" field.
func (m *DocMutation) ResetDate() {
	m.date = nil
}

// Where appends a list predicates to the DocMutation builder.
func (m *DocMutation) Where(ps ...predicate.Doc) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *DocMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Doc).
func (m *DocMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *DocMutation) Fields() []string {
	fields := make([]string, 0, 10)
	if m.wordInts != nil {
		fields = append(fields, doc.FieldWordInts)
	}
	if m.inputDocId != nil {
		fields = append(fields, doc.FieldInputDocId)
	}
	if m.userId != nil {
		fields = append(fields, doc.FieldUserId)
	}
	if m.businessId != nil {
		fields = append(fields, doc.FieldBusinessId)
	}
	if m.stars != nil {
		fields = append(fields, doc.FieldStars)
	}
	if m.useful != nil {
		fields = append(fields, doc.FieldUseful)
	}
	if m.funny != nil {
		fields = append(fields, doc.FieldFunny)
	}
	if m.cool != nil {
		fields = append(fields, doc.FieldCool)
	}
	if m.text != nil {
		fields = append(fields, doc.FieldText)
	}
	if m.date != nil {
		fields = append(fields, doc.FieldDate)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *DocMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case doc.FieldWordInts:
		return m.WordInts()
	case doc.FieldInputDocId:
		return m.InputDocId()
	case doc.FieldUserId:
		return m.UserId()
	case doc.FieldBusinessId:
		return m.BusinessId()
	case doc.FieldStars:
		return m.Stars()
	case doc.FieldUseful:
		return m.Useful()
	case doc.FieldFunny:
		return m.Funny()
	case doc.FieldCool:
		return m.Cool()
	case doc.FieldText:
		return m.Text()
	case doc.FieldDate:
		return m.Date()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *DocMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case doc.FieldWordInts:
		return m.OldWordInts(ctx)
	case doc.FieldInputDocId:
		return m.OldInputDocId(ctx)
	case doc.FieldUserId:
		return m.OldUserId(ctx)
	case doc.FieldBusinessId:
		return m.OldBusinessId(ctx)
	case doc.FieldStars:
		return m.OldStars(ctx)
	case doc.FieldUseful:
		return m.OldUseful(ctx)
	case doc.FieldFunny:
		return m.OldFunny(ctx)
	case doc.FieldCool:
		return m.OldCool(ctx)
	case doc.FieldText:
		return m.OldText(ctx)
	case doc.FieldDate:
		return m.OldDate(ctx)
	}
	return nil, fmt.Errorf("unknown Doc field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DocMutation) SetField(name string, value ent.Value) error {
	switch name {
	case doc.FieldWordInts:
		v, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetWordInts(v)
		return nil
	case doc.FieldInputDocId:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetInputDocId(v)
		return nil
	case doc.FieldUserId:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserId(v)
		return nil
	case doc.FieldBusinessId:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBusinessId(v)
		return nil
	case doc.FieldStars:
		v, ok := value.(float32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStars(v)
		return nil
	case doc.FieldUseful:
		v, ok := value.(int16)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUseful(v)
		return nil
	case doc.FieldFunny:
		v, ok := value.(int16)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFunny(v)
		return nil
	case doc.FieldCool:
		v, ok := value.(int16)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCool(v)
		return nil
	case doc.FieldText:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetText(v)
		return nil
	case doc.FieldDate:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDate(v)
		return nil
	}
	return fmt.Errorf("unknown Doc field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *DocMutation) AddedFields() []string {
	var fields []string
	if m.addstars != nil {
		fields = append(fields, doc.FieldStars)
	}
	if m.adduseful != nil {
		fields = append(fields, doc.FieldUseful)
	}
	if m.addfunny != nil {
		fields = append(fields, doc.FieldFunny)
	}
	if m.addcool != nil {
		fields = append(fields, doc.FieldCool)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *DocMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case doc.FieldStars:
		return m.AddedStars()
	case doc.FieldUseful:
		return m.AddedUseful()
	case doc.FieldFunny:
		return m.AddedFunny()
	case doc.FieldCool:
		return m.AddedCool()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *DocMutation) AddField(name string, value ent.Value) error {
	switch name {
	case doc.FieldStars:
		v, ok := value.(float32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddStars(v)
		return nil
	case doc.FieldUseful:
		v, ok := value.(int16)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUseful(v)
		return nil
	case doc.FieldFunny:
		v, ok := value.(int16)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddFunny(v)
		return nil
	case doc.FieldCool:
		v, ok := value.(int16)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCool(v)
		return nil
	}
	return fmt.Errorf("unknown Doc numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *DocMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *DocMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *DocMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Doc nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *DocMutation) ResetField(name string) error {
	switch name {
	case doc.FieldWordInts:
		m.ResetWordInts()
		return nil
	case doc.FieldInputDocId:
		m.ResetInputDocId()
		return nil
	case doc.FieldUserId:
		m.ResetUserId()
		return nil
	case doc.FieldBusinessId:
		m.ResetBusinessId()
		return nil
	case doc.FieldStars:
		m.ResetStars()
		return nil
	case doc.FieldUseful:
		m.ResetUseful()
		return nil
	case doc.FieldFunny:
		m.ResetFunny()
		return nil
	case doc.FieldCool:
		m.ResetCool()
		return nil
	case doc.FieldText:
		m.ResetText()
		return nil
	case doc.FieldDate:
		m.ResetDate()
		return nil
	}
	return fmt.Errorf("unknown Doc field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *DocMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *DocMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *DocMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *DocMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *DocMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *DocMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *DocMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Doc unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *DocMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Doc edge %s", name)
}
