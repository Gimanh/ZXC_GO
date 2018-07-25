package router

type Route struct {
    Method  string
    Path    string
    Handler string
    Regex   string
}

func (r *Route) Exec() {

}
