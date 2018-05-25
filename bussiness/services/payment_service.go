package services

import (
	"errors"

	"github.com/koneko096/openshop/models/datamodels"
	"github.com/koneko096/openshop/dao"
	"github.com/koneko096/openshop/bussiness/usecases"

	"encoding/binary"
	"encoding/base64"
	"os"
	"io"
	"mime/multipart"
)


func NewPaymentManager(dao dao.PaymentDAO) usecases.PaymentManager {
	return &paymentService{
		dao: dao,
	}
}

func NewPaymentProofManager(dao dao.PaymentDAO) usecases.PaymentProofManager {
	return &paymentService{
		dao: dao,
	}
}


type paymentService struct {
	dao dao.PaymentDAO
}



func (s *paymentService) GetAll() []datamodels.Payment {
	return s.dao.SelectMany(map[string]string{}, 0)
}

func (s *paymentService) GetByID(id int64) (datamodels.Payment, bool) {
	return s.dao.Select(map[string]string{
		"id": string(id),
	})
}

func (s *paymentService) GetByOrderID(orderId int64) (datamodels.Payment, bool) {
	return s.dao.Select(map[string]string{
		"order_id": string(orderId),
	})
}

func (s *paymentService) InsertOrUpdate(payment datamodels.Payment) (datamodels.Payment, error) {
	return s.dao.InsertOrUpdate(payment)
}

func (s *paymentService) UpdatePaymentProof(orderId int64, file multipart.File) (bool, error) {
	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs, uint64(orderId))
	filename := base64.StdEncoding.EncodeToString(bs)

	out, err := os.OpenFile(
		"./uploads/"+filename,
		os.O_WRONLY|os.O_CREATE,
		0666,
	)
	if err != nil {
		return false, errors.New(
			"error while saving\n" + err.Error(),
		)
	}
	defer out.Close()

	io.Copy(out, file)

	payment, found := s.dao.Select(map[string]string{
		"order_id": string(orderId),
	})
	if !found {
		return false, errors.New("order ID not found")
	}

	payment.Proof = filename
	_, err = s.InsertOrUpdate(payment)
	return true, err
}

func (s *paymentService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}
