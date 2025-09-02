package main

import "testing"

func Test_createUser(t *testing.T) {
	users = make(map[string]string) // Make data clear

	// Test inserting new user
	err := createUser("Adolf")
	t.Log(users)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if _, exists := users["Adolf"]; !exists {
		t.Error("Expected user 'Adolf' to be created")
	}

	// Test inserting existing user
	err = createUser("Adolf")
	if err == nil {
		t.Error("Expected error for existing user, got nil")
	}
	expectedErr := "User already exist!!"
	if err.Error() != expectedErr {
		t.Errorf("Expected error %q, got %q", expectedErr, err.Error())
	}
}
