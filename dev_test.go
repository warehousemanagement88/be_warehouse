package Fatwa_1214038

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/warehousemanagement88/be_warehouse/model"
	"github.com/warehousemanagement88/be_warehouse/module"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/argon2"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

var db = module.MongoConnect("MONGOSTRING", "warehouse_db")

func TestGetUserFromEmail(t *testing.T) {
	email := "admin@gmail.com"
	hasil, err := module.GetUserFromEmail(email, db)
	if err != nil {
		t.Errorf("Error TestGetUserFromEmail: %v", err)
	} else {
		fmt.Println(hasil)
	}
}

func TestInsertOneGudangA(t *testing.T) {
	var doc model.GudangA
	doc.Brand = "Adidas"
	doc.Name = "Beach Shorts"
	doc.Category = "Clothing"
	doc.QTY = "100"
	doc.SKU = "FJ5089"
	doc.SellingPrice = "675.000"
	doc.OriginalPrice = "592.000"
	doc.Availability = "InStock"
	doc.Color = "Black"
	doc.Breadcrumbs = "Women/Clothing"
	doc.Date = time.Now() // Inisialisasi dengan waktu saat ini

	if doc.Brand == "" || doc.Name == "" || doc.Category == "" || doc.QTY == "" ||
		doc.SKU == "" || doc.SellingPrice == "" || doc.OriginalPrice == "" ||
		doc.Availability == "" || doc.Color == "" || doc.Breadcrumbs == "" || doc.Date.IsZero() {
		t.Errorf("Mohon untuk melengkapi data")
	} else {
		insertedID, err := module.InsertOneDoc(db, "gudanga", doc)
		if err != nil {
			t.Errorf("Error inserting document: %v", err)
			fmt.Println("Data tidak berhasil disimpan")
		} else {
			fmt.Println("Data berhasil disimpan dengan id :", insertedID.Hex())
		}
	}
}

func TestInsertOneGudangB(t *testing.T) {
	var doc model.GudangB
	doc.Brand = "Adidas"
	doc.Name = "Beach Shorts"
	doc.Category = "Clothing"
	doc.QTY = "100"
	doc.SKU = "FJ5089"
	doc.SellingPrice = "675.000"
	doc.OriginalPrice = "592.000"
	doc.Availability = "InStock"
	doc.Color = "Black"
	doc.Breadcrumbs = "Women/Clothing"
	doc.Date = time.Now() // Inisialisasi dengan waktu saat ini

	if doc.Brand == "" || doc.Name == "" || doc.Category == "" || doc.QTY == "" ||
		doc.SKU == "" || doc.SellingPrice == "" || doc.OriginalPrice == "" ||
		doc.Availability == "" || doc.Color == "" || doc.Breadcrumbs == "" || doc.Date.IsZero() {
		t.Errorf("Mohon untuk melengkapi data")
	} else {
		insertedID, err := module.InsertOneDoc(db, "gudangb", doc)
		if err != nil {
			t.Errorf("Error inserting document: %v", err)
			fmt.Println("Data tidak berhasil disimpan")
		} else {
			fmt.Println("Data berhasil disimpan dengan id :", insertedID.Hex())
		}
	}
}

func TestInsertOneGudangC(t *testing.T) {
	var doc model.GudangC
	doc.Brand = "Adidas"
	doc.Name = "Beach Shorts"
	doc.Category = "Clothing"
	doc.QTY = "100"
	doc.SKU = "FJ5089"
	doc.SellingPrice = "675.000"
	doc.OriginalPrice = "592.000"
	doc.Availability = "InStock"
	doc.Color = "Black"
	doc.Breadcrumbs = "Women/Clothing"
	doc.Date = time.Now() // Inisialisasi dengan waktu saat ini

	if doc.Brand == "" || doc.Name == "" || doc.Category == "" || doc.QTY == "" ||
		doc.SKU == "" || doc.SellingPrice == "" || doc.OriginalPrice == "" ||
		doc.Availability == "" || doc.Color == "" || doc.Breadcrumbs == "" || doc.Date.IsZero() {
		t.Errorf("Mohon untuk melengkapi data")
	} else {
		insertedID, err := module.InsertOneDoc(db, "gudangc", doc)
		if err != nil {
			t.Errorf("Error inserting document: %v", err)
			fmt.Println("Data tidak berhasil disimpan")
		} else {
			fmt.Println("Data berhasil disimpan dengan id :", insertedID.Hex())
		}
	}
}


