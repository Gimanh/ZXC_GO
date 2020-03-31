package zxc_go

type ZXC struct {
	Router *Router
}

func GO() *ZXC {
	zxc := &ZXC{
		Router: NewRouter(),
	}
	return zxc
}
