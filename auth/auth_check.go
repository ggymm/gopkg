package auth

func (s *Service) CheckLogin(id int64) (bool, error) {
	var (
		err     error
		value   []byte
		session *Session
	)

	// 1、获取用户 session
	session, err = s.GetSession(id, false)
	if err != nil || session == nil {
		return false, err
	}

	// 2、清理过期的 token
	invalid := make([]string, 0)
	for _, v := range session.TokenList {
		// 1、获取用户 id
		value, err = s.store.Get(s.tokenId(v.Value))
		if err != nil || value == nil {
			invalid = append(invalid, v.Value)
			continue
		}
	}
	_ = session.removeToken(invalid)

	// 3、是有效的 token 数目大于 0，即为登录状态
	return len(session.TokenList) > 0, nil
}

func (s *Service) CheckToken(token string) (bool, error) {
	// 1、获取用户 id
	var tokenId = s.tokenId(token)
	value, err := s.store.Get(tokenId)
	if err != nil || value == nil {
		return false, err
	}
	userId, timeout, _ := s.parseTokenValue(value)

	// 2、更新 token 的活跃时间
	err = s.store.Update(tokenId, s.tokenValue(userId, timeout))
	if err != nil {
		return false, err
	}

	// 3、更新 token 的过期时间（续签）
	if timeout > 0 && s.autoRenewToken {
		err = s.store.UpdateTimeout(tokenId, timeout)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
