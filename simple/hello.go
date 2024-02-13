package simple

type SayHello interface {
	Hello(name string) string
}

type HelloService struct {
	SayHello SayHello
}

type SayHelloImpl struct {
}

func (s *SayHelloImpl) Hello(name string) string {
	return "Hello" + name
}

func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

func NewSayHelloService(SayHello SayHello) *HelloService {
	return &HelloService{
		SayHello: SayHello,
	}
}
