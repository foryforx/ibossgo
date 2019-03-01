package config

import (
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
)

// Configuration contains external config settings
type Configuration struct {
	ApplicationPort         string
	Host                    string
	Port                    string
	Name                    string
	User                    string
	Password                string
	Type                    string
	SSLMode                 string
	MaxDBConnections        int
	AcquireConnTimeout      int // in seconds
	MailAPIKey              string
	BaseURL                 string
	BasePort                string
	BaseProtocol            string
	S3Region                string
	S3Endpoint              string
	S3AccessKey             string
	S3SecretKey             string
	S3ReportBucket          string
	S3UserImageBucket       string
	S3UserFileBucket        string
	JwtSecret               string
	ModuleVersionExportKey  string
	ValidatorServerHostname string
	ValidatorDebounce       int
	DocxWorkDir             string
	TempFolderPath          string
}

// GetConfiguration instantiates Configuration
func GetConfiguration(isProduction bool) *Configuration {

	loadEnvOrExit := func(varName string) string {
		envVar := os.Getenv(varName)

		if envVar == "" {
			log.Fatal("Env var ", varName, " is required but not defined")
		}

		return envVar
	}

	loadEnvOrDefault := func(varName string, defaultValue string) string {
		envVar := os.Getenv(varName)

		if envVar == "" {
			log.Info("Using default value ", defaultValue, " for ", varName)

			return defaultValue
		}

		return envVar
	}

	conf := &Configuration{}

	conf.ApplicationPort = loadEnvOrDefault("PORT", "9090")

	conf.Host = loadEnvOrExit("PGHOST")
	conf.Port = loadEnvOrExit("PGPORT")
	conf.Name = loadEnvOrExit("PGDATABASE")
	conf.User = loadEnvOrExit("PGUSER")
	conf.Password = loadEnvOrDefault("PGPASSWORD", "")
	conf.Type = "postgresql"
	conf.SSLMode = loadEnvOrExit("PGSSLMODE")
	conf.MaxDBConnections, _ = strconv.Atoi(loadEnvOrDefault("PGMAXCONNECTIONS", "20"))
	conf.AcquireConnTimeout, _ = strconv.Atoi(loadEnvOrDefault("DB_CONN_ACQUIRE_TIMEOUT", "30"))

	// conf.MailAPIKey = loadEnvOrExit("SENDGRID_API_KEY")
	conf.BaseURL = loadEnvOrExit("BASE_URL")
	conf.BaseProtocol = loadEnvOrExit("BASE_PROTOCOL")

	// conf.S3Region = loadEnvOrExit("AWS_DEFAULT_REGION")
	// conf.S3Endpoint = loadEnvOrDefault("S3_ENDPOINT", fmt.Sprintf("https://s3.%v.amazonaws.com", conf.S3Region))
	// conf.S3AccessKey = loadEnvOrExit("AWS_ACCESS_KEY_ID")
	// conf.S3SecretKey = loadEnvOrExit("AWS_SECRET_ACCESS_KEY")
	// conf.S3ReportBucket = loadEnvOrExit("S3_REPORT_BUCKET")
	// conf.S3UserImageBucket = loadEnvOrExit("S3_USER_IMAGE_BUCKET")
	// conf.S3UserFileBucket = loadEnvOrExit("S3_USER_FILE_BUCKET")

	// conf.JwtSecret = loadEnvOrExit("JWT_SECRET")

	// conf.ModuleVersionExportKey = loadEnvOrExit("MODULE_VERSION_EXPORT_KEY")

	// conf.ValidatorServerHostname = loadEnvOrExit("VALIDATOR_HOSTNAME")
	// conf.ValidatorDebounce, _ = strconv.Atoi(loadEnvOrDefault("VALIDATOR_DEBOUNCE_MILLISECONDS", "500"))

	conf.TempFolderPath = "/tmp"
	if runtime.GOOS == "windows" {
		conf.TempFolderPath = os.Getenv("temp")
	} else if runtime.GOOS == "darwin" {
		conf.TempFolderPath = os.Getenv("TMPDIR")
	}

	return conf
}

// // GetS3Config returns an instance of S3Config
// func (c *Configuration) GetS3Config() *s3.Config {
// 	return &s3.Config{
// 		c.S3AccessKey,
// 		c.S3SecretKey,
// 		c.S3ReportBucket,
// 		c.S3UserImageBucket,
// 		c.S3UserFileBucket,
// 		c.S3Region,
// 	}
// }

// Print logs current configuration to stdout
func (c *Configuration) Print() {
	log.Info("Loaded configuration with settings")
	log.Info("Host: ", c.Host)
	log.Info("Port: ", c.Port)
	log.Info("User: ", c.User)
	log.Info("Password: ", len(c.Password), " characters")
	log.Info("Type: ", c.Type)
	log.Info("SSLMode: ", c.SSLMode)
	log.Info("Database name: ", c.Name)
	log.Info("S3 Default Region: ", c.S3Region)
	log.Info("S3 Endpoint: ", c.S3Endpoint)
	log.Info("S3 Reports Bucket: ", c.S3ReportBucket)
	log.Info("S3 User Images Bucket: ", c.S3UserImageBucket)
	log.Info("S3 User Files Bucket: ", c.S3UserFileBucket)
	log.Info("S3 User Access Key: ", len(c.S3AccessKey), "characters")
	log.Info("S3 User Secret Access Key: ", len(c.S3SecretKey), "characters")
	log.Info("Validate Server Url: ", c.ValidatorServerHostname)
	log.Info("Validator debounce time: ", c.ValidatorDebounce)
	log.Info("Sendgrid API Key: ", len(c.MailAPIKey), " characters")
	log.Info("Mail Base Url: ", c.BaseURL)
	log.Info("Max database connections:", strconv.Itoa(c.MaxDBConnections))
}
