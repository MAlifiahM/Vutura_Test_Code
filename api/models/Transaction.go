package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Transaction struct {
	ID			uint64		`gorm:"primary_key;auto_increment" json:"id"`
	IDUser		uint32		`gorm:"size:255;not null" json:"id_user"`
	Product		string		`gorm:"size:255;not null" json:"product"`
	Price		int			`gorm:"size:255;not null" json:"price"`
	CreatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt	time.Time	`gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (t *Transaction) Prepare() {
	t.ID 		= 0
	t.IDUser 	= 0
	t.Product 	= html.EscapeString(strings.TrimSpace(t.Product))
	t.Price		= 0
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()	
}

func (t *Transaction) Validate() error {
	if t.Product == "" {
		return errors.New("Required Product")
	}

	if t.Price <= 0 {
		return errors.New("Required Price")
	}

	if t.IDUser < 1 {
		return errors.New("Required User")
	}
	return nil
}

func (t *Transaction) SaveTransaction(db *gorm.DB) (*Transaction, error) {
	var err error

	err = db.Debug().Model(&Transaction{}).Create(&t).Error
	if err != nil {
		return &Transaction{}, err
	}
	return t, nil
}

func (t *Transaction) FindAllTransactions(db *gorm.DB) (*[]Transaction, error) {
	var err error
	transactions := []Transaction{}
	err = db.Debug().Model(&Transaction{}).Limit(100).Find(&transactions).Error
	if err != nil {
		return &[]Transaction{}, err
	}
	
	return &transactions, nil
}

func (t *Transaction) FindTransactionByID(db *gorm.DB, pid uint64) (*Transaction, error) {
	var err error

	err = db.Debug().Model(&Transaction{}).Where("id = ?", pid).Take(&t).Error
	if err != nil {
		return &Transaction{}, err
	}

	return t, nil
}

func (t *Transaction) UpdateATransaction(db *gorm.DB) (*Transaction, error) {
	var err error

	err = db.Debug().Model(&Transaction{}).Where("id = ?", t.ID).Updates(Transaction{Product: t.Product, Price: t.Price, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Transaction{}, err
	}

	return t, nil
}

func (t *Transaction) DeleteATransaction(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Transaction{}).Where("id = ? and id_user = ?", pid, uid).Take(&Transaction{}).Delete(&Transaction{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Transaction not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}


