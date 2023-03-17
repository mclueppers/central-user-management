package main

import (
	"flag"
	"fmt"
	"log"

	"cum/storage"
	"cum/types"
)

var (
	// Version is the current version of the application
	Version = "0.0.1"

	// Build is the build number of the application
	Build = "0"

	// BuildDate is the date when the application was built
	BuildDate = "1970-01-01T00:00:00Z"

	// GitCommit is the git commit hash of the application
	GitCommit = "0000000000000000000000000000000000000000"

	// GitBranch is the git branch of the application
	GitBranch = "master"

	// GitState is the git state of the application
	GitState = "clean"

	// GitSummary is the git summary of the application
	GitSummary = "0000000 master"

	// InMemory is a flag to use the in-memory storage
	InMemory = flag.Bool("in-memory", false, "Use the in-memory storage")

	// Postgres is a flag to use the PostgreSQL storage
	Postgres = flag.Bool("postgres", false, "Use the PostgreSQL storage")

	// PostgresHost is a flag to set the PostgreSQL host
	PostgresHost = flag.String("postgres-host", "localhost", "PostgreSQL host")

	// PostgresPort is a flag to set the PostgreSQL port
	PostgresPort = flag.Int("postgres-port", 5432, "PostgreSQL port")

	// PostgresUser is a flag to set the PostgreSQL user
	PostgresUser = flag.String("postgres-user", "postgres", "PostgreSQL user")

	// PostgresPassword is a flag to set the PostgreSQL password
	PostgresPassword = flag.String("postgres-password", "postgres", "PostgreSQL password")

	// PostgresDatabase is a flag to set the PostgreSQL database
	PostgresDatabase = flag.String("postgres-database", "cum", "PostgreSQL database")

	// PostgresSSLMode is a flag to set the PostgreSQL SSL mode
	PostgresSSLMode = flag.String("postgres-ssl-mode", "disable", "PostgreSQL SSL mode")

	// PostgresMaxIdleConnections is a flag to set the PostgreSQL max idle connections
	PostgresMaxIdleConnections = flag.Int("postgres-max-idle-connections", 10, "PostgreSQL max idle connections")

	// PostgresMaxOpenConnections is a flag to set the PostgreSQL max open connections
	PostgresMaxOpenConnections = flag.Int("postgres-max-open-connections", 10, "PostgreSQL max open connections")

	// VersionFlag is a flag to print the version of the application
	VersionFlag = flag.Bool("version", false, "Print the version of the application")

	// HelpFlag is a flag to print the help of the application
	HelpFlag = flag.Bool("help", false, "Print the help of the application")

	// postgresStorage is the PostgreSQL storage
	postgresStorage *storage.PostgresStorage
)

func version() {
	fmt.Printf("Version: %s Build: %s BuildDate: %s GitCommit: %s GitBranch: %s, GitSummary: %s", Version, Build, BuildDate, GitCommit, GitBranch, GitSummary)
}

