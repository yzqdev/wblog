package models

import (
	"github.com/gookit/color"
	"github.com/rs/xid"
	"testing"
)

func TestCreateAdmin(t *testing.T) {
	uid := xid.New()
	color.Redln(uid.Machine())
	color.Redln(uid.Pid())
	color.Redln(uid)
}
