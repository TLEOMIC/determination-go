package middleware

func isWebHttp(request Http,next Next) interface{}{
	if(request.W != nil && request.R != nil){
		return next(request)
	}
	return false
}
