package main

import "testing"

// This is just for forcing main to execute therefore showing a higher coverage.
func TestMain(t *testing.T) {
	main()
}

func TestGetUser(t *testing.T) {
	mdb := &MockDatastore{
		Users: map[int]User{
			2: {ID: 2, Name: "Jenny"},
		},
	}
	srvc := &Service{
		ds: mdb,
	}

	/* Case OK */
	uok, err := srvc.GetUser(2)
	if err != nil {
		t.Errorf("Got Error: %v", err)
	}
	if uok.Name != "Jenny" {
		t.Errorf("got %v, want %v", uok.Name, "Jenny")
	}

	/* Case KO */
	_, err = srvc.GetUser(3)
	if err == nil {
		t.Errorf("got %v, want %v", nil, "User ID: 3 not found")
	}
}

func TestSaveUser(t *testing.T) {
	mdb := &MockDatastore{
		Users: make(map[int]User),
	}
	srvc := &Service{
		ds: mdb,
	}
	u1 := User{
		ID:   4,
		Name: "Aragon",
	}

	/* Case OK */
	err := srvc.SaveUser(u1)
	if err != nil {
		t.Errorf("Got Error: %v", err)
	}

	/* Case KO */
	err = srvc.SaveUser(u1)
	if err == nil {
		t.Errorf("got %v, want %v", nil, "User ID: 4 already exists")
	}
}
