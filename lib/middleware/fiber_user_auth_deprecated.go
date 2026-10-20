package middleware

// func FiberUserAuth(
//	paramToken tokenpkg.IToken,
//	paramConfig *lib.Config,
//	paramRedis redis.Redis,
// ) func(c fiber.Ctx) error {
//	return func(c fiber.Ctx) error {
//		// Get Language [header]
//		thisLang := headerpkg.GetIndexAsString(c.GetReqHeaders()["S-Language"])
//
//		// Init Message
//		var thisMessage []messagepkg.Response
//		//fmt.Println("outside middleware")
//
//		// check Open Or Close Internal Service Connection
//		//var BreakingError = methods.InternalServiceConnectionIsOpenMiddleware(c)
//		var thisInternalConnection = methods.FiberCheckHeaderServiceInternalConnectionIsTrue(c)
//		if thisInternalConnection {
//			return c.Next()
//		}
//
//		// Get Authorization [header]
//		thisAuthorization := c.GetReqHeaders()["Authorization"]
//		// Extract Authorization String To Array || EX: (bearer AbcAbcAbcAbcAbcAbc) => ["bearer","AbcAbcAbcAbcAbcAbc"]
//		thisAuthorizationExtract := strings.Split(strings.Join(thisAuthorization, ""), " ")
//		if len(thisAuthorizationExtract) != 2 {
//			thisMessage = append(thisMessage, messagepkg.UnAuthorized(&thisLang, messagepkg.Keyword.Field["un_authorize"]))
//			return fiberpkg.ReturnAuto(c,
//				responsepkg.Code401(
//					thisMessage,
//					nil,
//				),
//			)
//		}
//
//		// Get Just Token Hash
//		thisAuthorizationToken := thisAuthorizationExtract[1]
//		fmt.Println("thisAuthorizationToken:", thisAuthorizationToken)
//		fmt.Println("paramConfig.Token.AccessSecret:", paramConfig.Token.AccessSecret)
//
//		// Check: Valid Token ?
//		thisAuthorized, err := paramToken.IsAuthorized(thisAuthorizationToken, paramConfig.Token.AccessSecret)
//		if !thisAuthorized {
//			thisMessage = append(thisMessage, messagepkg.UnAuthorized(&thisLang, messagepkg.Keyword.Field["un_authorize"]))
//			fmt.Println("err.Error():", err.Error())
//			return fiberpkg.ReturnAuto(c,
//				responsepkg.Code401(
//					thisMessage,
//					err.Error(), // TODO - Farsi error implement
//					//nil,
//				),
//			)
//		}
//		/*
//		*
//		*
//		* Token is Valid , continue Logic
//		*
//		*
//		 */
//		// Extract Token To Claims Struct || Use tokenPkg/claims.go Of The Claims
//		//userID, err := paramToken.ExtractIdFromToken(authToken, paramConfig.Token.AccessSecret)
//		thisUserClimes, err := paramToken.ParseToken(thisAuthorizationToken, paramConfig.Token.AccessSecret)
//		if err != nil {
//			thisMessage = append(thisMessage, messagepkg.UnAuthorized(&thisLang, messagepkg.Keyword.Field["un-authorize"]))
//			return fiberpkg.ReturnAuto(
//				c,
//				responsepkg.Code401(
//					thisMessage,
//					nil,
//				),
//			)
//		}
//
//		// Fetch: Get Token From Redis[redis] By key: "Consts.AccessTokenDefaultKey + thisUserClimes.UserId"
//		thisFetchCacheToken, err := paramRedis.Client().Get(
//			c.Context(),
//			//ghafpkg.BuildTokenKeyPattern(accessTokenDefaultKey, userID)).Result()
//			ghafpkg.BuildTokenKeyPattern(consts.AccessTokenDefaultKey, thisUserClimes.UserId),
//		).Result()
//		if err != nil { // Not exist on cache
//			thisMessage = append(thisMessage, messagepkg.UnAuthorized(&thisLang, messagepkg.Keyword.Field["un_authorize"]))
//			return fiberpkg.ReturnAuto(c,
//				responsepkg.Code401(
//					thisMessage,
//					nil,
//				))
//		}
//
//		// Check: Header token equal to the cached token ?
//		if thisAuthorizationToken != thisFetchCacheToken {
//			thisMessage = append(thisMessage, messagepkg.UnAuthorized(&thisLang, messagepkg.Keyword.Field["un_authorize"]))
//			return fiberpkg.ReturnAuto(c,
//				responsepkg.Code401(
//					thisMessage,
//					nil,
//				),
//			)
//		}
//
//		// Set [User.id & token] To Fiber Context
//		c.Locals(consts.ContextKeyUserId, thisUserClimes.UserId)
//		c.Locals(consts.ContextKeyToken, thisAuthorizationToken)
//
//		//fmt.Println(" climes.UserId", climes.UserId)
//		return c.Next()
//	}
// }
