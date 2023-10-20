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
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/warehousemanagement88/be_warehouse/model"
	"github.com/badoux/checkmail"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/argon2"
)

var MongoString string = os.Getenv("MONGOSTRING")

func MongoConnect() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MongoString))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
	}
	return client.Database("warehouse_db")
}

func GetUserFromEmail(email string, db *mongo.Database, col string) (result model.User, err error) {
	collection := db.Collection(col)
	filter := bson.M{"email": email}
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, fmt.Errorf("email tidak ditemukan")
		}
		return result, fmt.Errorf("kesalahan server")
	}
	return result, nil
}

func GetItemFromID(_id primitive.ObjectID) (doc model.Item, err error) {
	collection := MongoConnect().Collection("item")
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

func GetAllDocs(db *mongo.Database, col string, docs interface{}) interface{} {
	collection := db.Collection(col)
	filter := bson.M{}
	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error GetAllDocs in colecction", col, ":", err)
	}
	err = cursor.All(context.TODO(), &docs)
	if err != nil {
		fmt.Println(err)
	}
	return docs
}

func InsertOneDoc(db *mongo.Database, col string, doc interface{}) (insertedID primitive.ObjectID, err error) {
	result, err := db.Collection(col).InsertOne(context.Background(), doc)
	if err != nil {
		return insertedID, fmt.Errorf("kesalahan server")
	}
	insertedID = result.InsertedID.(primitive.ObjectID)
	return insertedID, nil
}

func UpdateOneDoc(db *mongo.Database, col string, id primitive.ObjectID, doc interface{}) (err error) {
	filter := bson.M{"_id": id}
	result, err := db.Collection(col).UpdateOne(context.Background(), filter, bson.M{"$set": doc})
	if err != nil {
		fmt.Printf("UpdatePresensi: %v\n", err)
		return 
	}
	if result.ModifiedCount == 0 {
		err = errors.New("no data has been changed with the specified id")
		return
	}
	return nil
}

func DeleteDocsByID(_id primitive.ObjectID, db *mongo.Database, col string) error {
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

func SignUp(db *mongo.Database, col string, insertedDoc model.User) (insertedID primitive.ObjectID, err error) {
	if insertedDoc.FirstName == "" || insertedDoc.LastName == "" || insertedDoc.Email == "" || insertedDoc.Password == "" {
		return insertedID, fmt.Errorf("mohon untuk melengkapi data")
	} 
	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return insertedID, fmt.Errorf("email tidak valid")
	} 
	userExists, _ := GetUserFromEmail(insertedDoc.Email, db, col)
	if insertedDoc.Email == userExists.Email {
		return insertedID, fmt.Errorf("email sudah terdaftar")
	} 
	if insertedDoc.Confirmpassword != insertedDoc.Password {
		return insertedID, fmt.Errorf("konfirmasi password salah")
	}
	if strings.Contains(insertedDoc.Password, " ") {
		return insertedID, fmt.Errorf("password tidak boleh mengandung spasi")
	}
	if len(insertedDoc.Password) < 8 {
		return insertedID, fmt.Errorf("password terlalu pendek")
	}
	salt := make([]byte, 16)
	_, err = rand.Read(salt)
	if err != nil {
		return insertedID, fmt.Errorf("kesalahan server")
	}
	hashedPassword := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
	insertedDoc.Password = hex.EncodeToString(hashedPassword)
	insertedDoc.Salt = hex.EncodeToString(salt)
	insertedDoc.Confirmpassword = ""
	return InsertOneDoc(db, col, insertedDoc)
}

func LogIn(db *mongo.Database, col string, insertedDoc model.User) (email string, err error) {
	if insertedDoc.Email == "" || insertedDoc.Password == "" {
		return email, fmt.Errorf("mohon untuk melengkapi data")
	} 
	if err = checkmail.ValidateFormat(insertedDoc.Email); err != nil {
		return email, fmt.Errorf("email tidak valid")
	} 
	existsDoc, err := GetUserFromEmail(insertedDoc.Email, db, col)
	if err != nil {
		return 
	}
	salt, err := hex.DecodeString(existsDoc.Salt)
	if err != nil {
		return email, fmt.Errorf("kesalahan server")
	}
	hash := argon2.IDKey([]byte(insertedDoc.Password), salt, 1, 64*1024, 4, 32)
	if hex.EncodeToString(hash) != existsDoc.Password {
		return email, fmt.Errorf("password salah")
	}
	return existsDoc.Email, nil
}

func GCFPostHandler(PASETOPRIVATEKEYENV, collectionname string, r *http.Request) string {
	var Response model.Credential
	Response.Status = false
	var datauser model.User
	err := json.NewDecoder(r.Body).Decode(&datauser)
	if err != nil {
		Response.Message = "error parsing application/json: " + err.Error()
		return GCFReturnStruct(Response)
	}
	email, err := LogIn(MongoConnect(), collectionname, datauser)
	if err != nil {
		Response.Message = "error LogIn: " + err.Error()
		return GCFReturnStruct(Response)
	}
	Response.Status = true
	tokenstring, err := Encode(email, os.Getenv(PASETOPRIVATEKEYENV))
	if err != nil {
		Response.Message = "Gagal Encode Token : " + err.Error()
	} else {
		Response.Message = "Selamat Datang"
		Response.Token = tokenstring
	}		
	return GCFReturnStruct(Response)
}

func Encode(id string, privateKey string) (string, error) {
	token := paseto.NewToken()
	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(2 * time.Hour))
	token.SetString("id", id)
	secretKey, err := paseto.NewV4AsymmetricSecretKeyFromHex(privateKey)
	return token.V4Sign(secretKey, nil), err
}

func GenerateKey() (privateKey, publicKey string) {
	secretKey := paseto.NewV4AsymmetricSecretKey() // don't share this!!!
	publicKey = secretKey.Public().ExportHex()     // DO share this one
	privateKey = secretKey.ExportHex()
	return privateKey, publicKey
}

func GCFReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}