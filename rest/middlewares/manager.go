package middleware

import (
	"net/http"
);

type Middleware func(http.Handler) http.Handler;

type Manager struct {
	globalMiddlewares []Middleware;
}

func NewManager()*Manager{
	return &Manager{
		globalMiddlewares: make([]Middleware,0),
	}
}

func (m *Manager) Use(middlewares ... Middleware) {
	m.globalMiddlewares = append(m.globalMiddlewares, middlewares...)
}

func(m *Manager)With(next http.Handler,middlewares...Middleware) http.Handler{

	n:=next;

		for _,mid:=range middlewares{
			n=mid(n);
		}

		// for _,mid:=range m.globalMiddlewares{
		// 	n=mid(n);
		// }

		return n;
		
	
}

func(m *Manager)WrapMux(next http.Handler,middlewares...Middleware) http.Handler{
    n:=next;
    for _,mid:=range m.globalMiddlewares{
	n=mid(n);
	}
    return n;
	}