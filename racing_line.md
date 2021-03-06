# Racing Line

1. creating a spline
2. at each spline control point, step out a particular distance on the direction perpendicular to the slope. 
    - call this distance track width
3. using those new positions to create 2 new splines

![](imgs/racingline_track_0.png)

This approach allows to use a single control spline along the middle,

PS. The track width along the path are not uniform !

How to fill the track ? Well we can just fill the triangle between adjacent points of the outer track boundray splines. For example 

![](imgs/racingline_track_1.png)

The downside is that triangles have straight edges, the track is going to look a bit clunky.

![](imgs/racingline_track_2.png)

Instead, I'm going to subdivide the spline into little steps. This allows to approximate the track with a lot more accuracy.

![](imgs/racingline_track_3.png)

![](imgs/racingline_track_4.png)


Now we've handled the drawing of the track. Let's talk about how we're going to handle the racing line. And I want to take the approach where I can store the racing line using the same number of nodes that I'm using to store the track. 

Along the normal of each control point, I'm going to store an additional variable, call it **displacement**.  What this means is somewhere along this normal I'm going to have a unique value which describes how far along here do I need to go from the middle. 

So my dispalcement value gives me an indication of whereabouts on this vector do I place the node. 

![](imgs/racingline_track_5.png)

So our racing line should be bounded by the control path , because the displacement component can't go beyond the track width.  

And I want to evolve my racing line towards a solution.  

So for each frame update, I'm going to assume that my racing line is at the same position as my control spline, but then I'm going to slightly alter the displacement values per iteration of evolution in an attempt to get this system to converge to a stable solution.

---

The first thing we should try is finding the shortest path around the track. 

Here I have 3 control nodes. If we're looking for the minimum path around here , we probably want something like this. 

![](imgs/racingline_track_6.png)

Which means from our control path we want to move this node downwards.

![](imgs/racingline_track_7.png)

We know we can get the normal relatively easily from a point , but we can also do it a separate way too. 

If we take 2 vector to left/right adjacent control nodes, and add them. 

![](imgs/racingline_track_8.png)

In order to minimize the curvature and perhaps also find the shortest distance,  we want to maximaize the angle θ for every node. 

For my first proposal is that we will move the middle control node a little step along this vector created by adding the two separate vectors to the left/right adjacent nodes. 






