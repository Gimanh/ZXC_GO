package ZXC_GO

type ZXC struct {
	Router *Router
}

func GO() *ZXC {
	zxc := &ZXC{
		Router: NewRouter(),
	}
	return zxc
}