func help() {
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if *VersionFlag {
		version()
		return
	}

	if *HelpFlag {
		help()
		return
	}

	// Initialize the storage
	var myStorage types.Storage
	var err error

	if *InMemory {
		inMemoryStorage := storage.NewInMemoryStorage()
		myStorage, err = types.NewStorage(inMemoryStorage)
		if err != nil {
			log.Fatalf("Failed to initialize the in-memory storage: %v", err)
		}
	} else if *Postgres {
		postgresStorage, err = storage.NewPostgresStorage(&storage.PostgresStorageConfig{
			Host:               *PostgresHost,
			Port:               *PostgresPort,
			User:               *PostgresUser,
			Password:           *PostgresPassword,
			Database:           *PostgresDatabase,
			SSLMode:            *PostgresSSLMode,
			MaxIdleConnections: *PostgresMaxIdleConnections,
			MaxOpenConnections: *PostgresMaxOpenConnections,
		})
		if err != nil {
			log.Fatalf("Failed to initialize the PostgreSQL storage: %v", err)
		}
		myStorage, err = types.NewStorage(postgresStorage)
		if err != nil {
			log.Fatalf("Failed to initialize the PostgreSQL storage: %v", err)
		}
	} else {
		log.Fatal("No storage specified")
	}

	// Close storage when the application exits
	defer myStorage.Close()

	// Create a new user
	user := &types.User{
		ID:       "user1",
		Username: "johndoe",
		Email:    "john@example.com",
		Password: "password",
	}
	err = myStorage.CreateUser(user)
	if err != nil {
		if err != types.ErrUserAlreadyExists {
			// Handle error
			log.Printf("Failed to create user: %v", err)
		}
	}

	// Create a new user
	user2 := &types.User{
		ID:       "user2",
		Username: "johndoe2",
		Email:    "johndoe@example.com",
	}

	err = myStorage.CreateUser(user2)
	if err != nil {
		log.Printf("Failed to create user: %v", err)
	}

	// Get a user by ID
	u, err := myStorage.GetUserByID("user1")
	if err != nil {
		log.Fatalf("Failed to get user by ID: %v", err)
	}
	fmt.Println(u)

	// Get a user by username
	u, err = myStorage.GetUserByUsername("johndoe")
	if err != nil {
		log.Fatalf("Failed to get user by username: %v", err)
	}
	fmt.Println(u)

	// Get a user by email
	u, err = myStorage.GetUserByEmail("john@example.com")
	if err != nil {
		log.Fatalf("Failed to get user by email: %v", err)
	}
	fmt.Println(u)

	// Update a user
	user.Email = "john.doe@example.com"
	err = myStorage.UpdateUser(user)
	if err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}

	// Get a user by ID after updating
	u, err = myStorage.GetUserByID("user1")
	if err != nil {
		log.Fatalf("Failed to get user by ID after updating: %v", err)
	}
	fmt.Println(u)

	// Create a new group
	group := &types.Group{
		ID:          "group1",
		Name:        "my-group",
		Description: "My group description",
	}
	err = myStorage.CreateGroup(group)
	if err != nil {
		log.Fatalf("Failed to create group: %v", err)
	}

	// Create a new group
	group2 := &types.Group{
		ID:   "group2",
		Name: "Developers",
	}
	err = myStorage.CreateGroup(group2)
	if err != nil {
		log.Fatalf("Failed to create group: %v", err)
	}

	// Get a group by ID
	g, err := myStorage.GetGroupByID("group1")
	if err != nil {
		log.Fatalf("Failed to get group by ID: %v", err)
	}
	fmt.Println(g)

	// Get a group by name
	g, err = myStorage.GetGroupByName("my-group")
	if err != nil {
		log.Fatalf("Failed to get group by name: %v", err)
	}
	fmt.Println(g)

	// Update a group
	group.Description = "My updated group description"
	err = myStorage.UpdateGroup(group)
	if err != nil {
		log.Fatalf("Failed to update group: %v", err)
	}

	// Get a group by ID after updating
	g, err = myStorage.GetGroupByID("group1")
	if err != nil {
		log.Fatalf("Failed to get group by ID after updating: %v", err)
	}
	fmt.Printf("Updated group: %v\n", g)

	// Assign user1 to group1
	err = myStorage.AddMemberToGroup(user, group.ID)
	if err != nil {
		log.Fatalf("Failed to add %s to %s: %v", user.ID, group.ID, err)
	}

	// Assign user2 to group2
	err = myStorage.AddMemberToGroup(user2, group2.ID)
	if err != nil {
		log.Fatalf("Failed to add %s to %s: %v", user2.ID, group2.ID, err)
	}

	// Assign user2 to group2
	err = myStorage.AddMemberToGroup(group2, group.ID)
	if err != nil {
		log.Fatalf("Failed to add %s to %s: %v", group2.ID, group.ID, err)
	}

	fmt.Println(myStorage)
	// Delete a user
	err = myStorage.DeleteUser("user1")
	if err != nil {
		log.Fatalf("Failed to delete user: %v", err)
	}

	// Delete a group
	err = myStorage.DeleteGroup(group)
	if err != nil {
		log.Fatalf("Failed to delete group: %v", err)
	}
	fmt.Println("Deleted group")

	fmt.Println(myStorage)
}
