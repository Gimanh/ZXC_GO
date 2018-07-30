package zxc

type Route struct {
    Method  string
    Path    string
    Handler Handler
    Regex   string
}

func (r *Route) Exec() {

}
