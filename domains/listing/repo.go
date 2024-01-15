package listing

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type I18nDetail struct {
	Locale string
	Slug   string
}

type Repository interface {
	Create(ctx context.Context, entity *Entity) (*Entity, *i18np.Error)
	Update(ctx context.Context, entity *Entity) *i18np.Error
	Delete(ctx context.Context, listingUUID string) *i18np.Error
	Restore(ctx context.Context, listingUUID string) *i18np.Error
	Disable(ctx context.Context, listingUUID string) *i18np.Error
	Enable(ctx context.Context, listingUUID string) *i18np.Error
	MarkValid(ctx context.Context, listingUUID string) *i18np.Error
	MarkInvalid(ctx context.Context, listingUUID string) *i18np.Error
	ReOrder(ctx context.Context, listingUUID string, order int) *i18np.Error
	View(ctx context.Context, detail I18nDetail) (*Entity, *i18np.Error)
	GetByUUID(ctx context.Context, listingUUID string) (*Entity, bool, *i18np.Error)
	AdminView(ctx context.Context, listingUUID string) (*Entity, *i18np.Error)
	Filter(ctx context.Context, filter FilterEntity, listConfig list.Config) (*list.Result[*Entity], *i18np.Error)
	FilterByBusiness(ctx context.Context, businessNickName string, filter FilterEntity, listConfig list.Config) (*list.Result[*Entity], *i18np.Error)
	ListMy(ctx context.Context, businessUUID string, listConfig list.Config) (*list.Result[*Entity], *i18np.Error)
}

type repo struct {
	factory    Factory
	collection *mongo.Collection
	helper     mongo2.Helper[*Entity, *Entity]
}

func NewRepo(collection *mongo.Collection, factory Factory) Repository {
	return &repo{
		factory:    factory,
		collection: collection,
		helper:     mongo2.NewHelper[*Entity, *Entity](collection, createEntity),
	}
}

func createEntity() **Entity {
	return new(*Entity)
}

func (r *repo) Create(ctx context.Context, e *Entity) (*Entity, *i18np.Error) {
	res, err := r.collection.InsertOne(ctx, e)
	if err != nil {
		return nil, r.factory.Errors.Failed("create")
	}
	e.UUID = res.InsertedID.(primitive.ObjectID).Hex()
	return e, nil
}

