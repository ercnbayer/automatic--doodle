package encryption

import (
	"automatic-doodle/types"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"sync"

	"github.com/google/uuid"
)

type Logger interface {
	Trace(format string, args ...any)
	Info(format string, args ...any)
	Warning(format string, args ...any)
	Error(format string, args ...any)
	Fatal(format string, args ...any)
}

type ConfigModule interface {
	GetConfig(string) (string, error)
	GetConfigInt(string) (int, error)
}

var (
	module     EncryptionModule
	moduleOnce sync.Once
)

func New(
	cfg ConfigModule,
	log Logger,
) *EncryptionModule {
	moduleOnce.Do(func() {
		module = EncryptionModule{
			cfg: cfg,
			log: log,
		}
	})

	return &module
}

type EncryptionModule struct {
	log Logger
	cfg ConfigModule
}

func (mod *EncryptionModule) EncryptTokens(
	tokens types.EncryptedTokenPayload,
) (string, error) {
	str, err := json.Marshal(tokens)
	if err != nil {
		mod.log.Error(
			"Encrypt tokens marshaling failed: %s",
			map[string]interface{}{
				"error":  err,
				"tokens": tokens,
			},
		)
		return "", err
	}

	return string(base64.RawStdEncoding.EncodeToString(str)), nil
}

func (mod *EncryptionModule) DecryptTokens(
	encryptedToken string,
) (types.EncryptedTokenPayload, error) {
	var res types.EncryptedTokenPayload
	decoded, err := base64.RawStdEncoding.DecodeString(encryptedToken)
	if err != nil {
		mod.log.Warning("Token decode failed: %s", map[string]interface{}{
			"error": err,
		})
		return res, err
	}

	if err = json.Unmarshal([]byte(decoded), &res); err != nil {
		mod.log.Error("Unmarshal tokens failed: %s", map[string]interface{}{
			"error": err,
		})
		return res, err
	}

	return res, nil
}

func (mod *EncryptionModule) GenerateSalt(length int) string {
	return uuid.NewString()[:length]
}

func (mod *EncryptionModule) HashPassword(
	password string,
	salt string,
) string {
	mixedPassword := password + salt
	hash := sha512.Sum512([]byte(mixedPassword))
	passwordHash := hex.EncodeToString(hash[:])
	return passwordHash
}

func (mod *EncryptionModule) CheckPasswordHash(
	password string,
	salt string,
	hashedPassword string,
) bool {
	mixedPassword := password + salt
	hash := sha512.Sum512([]byte(mixedPassword))
	passwordHash := hex.EncodeToString(hash[:])
	return passwordHash == hashedPassword
}
