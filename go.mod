module example/username/repo

go 1.21.5

require rsc.io/quote v1.5.2 // indirect -> 이 뜻은 직접적으로 이 패키지를 사용하지 않는다라는 것이라는데?

require (
	golang.org/x/text v0.14.0 // indirect go get golang.org/x/text
	rsc.io/sampler v1.3.1 // indirect -> go get rsc.io/sample@v1.3.1
)

//
