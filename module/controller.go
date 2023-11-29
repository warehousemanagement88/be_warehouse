package module

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/badoux/checkmail"
	"github.com/warehousemanagement88/be_warehouse/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/argon2"
)

// var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect(MongoString, warehouse_db string) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv(MongoString)))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database(warehouse_db)
}

// crud
func GetAllDocs(db *mongo.Database, col string, docs interface{}) interface{} {
	collection := db.Collection(col)
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error GetAllDocs %s: %s", col, err)
	}
	err = cursor.All(context.TODO(), &docs)
	if err != nil {
		return err
	}
	return docs
}

func InsertOneDoc(db *mongo.Database, col string, doc interface{}) (insertedID primitive.ObjectID, err error) {
	result, err := db.Collection(col).InsertOne(context.Background(), doc)
	if err != nil {
		return insertedID, fmt.Errorf("kesalahan server : insert")
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func UpdateOneDoc(id primitive.ObjectID, db *mongo.Database, col string, doc interface{}) (err error) {
	filter := bson.M{"_id": id}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, bson.M{"$set": doc})
	if err != nil {
		return fmt.Errorf("error update: %v", err)
	}
	if result.ModifiedCount == 0 {
		err = fmt.Errorf("tidak ada data yang diubah")
		return
	}
	return nil
}

// update password
func UpdatePasswordUser(iduser primitive.ObjectID, db *mongo.Database, insertedDoc model.Password) error {
	dataUser, err := GetUserFromID(iduser, db)
	if err != nil {
		return err
	}
	salt, err := hex.DecodeString(dataUser.Salt)
	if err != nil {
		return fmt.Errorf("kesalahan server : salt")
	}
	hash := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
	if hex.EncodeToString(hash) != dataUser.Password {
		return fmt.Errorf("password lama salah")
	}
	if insertedDoc.Newpassword == "" || insertedDoc.Confirmpassword == "" {
		return fmt.Errorf("mohon untuk melengkapi data")
	}
	if insertedDoc.Confirmpassword != insertedDoc.Newpassword {
		return fmt.Errorf("konfirmasi password salah")
	}
	if strings.Contains(insertedDoc.Newpassword, " ") {
		return fmt.Errorf("password tidak boleh mengandung spasi")
	}
	if len(insertedDoc.Newpassword) < 8 {
		return fmt.Errorf("password terlalu pendek")
	}
	salt = make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		return fmt.Errorf("kesalahan server : salt")
	}
	hashedPassword := argon2.IDKey([]byte(insertedDoc.Newpassword), salt, 1, 64*1024, 4, 32)
	user := bson.M{
		"email":    dataUser.Email,
		"password": hex.EncodeToString(hashedPassword),
		"salt":     hex.EncodeToString(salt),
		"role":     dataUser.Role,
	}
	err = UpdateOneDoc(iduser, db, "user", user)
	if err != nil {
		return err
	}
	return nil
}

func DeleteOneDoc(_id primitive.ObjectID, db *mongo.Database, col string) error {
	collection := db.Collection(col)
	filter := bson.M{"_id": _id}
	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return fmt.Errorf("error deleting data for ID %s: %s", _id, err.Error())
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("data with ID %s not found", _id)
	}

	return nil
}

