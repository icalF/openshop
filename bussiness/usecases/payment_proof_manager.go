package usecases

import "mime/multipart"

type PaymentProofManager interface {
	UpdatePaymentProof(orderId int64, file multipart.File) (bool, error)
}
