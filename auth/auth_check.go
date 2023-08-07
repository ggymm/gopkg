package auth

func (a *Auth) CheckLogin(id int64) bool {
	var (
		err     error
		value   []byte
		session *Session
	)

	// 1、获取用户 session
	session, err = a.GetSession(id, false)
	if err != nil || session == nil {
		return false
	}

	// 2、清理过期的 token
	invalid := make([]string, 0)
	for _, v := range session.TokenList {
		// 1、获取用户 id
		value, err = a.store.Get(a.tokenId(v.Value))
		if err != nil || value == nil {
			invalid = append(invalid, v.Value)
			continue
		}
	}
	_ = session.removeToken(invalid)

	// 3、是有效的 token 数目大于 0，即为登录状态
	return len(session.TokenList) > 0
}

func (a *Auth) CheckToken(token string) bool {
	// 1、获取用户 id
	var tokenId = a.tokenId(token)
	value, err := a.store.Get(tokenId)
	if err != nil || value == nil {
		return false
	}
	userId, timeout, _ := a.parseTokenValue(value)

	// 2、更新 token 的活跃时间
	err = a.store.Update(tokenId, a.tokenValue(userId, timeout))
	if err != nil {
		return false
	}

	// 3、更新 token 的过期时间（续签）
	if timeout > 0 && a.autoRenewToken {
		err = a.store.UpdateTimeout(tokenId, timeout)
		if err != nil {
			return false
		}
	}
	return true
}