type Userr struct {
	ID           	primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email  			string             `bson:"email,omitempty" json:"email,omitempty"`
	Role     		string			   `bson:"role,omitempty" json:"role,omitempty"`
}

func TestGetAllDoc(t *testing.T) {
	hasil := module.GetAllDocs(db, "user", []Userr{})
	fmt.Println(hasil)
}

// func TestUpdateOneDoc(t *testing.T) {
//  	var docs model.User
// 	id := "649063d3ad72e074286c61e8"
// 	objectId, _ := primitive.ObjectIDFromHex(id)
// 	docs.FirstName = "Aufah"
// 	docs.LastName = "Auliana"
// 	docs.Email = "aufa@gmail.com"
// 	docs.Password = "123456"
// 	if docs.FirstName == "" || docs.LastName == "" || docs.Email == "" || docs.Password == "" {
// 		t.Errorf("mohon untuk melengkapi data")
// 	} else {
// 		err := module.UpdateOneDoc(db, "user", objectId, docs)
// 		if err != nil {
// 			t.Errorf("Error inserting document: %v", err)
// 			fmt.Println("Data tidak berhasil diupdate")
// 		} else {
// 			fmt.Println("Data berhasil diupdate")
// 		}
// 	}
// }

// func TestGetLowonganFromID(t *testing.T){
// 	id := "64d0b1104255ba95ba588512"
// 	objectId, err := primitive.ObjectIDFromHex(id)
// 	if err != nil{
// 		t.Fatalf("error converting id to objectID: %v", err)
// 	}
// 	doc, err := module.GetLowonganFromID(objectId)
// 	if err != nil {
// 		t.Fatalf("error calling GetDocFromID: %v", err)
// 	}
// 	fmt.Println(doc)
// }

func TestInsertUser(t *testing.T) {
	var doc model.User
	doc.Email = "admin@gmail.com"
	password := "admin123"
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		t.Errorf("kesalahan server : salt")
	} else {
		hashedPassword := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)
		user := bson.M{
			"email": doc.Email,
			"password": hex.EncodeToString(hashedPassword),
			"salt": hex.EncodeToString(salt),
			"role": "admin",
		}
		_, err = module.InsertOneDoc(db, "user", user)
		if err != nil {
			t.Errorf("gagal insert")
		} else {
			fmt.Println("berhasil insert")
		}
	}
}

func TestGetUserByAdmin(t *testing.T) {
	id := "65473763d04dda3a8502b58f"
	idparam, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		t.Errorf("Error converting id to objectID: %v", err)
	}
	data, err := module.GetUserFromID(idparam, db)
	if err != nil {
		t.Errorf("Error getting document: %v", err)
	} else {
		if data.Role == "staff" {
			datastaff, err := module.GetStaffFromAkun(data.ID, db)
			if err != nil {
				t.Errorf("Error getting document: %v", err)
			} else {
				datastaff.Akun = data
				fmt.Println(datastaff) 
			}
		}
	}
}

func TestSignUpStaff(t *testing.T) {
	var doc model.Staff
	doc.NamaLengkap = "Al Novianti Ramadhani"
	doc.Jabatan = "Kepala Gudang"
	doc.JenisKelamin = "Perempuan"
	doc.Akun.Email = "novi@gmail.com"
	doc.Akun.Password = "87654321qwe"
	doc.Akun.Confirmpassword = "87654321qwe"
	err := module.SignUpStaff(db, doc)
	if err != nil {
		t.Errorf("Error inserting document: %v", err)
	} else {
	fmt.Println("Data berhasil disimpan dengan nama :", doc.NamaLengkap)
	}
}

// func TestSignUpIndustri(t *testing.T) {
// 	var doc model.User
// 	doc.FirstName = "Dimas"
// 	doc.LastName = "Ardianto"
// 	doc.Email = "dimas@gmail.com"
// 	doc.Password = "fghjkliow"
// 	doc.Confirmpassword = "fghjkliow"
// 	insertedID, err := module.SignUpIndustri(db, "user", doc)
// 	if err != nil {
// 		t.Errorf("Error inserting document: %v", err)
// 	} else {
// 	fmt.Println("Data berhasil disimpan dengan id :", insertedID.Hex())
// 	}
// }

