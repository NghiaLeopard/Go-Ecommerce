// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/sqlc-dev/pqtype"
)

type UsersStatus string

const (
	UsersStatus0 UsersStatus = "0"
	UsersStatus1 UsersStatus = "1"
)

func (e *UsersStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UsersStatus(s)
	case string:
		*e = UsersStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for UsersStatus: %T", src)
	}
	return nil
}

type NullUsersStatus struct {
	UsersStatus UsersStatus `json:"users_status"`
	Valid       bool        `json:"valid"` // Valid is true if UsersStatus is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUsersStatus) Scan(value interface{}) error {
	if value == nil {
		ns.UsersStatus, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UsersStatus.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUsersStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UsersStatus), nil
}

type UsersType string

const (
	UsersType1 UsersType = "1"
	UsersType2 UsersType = "2"
	UsersType3 UsersType = "3"
)

func (e *UsersType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UsersType(s)
	case string:
		*e = UsersType(s)
	default:
		return fmt.Errorf("unsupported scan type for UsersType: %T", src)
	}
	return nil
}

type NullUsersType struct {
	UsersType UsersType `json:"users_type"`
	Valid     bool      `json:"valid"` // Valid is true if UsersType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUsersType) Scan(value interface{}) error {
	if value == nil {
		ns.UsersType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UsersType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUsersType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.UsersType), nil
}

type City struct {
	ID       int64     `json:"_id"`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type DeliveryType struct {
	ID       int64     `json:"_id"`
	Name     string    `json:"name"`
	Price    int32     `json:"price"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"update_at"`
}

type PaymentType struct {
	ID       int64     `json:"_id"`
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	CreateAt time.Time `json:"createAt"`
	UpdateAt time.Time `json:"update_at"`
}

type Product struct {
	ID                int64     `json:"_id"`
	Name              string    `json:"name"`
	Image             string    `json:"image"`
	CountInStock      int32     `json:"countInStock"`
	Description       string    `json:"description"`
	Sold              int32     `json:"sold"`
	Discount          int32     `json:"discount"`
	DiscountStartDate time.Time `json:"discountStartDate"`
	DiscountEndDate   time.Time `json:"discountEndDate"`
	Type              int32     `json:"type"`
	Status            int32     `json:"status"`
	Slug              string    `json:"slug"`
	Price             int32     `json:"price"`
	Location          int32     `json:"location"`
	Views             int32     `json:"views"`
	CreateAt          time.Time `json:"create_at"`
}

type ProductLiked struct {
	ProductID int32     `json:"product_id"`
	UserID    int32     `json:"user_id"`
	LikeDate  time.Time `json:"like_date"`
}

type ProductType struct {
	ID       int64        `json:"_id"`
	Name     string       `json:"name"`
	Slug     string       `json:"slug"`
	CreateAt time.Time    `json:"create_at"`
	UpdateAt sql.NullTime `json:"update_at"`
}

type ProductUniqueView struct {
	ProductID int32     `json:"product_id"`
	UserID    int32     `json:"user_id"`
	ViewDate  time.Time `json:"view_date"`
}

type Role struct {
	ID         int64     `json:"_id"`
	Name       string    `json:"name"`
	Permission []string  `json:"permission"`
	CreateAt   time.Time `json:"create_at"`
	UpdateAt   time.Time `json:"update_at"`
}

type User struct {
	ID             int64                 `json:"_id"`
	Email          string                `json:"email"`
	Password       string                `json:"password"`
	UserType       UsersType             `json:"userType"`
	Status         UsersStatus           `json:"status"`
	Address        sql.NullString        `json:"address"`
	Avatar         sql.NullString        `json:"avatar"`
	Image          sql.NullString        `json:"image"`
	PhoneNumber    sql.NullString        `json:"phoneNumber"`
	Role           sql.NullInt64         `json:"role"`
	FirstName      sql.NullString        `json:"firstName"`
	LastName       sql.NullString        `json:"lastName"`
	MiddleName     sql.NullString        `json:"middleName"`
	City           sql.NullInt64         `json:"city"`
	LikeProducts   []int64               `json:"likeProducts"`
	ViewedProducts []int64               `json:"viewedProducts"`
	DeviceToken    []string              `json:"deviceToken"`
	Addresses      pqtype.NullRawMessage `json:"addresses"`
	CreateAt       time.Time             `json:"create_at"`
}
