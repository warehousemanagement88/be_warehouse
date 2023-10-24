package warehousemanagement88

import (
	"fmt"
	"testing"

	"github.com/warehousemanagement88/be_warehouse/model"
	"github.com/warehousemanagement88/be_warehouse/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var db = module.MongoConnect()

func TestGetUserFromEmail(t *testing.T) {
	email := "admin@gmail.com"
	hasil, err := module.GetUserFromEmail(email, db, "user")
	if err != nil {
		t.Errorf("Error TestGetUserFromEmail: %v", err)
	} else {
		fmt.Println(hasil)
	}
}

func TestInsertOneItem(t *testing.T) {
	var doc model.Item
	doc.Name = "Seblak"
	doc.Quantity = "10"
	doc.Category = "Makanan"
	if doc.Name == "" || doc.Quantity == "" || doc.Category == "" {
		t.Errorf("mohon untuk melengkapi data")
	} else {
		insertedID, err := module.InsertOneDoc(db, "item", doc)
		if err != nil {
			t.Errorf("Error inserting document: %v", err)
			fmt.Println("Data tidak berhasil disimpan")
		} else {
			fmt.Println("Data berhasil disimpan dengan id :", insertedID.Hex())
		}
	}
}

func TestGetAllDoc(t *testing.T) {
	var docs []model.Item
	hasil := module.GetAllDocs(db, "item", docs)
	fmt.Println(hasil)
}

func TestUpdateOneDoc(t *testing.T) {
	var docs model.User
	id := "649063d3ad72e074286c61e8"
	objectId, _ := primitive.ObjectIDFromHex(id)
	docs.FirstName = "Mickey"
	docs.LastName = "Mouse"
	docs.Email = "mickeymouse@gmail.com"
	docs.Password = "123456"
	if docs.FirstName == "" || docs.LastName == "" || docs.Email == "" || docs.Password == "" {
		t.Errorf("mohon untuk melengkapi data")
	} else {
		err := module.UpdateOneDoc(db, "user", objectId, docs)
		if err != nil {
			t.Errorf("Error inserting document: %v", err)
			fmt.Println("Data tidak berhasil diupdate")
		} else {
			fmt.Println("Data berhasil diupdate")
		}
	}
}

func TestGetItemFromID(t *testing.T) {
	id := "64d0b1104255ba95ba588512"
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Fatalf("error converting id to objectID: %v", err)
	}
	doc, err := module.GetItemFromID(objectId)
	if err != nil {
		t.Fatalf("error calling GetDocFromID: %v", err)
	}
	fmt.Println(doc)
}

func TestSignUp(t *testing.T) {
	var doc model.User
	doc.FirstName = "Agita"
	doc.LastName = "Nurfadillah"
	doc.Email = "agitanurfadillah1004@gmail.com"
	doc.Password = "987654321"
	doc.Confirmpassword = "987654321"
	insertedID, err := module.SignUp(db, "user", doc)
	if err != nil {
		t.Errorf("Error inserting document: %v", err)
	} else {
		fmt.Println("Data berhasil disimpan dengan id :", insertedID.Hex())
	}
}

func TestLogIn(t *testing.T) {
	var doc model.User
	doc.Email = "agitanurfadillah1003@gmail.com"
	doc.Password = "654321"
	user, err := module.LogIn(db, "user", doc)
	if err != nil {
		t.Errorf("Error getting document: %v", err)
	} else {
		fmt.Println("Welcome :", user)
	}
}

func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := module.GenerateKey()
	fmt.Println("ini private key :", privateKey)
	fmt.Println("ini public key :", publicKey)
	hasil, err := module.Encode("agitanurfadillah1003@gmail.com", privateKey)
	fmt.Println("ini hasil :", hasil, err)
}