func TestLogIn(t *testing.T) {
	var doc model.User
	doc.Email = "admin@gmail.com"
	doc.Password = "admin123"
	user, err := module.LogIn(db, doc)
	if err != nil {
		t.Errorf("Error getting document: %v", err)
	} else {
		fmt.Println("Welcome bang:", user)
	}
}

func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := module.GenerateKey()
	fmt.Println("ini private key :", privateKey)
	fmt.Println("ini public key :", publicKey)
	id := "64d0b1104255ba95ba588512"
	objectId, err := primitive.ObjectIDFromHex(id)
	role := "admin"
	if err != nil{
		t.Fatalf("error converting id to objectID: %v", err)
	}
	hasil, err := module.Encode(objectId, role, privateKey)
	fmt.Println("ini hasil :", hasil, err)
}

func TestUpdateStaff(t *testing.T) {
	var doc model.Staff
	id := "654a01bde89e6f232a62e41d"
	objectId, _ := primitive.ObjectIDFromHex(id)
	id2 := "654a01bce89e6f232a62e41b"
	userid, _ := primitive.ObjectIDFromHex(id2)
	doc.NamaLengkap = "Agita Nurfadillah"
	doc.Jabatan = "Kepala Gudang"
	doc.JenisKelamin = "Perempuan"
	if doc.NamaLengkap == "" || doc.Jabatan == "" || doc.JenisKelamin == "" {
		t.Errorf("mohon untuk melengkapi data")
	} else {
		err := module.UpdateStaff(objectId, userid, db, doc)
		if err != nil {
			t.Errorf("Error inserting document: %v", err)
			fmt.Println("Data tidak berhasil diupdate")
		} else {
			fmt.Println("Data berhasil diupdate")
		}
	}
}

// func TestWatoken2(t *testing.T) {
// 	var user model.User
// 	privateKey, publicKey := module.GenerateKey()
// 	fmt.Println("privateKey : ", privateKey)
// 	fmt.Println("publicKey : ", publicKey)
// 	id := "649063d3ad72e074286c61e8"
// 	objectId, _ := primitive.ObjectIDFromHex(id)
// 	user.FirstName = "Fatwa"
// 	user.LastName = "Fatahillah"
// 	user.Email = "fax@gmail.com"
// 	user.Role = "pelamar"
// 	tokenstring, err := module.Encode(objectId, privateKey)
// 	if err != nil {
// 		t.Errorf("Error getting document: %v", err)
// 	} else {
// 		body, err := module.Decode(publicKey, tokenstring)
// 		fmt.Println("signed : ", tokenstring)
// 		fmt.Println("isi : ", body)
// 		if err != nil {
// 			t.Errorf("Error getting document: %v", err)
// 		} else {
// 			fmt.Println("Berhasil yey!")
// 		}
// 	}
// }

func TestWatoken(t *testing.T) {
	body, err := module.Decode("f3248b509d9731ebd4e0ccddadb5a08742e083db01678e8a1d734ce81298868f", "v4.public.eyJlbWFpbCI6ImZheEBnbWFpbC5jb20iLCJleHAiOiIyMDIzLTEwLTIyVDAwOjQxOjQ1KzA3OjAwIiwiZmlyc3RuYW1lIjoiRmF0d2EiLCJpYXQiOiIyMDIzLTEwLTIxVDIyOjQxOjQ1KzA3OjAwIiwiaWQiOiI2NDkwNjNkM2FkNzJlMDc0Mjg2YzYxZTgiLCJsYXN0bmFtZSI6IkZhdGFoaWxsYWgiLCJuYmYiOiIyMDIzLTEwLTIxVDIyOjQxOjQ1KzA3OjAwIiwicm9sZSI6InBlbGFtYXIifR_Q4b9X7WC7up7dUUxz_Yki39M-ReovTIoTFfdJmFYRF5Oer0zQZx_ZQamkOsogJ6RuGJhxT3OxrXFS5p6dMg0")
	fmt.Println("isi : ", body, err)
}

