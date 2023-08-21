package auth

import (
	"time"

	"github.com/pkg/errors"
)

type LoginConfig struct {
	Device  string        // 设备信息
	Timeout time.Duration // 超时时间（单位s）
}

func (s *Service) Login(id int64, config LoginConfig) (string, error) {
	var (
		err     error
		token   string
		session *Session
	)

	// 1、校验参数
	if id <= 0 {
		return token, errors.New(InvalidLoginId)
	}

	// 2、创建 token 和 tokenId
	token, err = s.createToken(id, config)
	if err != nil {
		return token, err
	}

	// 3、保存 token 到 session
	session, err = s.GetSession(id, true)
	err = session.saveToken(token, config.Device)
	if err != nil {
		return token, err
	}

	// 4、保存 token -> id,timeout,activity time 的映射关系
	err = s.store.Put(s.tokenId(token), s.tokenValue(id, config.Timeout), config.Timeout)
	if err != nil {
		return token, err
	}

	// 5、更新 token 的过期时间（续签）
	if config.Timeout > 0 && s.autoRenewToken {
		err = s.store.UpdateTimeout(s.tokenId(token), config.Timeout)
		if err != nil {
			return token, err
		}
	}

	// 6、判断是否超过了最大登陆数
	if s.maxLoginCount > 0 && len(session.TokenList) > s.maxLoginCount {
		// TODO
		// 如果超过，选择以下策略删除 token
		// 1）删除最先登陆
		// 2) 删除最先过期
		// 3）删除最不活跃
	}

	// 7、返回 token
	return token, nil
}
