# Go - Pi using the Monte Carlo Method 
Implement with go routine to make the program parallel,


![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)
## [Monte Carlo Method](http://www.eveandersson.com/pi/monte-carlo-circle)

![](https://i.imgur.com/oZAtUD0.png)

We random put a lot of points in one square, and we can calculate that

![](https://latex.codecogs.com/gif.latex?%5Cdfrac%20%7B%20the%20%5C%20point%20%5C%20in%20%5C%20circle%7D%7Bthe%20%5C%20all%20%5C%20point%7D%24%20%3D%20%24%5Cdfrac%20%7B1/4%20%5Cpi%7D%20%7B1%7D)

![](https://latex.codecogs.com/gif.latex?%5Cpi%3D4%20*%20the%20%5C%20points%20%5C%20in%20%5C%20circle%20/%20the%20%5C%20all%20%5C%20points)


## How to implement 
```
go get github.com/kweisamx/Pi_go
cd $GOPATH/src/github.com/kweisamx/Pi_go
go build 
./pi 1000000
```

The pi will show after calculating!