// func TestWatoken3(t *testing.T) {
// 	var datauser model.User
// 	privateKey, publicKey := module.GenerateKey()
// 	fmt.Println("privateKey : ", privateKey)
// 	fmt.Println("publicKey : ", publicKey)
// 	datauser.Email = "fatwaff@gmail.com"
// 	datauser.Password = "fghjklio"
// 	user, err := module.LogIn(db, "user", datauser)
// 	fmt.Println("id : ", user.ID)
// 	fmt.Println("firstname : ", user.FirstName)
// 	fmt.Println("lastname : ", user.LastName)
// 	fmt.Println("email : ", user.Email)
// 	fmt.Println("role : ", user.Role)
// 	if err != nil {
// 		t.Errorf("Error getting document: %v", err)
// 	} else {
// 		tokenstring, err := module.Encode(user.ID, privateKey)
// 		if err != nil {
// 			t.Errorf("Error getting document: %v", err)
// 		} else {
// 			body, err := module.Decode(publicKey, tokenstring)
// 			fmt.Println("signed : ", tokenstring)
// 			fmt.Println("isi : ", body)
// 			if err != nil {
// 				t.Errorf("Error getting document: %v", err)
// 			} else {
// 				fmt.Println("Berhasil yey!")
// 			}
// 		}
// 	}
// }


// test Gudang A
func TestInsertGudangA(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "warehouse_db")
	payload, err := module.Decode("5c633cfc6243cec8c9b5dae4a8aae7b366126ad04ee4e5a90c7596e7f8b9a9d8", "v4.public.eyJleHAiOiIyMDIzLTExLTE0VDE2OjA0OjUyWiIsImlhdCI6IjIwMjMtMTEtMTRUMTQ6MDQ6NTJaIiwiaWQiOiJhZ2l0YW51cmZhZGlsbGFoMTAwM0BnbWFpbC5jb20iLCJuYmYiOiIyMDIzLTExLTE0VDE0OjA0OjUyWiJ9gfwp6wwUof8L9BYbblCGMxeiLQSbCVML7L71hatU1KXpXTaHFjpDVFtZ3FPkb3fdKJWWnovR6B4b7Gcr3wGoDQ")
	if err != nil {
		t.Errorf("Error decode token: %v", err)
	}
	// if payload.Role != "mitra" {
	// 	t.Errorf("Error role: %v", err)
	// }
	var datagudanga model.GudangA
	datagudanga.Brand = "Adidas"
	datagudanga.Name = "Five Ten Kestrel Lace Mountain Bike Shoes"
	datagudanga.Category = "Data Science"
	datagudanga.QTY = "75"
	datagudanga.SKU = "BC0770 	"
	datagudanga.SellingPrice = "2.220.000"
	datagudanga.OriginalPrice = "1.500.000"
	datagudanga.Availability = "InStock"
	datagudanga.Color = "Grey"
	datagudanga.Breadcrumbs = "Women/Shoes"
	datagudanga.Date = time.Now() // Inisialisasi dengan waktu saat ini
	err = module.InsertGudangA(payload.Id, conn, datagudanga)
	if err != nil {
		t.Errorf("Error insert : %v", err)
	} else {
		fmt.Println("Berhasil yey!")
	}
}