// signup
func SignUpStaff(db *mongo.Database, insertedDoc model.Staff) error {
	objectId := primitive.NewObjectID() 
	if insertedDoc.NamaLengkap == "" || insertedDoc.Jabatan == "" || insertedDoc.JenisKelamin == "" || insertedDoc.Akun.Email == "" || insertedDoc.Akun.Password == "" {
		return fmt.Errorf("mohon untuk melengkapi data")
	} 
	if err := checkmail.ValidateFormat(insertedDoc.Akun.Email); err != nil {
		return fmt.Errorf("email tidak valid")
	} 
	userExists, _ := GetUserFromEmail(insertedDoc.Akun.Email, db)
	if insertedDoc.Akun.Email == userExists.Email {
		return fmt.Errorf("email sudah terdaftar")
	} 
	if insertedDoc.Akun.Confirmpassword != insertedDoc.Akun.Password {
		return fmt.Errorf("konfirmasi password salah")
	}
	if strings.Contains(insertedDoc.Akun.Password, " ") {
		return fmt.Errorf("password tidak boleh mengandung spasi")
	}
	if len(insertedDoc.Akun.Password) < 8 {
		return fmt.Errorf("password terlalu pendek")
	} 
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return fmt.Errorf("kesalahan server : salt")
	}
	hashedPassword := argon2.IDKey([]byte(insertedDoc.Akun.Password), salt, 1, 64*1024, 4, 32)
	user := bson.M{
		"_id": objectId,
		"email": insertedDoc.Akun.Email,
		"password": hex.EncodeToString(hashedPassword),
		"salt": hex.EncodeToString(salt),
		"role": "staff",
	}
	staff := bson.M{
		"namalengkap": insertedDoc.NamaLengkap,
		"jabatan": insertedDoc.Jabatan,
		"jeniskelamin": insertedDoc.JenisKelamin,
		"akun": model.User {
			ID : objectId,
		},
	}
	_, err = InsertOneDoc(db, "user", user)
	if err != nil {
		return fmt.Errorf("kesalahan server")
	}
	_, err = InsertOneDoc(db, "staff", staff)
	if err != nil {
		return fmt.Errorf("kesalahan server")
	}
	return nil
}

// login
func LogIn(db *mongo.Database, insertedDoc model.User) (user model.User, err error) {
	if insertedDoc.Email == "" || insertedDoc.Password == "" {
		return user, fmt.Errorf("mohon untuk melengkapi data")
	} 
	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return user, fmt.Errorf("email tidak valid")
	} 
	existsDoc, err := GetUserFromEmail(insertedDoc.Email, db)
	if err != nil {
		return 
	}
	salt, err := hex.DecodeString(existsDoc.Salt)
	if err != nil {
		return user, fmt.Errorf("kesalahan server : salt")
	}
	hash := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
	if hex.EncodeToString(hash) != existsDoc.Password {
		return user, fmt.Errorf("password salah")
	}
	return existsDoc, nil
}

// func GetUserLogin(PASETOPUBLICKEYENV string, r *http.Request) (Payload, error) {
// 	tokenstring := r.Header.Get("Authorization")
// 	payload, err := Decode(os.Getenv(PASETOPUBLICKEYENV), tokenstring)
// 	if err != nil {
// 		return payload, err
// 	}
// 	return payload, nil
// }

//user
func UpdateUser(iduser primitive.ObjectID, db *mongo.Database, insertedDoc model.User) error {
	dataUser, err := GetUserFromID(iduser, db)
	if err != nil {
		return err
	}
	if insertedDoc.Email == "" || insertedDoc.Password == "" {
		return fmt.Errorf("mohon untuk melengkapi data")
	}
	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return fmt.Errorf("email tidak valid")
	}
	existsDoc, _ := GetUserFromEmail(insertedDoc.Email, db)
	if existsDoc.Email == insertedDoc.Email {
		return fmt.Errorf("email sudah terdaftar")
	}
	if insertedDoc.Confirmpassword != insertedDoc.Password {
		return fmt.Errorf("konfirmasi password salah")
	}
	if strings.Contains(insertedDoc.Password, " ") {
		return fmt.Errorf("password tidak boleh mengandung spasi")
	}
	if len(insertedDoc.Password) < 8 {
		return fmt.Errorf("password terlalu pendek")
	}
	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		return fmt.Errorf("kesalahan server : salt")
	}
	hashedPassword := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
	user := bson.M{
		"email": insertedDoc.Email,
		"password": hex.EncodeToString(hashedPassword),
		"salt": hex.EncodeToString(salt),
		"role": dataUser.Role,
	}
	err = UpdateOneDoc(iduser, db, "user", user)
	if err != nil {
		return err
	}
	return nil
}

func GetAllUser(db *mongo.Database) (user []model.User, err error) {
	collection := db.Collection("user")
	filter := bson.M{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return user, fmt.Errorf("error GetAllUser mongo: %s", err)
	}
	err = cursor.All(context.Background(), &user)
	if err != nil {
		return user, fmt.Errorf("error GetAllUser context: %s", err)
	}
	return user, nil
}

