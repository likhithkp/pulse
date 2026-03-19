package other

import (
	"errors"
	"math/rand"
	"pulse/utils/config"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Utils struct{}

func NewUtils(utils *config.Env) *Utils {
	return &Utils{}
}

func (o *Utils) Generate6DigitOtp() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	randInt := r.Intn(900000) + 100000
	return strconv.Itoa(randInt)
}

func (o *Utils) MongoErrToGrpcErr(err error, notFoundMessage string, otherErrCode codes.Code, otherErrMessage string) error {
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return status.Error(codes.NotFound, notFoundMessage)
		} else {
			return status.Error(otherErrCode, otherErrMessage)
		}
	}
	return nil
}

func (o *Utils) GrpcErr(code codes.Code, message string) error {
	return status.Error(code, message)
}

func (o *Utils) Contains(slice []string, target string) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// func (o *Utils) SendNotificationToUser(
// 	deviceToken string,
// 	deviceType string,
// 	title string,
// 	body string,
// 	data map[string]string,
// 	userId primitive.ObjectID,
// ) error {
// 	var ctx = context.Background()

// 	var message *messaging.Message = &messaging.Message{
// 		Data: data,
// 		Notification: &messaging.Notification{
// 			Title: title,
// 			Body:  body,
// 		},
// 		Token: deviceToken,
// 	}
// 	_, err := o.firebase.MessagingClient().Send(ctx, message)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (o *Utils) SendScheduledNotification() error {
// 	var ctx = context.Background()

// 	date := time.Now().UTC().AddDate(0, 0, 3)
// 	dayString := date.Format("02")

// 	fmt.Println("dayString", dayString)

// 	condition := "'premium-user' in topics && '" + dayString + "' in topics"

// 	message := &messaging.Message{
// 		Data: map[string]string{
// 			"title": "Period Reminder",
// 			"body":  "Your period is expected to start in 3 days.",
// 		},
// 		Notification: &messaging.Notification{
// 			Title: "Period Reminder",
// 			Body:  "Your period is expected to start in 3 days.",
// 		},
// 		Condition: condition,
// 	}
// 	response, err := o.firebase.MessagingClient().Send(ctx, message)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("Successfully sent message:", response)
// 	return nil
// }

// func (o *Utils) SendNotificationToPremiumUser(title string,
// 	body string) error {
// 	var ctx = context.Background()

// 	date := time.Now().UTC().AddDate(0, 0, 3)
// 	dayString := date.Format("02")

// 	fmt.Println("dayString", dayString)

// 	topic := "premium-user"

// 	message := &messaging.Message{
// 		Data: map[string]string{
// 			"title": title,
// 			"body":  body,
// 		},
// 		Notification: &messaging.Notification{
// 			Title: title,
// 			Body:  body,
// 		},
// 		Topic: topic,
// 	}
// 	response, err := o.firebase.MessagingClient().Send(ctx, message)
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("Successfully sent message:", response)
// 	return nil
// }

func (o *Utils) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