func TestUpdateGudangA(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "warehouse_db")
	payload, err := module.Decode("5c633cfc6243cec8c9b5dae4a8aae7b366126ad04ee4e5a90c7596e7f8b9a9d8", "v4.public.eyJleHAiOiIyMDIzLTExLTAxVDA2OjQ5OjQ0WiIsImlhdCI6IjIwMjMtMTEtMDFUMDQ6NDk6NDRaIiwiaWQiOiI2NTQwNjMyODI4NzY0ZDk2YzY0OWYyOWQiLCJuYmYiOiIyMDIzLTExLTAxVDA0OjQ5OjQ0WiJ92RxBGslXaHBoLQhvMJLQO7uEBG5c5FmkpZgakPjmk1aUFDdRkw3m3r-7BpkhDmCtByoARDr36X3DhjCL8HT8AQ")
	// payload, err := module.Decode("5c633cfc6243cec8c9b5dae4a8aae7b366126ad04ee4e5a90c7596e7f8b9a9d8", "v4.public.eyJleHAiOiIyMDIzLTExLTAxVDA2OjQ3OjMxWiIsImlhdCI6IjIwMjMtMTEtMDFUMDQ6NDc6MzFaIiwiaWQiOiI2NTNkZTllYjg5MzlmYjNjZjI3ZjZkMzciLCJuYmYiOiIyMDIzLTExLTAxVDA0OjQ3OjMxWiJ92YbTLQWznLupbH0Syb6GPKkj4ZW_JJLveVcFTfZElv8_jybZCMBnw8y-7SLZVMpRTq56PaArdEBwlvvSPQjtCg")
	if err != nil {
		t.Errorf("Error decode token: %v", err)
	}
	// if payload.Role != "mitra" {
	// 	t.Errorf("Error role: %v", err)
	// }
	var datagudanga model.GudangA
	datagudanga.Brand = "Adidas"
	datagudanga.Name = "Five Ten Kestrel Lace Mountain Bike Shoes"
	datagudanga.Category = "Data Science"
	datagudanga.QTY = "25"
	datagudanga.SKU = "BC0770 	"
	datagudanga.SellingPrice = "2.220.000"
	datagudanga.OriginalPrice = "1.500.000"
	datagudanga.Availability = "InStock"
	datagudanga.Color = "Grey"
	datagudanga.Breadcrumbs = "Women/Shoes"
	datagudanga.Date = time.Now() // Inisialisasi dengan waktu saat ini
	id := "65536f9a24e331f6ed254e56"
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		t.Fatalf("error converting id to objectID: %v", err)
	}
	err = module.UpdateGudangA(objectId, payload.Id, conn, datagudanga)
	if err != nil {
		t.Errorf("Error update : %v", err)
	} else {
		fmt.Println("Berhasil yey!")
	}
}

func TestDeleteGudangA(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "warehouse_db")
	payload, err := module.Decode("5c633cfc6243cec8c9b5dae4a8aae7b366126ad04ee4e5a90c7596e7f8b9a9d8", "v4.public.eyJleHAiOiIyMDIzLTExLTAxVDA2OjQ5OjQ0WiIsImlhdCI6IjIwMjMtMTEtMDFUMDQ6NDk6NDRaIiwiaWQiOiI2NTQwNjMyODI4NzY0ZDk2YzY0OWYyOWQiLCJuYmYiOiIyMDIzLTExLTAxVDA0OjQ5OjQ0WiJ92RxBGslXaHBoLQhvMJLQO7uEBG5c5FmkpZgakPjmk1aUFDdRkw3m3r-7BpkhDmCtByoARDr36X3DhjCL8HT8AQ")
	if err != nil {
		t.Errorf("Error decode token: %v", err)
	}
	// if payload.Role != "mitra" {
	// 	t.Errorf("Error role: %v", err)
	// }
	id := "65536f9a24e331f6ed254e56"
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		t.Fatalf("error converting id to objectID: %v", err)
	}
	err = module.DeleteGudangA(objectId, payload.Id, conn)
	if err != nil {
		t.Errorf("Error delete : %v", err)
	} else {
		fmt.Println("Berhasil yey!")
	}
}

func TestGetAllGudangA(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "warehouse_db")
	data, err := module.GetAllGudangA(conn)
	if err != nil {
		t.Errorf("Error get all : %v", err)
	} else {
		fmt.Println(data)
	}
}

func TestGetGudangAFromID(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "warehouse_db")
	id := "65406377996edfaee3ed9a19"
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		t.Fatalf("error converting id to objectID: %v", err)
	}
	gudanga, err := module.GetGudangAFromID(objectId, conn)
	if err != nil {
		t.Errorf("Error get gudanga : %v", err)
	} else {
		fmt.Println(gudanga)
	}
}

