
# Light Shadow

The shadow casting algorithm itself relies heavily on geometry and if we have to do geometry for all of the blocks every single frame, that's not efficient. 

So the first stage of the algorithm is to convert our tile map of blocks into a polygon map of edges.

![](imgs/lightshadow_polygon.png)

Once you move out of the blocky world of the tile map into the smooth world of the polygon there are lots of other advantages we get as well.  We could implement realistic physics and collision detection. We can have edges which don't align with the natual tile boundaries in the world. And of course we can have natural-looking gradients and slopes. 

I want to construct these polygon boundaries in the most efficient way possible. So ideally given a cutout of the large world I want to scan through it once to create the set of edges.

I start from the top left and I'm going to scan horizontally across the screen and when I get to end I start the next row. 

- Cell: each cell has properties:
    1. exist bool
    2. edge_exist[4]   ( N,S,E,W )
    3. edge_id[4]   (N,S,E,W)
        - id of a given edge on the north,south,east or west side from some edge pool that will create else where. 

- ![](imgs/lightshadow_scan_polygon_0.png)
    - Do I have a western neighbor ?
    - Clearly I don't have a western neighbor, so I definitely have a western edge.
    - But where do I get this edge from ? The only other place an edge could come from  is from a northern neighbor that also has a western edge, and in effect we'll take that edge and grow it downwards.
    - ![](imgs/lightshadow_scan_polygon_1.png)
    - ![](imgs/lightshadow_scan_polygon_2.png)
    - In this situation we don't have a northern neighbor.  So we're going to create a new edge.
    - So we'll add it, call edge A, to our edge pool. and our cell structure will link our edge_id to the location of the edge in the edge pool, 0.
    - ![](imgs/lightshadow_scan_polygon_3.png)

- We'll now systematically check the other sides of the tile.
- firstly we'll check the eastern edge. 
    - if I have an eastern neighbor ?  Yes. then cleary it has not an edge.
- Now we need to check our northern neighbor.
    - I don't have a northern neighbor in this case. 
    - So I need an edge. Where can I get one from?  Well I can either get 1 from my western neighbor , or I need to create one of my own.
    - ![](imgs/lightshadow_scan_polygon_4.png)
    - Since I don't have a western neighbor I'm going to have to create new edge B with id 1.
    - ![](imgs/lightshadow_scan_polygon_5.png)
- It's a similar for this southern edge too. We need to create a new edge C.
- Now we move to next cell.
    - ![](imgs/lightshadow_scan_polygon_6.png)
    - Do I need a easter / western edge ? No,
    - Do I have a northern neighbor? No, I need a edge. This time rather than creating a new one, my wetern neighbor currently already has a norther edge. So I going to extend that edge to suit my need.
    - ![](imgs/lightshadow_scan_polygon_7.png)
    - So do the southern edge.
    - ![](imgs/lightshadow_scan_polygon_8.png)
- Exactly the sam routine occurs for the 3rd cell.
- Let's run through it again for the corner cell.
    - ![](imgs/lightshadow_scan_polygon_9.png)
- Next cell is the standalone cell. 
    - ![](imgs/lightshadow_scan_polygon_10.png)

--- 

- ![](imgs/lightshadow_scan_polygon_11.png)
- Edge Pool:
    - 0:A, 1:B, 2:C, 3:D, 4:E, 5:F, 6:G, 7:H
    - 8:I, 9:J, 10:K, 11:L
- Once we've gone through all of the tiles we're interested in we'll have an edge pool that contains only the bounding edges of the shapes. 
    - And the edges are defined by a start and an end coordinate. 
    - It's very possible that edges will share a coordinate.
    - We've converted our tile map into a set of edges. **Not strictly a polygon**,  because we've not defined what's the inside and the outside of the polygon.
    - But we can assume that things don't pass through edges in our applicaiton.