func (r *repo) Update(ctx context.Context, e *Entity) *i18np.Error {
	id, err := mongo2.TransformId(e.UUID)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.CategoryUUIDs: e.CategoryUUIDs,
			fields.Images:        e.Images,
			fields.Meta:          e.Meta,
			fields.Features:      e.Features,
			fields.Prices:        e.Prices,
			fields.Location:      e.Location,
			fields.Boosts:        e.Boosts,
			fields.Validation:    e.Validation,
			fields.Order:         e.Order,
			fields.IsValid:       e.IsValid,
			fields.Currency:      e.Currency,
			fields.UpdatedAt:     time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Delete(ctx context.Context, listingUUID string) *i18np.Error {
	id, err := mongo2.TransformId(listingUUID)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
		fields.IsDeleted: bson.M{
			"$ne": true,
		},
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsDeleted: true,
			fields.IsActive:  false,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Restore(ctx context.Context, listingUUID string) *i18np.Error {
	id, err := mongo2.TransformId(listingUUID)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID:      id,
		fields.IsDeleted: true,
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsDeleted: false,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Disable(ctx context.Context, listingUUID string) *i18np.Error {
	id, err := mongo2.TransformId(listingUUID)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID:     id,
		fields.IsActive: true,
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsActive:  false,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) Enable(ctx context.Context, listingUUID string) *i18np.Error {
	id, err := mongo2.TransformId(listingUUID)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
		fields.IsActive: bson.M{
			"$ne": true,
		},
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsActive:  true,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) MarkValid(ctx context.Context, listingUUID string) *i18np.Error {
	id, err := mongo2.TransformId(listingUUID)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
		fields.IsValid: bson.M{
			"$ne": true,
		},
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsValid:   true,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) MarkInvalid(ctx context.Context, listingUUID string) *i18np.Error {
	id, err := mongo2.TransformId(listingUUID)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID:    id,
		fields.IsValid: true,
	}
	update := bson.M{
		"$set": bson.M{
			fields.IsValid:   false,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) ReOrder(ctx context.Context, listingUUID string, order int) *i18np.Error {
	id, err := mongo2.TransformId(listingUUID)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	update := bson.M{
		"$set": bson.M{
			fields.Order:     order,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, update)
}

func (r *repo) View(ctx context.Context, detail I18nDetail) (*Entity, *i18np.Error) {
	filter := bson.M{
		metaField(detail.Locale, metaFields.Slug): detail.Slug,
		fields.IsDeleted: bson.M{
			"$ne": true,
		},
		fields.IsActive: true,
		fields.IsValid:  true,
	}
	e, exist, err := r.helper.GetFilter(ctx, filter, r.viewOptions())
	if err != nil {
		return nil, r.factory.Errors.Failed("get")
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return *e, nil
}

func (r *repo) GetByUUID(ctx context.Context, uuid string) (*Entity, bool, *i18np.Error) {
	id, _err := mongo2.TransformId(uuid)
	if _err != nil {
		return nil, false, r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
		fields.IsDeleted: bson.M{
			"$ne": true,
		},
		fields.IsActive: true,
		fields.IsValid:  true,
	}
	e, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, false, r.factory.Errors.Failed("get")
	}
	if !exist {
		return nil, false, nil
	}
	return *e, true, nil
}

func (r *repo) AdminView(ctx context.Context, listingUUID string) (*Entity, *i18np.Error) {
	id, _err := mongo2.TransformId(listingUUID)
	if _err != nil {
		return nil, r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	e, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, r.factory.Errors.Failed("get")
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return *e, nil
}

func (r *repo) FilterByBusiness(ctx context.Context, businessNickName string, filter FilterEntity, listConfig list.Config) (*list.Result[*Entity], *i18np.Error) {
	businessFilter := r.filterToBson(filter, businessNickName)
	l, err := r.helper.GetListFilter(ctx, businessFilter, r.sort(r.filterOptions(listConfig), filter))
	if err != nil {
		return nil, err
	}
	filtered, _err := r.helper.GetFilterCount(ctx, businessFilter)
	if _err != nil {
		return nil, _err
	}
	total, _err := r.helper.GetFilterCount(ctx, r.businessFilter(businessNickName))
	if _err != nil {
		return nil, _err
	}
	return &list.Result[*Entity]{
		IsNext:        filtered > listConfig.Offset+listConfig.Limit,
		IsPrev:        listConfig.Offset > 0,
		FilteredTotal: filtered,
		Total:         total,
		Page:          listConfig.Offset/listConfig.Limit + 1,
		List:          l,
	}, nil
}

func (r *repo) Filter(ctx context.Context, filter FilterEntity, listConfig list.Config) (*list.Result[*Entity], *i18np.Error) {
	filters := r.filterToBson(filter, "")
	l, err := r.helper.GetListFilter(ctx, filters, r.sort(r.filterOptions(listConfig), filter))
	if err != nil {
		return nil, err
	}
	filtered, _err := r.helper.GetFilterCount(ctx, filters)
	if _err != nil {
		return nil, _err
	}
	total, _err := r.helper.GetFilterCount(ctx, r.baseFilter())
	if _err != nil {
		return nil, _err
	}
	return &list.Result[*Entity]{
		IsNext:        filtered > listConfig.Offset+listConfig.Limit,
		IsPrev:        listConfig.Offset > 0,
		FilteredTotal: filtered,
		Total:         total,
		Page:          listConfig.Offset/listConfig.Limit + 1,
		List:          l,
	}, nil
}

func (r *repo) ListMy(ctx context.Context, businessUUID string, listConfig list.Config) (*list.Result[*Entity], *i18np.Error) {
	filter := bson.M{
		businessField(businessFields.UUID): businessUUID,
	}
	l, err := r.helper.GetListFilter(ctx, filter, r.businessListOptions(listConfig))
	if err != nil {
		return nil, err
	}
	filtered, _err := r.helper.GetFilterCount(ctx, filter)
	if _err != nil {
		return nil, _err
	}
	total, _err := r.helper.GetFilterCount(ctx, filter)
	if _err != nil {
		return nil, _err
	}
	return &list.Result[*Entity]{
		IsNext:        filtered > listConfig.Offset+listConfig.Limit,
		IsPrev:        listConfig.Offset > 0,
		FilteredTotal: filtered,
		Total:         total,
		Page:          listConfig.Offset/listConfig.Limit + 1,
		List:          l,
	}, nil
}

func (r *repo) businessListOptions(listConfig list.Config) *options.FindOptions {
	opts := &options.FindOptions{}
	opts.SetProjection(bson.M{
		fields.UUID:      1,
		fields.Images:    1,
		fields.Meta:      1,
		fields.Location:  1,
		fields.Boosts:    1,
		fields.Order:     1,
		fields.IsDeleted: 1,
		fields.IsActive:  1,
		fields.Currency:  1,
		fields.IsValid:   1,
		fields.CreatedAt: 1,
	}).SetSort(bson.D{{Key: fields.Order, Value: 1}}).SetSkip(listConfig.Offset).SetLimit(listConfig.Limit)
	return opts
}

func (r *repo) adminListOptions(listConfig list.Config) *options.FindOptions {
	opts := &options.FindOptions{}
	opts.SetProjection(bson.M{
		fields.UUID:          1,
		fields.Business:      1,
		fields.Images:        1,
		fields.Meta:          1,
		fields.CategoryUUIDs: 1,
		fields.Features:      1,
		fields.Prices:        1,
		fields.Location:      1,
		fields.Boosts:        1,
		fields.Validation:    1,
		fields.Order:         1,
		fields.IsDeleted:     1,
		fields.Currency:      1,
		fields.IsActive:      1,
		fields.IsValid:       1,
		fields.CreatedAt:     1,
	}).SetSort(bson.D{{Key: fields.Order, Value: 1}}).SetSkip(listConfig.Offset).SetLimit(listConfig.Limit)
	return opts
}

func (r *repo) filterOptions(listConfig list.Config) *options.FindOptions {
	opts := &options.FindOptions{}
	opts.SetProjection(bson.M{
		fields.UUID:          1,
		fields.Business:      1,
		fields.Images:        1,
		fields.Meta:          1,
		fields.CategoryUUIDs: 1,
		fields.Validation:    1,
		fields.Features:      1,
		fields.Prices:        1,
		fields.Location:      1,
		fields.Currency:      1,
		fields.Boosts:        1,
	}).SetSkip(listConfig.Offset).SetLimit(listConfig.Limit)
	return opts
}

func (r *repo) viewOptions() *options.FindOneOptions {
	opts := &options.FindOneOptions{}
	opts.SetProjection(bson.M{
		fields.UUID:          1,
		fields.Business:      1,
		fields.Images:        1,
		fields.Meta:          1,
		fields.CategoryUUIDs: 1,
		fields.Validation:    1,
		fields.Features:      1,
		fields.Prices:        1,
		fields.Location:      1,
		fields.Currency:      1,
		fields.Boosts:        1,
		fields.UpdatedAt:     1,
		fields.CreatedAt:     1,
	})
	return opts
}
