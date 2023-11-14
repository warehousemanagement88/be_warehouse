package model

import (
	"time"

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
	Jabatan      string             `bson:"jabatan,omitempty" json:"jabatan,omitempty"`
	JenisKelamin string             `bson:"jeniskelamin,omitempty" json:"jeniskelamin,omitempty"`
	Akun         User               `bson:"akun,omitempty" json:"akun,omitempty"`
}

type GudangA struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Brand         string             `bson:"brand,omitempty" json:"brand,omitempty"`
	Name          string             `bson:"name,omitempty" json:"name,omitempty"`
	Category      string             `bson:"category,omitempty" json:"category,omitempty"`
	QTY           string             `bson:"qty,omitempty" json:"qty,omitempty"`
	SKU           string             `bson:"sku,omitempty" json:"sku,omitempty"`
	SellingPrice  string             `bson:"sellingprice,omitempty" json:"sellingprice,omitempty"`
	OriginalPrice string             `bson:"originalprice,omitempty" json:"originalprice,omitempty"`
	Availability  string             `bson:"availability,omitempty" json:"availability,omitempty"`
	Color         string             `bson:"color,omitempty" json:"color,omitempty"`
	Breadcrumbs   string             `bson:"breadcrumbs,omitempty" json:"breadcrumbs,omitempty"`
	Date          time.Time          `bson:"date,omitempty" json:"date,omitempty"`
}
type GudangB struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Brand         string             `bson:"brand,omitempty" json:"brand,omitempty"`
	Name          string             `bson:"name,omitempty" json:"name,omitempty"`
	Category      string             `bson:"category,omitempty" json:"category,omitempty"`
	QTY           string             `bson:"qty,omitempty" json:"qty,omitempty"`
	SKU           string             `bson:"sku,omitempty" json:"sku,omitempty"`
	SellingPrice  string             `bson:"sellingprice,omitempty" json:"sellingprice,omitempty"`
	OriginalPrice string             `bson:"originalprice,omitempty" json:"originalprice,omitempty"`
	Availability  string             `bson:"availability,omitempty" json:"availability,omitempty"`
	Color         string             `bson:"color,omitempty" json:"color,omitempty"`
	Breadcrumbs   string             `bson:"breadcrumbs,omitempty" json:"breadcrumbs,omitempty"`
	Date          time.Time          `bson:"date,omitempty" json:"date,omitempty"`
}
type GudangC struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Brand         string             `bson:"brand,omitempty" json:"brand,omitempty"`
	Name          string             `bson:"name,omitempty" json:"name,omitempty"`
	Category      string             `bson:"category,omitempty" json:"category,omitempty"`
	QTY           string             `bson:"qty,omitempty" json:"qty,omitempty"`
	SKU           string             `bson:"sku,omitempty" json:"sku,omitempty"`
	SellingPrice  string             `bson:"sellingprice,omitempty" json:"sellingprice,omitempty"`
	OriginalPrice string             `bson:"originalprice,omitempty" json:"originalprice,omitempty"`
	Availability  string             `bson:"availability,omitempty" json:"availability,omitempty"`
	Color         string             `bson:"color,omitempty" json:"color,omitempty"`
	Breadcrumbs   string             `bson:"breadcrumbs,omitempty" json:"breadcrumbs,omitempty"`
	Date          time.Time          `bson:"date,omitempty" json:"date,omitempty"`
}

type Credential struct {
	Status  bool   `json:"status" bson:"status"`
	Token   string `json:"token,omitempty" bson:"token,omitempty"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
	Role    string `json:"role,omitempty" bson:"role,omitempty"`
}

type Response struct {
	Status  bool   `json:"status" bson:"status"`
	Message string `json:"message,omitempty" bson:"message,omitempty"`
}

type Payload struct {
	Id   primitive.ObjectID `json:"id"`
	Role string             `json:"role"`
	Exp  time.Time          `json:"exp"`
	Iat  time.Time          `json:"iat"`
	Nbf  time.Time          `json:"nbf"`
}
