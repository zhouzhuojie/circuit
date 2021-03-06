package database

import (
	"errors"

	"github.com/codeamp/circuit/plugins"
	log "github.com/codeamp/logger"
	"github.com/codeamp/transistor"
)

// Database
type Database struct {
	events chan transistor.Event
}

// init
func init() {
	transistor.RegisterPlugin("database", func() transistor.Plugin {
		return &Database{}
	}, plugins.ProjectExtension{})
}

// Description
func (x *Database) Description() string {
	return "The plugin to install databases"
}

// SampleConfig
func (x *Database) SampleConfig() string {
	return ` `
}

// Start
func (x *Database) Start(e chan transistor.Event) error {
	x.events = e
	log.Info("Started Database")
	return nil
}

// Stop
func (x *Database) Stop() {
	log.Info("Stopping Database")
}

// Subscribe
func (x *Database) Subscribe() []string {
	return []string{
		"project:database:create",
		"project:database:delete",
	}
}

func (x *Database) sendFailedStatusEvent(err error) {
	ev := transistor.NewEvent(plugins.GetEventName("project:database"), transistor.GetAction("status"), nil)
	ev.State = transistor.GetState("failed")
	ev.StateMessage = err.Error()
	x.events <- ev
}

// Process
func (x *Database) Process(e transistor.Event) error {
	log.Debug("Processing database event")

	// confirm all required structured inputs (Project slug, Environment)
	if e.PayloadModel != "plugins.ProjectExtension" {
		x.sendFailedStatusEvent(errors.New("invalid payload model"))
		return nil
	}

	projectExtensionEvent := e.Payload.(plugins.ProjectExtension)

	instanceEndpoint, err := e.GetArtifact("SHARED_DATABASE_HOST")
	if err != nil {
		x.sendFailedStatusEvent(err)
		return nil
	}

	instanceUsername, err := e.GetArtifact("SHARED_DATABASE_ADMIN_USERNAME")
	if err != nil {
		x.sendFailedStatusEvent(err)
		return nil
	}

	instancePassword, err := e.GetArtifact("SHARED_DATABASE_ADMIN_PASSWORD")
	if err != nil {
		x.sendFailedStatusEvent(err)
		return nil
	}

	instancePort, err := e.GetArtifact("SHARED_DATABASE_PORT")
	if err != nil {
		x.sendFailedStatusEvent(err)
		return nil
	}

	dbType, err := e.GetArtifact("DB_TYPE")
	if err != nil {
		x.sendFailedStatusEvent(err)
		return nil
	}

	dbInstance, err := initDBInstance(dbType.String(), instanceEndpoint.String(), instanceUsername.String(), instancePassword.String(), instancePort.String())
	if err != nil {
		x.sendFailedStatusEvent(err)
		return nil
	}

	// Create DB within shared instance of the correct db variant (postgres/mysql)
	switch e.Action {
	case transistor.GetAction("create"):
		dbUsername := genDBUser(projectExtensionEvent)
		dbName := genDBName(projectExtensionEvent)
		dbPassword := genDBPassword()

		dbMetadata, err := (*dbInstance).CreateDatabaseAndUser(dbName, dbUsername, dbPassword)
		if err != nil {
			x.sendFailedStatusEvent(err)
			return nil
		}

		// store db metadata into instance
		respEvent := transistor.NewEvent(e.Name, transistor.GetAction("status"), nil)
		respEvent.State = transistor.GetState("complete")

		respEvent.AddArtifact("DB_USER", dbMetadata.Credentials.Username, false)
		respEvent.AddArtifact("DB_PASSWORD", dbMetadata.Credentials.Password, false)
		respEvent.AddArtifact("DB_NAME", dbMetadata.Name, false)
		respEvent.AddArtifact("DB_ENDPOINT", (*dbInstance).GetInstanceMetadata().Endpoint, false)
		respEvent.AddArtifact("DB_PORT", (*dbInstance).GetInstanceMetadata().Port, false)

		x.events <- respEvent
	case transistor.GetAction("delete"):
		dbName, err := e.GetArtifact("DB_NAME")
		if err != nil {
			x.sendFailedStatusEvent(err)
			return nil
		}

		dbUser, err := e.GetArtifact("DB_USER")
		if err != nil {
			x.sendFailedStatusEvent(err)
			return nil
		}

		if err = (*dbInstance).DeleteDatabaseAndUser(dbName.String(), dbUser.String()); err != nil {
			x.sendFailedStatusEvent(err)
			return nil
		}

		// store db metadata into instance
		respEvent := transistor.NewEvent(e.Name, transistor.GetAction("status"), nil)
		respEvent.State = transistor.GetState("complete")

		x.events <- respEvent
	}

	return nil
}
