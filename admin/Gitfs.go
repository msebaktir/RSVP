package admin

import (
	models "msebaktir.com/LCV/Models"
	repository "msebaktir.com/LCV/utils"
)

func CreateGift(Name string, Description string, Price float64) int {
	db, err := repository.Connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var gift models.Gift
	gift.Description = Description
	gift.Name = Name
	gift.Price = Price

	err, giftId := repository.CreateGift(db, gift)

	if err != nil {
		panic(err)
	}
	return giftId
}
