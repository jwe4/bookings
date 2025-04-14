package render

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/jwe4/bookings/internal/config"
	"github.com/jwe4/bookings/internal/models"
	"net/http"
	"os"
	"testing"
	"time"
)

var testSession *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	gob.Register(models.Reservation{})

	// change this to true when in production
	testApp.InProduction = false

	// set up the session
	testSession = scs.New()
	testSession.Lifetime = 24 * time.Hour
	testSession.Cookie.Persist = true
	testSession.Cookie.SameSite = http.SameSiteLaxMode
	testSession.Cookie.Secure = false

	testApp.Session = testSession

	app = &testApp

	os.Exit(m.Run())
}
