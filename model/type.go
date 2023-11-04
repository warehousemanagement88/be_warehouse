package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	FirstName       string             `bson:"firstname,omitempty" json:"firstname,omitempty"`
	LastName        string             `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Email           string             `bson:"email,omitempty" json:"email,omitempty"`
	Password        string             `bson:"password,omitempty" json:"password,omitempty"`
	Confirmpassword string             `bson:"confirmpass,omitempty" json:"confirmpass,omitempty"`
	Salt            string             `bson:"salt,omitempty" json:"salt,omitempty"`
	Role            string             `bson:"role,omitempty" json:"role,omitempty" `
}

type Staff struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaLengkap  string             `bson:"namalengkap,omitempty" json:"namalengkap,omitempty"`
	TanggalLahir string             `bson:"tanggallahir,omitempty" json:"tanggallahir,omitempty"`
	JenisKelamin string             `bson:"jeniskelamin,omitempty" json:"jeniskelamin,omitempty"`
	Akun         User               `bson:"akun,omitempty" json:"akun,omitempty"`
}
type Karyawan struct {
	ID           	primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	NamaLengkap  	string             `bson:"namalengkap,omitempty" json:"namalengkap,omitempty"`
	TanggalLahir	string             `bson:"tanggallahir,omitempty" json:"tanggallahir,omitempty"`
	JenisKelamin  	string             `bson:"jeniskelamin,omitempty" json:"jeniskelamin,omitempty"`
	Akun     		User			   `bson:"akun,omitempty" json:"akun,omitempty"`
	Akunadmin     		User			   `bson:"akunadmin,omitempty" json:"akunadmin,omitempty"`
}

type Item struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string             `bson:"name,omitempty" json:"name,omitempty"`
	Quantity string             `bson:"quantity,omitempty" json:"quantity,omitempty"`
	Category string             `bson:"category,omitempty" json:"category,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}
