module example

go 1.13

require (
	github.com/go-gl/glfw v0.0.0-20200420212212-258d9bec320e
	github.com/mebusy/simpleui v0.0.0-20200606090521-6757838e81ad
    spline v0.0.0
)

replace (
    spline v0.0.0 => ./spline
)