// test Gudang B
func TestInsertGudangB(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "warehouse_db")
	payload, err := module.Decode("5c633cfc6243cec8c9b5dae4a8aae7b366126ad04ee4e5a90c7596e7f8b9a9d8", "v4.public.eyJleHAiOiIyMDIzLTExLTE0VDE2OjA0OjUyWiIsImlhdCI6IjIwMjMtMTEtMTRUMTQ6MDQ6NTJaIiwiaWQiOiJhZ2l0YW51cmZhZGlsbGFoMTAwM0BnbWFpbC5jb20iLCJuYmYiOiIyMDIzLTExLTE0VDE0OjA0OjUyWiJ9gfwp6wwUof8L9BYbblCGMxeiLQSbCVML7L71hatU1KXpXTaHFjpDVFtZ3FPkb3fdKJWWnovR6B4b7Gcr3wGoDQ")
	if err != nil {
		t.Errorf("Error decode token: %v", err)
	}
	// if payload.Role != "mitra" {
	// 	t.Errorf("Error role: %v", err)
	// }
	var datagudangb model.GudangB
	datagudangb.Brand = "Adidas"
	datagudangb.Name = "Five Ten Kestrel Lace Mountain Bike Shoes"
	datagudangb.Category = "Data Science"
	datagudangb.QTY = "75"
	datagudangb.SKU = "BC0770 	"
	datagudangb.SellingPrice = "2.220.000"
	datagudangb.OriginalPrice = "1.500.000"
	datagudangb.Availability = "InStock"
	datagudangb.Color = "Grey"
	datagudangb.Breadcrumbs = "Women/Shoes"
	datagudangb.Date = time.Now() // Inisialisasi dengan waktu saat ini
	err = module.InsertGudangB(payload.Id, conn, datagudangb)
	if err != nil {
		t.Errorf("Error insert : %v", err)
	} else {
		fmt.Println("Berhasil yey!")
	}
}

func TestUpdateGudangB(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "warehouse_db")
	payload, err := module.Decode("5c633cfc6243cec8c9b5dae4a8aae7b366126ad04ee4e5a90c7596e7f8b9a9d8", "v4.public.eyJleHAiOiIyMDIzLTExLTAxVDA2OjQ5OjQ0WiIsImlhdCI6IjIwMjMtMTEtMDFUMDQ6NDk6NDRaIiwiaWQiOiI2NTQwNjMyODI4NzY0ZDk2YzY0OWYyOWQiLCJuYmYiOiIyMDIzLTExLTAxVDA0OjQ5OjQ0WiJ92RxBGslXaHBoLQhvMJLQO7uEBG5c5FmkpZgakPjmk1aUFDdRkw3m3r-7BpkhDmCtByoARDr36X3DhjCL8HT8AQ")
	// payload, err := module.Decode("5c633cfc6243cec8c9b5dae4a8aae7b366126ad04ee4e5a90c7596e7f8b9a9d8", "v4.public.eyJleHAiOiIyMDIzLTExLTAxVDA2OjQ3OjMxWiIsImlhdCI6IjIwMjMtMTEtMDFUMDQ6NDc6MzFaIiwiaWQiOiI2NTNkZTllYjg5MzlmYjNjZjI3ZjZkMzciLCJuYmYiOiIyMDIzLTExLTAxVDA0OjQ3OjMxWiJ92YbTLQWznLupbH0Syb6GPKkj4ZW_JJLveVcFTfZElv8_jybZCMBnw8y-7SLZVMpRTq56PaArdEBwlvvSPQjtCg")
	if err != nil {
		t.Errorf("Error decode token: %v", err)
	}
	// if payload.Role != "mitra" {
	// 	t.Errorf("Error role: %v", err)
	// }
	var datagudangb model.GudangB
	datagudangb.Brand = "Adidas"
	datagudangb.Name = "Five Ten Kestrel Lace Mountain Bike Shoes"
	datagudangb.Category = "Data Science"
	datagudangb.QTY = "25"
	datagudangb.SKU = "BC0770 	"
	datagudangb.SellingPrice = "2.220.000"
	datagudangb.OriginalPrice = "1.500.000"
	datagudangb.Availability = "InStock"
	datagudangb.Color = "Grey"
	datagudangb.Breadcrumbs = "Women/Shoes"
	datagudangb.Date = time.Now() // Inisialisasi dengan waktu saat ini
	id := "65406377996edfaee3ed9a19"
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		t.Fatalf("error converting id to objectID: %v", err)
	}
	err = module.UpdateGudangB(objectId, payload.Id, conn, datagudangb)
	if err != nil {
		t.Errorf("Error update : %v", err)
	} else {
		fmt.Println("Berhasil yey!")
	}
}

