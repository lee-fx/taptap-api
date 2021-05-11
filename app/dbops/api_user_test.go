package dbops

import "testing"

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testAddUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("lfxpupa", "123")
	if err != nil {
		t.Errorf("Error of AddUser : %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("lfxpupa")
	if pwd != "123" && err != nil {
		t.Errorf("Error of GetUser ")
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("lfxpupa", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser : %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("lfxpupa")
	if err != nil {
		t.Errorf("Error of RegetUser : %v", err)
	}

	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}