func GetUserFromID(_id primitive.ObjectID, db *mongo.Database) (doc model.User, err error) {
	collection := db.Collection("user")
	filter := bson.M{"_id": _id}
	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return doc, fmt.Errorf("no data found for ID %s", _id)
		}
		return doc, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return doc, nil
}

func GetUserFromEmail(email string, db *mongo.Database) (doc model.User, err error) {
	collection := db.Collection("user")
	filter := bson.M{"email": email}
	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return doc, fmt.Errorf("email tidak ditemukan")
		}
		return doc, fmt.Errorf("kesalahan server")
	}
	return doc, nil
}

// staff
func UpdateStaff(idparam, iduser primitive.ObjectID, db *mongo.Database, insertedDoc model.Staff) error {
	staff, err := GetStaffFromAkun(iduser, db)
	if err != nil {
		return err
	}
	if staff.ID != idparam {
		return fmt.Errorf("kamu bukan pemilik data ini")
	}
	if insertedDoc.NamaLengkap == "" || insertedDoc.Jabatan == "" || insertedDoc.JenisKelamin == ""  {
		return fmt.Errorf("mohon untuk melengkapi data")
	}
	stf := bson.M{
		"namalengkap": insertedDoc.NamaLengkap,
		"jabatan": insertedDoc.Jabatan,
		"jeniskelamin": insertedDoc.JenisKelamin,
		"akun": model.User {
			ID : staff.Akun.ID,
		},
	}
	err = UpdateOneDoc(idparam, db, "staff", stf)
	if err != nil {
		return err
	}
	return nil
}

func GetAllStaff(db *mongo.Database) (staff []model.Staff, err error) {
	collection := db.Collection("staff")
	filter := bson.M{}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return staff, fmt.Errorf("error GetAllStaff mongo: %s", err)
	}
	err = cursor.All(context.Background(), &staff)
	if err != nil {
		return staff, fmt.Errorf("error GetAllStaff context: %s", err)
	}
	return staff, nil
}

func GetStaffFromID(_id primitive.ObjectID, db *mongo.Database) (doc model.Staff, err error) {
	collection := db.Collection("staff")
	filter := bson.M{"_id": _id}
	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return doc, fmt.Errorf("no data found for ID %s", _id)
		}
		return doc, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	user, err := GetUserFromID(doc.Akun.ID, db)
	if err != nil {
		return doc, fmt.Errorf("kesalahan server")
	}
	akun := model.User{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}
	doc.Akun = akun
	return doc, nil
}

func GetStaffFromAkun(id_akun primitive.ObjectID, db *mongo.Database) (doc model.Staff, err error) {
	collection := db.Collection("staff")
	filter := bson.M{"akun._id": id_akun}
	err = collection.FindOne(context.TODO(), filter).Decode(&doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return doc, fmt.Errorf("staff tidak ditemukan")
		}
		return doc, fmt.Errorf("kesalahan server")
	}
	user, err := GetUserFromID(doc.Akun.ID, db)
	if err != nil {
		return doc, fmt.Errorf("kesalahan server")
	}
	akun := model.User{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}
	doc.Akun = akun
	return doc, nil
}

//email
func UpdateEmailUser(iduser primitive.ObjectID, db *mongo.Database, insertedDoc model.User) error {
	dataUser, err := GetUserFromID(iduser, db)
	if err != nil {
		return err
	}
	if insertedDoc.Email == "" {
		return fmt.Errorf("mohon untuk melengkapi data")
	}
	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return fmt.Errorf("email tidak valid")
	}
	existsDoc, _ := GetUserFromEmail(insertedDoc.Email, db)
	if existsDoc.Email == insertedDoc.Email {
		return fmt.Errorf("email sudah terdaftar")
	}
	user := bson.M{
		"email":    insertedDoc.Email,
		"password": dataUser.Password,
		"salt":     dataUser.Salt,
		"role":     dataUser.Role,
	}
	err = UpdateOneDoc(iduser, db, "user", user)
	if err != nil {
		return err
	}
	return nil
}

