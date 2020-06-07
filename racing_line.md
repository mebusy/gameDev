# Racing Line

1. creating a spline
2. at each spline control point, step out a particular distance on the direction perpendicular to the slope. 
    - call this distance track width
3. using those new positions to create 2 new splines

![](imgs/racingline_track_0.png)

This approach allows to use a single control spline along the middle,

How to fill the track ? Well we can just fill the triangle between adjacent points of the outer track boundray splines. For example 

![](imgs/racingline_track_1.png)

The downside is that triangles have straight edges, the track is going to look a bit clunky.

![](imgs/racingline_track_2.png)

Instead, I'm going to subdivide the spline into little steps. This allows to approximate the track with a lot more accuracy.

![](imgs/racingline_track_3.png)





