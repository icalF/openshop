package usecases

type PaymentProofManager interface {
	UpdatePaymentProof(orderId int64, filename string) (bool, error)
}