// Gudang A
// insert gudang a
func InsertGudangA(_id primitive.ObjectID, db *mongo.Database, insertedDoc model.GudangA) error {
	if insertedDoc.Brand == "" || insertedDoc.Name == "" || insertedDoc.Category == "" || insertedDoc.QTY == "" ||
	insertedDoc.SKU == "" || insertedDoc.SellingPrice == "" || insertedDoc.OriginalPrice == "" ||
	insertedDoc.Availability == "" || insertedDoc.Color == "" || insertedDoc.Breadcrumbs == "" {
		return fmt.Errorf("mohon untuk melengkapi data")
	}
	// gudanga := bson.M{
	// 	"brand": insertedDoc.Brand,
	// 	"name": insertedDoc.Name,
	// 	"category": insertedDoc.Category,
	// 	"qty": insertedDoc.QTY,
	// 	"sku": insertedDoc.SKU,
	// 	"sellingprice": insertedDoc.SellingPrice,
	// 	"originalprice": insertedDoc.OriginalPrice,
	// 	"breadcrumbs": insertedDoc.Breadcrumbs,
	// 	"date": insertedDoc.Date,
	// }
	_, err := InsertOneDoc(db, "gudanga", insertedDoc)
	if err != nil {
		fmt.Printf("InsertGudangA: %v\n", err)
	}
	return nil
}

// func InsertTodo(db *mongo.Database, col string, todo model.Todo) (insertedID primitive.ObjectID, err error) {
// 	todo.TimeStamp.CreatedAt = time.Now()
// 	todo.TimeStamp.UpdatedAt = time.Now()

// 	insertedID, err = InsertOneDoc(db, col, todo)
// 	if err != nil {
// 		fmt.Printf("InsertTodo: %v\n", err)
// 	}
// 	return insertedID, nil
// }

// update gudang a
func UpdateGudangA(idparam, iduser primitive.ObjectID, db *mongo.Database, insertedDoc model.GudangA) error {
	_, err := GetGudangAFromID(idparam, db)
	if err != nil {
		return err
	}
	if insertedDoc.Brand == "" || insertedDoc.Name == "" || insertedDoc.Category == "" || insertedDoc.QTY == "" ||
	insertedDoc.SKU == "" || insertedDoc.SellingPrice == "" || insertedDoc.OriginalPrice == "" ||
	insertedDoc.Availability == "" || insertedDoc.Color == "" || insertedDoc.Breadcrumbs == "" ||
	insertedDoc.Date.IsZero() {
		return fmt.Errorf("mohon untuk melengkapi data")
	}

	gudanga := bson.M{
		"brand": insertedDoc.Brand,
		"name": insertedDoc.Name,
		"category": insertedDoc.Category,
		"qty": insertedDoc.QTY,
		"sku": insertedDoc.SKU,
		"sellingprice": insertedDoc.SellingPrice,
		"originalprice": insertedDoc.OriginalPrice,
		"breadcrumbs": insertedDoc.Breadcrumbs,
		"date": insertedDoc.Date,
	}

	err = UpdateOneDoc(idparam, db, "gudanga", gudanga)
	if err != nil {
		return err
	}
	return nil
}

// delete gudang a
func DeleteGudangA(idparam, iduser primitive.ObjectID, db *mongo.Database) error {
	_, err := GetGudangAFromID(idparam, db)
	if err != nil {
		return err
	}
	err = DeleteOneDoc(idparam, db, "gudanga")
	if err != nil {
		return err
	}
	return nil
}

// func DeleteTodo(db *mongo.Database, col string, _id primitive.ObjectID) (status bool, err error) {
// 	cols := db.Collection(col)
// 	filter := bson.M{"_id": _id}
// 	result, err := cols.DeleteOne(context.Background(), filter)
// 	if err != nil {
// 		return false, err
// 	}
// 	if result.DeletedCount == 0 {
// 		err = fmt.Errorf("Data tidak berhasil dihapus")
// 		return false, err
// 	}
// 	return true, nil
// }

// get all gudang a
func GetAllGudangA(db *mongo.Database) (gudanga []model.GudangA, err error) {
	collection := db.Collection("gudanga")
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return gudanga, fmt.Errorf("error GetAllGudangA mongo: %s", err)
	}
	err = cursor.All(context.TODO(), &gudanga)
	if err != nil {
		return gudanga, fmt.Errorf("error GetAllGudangA context: %s", err)
	}
	return gudanga, nil
}

