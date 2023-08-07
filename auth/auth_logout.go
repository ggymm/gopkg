package auth

func (a *Auth) Logout(id int64, device ...string) error {
	var (
		err       error
		session   *Session
		tokenList []string
	)

	// 1、获取 session
	session, err = a.GetSession(id, false)
	if err != nil {
		return err
	}

	// 2、获取 token 列表
	tokenList = session.selectToken(device...)

	// 3、从 session 中删除 token
	err = session.removeToken(tokenList)
	if err != nil {
		return err
	}

	// 4、移除 token -> id 的映射关系
	for _, token := range tokenList {
		err = a.store.Delete(a.tokenId(token))
		if err != nil {
			return err
		}
	}

	// 5、判断是否需要删除 session
	if len(session.TokenList) == 0 {
		err = a.store.Delete(session.SessionId) // sessionId == a.sessionId(id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Auth) LogoutByToken(token string) (err error) {
	var (
		value   []byte
		userId  int64
		session *Session
	)

	// 1、获取 token 信息
	tokenId := a.tokenId(token)
	value, err = a.store.Get(tokenId)
	if err != nil {
		return err
	}
	if value == nil {
		return nil
	}
	userId, _, _ = a.parseTokenValue(value)

	// 2、移除 token -> id 的映射关系
	err = a.store.Delete(tokenId)
	if err != nil {
		return err
	}

	// 3、清理 session 中的 token
	session, err = a.GetSession(userId, false)
	if err != nil {
		return err
	}
	err = session.removeToken([]string{token})
	if err != nil {
		return err
	}

	// 4、判断是否需要删除 session
	if len(session.TokenList) == 0 {
		err = a.store.Delete(session.SessionId) // sessionId == a.sessionId(id)
		if err != nil {
			return err
		}
	}

	return nil
}
