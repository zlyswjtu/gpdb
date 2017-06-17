package commands

import (
	"database/sql/driver"
	"gp_upgrade/db"
	"gp_upgrade/test_utils"

	"encoding/json"
	"errors"
	"gp_upgrade/config"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	"gp_upgrade/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var _ = Describe("check tests", func() {

	var (
		mock          sqlmock.Sqlmock
		dbConn        *db.DBConn
		subject       CheckCommand
		save_home_dir string
		fixture_path  string
	)

	BeforeEach(func() {
		_, this_file_path, _, _ := runtime.Caller(0)
		fixture_path = path.Join(path.Dir(this_file_path), "fixtures")

		save_home_dir = test_utils.ResetTempHomeDir()
		dbConn, mock = test_utils.CreateMockDBConn("localhost", 5432)
		subject = CheckCommand{}
	})

	AfterEach(func() {
		os.Setenv("HOME", save_home_dir)
	})

	Describe("check", func() {
		Describe("happy: the database is running, master-host is provided, and connection is successful", func() {
			It("writes a file to ~/.gp_upgrade/cluster_config.json with correct json", func() {
				setupSegmentConfigInDB(mock)
				buffer := gbytes.NewBuffer()

				err := subject.execute(dbConn, buffer, config.NewWriter())

				buffer.Close()
				Expect(err).ToNot(HaveOccurred())
				Expect(dbConn.Conn.Stats().OpenConnections).To(Equal(0))
				content, err := ioutil.ReadFile(config.GetConfigFilePath())
				test_utils.Check("cannot read file", err)

				resultData := make([]map[string]interface{}, 0)
				expectedData := make([]map[string]interface{}, 0)
				json.Unmarshal(content, resultData)
				json.Unmarshal([]byte(expected_check_configuration_output), expectedData)
				Expect(expectedData).To(Equal(resultData))
			})
		})

		Describe("errors", func() {
			Describe("when the query fails on AO table count", func() {

				It("returns an error", func() {
					mock.ExpectQuery(SELECT_SEGMENT_CONFIG_QUERY).WillReturnError(errors.New("the query has failed"))

					err := subject.execute(dbConn, nil, config.NewWriter())

					Expect(dbConn.Conn.Stats().OpenConnections).To(Equal(0))
					Expect(err).To(HaveOccurred())
				})
			})

			Describe("when the db dbConn fails", func() {
				It("returns an error", func() {
					subject.Database_name = "invalidDBthatnobodywouldeverhave"
					subject.Master_host = "localhost"

					err := subject.Execute(nil)

					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("pq: database \"invalidDBthatnobodywouldeverhave\" does not exist"))
				})
			})
			Describe("when the home directory is not writable", func() {
				It("returns an error", func() {
					setupSegmentConfigInDB(mock)
					buffer := gbytes.NewBuffer()
					err := os.MkdirAll(config.GetConfigDir(), 0500)
					test_utils.Check("cannot chmod: ", err)
					subject.Master_host = "localhost"

					err = subject.execute(dbConn, buffer, config.NewWriter())

					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("open /tmp/gp_upgrade_test_temp_home_dir/.gp_upgrade/cluster_config.json: permission denied"))
				})
			})

			Describe("when db result cannot be parsed", func() {
				It("returns an error", func() {
					setupSegmentConfigInDB(mock)
					mock.ExpectQuery(SELECT_SEGMENT_CONFIG_QUERY).WillReturnError(errors.New("the query has failed"))
					buffer := gbytes.NewBuffer()
					subject.Master_host = "localhost"

					fake := FakeWriter{}
					err := subject.execute(dbConn, buffer, fake)

					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("I always fail"))
				})
			})
		})
	})
})

type FakeWriter struct{}

func (writer FakeWriter) Load(rows utils.RowsWrapper) error {
	return errors.New("I always fail")
}

func (writer FakeWriter) Write() error {
	return errors.New("I always fail")
}
func setupSegmentConfigInDB(mock sqlmock.Sqlmock) {
	header := []string{"dbid", "content", "role", "preferred_role", "mode", "status", "port",
		"hostname", "address", "datadir"}
	fakeConfigRow := []driver.Value{1, -1, 'p', 'p', 's', 'u', 15432, "office-5-231.pa.pivotal.io",
		"office-5-231.pa.pivotal.io", nil}
	fakeConfigRow2 := []driver.Value{2, 0, 'p', 'p', 's', 'u', 25432, "office-5-231.pa.pivotal.io",
		"office-5-231.pa.pivotal.io", nil}
	rows := sqlmock.NewRows(header)
	heapfakeResult := rows.AddRow(fakeConfigRow...).AddRow(fakeConfigRow2...)
	mock.ExpectQuery(SELECT_SEGMENT_CONFIG_QUERY).WillReturnRows(heapfakeResult)
}

const (
	expected_check_configuration_output = `[
	{
	  "address": "office-5-231.pa.pivotal.io",
	  "content": -1,
	  "datadir": null,
	  "dbid": 1,
	  "hostname": "office-5-231.pa.pivotal.io",
	  "mode": 115,
	  "port": 15432,
	  "preferred_role": 112,
	  "role": 112,
	  "status": 117
	},
	{
	  "address": "office-5-231.pa.pivotal.io",
	  "content": 0,
	  "datadir": null,
	  "dbid": 2,
	  "hostname": "office-5-231.pa.pivotal.io",
	  "mode": 115,
	  "port": 25432,
	  "preferred_role": 112,
	  "role": 112,
	  "status": 117
	}
	]`

	SELECT_SEGMENT_CONFIG_QUERY = "select dbid, content.*"
)
