package env

var OrderCreatedQueueName string

func Initialise() {
	// OrderCreatedQueueName = os.Getenv("SQS_ORDER_CREATED_QUEUE_URL")
	OrderCreatedQueueName = "checkout_order_created"

}