// Get GudangA FromID
func GetGudangAFromID(_id primitive.ObjectID, db *mongo.Database) (gudanga model.GudangA, err error) {
	collection := db.Collection("gudanga")
	filter := bson.M{"_id": _id}
	err = collection.FindOne(context.TODO(), filter).Decode(&gudanga)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return gudanga, fmt.Errorf("no data found for ID %s", _id)
		}
		return gudanga, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return gudanga, nil
}

// Gudang B
// insert gudang b
func InsertGudangB(_id primitive.ObjectID, db *mongo.Database, insertedDoc model.GudangB) error {
	if insertedDoc.Brand == "" || insertedDoc.Name == "" || insertedDoc.Category == "" || insertedDoc.QTY == "" ||
	insertedDoc.SKU == "" || insertedDoc.SellingPrice == "" || insertedDoc.OriginalPrice == "" ||
	insertedDoc.Availability == "" || insertedDoc.Color == "" || insertedDoc.Breadcrumbs == "" ||
	insertedDoc.Date.IsZero() {
		return fmt.Errorf("mohon untuk melengkapi data")
	}
	gudangb := bson.M{
		"brand": insertedDoc.Brand,
		"name": insertedDoc.Name,
		"category": insertedDoc.Category,
		"qty": insertedDoc.QTY,
		"sku": insertedDoc.SKU,
		"sellingprice": insertedDoc.SellingPrice,
		"originalprice": insertedDoc.OriginalPrice,
		"breadcrumbs": insertedDoc.Breadcrumbs,
		"date": insertedDoc.Date,
	}
	_, err := InsertOneDoc(db, "gudangb", gudangb)
	if err != nil {
		return err
	}
	return nil
}

// update gudang b
func UpdateGudangB(idparam, iduser primitive.ObjectID, db *mongo.Database, insertedDoc model.GudangB) error {
	_, err := GetGudangBFromID(idparam, db)
	if err != nil {
		return err
	}
	if insertedDoc.Brand == "" || insertedDoc.Name == "" || insertedDoc.Category == "" || insertedDoc.QTY == "" ||
	insertedDoc.SKU == "" || insertedDoc.SellingPrice == "" || insertedDoc.OriginalPrice == "" ||
	insertedDoc.Availability == "" || insertedDoc.Color == "" || insertedDoc.Breadcrumbs == "" ||
	insertedDoc.Date.IsZero() {
		return fmt.Errorf("mohon untuk melengkapi data")
	}

	gudangb := bson.M{
		"brand": insertedDoc.Brand,
		"name": insertedDoc.Name,
		"category": insertedDoc.Category,
		"qty": insertedDoc.QTY,
		"sku": insertedDoc.SKU,
		"sellingprice": insertedDoc.SellingPrice,
		"originalprice": insertedDoc.OriginalPrice,
		"breadcrumbs": insertedDoc.Breadcrumbs,
		"date": insertedDoc.Date,
	}

	err = UpdateOneDoc(idparam, db, "gudangb", gudangb)
	if err != nil {
		return err
	}
	return nil
}

// delete gudang b
func DeleteGudangB(idparam, iduser primitive.ObjectID, db *mongo.Database) error {
	_, err := GetGudangBFromID(idparam, db)
	if err != nil {
		return err
	}
	err = DeleteOneDoc(idparam, db, "gudangb")
	if err != nil {
		return err
	}
	return nil
}

// get all gudang b
func GetAllGudangB(db *mongo.Database) (gudangb []model.GudangB, err error) {
	collection := db.Collection("gudangb")
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return gudangb, fmt.Errorf("error GetAllGudangB mongo: %s", err)
	}
	err = cursor.All(context.TODO(), &gudangb)
	if err != nil {
		return gudangb, fmt.Errorf("error GetAllGudangB context: %s", err)
	}
	return gudangb, nil
}

func GetGudangBFromID(_id primitive.ObjectID, db *mongo.Database) (gudangb model.GudangB, err error) {
	collection := db.Collection("gudangb")
	filter := bson.M{"_id": _id}
	err = collection.FindOne(context.TODO(), filter).Decode(&gudangb)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return gudangb, fmt.Errorf("no data found for ID %s", _id)
		}
		return gudangb, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return gudangb, nil
}


