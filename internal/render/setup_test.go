package render

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"github.com/jwe4/bookings/internal/config"
	"github.com/jwe4/bookings/internal/models"
	"log"
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
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERRORt", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog

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

type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {}

func (tw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}
