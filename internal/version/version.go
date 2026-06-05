package version

const Name = "creamcal" // TODO: rename

// Version current version of the app.
// NOTE: propagated during complication:
//
//	go build -ldflags="-X 'codeberg.org/liampas/calendar/internal/version.Version=v1.0.0'"
const Version = "dev"
