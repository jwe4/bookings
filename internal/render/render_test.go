package render

import (
	"fmt"
	"github.com/jwe4/bookings/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Fatalf("Failed to get session: %v", err)
	}

	testSession.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)
	if result.Flash != "123" {
		t.Error("flash value of 123 not found in session")
	}

}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, err = testSession.Load(ctx, r.Header.Get("X-Session"))
	if err != nil {
		return nil, fmt.Errorf("failed to load session: %w", err)
	}
	r = r.WithContext(ctx)
	return r, nil
}