func TestDeleteGudangB(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "warehouse_db")
	payload, err := module.Decode("5c633cfc6243cec8c9b5dae4a8aae7b366126ad04ee4e5a90c7596e7f8b9a9d8", "v4.public.eyJleHAiOiIyMDIzLTExLTAxVDA2OjQ5OjQ0WiIsImlhdCI6IjIwMjMtMTEtMDFUMDQ6NDk6NDRaIiwiaWQiOiI2NTQwNjMyODI4NzY0ZDk2YzY0OWYyOWQiLCJuYmYiOiIyMDIzLTExLTAxVDA0OjQ5OjQ0WiJ92RxBGslXaHBoLQhvMJLQO7uEBG5c5FmkpZgakPjmk1aUFDdRkw3m3r-7BpkhDmCtByoARDr36X3DhjCL8HT8AQ")
	if err != nil {
		t.Errorf("Error decode token: %v", err)
	}
	// if payload.Role != "mitra" {
	// 	t.Errorf("Error role: %v", err)
	// }
	id := "65406377996edfaee3ed9a19"
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		t.Fatalf("error converting id to objectID: %v", err)
	}
	err = module.DeleteGudangB(objectId, payload.Id, conn)
	if err != nil {
		t.Errorf("Error delete : %v", err)
	} else {
		fmt.Println("Berhasil yey!")
	}
}

func TestGetAllGudangB(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "warehouse_db")
	data, err := module.GetAllGudangB(conn)
	if err != nil {
		t.Errorf("Error get all : %v", err)
	} else {
		fmt.Println(data)
	}
}

func TestGetGudangBFromID(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "warehouse_db")
	id := "65406377996edfaee3ed9a19"
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		t.Fatalf("error converting id to objectID: %v", err)
	}
	gudangb, err := module.GetGudangBFromID(objectId, conn)
	if err != nil {
		t.Errorf("Error get gudangb : %v", err)
	} else {
		fmt.Println(gudangb)
	}
}

// test Gudang C
func TestInsertGudangC(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "warehouse_db")
	payload, err := module.Decode("5c633cfc6243cec8c9b5dae4a8aae7b366126ad04ee4e5a90c7596e7f8b9a9d8", "v4.public.eyJleHAiOiIyMDIzLTExLTE0VDE2OjA0OjUyWiIsImlhdCI6IjIwMjMtMTEtMTRUMTQ6MDQ6NTJaIiwiaWQiOiJhZ2l0YW51cmZhZGlsbGFoMTAwM0BnbWFpbC5jb20iLCJuYmYiOiIyMDIzLTExLTE0VDE0OjA0OjUyWiJ9gfwp6wwUof8L9BYbblCGMxeiLQSbCVML7L71hatU1KXpXTaHFjpDVFtZ3FPkb3fdKJWWnovR6B4b7Gcr3wGoDQ")
	if err != nil {
		t.Errorf("Error decode token: %v", err)
	}
	// if payload.Role != "mitra" {
	// 	t.Errorf("Error role: %v", err)
	// }
	var datagudangc model.GudangC
	datagudangc.Brand = "Adidas"
	datagudangc.Name = "Five Ten Kestrel Lace Mountain Bike Shoes"
	datagudangc.Category = "Data Science"
	datagudangc.QTY = "75"
	datagudangc.SKU = "BC0770 	"
	datagudangc.SellingPrice = "2.220.000"
	datagudangc.OriginalPrice = "1.500.000"
	datagudangc.Availability = "InStock"
	datagudangc.Color = "Grey"
	datagudangc.Breadcrumbs = "Women/Shoes"
	datagudangc.Date = time.Now() // Inisialisasi dengan waktu saat ini
	err = module.InsertGudangC(payload.Id, conn, datagudangc)
	if err != nil {
		t.Errorf("Error insert : %v", err)
	} else {
		fmt.Println("Berhasil yey!")
	}
}

