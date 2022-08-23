package utils

import (
	"database/sql"

	models "msebaktir.com/LCV/Models"
	// _ "go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil

}

func GetAllCustomers(db *sql.DB) ([]models.Customer, error) {
	rows, err := db.Query("SELECT * FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	customers := []models.Customer{}
	for rows.Next() {
		var customer models.Customer
		err := rows.Scan(&customer.ID, &customer.UserName, &customer.Status)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}
func UpdateCustomer(db *sql.DB, customer models.Customer) error {
	_, err := db.Exec("UPDATE customers SET Status = ?,UserName = ?,Password = ?, EventId = ? WHERE Id = ?", customer.Status, customer.UserName, customer.Password, customer.Event.Id, customer.ID)
	if err != nil {
		return err
	}
	return nil
}
func CreateCustomer(db *sql.DB, customer models.Customer) (error, int) {
	data, err := db.Exec("INSERT INTO customers (UserName,Password,Status,EventId) VALUES (?,?,?,?)", customer.UserName, customer.Password, customer.Status, customer.Event.Id)
	if err != nil {
		return err, 0
	}
	id, err := data.LastInsertId()
	if err != nil {
		return err, 0
	}
	return nil, int(id)
}
func DeleteCustomer(db *sql.DB, customer models.Customer) error {
	_, err := db.Exec("DELETE FROM customers WHERE Id = ?", customer.ID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllEvents(db *sql.DB) ([]models.Event, error) {
	rows, err := db.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	events := []models.Event{}
	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.Id, &event.Title, &event.Description, &event.StartDate, &event.LastCall, &event.Status)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
func UpdateEvent(db *sql.DB, event models.Event) error {
	_, err := db.Exec("UPDATE events SET Title = ?,Description = ?,StartDate = ?,LastCall = ?,Status = ? WHERE Id = ?", event.Title, event.Description, event.StartDate, event.LastCall, event.Status, event.Id)
	if err != nil {
		return err
	}
	return nil
}
func CreateEvent(db *sql.DB, event models.Event) (error, int) {
	data, err := db.Exec("INSERT INTO events (Title,Description,StartDate,LastCall,Status) VALUES (?,?,?,?,?)", event.Title, event.Description, event.StartDate, event.LastCall, event.Status)
	if err != nil {
		return err, 0
	}
	id, err := data.LastInsertId()
	if err != nil {
		return err, 0
	}
	return nil, int(id)
}
func DeleteEvent(db *sql.DB, event models.Event) error {
	_, err := db.Exec("DELETE FROM events WHERE Id = ?", event.Id)
	if err != nil {
		return err
	}
	return nil
}

func GetAllGuests(db *sql.DB) ([]models.Guest, error) {
	rows, err := db.Query("SELECT * FROM guests")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	guests := []models.Guest{}
	for rows.Next() {
		var guest models.Guest
		err := rows.Scan(&guest.Id, &guest.Name, &guest.Email, &guest.Phone, &guest.HowManyPeople, &guest.HowManyChildren, &guest.Customer.ID, &guest.Event.Id, &guest.Gift.Id, &guest.Status)
		if err != nil {
			return nil, err
		}
		guests = append(guests, guest)
	}
	return guests, nil
}
func UpdateGuest(db *sql.DB, guest models.Guest) error {
	_, err := db.Exec("UPDATE guests SET Name = ?,Email = ?,Phone = ?,HowManyPeople = ?,HowManyChildren = ?,CustomerId = ?,EventId = ?,GiftId = ?,Status = ? WHERE Id = ?", guest.Name, guest.Email, guest.Phone, guest.HowManyPeople, guest.HowManyChildren, guest.Customer.ID, guest.Event.Id, guest.Gift.Id, guest.Status, guest.Id)
	if err != nil {
		return err
	}
	return nil
}
func CreateGuest(db *sql.DB, guest models.Guest) (error, int) {
	data, err := db.Exec("INSERT INTO guests (Name,Email,Phone,HowManyPeople,HowManyChildren,CustomerId,EventId,GiftId,Status) VALUES (?,?,?,?,?,?,?,?,?)", guest.Name, guest.Email, guest.Phone, guest.HowManyPeople, guest.HowManyChildren, guest.Customer.ID, guest.Event.Id, guest.Gift.Id, guest.Status)
	if err != nil {
		return err, 0
	}
	id, err := data.LastInsertId()
	if err != nil {
		return err, 0
	}
	return nil, int(id)
}
func DeleteGuest(db *sql.DB, guest models.Guest) error {
	_, err := db.Exec("DELETE FROM guests WHERE Id = ?", guest.Id)
	if err != nil {
		return err
	}
	return nil
}

func GetAllGifts(db *sql.DB) ([]models.Gift, error) {
	rows, err := db.Query("SELECT * FROM gifts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	gifts := []models.Gift{}
	for rows.Next() {
		var gift models.Gift
		err := rows.Scan(&gift.Id, &gift.Name, &gift.Description, &gift.Price, &gift.Image)
		if err != nil {
			return nil, err
		}
		gifts = append(gifts, gift)
	}
	return gifts, nil
}
func UpdateGift(db *sql.DB, gift models.Gift) error {
	_, err := db.Exec("UPDATE gifts SET Name = ?,Description = ?,Price = ?,Image = ? WHERE Id = ?", gift.Name, gift.Description, gift.Price, gift.Image, gift.Id)
	if err != nil {
		return err
	}
	return nil
}
func CreateGift(db *sql.DB, gift models.Gift) (error, int) {
	data, err := db.Exec("INSERT INTO gifts (Name,Description,Price,Image) VALUES (?,?,?,?)", gift.Name, gift.Description, gift.Price, gift.Image)
	if err != nil {
		return err, 0
	}
	id, err := data.LastInsertId()
	if err != nil {
		return err, 0
	}
	return nil, int(id)
}
func DeleteGift(db *sql.DB, gift models.Gift) error {
	_, err := db.Exec("DELETE FROM gifts WHERE Id = ?", gift.Id)
	if err != nil {
		return err
	}
	return nil
}
