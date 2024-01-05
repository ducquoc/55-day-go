// main.go
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/appleboy/go-fcm"
	//"github.com/OneSignal/onesignal-go-api"
	"github.com/go-chi/chi/v5"
)

const APP_ID = "b40bf034-b84b-4122-a7e6-865e9142f5ee"              // "910e2386-d779-408e-bb95-d0ed80453626"
const API_KEY = "MWI0MzJhNmItZDZmNi00NTYwLWI4MTctM2ZkYTNlYTc1OWZi" // "ZGQwNjkwOGEtNjgzNS00YjkyLTlmMzMtOGUyZTU2OGJjNmJj"
const ABTEST_ID = "910e2386-d779-408e-bb95-d0ed80453626"           // AB Testing
const ABTEST = "ZGQwNjkwOGEtNjgzNS00YjkyLTlmMzMtOGUyZTU2OGJjNmJj"  // AB Testing
const F_SERVER_KEY = "AAAAWZeZA2Q:APA91bF87EjWJxE5v6wJHrwS_5XayelaLP_pjJJ0_s6w_QHYSuEe18OBbKLsJD-n0dXpEiRpNHYQ9mTZsDbeqf6yFO3Rso9rVqUgG2vJnPUj3hj1-y_ScS8ANP_pJQeNwIK0T7X6bo8o"
const F_REG_TOKEN = "dXWz4A3SQpGV4n9VJbz5zN:APA91bEHcpsIay8mRMjTF4mHXFATRElq5DDb-L1oy4WDb86cuGkcFocTRWU0BFCL8wauu5u-KZHiJHh5g2x8cbDtr7K6gicCFz7U6MmRBVMtlQ3zz1XT8DS509j6P4oKMjnoffDSTLX8"
const F_API_KEY = "AIzaSyBUyXtDLB657j7bAKdEPFtO0nEm7xqiXLE"
const F_OAUTH2_TOKEN = "AAAAWZeZA2Q:APA91bH-4mUl6iIk8-I7pgWZaEjwchSNkFkRYe3yn5JYkmnxkioJr7VTcqdqXr0JpKma2n8X66IPrDMcXlFsUzSk-C3HKGZW9A9SztbM2onhSDBF6mw_IOeMSDDySKorm5emeC-bVMpY"

func main() {
	r := chi.NewRouter()

	r.Post("/notifications/push", pushAdditionalData)
	r.Post("/notifications/push/silent", pushSilent)

	port := ":8080"
	fmt.Printf("Server listening on port %s\n", port)
	http.ListenAndServe(port, r)
}

func pushAdditionalData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	notification := createBaseNotification()
	notification.Notification.Body = "Data Test at " + time.Now().String()
	notification.Data = map[string]interface{}{
		"job":        "sync-certificates", // flag to display on mobile app
		"student_id": "910e2386-d779-408e-bb95-d0ed80453626",
		"email":      "ducquoc.vn@gmail.com",
	}
	internalPushNotification(notification)
}

func createBaseNotification() *fcm.Message {
	msg := &fcm.Message{
		To: F_REG_TOKEN,
		Notification: &fcm.Notification{
			Title: "DTest title",
			Body:  "DTest body",
		},
		Data: map[string]interface{}{
			//"priority": "high",
		},
	}
	return msg
}

func internalPushNotification(msg *fcm.Message) {
	// Create a FCM client to send the message.
	client, err := fcm.NewClient(F_SERVER_KEY)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
		fmt.Fprintf(os.Stdout, "Error creating FCM client: %v\n", err)
	}

	// Send the message and receive the response without retries.
	response, err := client.Send(msg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending to FCM: %v\n", err)
	}
	fmt.Fprintf(os.Stdout, "Response from FCM API: %v\n", response)
}

func pushSilent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	notification := createBaseNotification() // empty data
	//if notification.Notification != nil {
	//	notification.Notification.Sound = "true"
	//}
	notification.Notification.Body = "Generic noti test at " + time.Now().String()
	internalPushNotification(notification)
}

// ToPointer is a helper function to create a pointer to a value.
// x := &5 doesn't compile
// x := ToPointer(5) good.
func toPointer[T any](p T) *T {
	return &p
}
