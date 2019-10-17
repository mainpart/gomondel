package templates

func (t *Template) parentMethod() string {
	return `func ({{.ModelSign}} {{.ModelName}}) {{.ParentModelName}}() ({{.ParentModelSign}} *{{.ParentModelName}}, err error) {
	{{.ParentModelSign}}, err = {{.ParentModelName}}{}.FindById({{.ModelSign}}.{{.ParentModelName}}Id)
	return
}`
}

func (t *Template) parentField() string {
	return `{{.ParentModelName}}Id primitive.ObjectID ` + "`" + `bson:"{{.ParentModelNameSmall}}_id"` + "`"
}

func (t *Template) model() string {
	return `package models
import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type {{.ModelName}} struct {
	Id primitive.ObjectID ` + "`" + `bson:"_id,omitempty"` + "`" + `
	{{.ParentField}}
}

func ({{.ModelSign}} {{.ModelName}}) New() (n{{.ModelSign}} *{{.ModelName}}, err error) {
	n{{.ModelSign}} = &{{.ModelName}}{Id: primitive.NewObjectID()}
	_, err = DB.Collection("{{.TableName}}").InsertOne(NewContext(), &n{{.ModelSign}})
	return
}

func ({{.ModelSign}} *{{.ModelName}}) Create() (err error) {
	_, err = DB.Collection("{{.TableName}}").InsertOne(NewContext(), &{{.ModelSign}})
	return
}

func ({{.ModelSign}} {{.ModelName}}) CreateCustom(query bson.M) (err error) {
	_, err = DB.Collection("{{.TableName}}").InsertOne(NewContext(), query)
	return
}

func ({{.ModelSign}} {{.ModelName}}) FindById(id primitive.ObjectID) (f{{.ModelSign}} *{{.ModelName}}, err error) {
	err = DB.Collection("{{.TableName}}").FindOne(NewContext(), bson.M{"_id": id}).Decode(&f{{.ModelSign}})
	return
}

func ({{.ModelSign}} {{.ModelName}}) FindOne() (n{{.ModelSign}} *{{.ModelName}}, err error) {
	n{{.ModelSign}}, err = {{.ModelSign}}.FindOneWithFilter(bson.M{})
	return
}

func ({{.ModelSign}} {{.ModelName}}) FindOneWithFilter(filter bson.M) (n{{.ModelSign}} *{{.ModelName}}, err error) {
	err = DB.Collection("{{.TableName}}").FindOne(NewContext(), filter).Decode(&n{{.ModelSign}})
	return
}

func ({{.ModelSign}} {{.ModelName}}) FindWithFilter(filter bson.M, findOptions ...*options.FindOptions) ({{.ModelSign}}s []{{.ModelName}}, err error) {
	ctx := NewContext()
	var cur *mongo.Cursor
	cur, err = DB.Collection("{{.TableName}}").Find(ctx, filter, findOptions...)
	if err != nil {
		return
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result {{.ModelName}}
		err = cur.Decode(&result)
		if err != nil {
			return
		}
		{{.ModelSign}}s = append({{.ModelSign}}s, result)
	}
	if err = cur.Err(); err != nil {
		return
	}
	return
}

func ({{.ModelSign}} {{.ModelName}}) Find() ({{.ModelSign}}s []{{.ModelName}}, err error) {
	{{.ModelSign}}s, err = {{.ModelSign}}.FindWithFilter(bson.M{})
	return
}

func ({{.ModelSign}} {{.ModelName}}) Count() (count int64, err error) {
	count, err = {{.ModelSign}}.CountWithFilter(bson.M{})
	return
}

func ({{.ModelSign}} {{.ModelName}}) CountWithFilter(filter bson.M) (count int64, err error) {
	count, err = DB.Collection("{{.TableName}}").CountDocuments(NewContext(), bson.M{})
	return
}

func ({{.ModelSign}} {{.ModelName}}) DistinctWithFilter(field string, filter bson.M, options ...*options.DistinctOptions) (ids []primitive.ObjectID, err error) {
	var result []interface{}
	result, err = DB.Collection("{{.TableName}}").Distinct(NewContext(), field, filter, options...)
	for _, v := range result {
		ids = append(ids, v.(primitive.ObjectID))
	}
	return
}

func ({{.ModelSign}} {{.ModelName}}) Distinct(field string, filter bson.M, options ...*options.DistinctOptions) (ids []primitive.ObjectID, err error) {
	ids, err = {{.ModelSign}}.DistinctWithFilter(field, bson.M{}, options...)
	return
}

func ({{.ModelSign}} *{{.ModelName}}) Save() (err error) {
	_, err = DB.Collection("{{.TableName}}").UpdateOne(NewContext(), bson.M{"_id": {{.ModelSign}}.Id}, bson.M{"$set": &{{.ModelSign}}})
	return
}

func ({{.ModelSign}} *{{.ModelName}}) SaveCustom(query bson.M) (err error) {
	_, err = DB.Collection("{{.TableName}}").UpdateOne(NewContext(), bson.M{"_id": {{.ModelSign}}.Id}, query)
	return
}

func ({{.ModelSign}} *{{.ModelName}}) Remove() (err error) {
	_, err = DB.Collection("{{.TableName}}").DeleteOne(NewContext(), bson.M{"_id": {{.ModelSign}}.Id})
	return
}

{{.ParentMethod}}
`
}
