package orders

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"
	OrderStatusPaid      OrderStatus = "PAID"
	OrderStatusShipped   OrderStatus = "SHIPPED"
	OrderStatusDelivered OrderStatus = "DELIVERED"
	OrderStatusCancelled OrderStatus = "CANCELLED"
)

type PaymentMethod string

const (
	PaymentMethodCreditCard   PaymentMethod = "CREDIT_CARD"
	PaymentMethodDebitCard    PaymentMethod = "DEBIT_CARD"
	PaymentMethodBankTransfer PaymentMethod = "BANK_TRANSFER"
)

type PaymentStatus string

const (
	PaymentStatusPending PaymentStatus = "PENDING"
	PaymentStatusSuccess PaymentStatus = "SUCCESS"
	PaymentStatusFailed  PaymentStatus = "FAILED"
)
