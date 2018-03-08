package TomlConfiguration

import (
	"testing"
)

type Server struct {
	Name     string `default:"123123"`
	Port     int    `default:"6060"`
	Enabled  bool
	Users    []string
	Postgres Postgres
}
type Postgres struct {
	Enabled           bool
	Port              int
	Hosts             []string
	DBName            string
	AvailabilityRatio float64
}

func getDefaultServer() *Server {
	return &Server{
		Name:    "koding",
		Port:    6060,
		Enabled: true,
		Users:   []string{"ankara", "istanbul"},
		Postgres: Postgres{
			Enabled:           true,
			Port:              5432,
			Hosts:             []string{"192.168.2.1", "192.168.2.2", "192.168.2.3"},
			DBName:            "configdb",
			AvailabilityRatio: 8.23,
		},
	}
}
func TestToml(t *testing.T) {
	m := TOMLLoader{Path: "conf/conf.toml"}

	s := &Server{}
	if err := m.Load(s); err != nil {
		t.Error(err)
	}
	testStruct(t, s, getDefaultServer())
}

func testStruct(t *testing.T, s *Server, d *Server) {
	if s.Name != d.Name {
		t.Errorf("Name value is wrong: %s, want: %s", s.Name, d.Name)
	}

	if s.Port != d.Port {
		t.Errorf("Port value is wrong: %d, want: %d", s.Port, d.Port)
	}

	if s.Enabled != d.Enabled {
		t.Errorf("Enabled value is wrong: %t, want: %t", s.Enabled, d.Enabled)
	}

	if len(s.Users) != len(d.Users) {
		t.Errorf("Users value is wrong: %d, want: %d", len(s.Users), len(d.Users))
	} else {
		for i, user := range d.Users {
			if s.Users[i] != user {
				t.Errorf("User is wrong for index: %d, user: %s, want: %s", i, s.Users[i], user)
			}
		}
	}

	testPostgres(t, s.Postgres, d.Postgres)
}
func testPostgres(t *testing.T, s Postgres, d Postgres) {
	if s.Enabled != d.Enabled {
		t.Errorf("Postgres enabled is wrong %t, want: %t", s.Enabled, d.Enabled)
	}

	if s.Port != d.Port {
		t.Errorf("Postgres Port value is wrong: %d, want: %d", s.Port, d.Port)
	}

	if s.DBName != d.DBName {
		t.Errorf("DBName is wrong: %s, want: %s", s.DBName, d.DBName)
	}

	if s.AvailabilityRatio != d.AvailabilityRatio {
		t.Errorf("AvailabilityRatio is wrong: %f, want: %f", s.AvailabilityRatio, d.AvailabilityRatio)
	}

	if len(s.Hosts) != len(d.Hosts) {
		// do not continue testing if this fails, because others is depending on this test
		t.Fatalf("Hosts len is wrong: %v, want: %v", s.Hosts, d.Hosts)
	}

	for i, host := range d.Hosts {
		if s.Hosts[i] != host {
			t.Fatalf("Hosts number %d is wrong: %v, want: %v", i, s.Hosts[i], host)
		}
	}
}
