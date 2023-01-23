package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gorilla/mux"
)

func createUserPool(w http.ResponseWriter, r *http.Request) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := cognitoidentityprovider.New(sess)

	input := &cognitoidentityprovider.CreateUserPoolInput{
		PoolName: aws.String("hello"),
	}

	result, err := svc.CreateUserPool(input)
	if err != nil {
		fmt.Println("Error creating user pool:", err)
		return
	}

	fmt.Println("Successfully created user pool:", result)
}

func readUserPool(w http.ResponseWriter, r *http.Request) {
	// Create a new Cognito client
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)
	cognitoClient := cognitoidentityprovider.New(sess)

	// Get the user pool ID from the request's path variables
	vars := mux.Vars(r)
	userPoolID := vars["userPoolID"]

	// Define input parameters for the DescribeUserPool function
	input := &cognitoidentityprovider.DescribeUserPoolInput{
		UserPoolId: aws.String(userPoolID),
	}

	// Call the DescribeUserPool function on the Cognito client
	result, err := cognitoClient.DescribeUserPool(input)
	if err != nil {
		fmt.Fprintf(w, "Error reading User Pool: %v", err)
		return
	}

	json.NewEncoder(w).Encode(result.UserPool)
}

func updateUserPool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userPoolID := vars["userPoolID"]

	// Create a new session and a new Cognito Identity Provider client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := cognitoidentityprovider.New(sess)

	// Update the user pool with the new settings
	input := &cognitoidentityprovider.UpdateUserPoolInput{
		UserPoolId: aws.String(userPoolID),
		// Add the new settings here
	}

	_, err := svc.UpdateUserPool(input)
	if err != nil {
		http.Error(w, "Failed to update user pool: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User pool updated successfully"))
}

func deleteUserPool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userPoolID := vars["userPoolID"]

	// Create a new session and a new Cognito Identity Provider client
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := cognitoidentityprovider.New(sess)

	// Delete the user pool
	input := &cognitoidentityprovider.DeleteUserPoolInput{
		UserPoolId: aws.String(userPoolID),
	}

	_, err := svc.DeleteUserPool(input)
	if err != nil {
		http.Error(w, "Failed to delete user pool: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User pool deleted successfully"))
}
