package repo

import (
	"GrabOrders/database"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Order struct {
	ID        int64     `gorm:"column:id" json:"id"`
	Amount    int64     `gorm:"column:amount" json:"amount"`
	CSID      int64     `gomr:"column:cs_id" json:"csID"`
	CreatedAt time.Time `gomr:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time `gomr:"column:updated_at" json:"updatedAt"`
}

func (*Order) TableName() string {
	return "orders"
}

func GenOrders(n int64) {
	db := database.GetConnection()
	defer db.Close()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	amounts := make([]string, n)
	for i := range amounts {
		amounts[i] = strconv.Itoa(r1.Intn(50000))
	}

	sql := fmt.Sprintf(`INSERT INTO orders (amount) VALUES (%s)`, strings.Join(amounts, "),("))

	err := db.Exec(sql).Error
	if err != nil {
		log.Printf("%#v \n", err)
	}
}

func GetFirstUngrabOrders(csID int64) int64 {
	db := database.GetConnection()
	defer db.Close()

	order := &Order{}
	err := db.Model(&Order{}).Where("cs_id = 0").First(order).Error
	if err != nil {
		log.Println(err)
		return 0
	}
	log.Println(csID, " get id: ", order.ID)
	return order.ID
}

func GrabOrder(csID, id int64) error {
	db := database.GetConnection()
	defer db.Close()

	//BEGIN;
	//SELECT * FROM orders WHERE id = 1 AND cs_id = 0 FOR UPDATE;
	//UPDATE orders SET cs_id = 2 WHERE id = 1 AND cs_id = 0;
	//COMMIT;

	tx := db.Begin()
	order := &Order{}
	rowsAffected := tx.Raw("SELECT * FROM orders WHERE id = ? AND cs_id = 0 FOR UPDATE;", id).Find(&order).RowsAffected
	if rowsAffected != 0 {
		err := tx.Exec("UPDATE orders SET cs_id = ? WHERE id = ? AND cs_id = 0;", csID, id).Error
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return err
		}
		tx.Commit()
		return nil
	} else {
		tx.Commit()
	}
	return errors.New("this order is grabbed")
}