// Gudang C
// insert gudang c
func InsertGudangC(_id primitive.ObjectID, db *mongo.Database, insertedDoc model.GudangC) error {
	if insertedDoc.Brand == "" || insertedDoc.Name == "" || insertedDoc.Category == "" || insertedDoc.QTY == "" ||
	insertedDoc.SKU == "" || insertedDoc.SellingPrice == "" || insertedDoc.OriginalPrice == "" ||
	insertedDoc.Availability == "" || insertedDoc.Color == "" || insertedDoc.Breadcrumbs == "" ||
	insertedDoc.Date.IsZero() {
		return fmt.Errorf("mohon untuk melengkapi data")
	}
	gudangc := bson.M{
		"brand": insertedDoc.Brand,
		"name": insertedDoc.Name,
		"category": insertedDoc.Category,
		"qty": insertedDoc.QTY,
		"sku": insertedDoc.SKU,
		"sellingprice": insertedDoc.SellingPrice,
		"originalprice": insertedDoc.OriginalPrice,
		"breadcrumbs": insertedDoc.Breadcrumbs,
		"date": insertedDoc.Date,
	}
	_, err := InsertOneDoc(db, "gudangc", gudangc)
	if err != nil {
		return err
	}
	return nil
}

// update gudang b
func UpdateGudangC(idparam, iduser primitive.ObjectID, db *mongo.Database, insertedDoc model.GudangC) error {
	_, err := GetGudangCFromID(idparam, db)
	if err != nil {
		return err
	}
	if insertedDoc.Brand == "" || insertedDoc.Name == "" || insertedDoc.Category == "" || insertedDoc.QTY == "" ||
	insertedDoc.SKU == "" || insertedDoc.SellingPrice == "" || insertedDoc.OriginalPrice == "" ||
	insertedDoc.Availability == "" || insertedDoc.Color == "" || insertedDoc.Breadcrumbs == "" ||
	insertedDoc.Date.IsZero() {
		return fmt.Errorf("mohon untuk melengkapi data")
	}

	gudangc := bson.M{
		"brand": insertedDoc.Brand,
		"name": insertedDoc.Name,
		"category": insertedDoc.Category,
		"qty": insertedDoc.QTY,
		"sku": insertedDoc.SKU,
		"sellingprice": insertedDoc.SellingPrice,
		"originalprice": insertedDoc.OriginalPrice,
		"breadcrumbs": insertedDoc.Breadcrumbs,
		"date": insertedDoc.Date,
	}

	err = UpdateOneDoc(idparam, db, "gudangc", gudangc)
	if err != nil {
		return err
	}
	return nil
}

// delete gudang c
func DeleteGudangC(idparam, iduser primitive.ObjectID, db *mongo.Database) error {
	_, err := GetGudangCFromID(idparam, db)
	if err != nil {
		return err
	}
	err = DeleteOneDoc(idparam, db, "gudangc")
	if err != nil {
		return err
	}
	return nil
}

// get all gudang c
func GetAllGudangC(db *mongo.Database) (gudangc []model.GudangC, err error) {
	collection := db.Collection("gudangc")
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return gudangc, fmt.Errorf("error GetAllGudangC mongo: %s", err)
	}
	err = cursor.All(context.TODO(), &gudangc)
	if err != nil {
		return gudangc, fmt.Errorf("error GetAllGudangC context: %s", err)
	}
	return gudangc, nil
}

func GetGudangCFromID(_id primitive.ObjectID, db *mongo.Database) (gudangc model.GudangC, err error) {
	collection := db.Collection("gudangc")
	filter := bson.M{"_id": _id}
	err = collection.FindOne(context.TODO(), filter).Decode(&gudangc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return gudangc, fmt.Errorf("no data found for ID %s", _id)
		}
		return gudangc, fmt.Errorf("error retrieving data for ID %s: %s", _id, err.Error())
	}
	return gudangc, nil
}


// get user login
func GetUserLogin(PASETOPUBLICKEYENV string, r *http.Request) (model.Payload, error) {
	tokenstring := r.Header.Get("Authorization")
	payload, err := Decode(os.Getenv(PASETOPUBLICKEYENV), tokenstring)
	if err != nil {
		return payload, err
	}
	return payload, nil
}

// get id
func GetID(r *http.Request) string {
    return r.URL.Query().Get("id")
}

// return struct
func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}