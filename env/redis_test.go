// +build redis

package env_test

import (
	"os"
	"testing"

	"github.com/go-redis/redis"
	"github.com/lestrrat-go/config/env"
	envload "github.com/lestrrat-go/envload"
	"github.com/stretchr/testify/assert"
)

func TestRedisOptions(t *testing.T) {
	l := envload.New()
	defer l.Restore()

	os.Setenv("REDIS_ADDR", "redis:16379")

	var options redis.Options
	if err := env.NewDecoder(env.System).Prefix("REDIS").Decode(&options); !assert.NoError(t, err, "Decode should succeed") {
		return
	}

	if !assert.Nil(t, options.TLSConfig, "TLSConfig should be nil") {
		return
	}

	t.Logf("%#v", options)
}
