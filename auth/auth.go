package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/ggymm/gopkg/log"
	"github.com/ggymm/gopkg/utils"
	"github.com/ggymm/gopkg/utils/cast"
)

const (
	ErrAuthNotInit = "auth not init"
	InvalidTimeout = "invalid timeout"
	InvalidLoginId = "invalid login id"
)

var auth *Auth

type Auth struct {
	log   zerolog.Logger
	store store

	concurrent    bool // 是否允许并发登录
	shareToken    bool // 是否允许共享 token
	maxLoginCount int  // 最大登录数（允许并发登陆，非共享 token）

	tokenName      string        // token 名称（like: ninelock-token）
	tokenTimeout   time.Duration // token 过期时间（秒）
	autoRenewToken bool          // 是否自动更新 token 的过期时间（续签）
}

type Config struct {
	LogPath string

	Store       storeType
	LocalConfig LocalConfig
	RedisConfig RedisConfig

	Concurrent    bool // 是否允许并发登录
	ShareToken    bool // 是否允许共享 token
	MaxLoginCount int  // 最大登录数（允许并发登陆，非共享 token）

	TokenName      string        // token 名称（like: ninelock-token）
	TokenTimeout   time.Duration // token 过期时间（秒）
	AutoRenewToken bool          // 是否自动更新 token 的过期时间（续签）
}

func Init(c Config) (err error) {
	auth = &Auth{
		concurrent:    c.Concurrent,
		shareToken:    c.ShareToken,
		maxLoginCount: c.MaxLoginCount,

		tokenName:      c.TokenName,
		tokenTimeout:   c.TokenTimeout,
		autoRenewToken: c.AutoRenewToken,
	}

	// 初始化日志
	auth.log = log.InitCustom(c.LogPath)

	switch c.Store {
	case Local:
		auth.store, err = newLocal(c.LocalConfig, auth.log)
	case Redis:
		auth.store, err = newRedis(c.RedisConfig, auth.log)
	}
	if err != nil {
		return err
	}
	return nil
}

// 生成 sessionId
func (a *Auth) sessionId(id int64) []byte {
	return []byte(fmt.Sprintf("%s:session:%d", a.tokenName, id))
}

// 生成 tokenId
func (a *Auth) tokenId(token string) []byte {
	return []byte(fmt.Sprintf("%s:token:%s", a.tokenName, token))
}

// 构造 token value
// value: id,timeout(second),activity time(timestamp)
func (a *Auth) tokenValue(id int64, timeout time.Duration) []byte {
	return []byte(fmt.Sprintf("%d,%d,%d", id, timeout, utils.Now()))
}

// 解析 token value
func (a *Auth) parseTokenValue(token []byte) (int64, time.Duration, int64) {
	split := strings.Split(string(token), ",")
	if len(split) < 3 {
		return 0, 0, 0
	}
	return cast.ToInt64(split[0]), time.Duration(cast.ToInt64(split[1])), cast.ToInt64(split[2])
}

// 创建 token
func (a *Auth) createToken() string {
	// 判断是否允许重复登录
	if !a.concurrent {
		// TODO
		// 需要将其他账号的登陆状态设置为无效
	} else {
		// 在允许重复登录的情况下
		// 需要判断是否允许共享 token
		if a.shareToken {
			// TODO
			// 查询是否有可用的 token
			// 如果查询到可用的 token。直接返回
		}
	}

	// 生成新的 token
	return uuid.New().String()
}
