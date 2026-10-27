//go:build !js

// Package mobile is the gomobile binding of api.Api for Android and iOS.
//
// It differs from api.Api in three ways that native apps need:
//   - NewCore takes the data directory, because Android has no home directory
//     (Go would fall back to /sdcard, which apps cannot write).
//   - Errors are JSON objects {"code", "calendar", "message"}, because gomobile
//     only transports an error's message and would drop its errcode details.
//   - The time zone database is embedded, so iCalendar TZIDs always resolve.
//
// The WebAssembly build does not use this package.
package mobile

import (
	_ "time/tzdata" // DTSTART;TZID and iCalendar zones need a zone database

	"github.com/git-calendar/core/pkg/api"
	"github.com/git-calendar/core/pkg/filesystem"
)

// Core is the mobile-facing wrapper around api.Api.
type Core struct {
	inner *api.Api
}

// NewCore creates a core that stores its data under dataDir/.git-calendar-data.
// On Android, pass the app's files directory.
func NewCore(dataDir string) (*Core, error) {
	fs, err := filesystem.GetFSAt(dataDir)
	if err != nil {
		return nil, wrapErr(err)
	}
	return &Core{inner: api.NewApiWithFS(fs)}, nil
}

// ----------------------------------------------- Calendars ------------------------------------------------

func (c *Core) CreateCalendar(name, password string) error {
	return wrapErr(c.inner.CreateCalendar(name, password))
}

func (c *Core) RemoveCalendar(name string) error { return wrapErr(c.inner.RemoveCalendar(name)) }

func (c *Core) RenameCalendar(oldName, newName string) error {
	return wrapErr(c.inner.RenameCalendar(oldName, newName))
}

func (c *Core) LoadCalendars() error { return wrapErr(c.inner.LoadCalendars()) }

// SetTimeZone sets the IANA zone for date-only iCalendar events; Go's local zone is UTC on Android.
func (c *Core) SetTimeZone(name string) error { return wrapErr(c.inner.SetTimeZone(name)) }

func (c *Core) ListCalendars() (string, error) { return wrapStr(c.inner.ListCalendars()) }

func (c *Core) UpdateRemote(calendar, remoteURL string, readonly bool) error {
	return wrapErr(c.inner.UpdateRemote(calendar, remoteURL, readonly))
}

func (c *Core) CloneCalendar(repoURL, password string, readonly bool) error {
	return wrapErr(c.inner.CloneCalendar(repoURL, password, readonly))
}

func (c *Core) SyncAll() error { return wrapErr(c.inner.SyncAll()) }

// ----------------------------------------------- iCalendar ------------------------------------------------

func (c *Core) ImportICalURL(name, rawURL string) error {
	return wrapErr(c.inner.ImportICalURL(name, rawURL))
}

func (c *Core) UpdateICalURL(name, rawURL string) error {
	return wrapErr(c.inner.UpdateICalURL(name, rawURL))
}

func (c *Core) ImportICalFile(calendar, tagID, data string) error {
	return wrapErr(c.inner.ImportICalFile(calendar, tagID, data))
}

func (c *Core) ExportICal(calendar string) ([]byte, error) {
	return wrapBytes(c.inner.ExportICal(calendar))
}

// ----------------------------------------------- Backups --------------------------------------------------

func (c *Core) ExportZip(calendar string) ([]byte, error) {
	return wrapBytes(c.inner.ExportZip(calendar))
}

func (c *Core) RestoreZip(data []byte) error { return wrapErr(c.inner.RestoreZip(data)) }

// ------------------------------------------------ Events --------------------------------------------------

func (c *Core) CreateEvent(eventJSON string) (string, error) {
	return wrapStr(c.inner.CreateEvent(eventJSON))
}

func (c *Core) UpdateEvent(eventJSON string) (string, error) {
	return wrapStr(c.inner.UpdateEvent(eventJSON))
}

func (c *Core) UpdateRepeatingEvent(oldEventJSON, newEventJSON string, strategy int) (string, error) {
	return wrapStr(c.inner.UpdateRepeatingEvent(oldEventJSON, newEventJSON, strategy))
}

func (c *Core) RemoveEvent(eventJSON string) error { return wrapErr(c.inner.RemoveEvent(eventJSON)) }

func (c *Core) RemoveRepeatingEvent(eventJSON string, strategy int) error {
	return wrapErr(c.inner.RemoveRepeatingEvent(eventJSON, strategy))
}

func (c *Core) GetEvent(id string) (string, error) { return wrapStr(c.inner.GetEvent(id)) }

func (c *Core) GetEvents(from, to, filterJSON string) (string, error) {
	return wrapStr(c.inner.GetEvents(from, to, filterJSON))
}

func (c *Core) GetNotifications(from, to, filterJSON string) (string, error) {
	return wrapStr(c.inner.GetNotifications(from, to, filterJSON))
}

// ------------------------------------------------- Tags ---------------------------------------------------

func (c *Core) CreateTag(calendar, tagJSON string) (string, error) {
	return wrapStr(c.inner.CreateTag(calendar, tagJSON))
}

func (c *Core) UpdateTag(calendar, tagJSON string) (string, error) {
	return wrapStr(c.inner.UpdateTag(calendar, tagJSON))
}

func (c *Core) RemoveTag(calendar, id string) error { return wrapErr(c.inner.RemoveTag(calendar, id)) }