func TestUpdateGudangC(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "warehouse_db")
	payload, err := module.Decode("5c633cfc6243cec8c9b5dae4a8aae7b366126ad04ee4e5a90c7596e7f8b9a9d8", "v4.public.eyJleHAiOiIyMDIzLTExLTAxVDA2OjQ5OjQ0WiIsImlhdCI6IjIwMjMtMTEtMDFUMDQ6NDk6NDRaIiwiaWQiOiI2NTQwNjMyODI4NzY0ZDk2YzY0OWYyOWQiLCJuYmYiOiIyMDIzLTExLTAxVDA0OjQ5OjQ0WiJ92RxBGslXaHBoLQhvMJLQO7uEBG5c5FmkpZgakPjmk1aUFDdRkw3m3r-7BpkhDmCtByoARDr36X3DhjCL8HT8AQ")
	// payload, err := module.Decode("5c633cfc6243cec8c9b5dae4a8aae7b366126ad04ee4e5a90c7596e7f8b9a9d8", "v4.public.eyJleHAiOiIyMDIzLTExLTAxVDA2OjQ3OjMxWiIsImlhdCI6IjIwMjMtMTEtMDFUMDQ6NDc6MzFaIiwiaWQiOiI2NTNkZTllYjg5MzlmYjNjZjI3ZjZkMzciLCJuYmYiOiIyMDIzLTExLTAxVDA0OjQ3OjMxWiJ92YbTLQWznLupbH0Syb6GPKkj4ZW_JJLveVcFTfZElv8_jybZCMBnw8y-7SLZVMpRTq56PaArdEBwlvvSPQjtCg")
	if err != nil {
		t.Errorf("Error decode token: %v", err)
	}
	// if payload.Role != "mitra" {
	// 	t.Errorf("Error role: %v", err)
	// }
	var datagudangc model.GudangC
	datagudangc.Brand = "Adidas"
	datagudangc.Name = "Five Ten Kestrel Lace Mountain Bike Shoes"
	datagudangc.Category = "Data Science"
	datagudangc.QTY = "25"
	datagudangc.SKU = "BC0770 	"
	datagudangc.SellingPrice = "2.220.000"
	datagudangc.OriginalPrice = "1.500.000"
	datagudangc.Availability = "InStock"
	datagudangc.Color = "Grey"
	datagudangc.Breadcrumbs = "Women/Shoes"
	datagudangc.Date = time.Now() // Inisialisasi dengan waktu saat ini
	id := "65406377996edfaee3ed9a19"
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		t.Fatalf("error converting id to objectID: %v", err)
	}
	err = module.UpdateGudangC(objectId, payload.Id, conn, datagudangc)
	if err != nil {
		t.Errorf("Error update : %v", err)
	} else {
		fmt.Println("Berhasil yey!")
	}
}

func TestDeleteGudangC(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "warehouse_db")
	payload, err := module.Decode("5c633cfc6243cec8c9b5dae4a8aae7b366126ad04ee4e5a90c7596e7f8b9a9d8", "v4.public.eyJleHAiOiIyMDIzLTExLTAxVDA2OjQ5OjQ0WiIsImlhdCI6IjIwMjMtMTEtMDFUMDQ6NDk6NDRaIiwiaWQiOiI2NTQwNjMyODI4NzY0ZDk2YzY0OWYyOWQiLCJuYmYiOiIyMDIzLTExLTAxVDA0OjQ5OjQ0WiJ92RxBGslXaHBoLQhvMJLQO7uEBG5c5FmkpZgakPjmk1aUFDdRkw3m3r-7BpkhDmCtByoARDr36X3DhjCL8HT8AQ")
	if err != nil {
		t.Errorf("Error decode token: %v", err)
	}
	// if payload.Role != "mitra" {
	// 	t.Errorf("Error role: %v", err)
	// }
	id := "65406377996edfaee3ed9a19"
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		t.Fatalf("error converting id to objectID: %v", err)
	}
	err = module.DeleteGudangC(objectId, payload.Id, conn)
	if err != nil {
		t.Errorf("Error delete : %v", err)
	} else {
		fmt.Println("Berhasil yey!")
	}
}

func TestGetAllGudangC(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "warehouse_db")
	data, err := module.GetAllGudangC(conn)
	if err != nil {
		t.Errorf("Error get all : %v", err)
	} else {
		fmt.Println(data)
	}
}

func TestGetGudangCFromID(t *testing.T) {
	conn := module.MongoConnect("MONGOSTRING", "warehouse_db")
	id := "65406377996edfaee3ed9a19"
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		t.Fatalf("error converting id to objectID: %v", err)
	}
	gudangc, err := module.GetGudangCFromID(objectId, conn)
	if err != nil {
		t.Errorf("Error get gudangc : %v", err)
	} else {
		fmt.Println(gudangc)
	}
}

func TestReturnStruct(t *testing.T){
	// var user model.User
	// user.Email = "fatwa"
	id := "65473763d04dda3a8502b58f"
	objectId, _ := primitive.ObjectIDFromHex(id)
	user, _ := module.GetUserFromID(objectId, db)
	data := model.User{ 
		ID : user.ID,
		Email: user.Email,
		Role : user.Role,
	}
	hasil := module.GCFReturnStruct(data)
	fmt.Println(hasil)
}